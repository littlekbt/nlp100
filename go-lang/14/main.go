package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	head, _ := strconv.Atoi(os.Args[1])

	file, err := os.Open("../hightemp.txt")
	if err != nil {

	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	i := 0
	for ; sc.Scan() && i < head; i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
		fmt.Println(sc.Text())
	}

}
