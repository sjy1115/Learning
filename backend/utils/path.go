package utils

import (
	"os"
)

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetStartPage(page, pageSize int) int {
	if page >= 1 {
		page = page - 1
	}

	return page * pageSize
}
