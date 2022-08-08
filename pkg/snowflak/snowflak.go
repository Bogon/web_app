package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// Init > Initialize the snowflake generator with the given start time and machine ID
func Init(startTime string, machineID int64) (err error) {

	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

// GetId It generates a unique ID.
func GetId() int64 {
	return node.Generate().Int64()
}

func main() {

	if err := Init("2022-08-08", 1); err != nil {
		fmt.Println("init failed, error:", err)
		return
	}
	id := GetId()
	fmt.Printf("id: %v\n", id)
}
