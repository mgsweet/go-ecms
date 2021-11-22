package main

import "github.com/go-ecms/generator"

func main() {
	generator.GenerateGoCode("platforms/", "generator/constant.go.tpl", "")
}
