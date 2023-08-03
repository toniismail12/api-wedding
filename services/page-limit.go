package services

import "strconv"

func PageLimit(Page, Limit string) (int, int) {

	page, _ := strconv.Atoi(Page)
	limit, _ := strconv.Atoi(Limit)

	return page, limit
}
