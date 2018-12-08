package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	for i := 1; i < 100; i++ {
		url := "http://ldzl.people.com.cn/dfzlk/front/personPage"
		if i > 0 {
			url = url + fmt.Sprintf("%d", i) + ".htm"
			log.Printf(">>>>>%s\n", url)

			if i%5 == 0 {
				fmt.Printf("-----")
				time.Sleep(5 * time.Second)
			}
		}
	}

}
