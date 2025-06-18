package email

import "regexp"

func IsValidFormat(email string) bool {
	re := `^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`
	return regexp.MustCompile(re).MatchString(email)
}
