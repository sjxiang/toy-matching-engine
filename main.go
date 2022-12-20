package main

import (
	"fmt"

	"toy-matching-engine/conf"
)


func main() {
	conf.Load()
	fmt.Println(conf.Cfg.Log.FileDir)
}