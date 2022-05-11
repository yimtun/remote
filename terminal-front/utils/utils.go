package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func StrFartory(str string) ([]string, int) {

	b := strings.TrimSpace(str)
	a := regexp.MustCompile(" +").Split(b, -1)
	if a[0] == "rs" {

	}
	return a, len(a)

}

func IsIP(host string) (b bool) {
	/*	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
					}
							return true
	*/

	parts := strings.Split(host, ".")

	if len(parts) < 4 {
		return false
	}

	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}

	}
	return true
}
