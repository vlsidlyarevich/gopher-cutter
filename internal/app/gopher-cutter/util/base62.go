package util

import (
	"log"
	. "math"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base     = len(alphabet)
)

func Encode(num int) (encStr string) {
	//base62.StdEncoding.EncodeToString([]byte(url))
	var result string
	for num > 0 {
		result += string(alphabet[num%base])
		num /= base
	}

	return reverse(result)
}

func Decode(encStr string) (num int) {
	var res int

	for i, c := range encStr {
		index := strings.Index(alphabet, string(c))
		if index == -1 {
			log.Panicf("Character %s not supported", string(c))
		}
		res += index * int(Pow(float64(base), float64(len(encStr)-1-i)))
	}

	return res
}
