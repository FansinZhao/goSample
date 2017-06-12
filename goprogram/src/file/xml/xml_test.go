package main

import (
	"testing"
)

func TestIfelse(t *testing.T) {
	Ifelse()
	defer func() {
		if r := recover(); r != nil {
			//TODO
		}
	}()
}

func Ifelse() bool {
	var a = 5
	if a == 5 {
		return true
	} else {
		return false
	}
}
