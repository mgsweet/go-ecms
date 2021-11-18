package errcode

// Error code should start from 10 000 000 to ensure equal length.
// Format: AA BBB CCC
// AA: Platform (10 is for system error. Business related errors should start at 20)
// BBB: Module
// CCC: Specific error
const (
	SUCCESS = 0

	ReqParamIllegal = 2
	Unauthorized    = 3
	Forbidden       = 4
)

// error code default description
var codeDefaultDesc = map[int32]string{
	SUCCESS: "success",

	ReqParamIllegal: "Illegal request parameter",
	Unauthorized:    "Unauthorized",
	Forbidden:       "Forbidden",
}

// CodeDefaultDesc returns the default description for the given error code.
func CodeDefaultDesc(code int32) string {
	if v, ok := codeDefaultDesc[code]; ok {
		return v
	} else {
		return ""
	}
}
