package utils

import "time"

func IsExistsInList(list []string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func ParsingDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
