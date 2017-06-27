package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file1, _ := os.Open("../12/col1.txt")
	file2, _ := os.Open("../12/col2.txt")

	defer file1.Close()
	defer file2.Close()

	sc1 := bufio.NewScanner(file1)
	sc2 := bufio.NewScanner(file2)

	col1 := make([]string, 24)
	for i := 0; sc1.Scan(); i++ {
		if err := sc1.Err(); err != nil {
			// エラー処理
			break
		}

		col1[i] = sc1.Text()
	}

	col2 := make([]string, 24)
	for i := 0; sc2.Scan(); i++ {
		if err := sc2.Err(); err != nil {
			// エラー処理
			break
		}

		col2[i] = sc2.Text()
	}

	var out string
	for i := 0; len(col1) > i; i++ {
		out += fmt.Sprintf("%s\t%s\n", col1[i], col2[i])
	}
	ioutil.WriteFile("col1col2.txt", []byte(out), os.ModePerm)
}
