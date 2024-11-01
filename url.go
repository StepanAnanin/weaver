package weaver

import (
	"errors"
	"net/url"
	"strings"
)

func GetQueryParams(url *url.URL, paramsKeys ...string) ([]string, error) {
	res := []string{}
	misses := []string{}

	for _, key := range paramsKeys {
		param := url.Query().Get(key)

		if param != "" {
			res = append(res, param)
		} else {
			misses = append(misses, param)
		}
	}

	if len(misses) > 0 {
		return res, errors.New("missing query params: " + strings.Join(misses, ", "))
	}

	return res, nil
}
