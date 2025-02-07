package utils

import "testing"

func TestGenShortPath(t *testing.T) {
	for i := 0; i < 1000; i++ {
		shortPath := GenShortPath(10)
		t.Log(shortPath)
	}
}
