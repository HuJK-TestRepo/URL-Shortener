package Tools

import (
	"math/rand"
	"strings"
	"time"
)

func RandomString(length int) string {
	rand.Seed(time.Now().Unix())

	var output strings.Builder

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJULMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return (output.String())
}

func CheckLinkValid(Link string) bool {
	if strings.HasPrefix(Link, "http://") || strings.HasPrefix(Link, "https://") {
		return true
	} else {
		return false
	}
}
