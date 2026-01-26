// presentation/http/server/constant.go
package server

import "regexp"

var PARAM_REGEX = regexp.MustCompile(`:([a-zA-Z0-9_]+)`)
