package main

import "fmt"

// 引数x, y, zを受け取り「x時のyはz」という文字列を返す関数を実装せよ．さらに，x=12, y="気温", z=22.4として，実行結果を確認せよ．

func mkWord(x int64, y string, z float64) string {
	return fmt.Sprintf("%d時の%sは%.1f", x, y, z)
}
func main() {
	fmt.Println(mkWord(12, "気温", 22.4))
}
