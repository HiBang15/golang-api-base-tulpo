package utils

import (
	"regexp"
	"unicode"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_\\`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)+$")
//var passwordRegex = regexp.MustCompile("^(\\?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\_\\@\\&\\#\\&\\^\\*\\!\\d]{8,}$")
//var passwordRegex = regexp.MustCompile("^(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$")
var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9\\_\\.]{8,}$")

func IsEmailValid(e string) bool {
	if len(e) < 3 || len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

//func IsPassword(e string) bool {
//	if len(e) < 8 {
//		return false
//	}
//	return passwordRegex.MatchString(e)
//}

func IsPassword(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

func IsUsername(e string) bool {
	if len(e) < 8 || len(e) > 20 {
		return false
	}
	return usernameRegex.MatchString(e)
}

func IsPrice(price float32) bool  {
	if price <= 0 {
		return false
	}
	return true
}

func IsQuantity(quantity int32) bool {
	if quantity <= 0 {
		return false
	}
	return true
}
