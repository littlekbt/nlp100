package main

import (
	"fmt"
	"strings"
)

// 転置インデックスのキーの切り出しを、辞書や構文解析に基づくのではなく、単に一定の文字数で切り出した語を入れることで作る方式のことをN-gramという。
// Nとは、切り出す文字の単位が複数ありえるということを意味する。
// 与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．

func ngram(str string, n int64) ([][]string, [][]string) {
	// 単語で分割
	split_word := func(s string) []string {
		return strings.Split(s, " ")
	}

	// 文字で分割
	split_char := func(s string) []string {
		return strings.Split(strings.Replace(s, " ", "_", -1), "")
	}

	splited_word := split_word(str)
	splited_char := split_char(str)

	word_ngram := make([][]string, 0)
	char_ngram := make([][]string, 0)

	// どっちもnの単位でn-gramを作る。
	for i, _ := range splited_word {
		if i == (len(splited_word) - (int(n) - 1)) {
			break
		}
		word_ngram = append(word_ngram, splited_word[i:int64(i)+n])
	}

	for i, _ := range splited_char {
		if i == (len(splited_char) - (int(n) - 1)) {
			break
		}
		char_ngram = append(char_ngram, splited_char[i:int64(i)+n])
	}

	return word_ngram, char_ngram
}

func main() {
	msg := "I am an NLPer"
	word_ngram, char_ngram := ngram(msg, 2)

	fmt.Println("word_gram: ", word_ngram)
	fmt.Println("char_gram: ", char_ngram)
}
