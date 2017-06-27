package main

import (
	//	"bufio"
	//	"io/ioutil"
	//	"os"
	//	"strconv"
	"fmt"
)

// aa -> ab -> ac ... az -> ba ... -> zz -> aaa ... -> aaz -> aba -> abb
// 最後の文字がz以外だったらコードポイントを1増やす。
// 最後の文字がzだったら左隣のコードポイントを1増やして、自分はaに戻る。
// 左隣がzだったら、桁数を１つ増やしてaに戻る。
func mkFileNamePart(part string, pos int) string {

	return part[0:len(part)-1] + string(lastPart)
}

func main() {
	b := "aa"
	for i := 0; i < 680; i++ {
		fmt.Println(b)
		b = mkFileNamePart(b)
	}
}

// func main() {
// 	n, _ := strconv.Atoi(os.Args[1])
// 	file, _ := os.Open("../hightemp.txt")
// 	defer file.Close()
//
// 	sc := bufio.NewScanner(file)
//
// 	buf := ""
// 	fileNamePart := "aa"
// 	for i := 1; sc.Scan(); i++ {
// 		if err := sc.Err(); err != nil {
// 			// エラー処理
// 			break
// 		}
//
// 		buf += sc.Text() + "\n"
//
// 		if i%n == 0 {
// 			ioutil.WriteFile("x"+fileNamePart+".txt", []byte(buf), os.ModePerm)
// 			fileNamePart = mkFileNamePart(fileNamePart)
// 			buf = ""
// 		}
// 	}
// }
