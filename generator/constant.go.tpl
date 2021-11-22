package errcode

// Error code should start from 10 000 000 to ensure equal length.
// Format: AA BBB CCC
// AA: Platform (10 is for system error. Business related errors should start at 20)
// BBB: Module
// CCC: Specific error
const (
    SUCCESS = 0

    {{range $platform := .}}// Platform: "{{ $platform.Name }}"
    {{range $module := $platform.Modules}}// Module: "{{ $module.Name }}", Code format: {{ $platform.Code }} {{ $module.Code }} xxx
    {{range $specificError := $module.SpecificErrors}}const {{ $platform.Prefix }}{{ $module.Prefix }}{{ $specificError.Name }} = {{ $platform.Code }}{{ $module.Code }}{{ $specificError.Code }} // {{$specificError.Description}}
    {{end}}{{end}}{{end}}
)

// error code default description
var codeDefaultDesc = map[int32]string{
    SUCCESS: "success",

    {{range $platform := .}}// Platform: "{{ $platform.Name }}"
    {{range $module := $platform.Modules}}// Module: "{{ $module.Name }}"
    {{range $specificError := $module.SpecificErrors}}const {{ $platform.Prefix }}{{ $module.Prefix }}{{ $specificError.Name }} = "{{ $specificError.Description }}"
    {{end}}{{end}}{{end}}
}

// CodeDefaultDesc returns the default description for the given error code.
func CodeDefaultDesc(code int32) string {
    if desc, ok := codeDefaultDesc[code]; ok {
        return desc
    } else {
        return ""
    }
}
