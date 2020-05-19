package utils

import (
	"errors"
	"strings"
)

// TransformQueryGetAll ...
func TransformQueryGetAll(queryStr string) (query map[string]string, err error) {
	if queryStr != "" {
		query = make(map[string]string) // queryStr: k:v,k:v
		for _, cond := range strings.Split(queryStr, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				err = errors.New("Error: invalid query key/value pair")
				return
			}
			// kv[0], rewrite dot-notation to Object__Attribute
			k, v := strings.Replace(kv[0], ".", "__", -1), kv[1]
			query[k] = v
		}
	}
	return
}
