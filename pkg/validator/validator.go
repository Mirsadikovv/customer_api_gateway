package validator

import "regexp"

func ValidateGmail(gmail string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail.com$`).MatchString(gmail)
}
