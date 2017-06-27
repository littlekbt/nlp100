package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// タブ1文字につきスペース1文字に置換せよ．確認にはsedコマンド，trコマンド，もしくはexpandコマンドを用いよ．

func main() {
	bytes, e := ioutil.ReadFile("../hightemp.txt")
	if e != nil {
		panic(e)
	}

	str := strings.Replace(string(bytes), "\t", " ", -1)
	fmt.Println(str)
}
