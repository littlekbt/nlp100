package main

import "fmt"
import "golang.org/x/exp/utf8string"

// 「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．

func main() {
	// utf8を扱うためにはこれ必要
	str := utf8string.NewString("パタトクカシーー")
	fmt.Printf("%s%s%s%s\n", str.Slice(0, 1), str.Slice(2, 3), str.Slice(4, 5), str.Slice(6, 7))
}
