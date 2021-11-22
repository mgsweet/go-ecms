---
title: "错误总览"
date: {{.TimeStr}}
---
{{range $platform := .Platforms }}## {{$platform.Name}} ({{$platform.Code}})
{{range $module := $platform.Modules}}### {{$module.Name}} ({{$module.Code}})
| Name                       | Code     | Description                     |
|----------------------------|----------|---------------------------------|
{{range $specificError := $module.SpecificErrors}}| {{ $platform.Prefix }}{{ $module.Prefix }}{{ $specificError.Name }} | {{ $platform.Code }}{{ $module.Code }}{{ $specificError.Code }} | {{$specificError.Desc}} |
{{end}}{{end}}{{end}}