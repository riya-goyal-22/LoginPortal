package utils

import "strings"

var Nums = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var Specials = []string{"@", "#", "$", "%", "&", "*"}

type Person struct {
	Name string
	Pass string
}

func IsPassCorrect(s string) bool {
	if len(s) >= 13 {
		numFlag := 0
		SpecialFlag := 0
		for _, i := range Nums {
			if strings.Contains(s, i) {
				numFlag++
			}
		}
		for _, i := range Specials {
			if strings.Contains(s, i) {
				SpecialFlag++
			}
		}
		if numFlag != 0 && SpecialFlag != 0 {
			return true
		} else {
			return false
		}
	}
	return false
}
