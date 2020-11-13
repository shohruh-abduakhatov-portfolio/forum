package internal

import (
	"net/http"
	"strconv"
	"strings"
)

func getCode(r *http.Request, prefix string) (code int64, validPrefix bool) {
	rest := strings.Replace(r.URL.Path, prefix, "", -1)
	p := strings.Split(rest, "/")
	if len(p) > 0 {
		resp, err := strconv.Atoi(p[0])
		if err == nil {
			code = int64(resp)
			validPrefix = true
		}
	}
	return code, validPrefix
}
