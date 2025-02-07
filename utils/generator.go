package utils

import "math/rand/v2"

const letters string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenShortPath(length int) string {
	var lettersLen int = len(letters)
	shortPath := make([]byte, 0, length)

	for i := 0; i < length; i++ {
		shortPath = append(shortPath, letters[rand.IntN(lettersLen)])
	}

	return string(shortPath)
}
