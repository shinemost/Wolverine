package main

import (
	"fmt"
	"hjfu/Wolverine/domain"

	"github.com/spf13/viper"
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

func main() {
	// loggers.MainMethod()
	viper.Set("fileDr", "./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("name"))

	// viper.Set("age", "181")
	// viper.WriteConfigAs("config.json")
	// viper.WatchConfig()
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	fmt.Println("config changed", in.Name)
	// })

	if !viper.IsSet("house") {
		fmt.Println("no house key")
	}
	// fmt.Println(viper.Get("address.loaction"))

	var config domain.Config2
	viper.Unmarshal(&config)
	fmt.Println(config)

}
