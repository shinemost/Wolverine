package main

import (
	"context"
	"fmt"
	"hjfu/Wolverine/route"
	"hjfu/Wolverine/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//func main() {

// arr := strings.Split("a,b,c", ",")
// arr2 := strings.SplitN("a,b,cddd", ",", 2)
// fmt.Println(arr)
// fmt.Println(arr2)

// sce := []string{"AA", "BB", "CC"}
// ll := strings.Join(sce, "--")
// fmt.Println(ll)

// str := "HELLO 22 13213 "
// reg := regexp.MustCompile("2")
// ss := reg.FindAllString(str, -1)
// fmt.Println(ss)

// t := time.Now()
// fmt.Println(t)
// fmt.Printf("当前的时间是：%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

// ls := t.Format("2006-01-02 15:04:05")
// fmt.Println("---------------------")
// fmt.Println(ls)

// mysql.InitMysqlX()
// mysql.BetterInsert()

// users := []domain.SysUserInfo{
// 	{
// 		UserName:  "蔡建军",
// 		LoginName: "jjcai5",
// 	},
// 	{
// 		UserName:  "陈霄峰",
// 		LoginName: "xfchen8",
// 	},
// 	{
// 		UserName:  "秦坤",
// 		LoginName: "kunqin",
// 	},
// }
// mysql.InsertMoreUsersX(users)
// mysql.SearchByIDsX([]int{1, 105, 106})']
// redis.InitClient()
// redis.RedisDemo2()
// redis.WatchDemo()

// if err != nil {
// 	panic(err)
// }

// mysql.QueryUser(db)
// mysql.QueryUserRows(db)\
// mysql.InsertUser(db)
// mysql.PrepareTest(db)
// mysql.TransactionDemo(db)

// defer db.Close()
// fmt.Println("db connect success")

//}

// func main() {
// 	// loggers.MainMethod()
// 	viper.Set("fileDr", "./")
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("json")
// 	viper.AddConfigPath(".")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(viper.Get("name"))

// 	// viper.Set("age", "181")
// 	// viper.WriteConfigAs("config.json")
// 	// viper.WatchConfig()
// 	// viper.OnConfigChange(func(in fsnotify.Event) {
// 	// 	fmt.Println("config changed", in.Name)
// 	// })

// 	if !viper.IsSet("house") {
// 		fmt.Println("no house key")
// 	}
// 	// fmt.Println(viper.Get("address.loaction"))

// 	var config domain.Config2
// 	viper.Unmarshal(&config)
// 	fmt.Println(config)

// }

func main() {
	// 加载配置文件
	if err := setting.InitConfig(); err != nil {
		fmt.Printf("init settings failed:%s \n", err)
		return
	}
	// 初始化日志
	if err := setting.InitLogger(); err != nil {
		fmt.Printf("init settings failed:%s \n", err)
		return
	}
	zap.L().Debug("logger init success")
	defer zap.L().Sync()

	// 初始化mysql
	if err := setting.InitMysql(setting.Config.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed:%s \n", err)
		return
	}
	zap.L().Debug("mysql init success")
	// 初始化redis
	if err := setting.InitRedis(setting.Config.RedisConfig); err != nil {
		fmt.Printf("init redis failed:%s \n", err)
		return
	}
	zap.L().Debug("redis init success")

	// 不要遗漏2个 db的close
	defer setting.CloseMysql()
	defer setting.CloseRedis()

	// 注册路由
	r := route.Setup()
	// r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))

	// 启动服务 （优雅关机）
	// 启动服务 （优雅关机）

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown server")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")

}
