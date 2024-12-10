package utils

import (
	"math/rand"
	"strings"
)

func RandomString(noOfChars int) string {
	alphabets := "abcdefghijklmnopqrstuvwxyz"

	var sb strings.Builder

	for i := 0; i < noOfChars; i++ {
		c := alphabets[rand.Intn(len(alphabets))]
		sb.WriteByte(c)
	}
	return sb.String()
}