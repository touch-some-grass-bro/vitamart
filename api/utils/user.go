package utils

import (
	"strconv"
	"strings"
)

func GetUserJoinYear(email string) int32 {
	at_index := strings.Index(email, "@")
	year := email[at_index-4 : at_index]
	year_int, err := strconv.Atoi(year)

	if err != nil {
		return 0
	}

	return int32(year_int)
}
