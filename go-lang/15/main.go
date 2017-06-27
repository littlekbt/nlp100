package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func newScanner() *bufio.Scanner {
	file, err := os.Open("../hightemp.txt")
	if err != nil {

	}

	return bufio.NewScanner(file)
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	sc := newScanner()
	ln := 0
	for ; sc.Scan(); ln++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
	}

	sc = newScanner()
	tail := ln - n
	for i := 0; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}

		if i < tail {
			continue
		}

		fmt.Println(sc.Text())
	}
}
