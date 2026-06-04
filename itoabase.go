package main

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}

	if value == 0 {
		return "0"
	}

	digits := "0123456789ABCDEF"
	result := ""

	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value /= base
	}

	return sign + result
}
