package utils

import "github.com/microcosm-cc/bluemonday"

func AvoidXSS(s string) string {
	return bluemonday.UGCPolicy().Sanitize(s)
}
