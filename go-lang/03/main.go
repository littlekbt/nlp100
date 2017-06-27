package main

import (
	"fmt"
	"strings"
)

// "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．
func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	strs := strings.Split(str, " ")
	slice := make([]int, len(strs))
	for i, s := range strs {
		slice[i] = len(s)
	}

	fmt.Println(slice)
}
