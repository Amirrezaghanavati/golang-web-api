package common

import "regexp"

var iranianMobileNumberPattern = regexp.MustCompile(`^(0|\+98|0098)9[0-9]{9}$`)

func IsIranianMobileNumber(number string) bool {
	return iranianMobileNumberPattern.MatchString(number)
}
