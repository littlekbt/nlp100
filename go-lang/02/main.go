package main

import "fmt"
import "golang.org/x/exp/utf8string"

func main() {
	str1 := utf8string.NewString("パトカー")
	str2 := utf8string.NewString("タクシー")

	str := ""
	j := 0
	for i := 0; i < str1.RuneCount()*2; i++ {
		if i%2 == 0 {
			str += str1.Slice(j, j+1)
		} else {
			str += str2.Slice(j, j+1)
			j++
		}
	}

	fmt.Println(str)
}
