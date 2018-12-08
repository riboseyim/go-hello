package main

import (
	"flag"
	"fmt"
)

const (
	VERSION = "-1.0-release-201802"
	AUTHOR  = "@RiboseYim"
)

func main() {

	task := flag.String("task", "", "batch_cmd.eg")
	cmdname := flag.String("cmdname", "", "cmd_test.eg")

	flag.Parse()
	fmt.Printf("Welcome to [ Ops System %s ] Author:%s \n -act:%s \n\n", VERSION, AUTHOR, *task)

	switch *task {
	case "batchcmd":
		NewBatchCmdTask(*task, *cmdname)
	case "consumer":

	}

}
