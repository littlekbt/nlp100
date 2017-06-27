package main

import (
	"fmt"
	"math/rand"
	"time"
)

// スペースで区切られた単語列に対して，各単語の先頭と末尾の文字は残し，それ以外の文字の順序をランダムに並び替えるプログラムを作成せよ．
// ただし，長さが４以下の単語は並び替えないこととする．
// 適当な英語の文（例えば"I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."）を与え，その実行結果を確認せよ．

func head(str string) (headStr string, others string) {
	headStr = string(str[0])
	others = str[1:len(str)]

	return
}

func tail(str string) (others string, tailStr string) {
	tailStr = string(str[len(str)-1])
	others = str[0 : len(str)-1]

	return
}

func randStr(str string) (result string) {
	headStr, others := head(str)
	others, tailStr := tail(others)

	strNum := len([]rune(others))
	rand.Seed(time.Now().UnixNano())

	for {
		// 最後の文字なら終わり。
		if strNum == 0 {
			result += others
			break
		}

		n := rand.Intn(strNum)
		result += string(others[n])

		// othersから文字を抜いて次のループへ
		// 文字列を抜き出して行けば、重複して文字を取り出すことはない。
		if n == 0 {
			// 最初の文字を抜き出したら
			others = others[1:strNum]
		} else if n == strNum {
			// 最後の文字を抜き出したら
			others = others[0 : strNum-1]
		} else {
			// 間の文字を抜き出したら
			// 取り出したところまで + それ以降
			// "hoge"[0:2] => ho
			others = others[0:n] + others[n+1:strNum]
		}

		strNum -= 1
	}

	result = headStr + result + tailStr
	return
}

func main() {
	// あえてスプリットしてforしないやり方。いつものやり方だから。

	tmp := ""
	result := ""
	str := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	count := 0
	for _, e := range str {
		s := string(e)
		// 空白だったら区切る。
		if s == " " {
			// 4以上だったら、先頭と末尾以外をランダムにする。
			if 4 <= count {
				result += (randStr(tmp) + " ")
			} else {
				result += (tmp + " ")
			}
			// 初期化
			tmp = ""
			count = 0
		} else {
			tmp += s
			count++
		}
	}

	result += tmp
	fmt.Println(str)
	fmt.Println(result)
}
