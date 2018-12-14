// common
package common

import (
	"regexp"
	"strconv"
	"strings"
)

const buToMoDecimals int64 = 8
const maxLen = 21
const maxDecimals = 18

func BU2MO(amount string) string {
	return UnitWithDecimals(amount, 8)
}
func MO2BU(amount string) string {
	return UnitWithoutDecimals(amount, 8)
}
func UnitWithDecimals(amount string, decimals int) string {
	if decimals > maxDecimals || decimals < 0 {
		return ""
	}
	match, err := regexp.MatchString("(^0(.[0-9]{0,"+strconv.FormatInt(int64(decimals), 10)+"}[1-9])?$)|"+"(^[1-9][0-9]{0,"+strconv.FormatInt(int64(maxDecimals-decimals), 10)+"}(.[0-9]{0,"+strconv.FormatInt(int64(decimals), 10)+"}[1-9])?$)", amount)
	if err != nil || match == false {
		return ""
	}
	if decimals == 0 {
		return amount
	}
	var addStr string
	for i := 0; i < decimals; i++ {
		addStr += "0"
	}
	var beforeStr string
	var afterStr string
	var tempArray []string = strings.Split(amount, ".")
	if len(tempArray) > 1 {
		beforeStr = tempArray[0]
		afterStr = tempArray[1]
	} else {
		amount = amount + addStr
		_, err := strconv.ParseInt(amount, 10, 64)
		if err != nil {
			return ""
		}
		return amount
	}
	endIndex := len(afterStr)
	if endIndex > decimals {
		endIndex = decimals
		afterStr = afterStr[0:decimals]
	} else {
		addZero := ""
		for i := 0; i < decimals-len(afterStr); i++ {
			addZero += "0"
		}
		afterStr += addZero
	}
	amount = delStartZero(beforeStr + afterStr)
	_, err = strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return ""
	}
	return amount
}

func UnitWithoutDecimals(amount string, decimals int) string {
	_, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return ""
	}
	match, err := regexp.MatchString("(^0?$)|(^[1-9][0-9]{0,18}?$)", amount)
	if err != nil || match == false {
		return ""
	}
	if decimals > maxDecimals || decimals < 0 {
		return ""
	}
	if decimals == 0 {
		return amount
	}
	var beforeStr string
	var afterStr string
	if len(amount) > decimals {
		afterStr = "." + amount[len(amount)-decimals:]
		beforeStr = amount[0 : len(amount)-decimals]
	} else {
		addZero := ""
		for i := 0; i < decimals-len(amount); i++ {
			addZero += "0"
		}
		afterStr = "0." + addZero + amount
	}
	amount = delEndsZero(beforeStr + afterStr)
	if amount[len(amount)-1:] == "." {
		amount = amount[:len(amount)-1]
	}
	return amount
}
func delEndsZero(src string) string {
	if src[len(src)-1:len(src)] == "0" {
		return delEndsZero(src[0 : len(src)-1])
	} else {
		return src
	}
}
func delStartZero(src string) string {
	if src[0:1] == "0" && len(src) != 1 {
		return delStartZero(src[1:])
	} else {
		return src
	}
}
