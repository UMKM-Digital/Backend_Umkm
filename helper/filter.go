package helper

import (
	"net/url"
	"strconv"
)

func ExtractFilterSort(params url.Values) (string, string, int, int) {
	// Ambil filter tunggal dari URL
	filter := params.Get("filter")
	status := params.Get("status")
	// status, _:= strconv.Atoi(params.Get("status"))
	limit, _ := strconv.Atoi(params.Get("limit"))
	page, _ := strconv.Atoi(params.Get("page"))

	return filter, status,limit, page
}
