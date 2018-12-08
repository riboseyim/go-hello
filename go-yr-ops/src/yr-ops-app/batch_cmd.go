package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func NewBatchCmdTask(task string, cmdname string) {
	log.Println("-----Exec NewBatchCmdTask(%s,%s)", task, cmdname)
	cmdfile := "./batch_cmd/" + cmdname + ".sh"
	log.Println("-----Load cmdfile:%s", cmdfile)
	data, err := ReadLine(cmdfile)
	if data != nil {
		for i := range data {
			cmd := data[i]
			out, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				log.Printf("Failed to execute command: %s", cmd)
			}
			log.Println(string(out))
		}
	} else {
		log.Println("-----load cmd err:%s", err)
	}
}

func ReadLine(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	var result []string
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				return result, nil
			}
			return nil, err
		}
		log.Println("-----ReadLine:%s", line)
		result = append(result, line)
	}
	return result, nil
}
