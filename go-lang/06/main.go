package main

import (
	"fmt"
	"strings"
)

func ngram(str string, n int64) ([][]string, []string) {
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
	char_ngram := make([]string, 0)

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
		char_ngram = append(char_ngram, splited_char[i]+splited_char[int64(i)+1])
	}

	return word_ngram, char_ngram
}

func exist(elem string, array []string) (exist bool) {
	exist = false
	for _, e := range array {
		if e == elem {
			exist = true
			return
		}
	}
	return
}

func uniq(array []string) (uniq_array []string) {
	for _, e := range array {
		if !exist(e, uniq_array) {
			uniq_array = append(uniq_array, e)
		}
	}

	return
}

// 和集合
// 集合Bになかったら集合Aの要素をCに加える
// 集合BをすべてCに加える
func union(array1 []string, array2 []string) (union_array []string) {
	for _, e := range array1 {
		if !exist(e, array2) {
			union_array = append(union_array, e)
		}
	}
	union_array = append(union_array, array2...)
	union_array = uniq(union_array)
	return
}

// 積集合
// 集合Bの要素なら、集合Aの要素をCに加える。
func intersection(array1 []string, array2 []string) (intersection_array []string) {
	for _, e := range array1 {
		if exist(e, array2) {
			intersection_array = append(intersection_array, e)
		}
	}

	intersection_array = uniq(intersection_array)

	return
}

// 差
func difference(array1 []string, array2 []string) (difference_array []string) {
	for _, e := range array1 {
		if !exist(e, array2) {
			difference_array = append(difference_array, e)
		}
	}

	difference_array = uniq(difference_array)

	return
}

func main() {
	msg1 := "paraparaparadise"
	msg2 := "paragraph"
	_, char_ngram1 := ngram(msg1, 2)
	_, char_ngram2 := ngram(msg2, 2)

	fmt.Println("char_gram1: ", char_ngram1)
	fmt.Println("char_gram2: ", char_ngram2)

	fmt.Println("union: ", union(char_ngram1, char_ngram2))
	fmt.Println("intersection: ", intersection(char_ngram1, char_ngram2))
	fmt.Println("difference: ", difference(char_ngram1, char_ngram2))
}
