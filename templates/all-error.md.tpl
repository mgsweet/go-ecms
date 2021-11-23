---
title: "错误总览"
date: 2021-11-22T17:39:10+08:00
---
{{range $platform := .Platforms }}## {{$platform.Name}} ({{$platform.Code}})
{{range $module := $platform.Modules}}### {{$module.Name}} ({{$module.Code}})
| Name                       | Code     | Description                     |
|----------------------------|----------|---------------------------------|
{{range $specificError := $module.SpecificErrors}}| {{ $platform.Prefix }}{{ $module.Prefix }}{{ $specificError.Suffix }} | {{ $platform.Code }}{{ $module.Code }}{{ $specificError.Code }} | {{$specificError.Desc}} |
{{end}}{{end}}{{end}}