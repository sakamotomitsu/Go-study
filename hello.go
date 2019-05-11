package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!!")
}

/*

// helloパッケージに所属
package hello

// 別パッケージから使用できる関数名 先頭が大文字になっているもの
func Hello() {
	// ...
}

// 別パッケージから使用できない関数名
func hello() {
	// ...
}



// 変数の型　大きく分けると「値型」「参照型」「ポインタ型」の 3 種類
「参照型」は少し特殊で、Go では「スライス」「マップ」「チャネル」という 3 つのデータ構造のいずれか

// 明示的な意義
	// int型の変数nを定義する
	var n int
	// 複数定義もできる
	var x, y, z int
	// まとめて指定ができる
	// int型の変数x, yとstring型の変数nameを定義する
	var (
		x,y int
		name string
	)
// 暗黙的な定義
	// int型の変数aに1を代入
	var a = 1
	// これでもええで
	a := 1
*/
