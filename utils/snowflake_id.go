package utils

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func GenId() int64 {
	return node.Generate().Int64()
}

func Init(machineId int64) (err error) {
	// var st time.Time
	// st, err = time.Parse("2006-01-02", startTime)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("st:", st)
	// snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return err
}
