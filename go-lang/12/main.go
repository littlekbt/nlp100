package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// 各行の1列目だけを抜き出したものをcol1.txtに，2列目だけを抜き出したものをcol2.txtとしてファイルに保存せよ．確認にはcutコマンドを用いよ．

func main() {
	file, err := os.Open("../hightemp.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	var col1 string
	var col2 string

	for sc.Scan() {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
		f := strings.Fields(sc.Text())
		col1 += f[0] + "\n"
		col2 += f[1] + "\n"
	}

	ioutil.WriteFile("col1.txt", []byte(col1), os.ModePerm)
	ioutil.WriteFile("col2.txt", []byte(col2), os.ModePerm)
}
