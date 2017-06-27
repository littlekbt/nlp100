package main

import (
	"fmt"
)

// 与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ．
//
// 英小文字ならば(219 - 文字コード)の文字に置換
// その他の文字はそのまま出力
// この関数を用い，英語のメッセージを暗号化・復号化せよ．

func cipher(str string) (newStr string) {
	for _, runeValue := range str {
		if runeValue >= 97 && runeValue <= 122 {
			newStr += string(219 - runeValue)
		} else {
			newStr += string(runeValue)
		}
	}
	return
}

func main() {
	encryption := cipher("abcdefghijklmnopqrstuvwxyz-@.")
	fmt.Println(encryption)
	decryption := cipher(encryption)
	fmt.Println(decryption)
}
