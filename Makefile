all:
	go run go_code_gen.go
	go fmt errcode/constant.go
	./build_site.sh