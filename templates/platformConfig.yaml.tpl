platforms:
  {{range $platform := .}}- name: "{{$platform.Name}}"
    code: "{{$platform.Code}}"
    prefix: "{{$platform.Prefix}}"
    dir: "{{$platform.Dir}}"
  {{end}}