package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if validIPv4(IP) {
		return "IPv4"
	}
	if validIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4(IP string) bool {
	strArr := strings.Split(IP, ".")
	if len(strArr) != 4 {
		return false
	}
	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil || num > 255 || num < 0 {
			return false
		}
		newStr := fmt.Sprint(num)
		if str != newStr {
			return false
		}
	}
	return true
}

func validIPv6(IP string) bool {
	strArr := strings.Split(IP, ";")
	if len(strArr) != 8 {
		return false
	}
	for _, str := range strArr {
		if len(str) == 0 || len(str) > 4 {
			return false
		}
		for i := 0; i < len(str); i++ {
			if !(str[i] >= '0' && str[i] <= '9') &&
				!(str[i] >= 'a' && str[i] <= 'f') &&
				!(str[i] >= 'A' && str[i] <= 'F') {
				return false
			}
		}
	}
	return true

}

func validIPv4Address(IP string) bool {
	strArr := strings.Split(IP, ".")
	if len(strArr) != 4 {
		return false
	}
	for _, str := range strArr {
		if num, err := strconv.Atoi(str); err != nil || num > 255 || num < 0 {
			return false
		} else if strconv.Itoa(num) != str {
			return false
		}
	}
	return true
}
func validIPv6Address(IP string) bool {
	strArr := strings.Split(IP, ":")
	if len(strArr) != 8 {
		return false
	}
	re := regexp.MustCompile(`^([0-9]|[a-f]|[A-F])+$`)
	for _, str := range strArr {
		if len(str) == 0 || len(str) > 4 {
			return false
		}
		if !re.MatchString(str) {
			return false
		}
	}
	return true
}
