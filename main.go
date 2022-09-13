package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// ファジングを試す。ファジングとは？
// プログラムへの入力を継続的に操作して、バグや脆弱性を見つけるソフトウェアテスト手法の一つ。
// セミランダムな入力を与えることで、予期しないエッジケースの不具合を見つけるのに有効な手段

func main() {
	input := "I Prepared a Sentence for the test"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("オリジナル文章： %q\n", input)
	fmt.Printf("逆： %q\n, エラー： %v\n", rev, revErr)
	fmt.Printf("再リバース： %q\n, エラー： %v\n", doubleRev, doubleRevErr)
}

// 引数はstringの文字rつ、文字列を受け取ったら一度に1バイトずつループし、最後に逆の文字列を返す
// func Reverse(s string) string {
// 	b := []byte(s)
// 	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
// 		b[i], b[j] = b[j], b[i]
// 	}
// 	return string(b)
// }

// 関数の修正
// func Reverse(s string) string {
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// 関数の再修正
// func Reverse(s string) string {
// 	fmt.Printf("input: %q\n", s)
// 	r := []rune(s)
// 	fmt.Printf("runes: %q\n", r)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// 関数の再再修正
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
