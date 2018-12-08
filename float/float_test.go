package test

import (
	"log"
	"strconv"
	"testing"
)

func Test_float_to_string(t *testing.T) {
	a := float64(100.01)
	target := strconv.FormatFloat(a, 'f', -1, 32)
	log.Println(a, "--to:", target)
}
