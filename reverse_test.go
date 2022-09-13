package main

import (
	"testing"
	"unicode/utf8"
)

// 一般的な単体テスト
// func TestReverse(t *testing.T) {
// 	testcase := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}
// 	for _, tc := range testcase {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("逆： %q, want %q", rev, tc.want)
// 		}
// 	}
// }

// ファジング
// ファジングは入力と出力を固定していないので、出力が期待した出力かどうかを確認できない。
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		// f.Add を使用してシードコーパスを提供する
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}

		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}

		t.Logf(
			"ルーンの数: orig=%d, rev=%d, doubleRev=%d",
			utf8.RuneCountInString(orig),
			utf8.RuneCountInString(rev),
			utf8.RuneCountInString(doubleRev),
		)
		// このファジングでチェックされる項目は２つ
		// 2回反転すると、元の文字列になっているか？
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		// 逆文字列はUTF-8形式かどうか
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
