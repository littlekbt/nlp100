package main

import (
	"fmt"
)

// 文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．

func main() {
	str := "stressed"
	str2 := ""
	// stringをrangeで回すと、人文字単位のruneが取れる。
	for _, s := range str {
		str2 = fmt.Sprintf("%s%s", string(rune(s)), str2)
	}

	fmt.Println(str2)
}
