package pkg

import (
	"math/rand"
	"os"
	"strings"
)

func GenerateRandomString(length int) string {
	const CHARACTER_POOL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890-_=+:;?><|"
	key := ""
	for i := 0; i < length; i++ {
		idx := rand.Intn(len(CHARACTER_POOL))
		key += string(CHARACTER_POOL[idx])
	}
	return key
}

func IsValidDir(path string) bool {
	if path == "." {
		return true
	}
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return info.IsDir()
}

func TrimNewLineChar(s *string) {
	*s = strings.TrimSuffix(*s, "\n")
}
