package main

import "testing"

func TestMainFunc(t *testing.T) {
	go main()
	// 実行できるが終了は指示できない
}
