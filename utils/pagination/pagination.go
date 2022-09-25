package pagination

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Pagination : Extract params of pagination in http request
func Pagination(r *http.Request) (int, int) {
	var err error
	var limit int
	var offset int

	params := mux.Vars(r)

	limit, err = strconv.Atoi(params["limit"])
	if err != nil {
		limit = 100
	}

	offset, err = strconv.Atoi(params["offset"])
	if err != nil {
		offset = 0
	}

	return limit, offset
}
