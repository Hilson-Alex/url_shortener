package base62

var digits = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const base int64 = 62

func Encode(num int64) string {
	var digit = num
	var result = ""
	for digit/base > 0 {
		result = string(digits[digit%base]) + result
		digit = digit / base
	}
	return result
}
