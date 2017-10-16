package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(myAtoi("  98 "))
}

func myAtoi(str string) int {
	var started bool = false
	var sign bool = true
	var out int = 0
	var size int = 0
	var end bool = false
	for _, v := range str {
		switch v {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if (!started) {
				started = true
				sign = true
			}
			out = out * 10 + int(v - '0')
			size ++
			if (size > 10) {
				end = true
			}
			break

		case '-' :
			if (!started) {
				sign = false
				started = true
			} else {
				end = true
			}
			break
		case '+':
			if (!started) {
				sign = true
				started = true
			} else {
				end = true
			}
			break
		case ' ', '\n', '\t', '\r' :
			if (started) {
				end = true
			}
			break
		default:
			end = true
		}

		if (end) {
			break
		}

	}
	if (sign) {
		if (out > math.MaxInt32) {
			return math.MaxInt32
		}
		return out
	} else {
		if (-out < math.MinInt32) {
			return math.MinInt32
		}
		return -out
	}
}
