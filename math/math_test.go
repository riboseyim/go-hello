package test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCountInt(t *testing.T) {
	aa := 100
	bb := 200
	cc := aa / bb
	log.Printf("aa / bb = %d", cc)
}

func TestCountFloat(t *testing.T) {
	aa := 100.0
	bb := 200.0
	cc := aa / bb
	log.Printf("aa / bb = %v", cc)
}

func Test_strings_1(t *testing.T) {
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
