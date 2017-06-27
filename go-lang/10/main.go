package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../hightemp.txt")
	if err != nil {

	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	i := 0
	for ; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
	}

	fmt.Println(i)
}
