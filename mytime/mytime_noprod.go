// +build !prod

package mytime

import "time"

// FixTime は Now() で返す値を、引数に受け取った time.Time の値に固定します
func FixTime(t time.Time) {
	fixTime = &t
}

// Reset は fixTime を 初期化します
func Reset() {
	fixTime = nil
}
