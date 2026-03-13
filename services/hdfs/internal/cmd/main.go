package main

import (
	def "github.com/rgglez/go-storage/v5/definitions"
)

func main() {
	def.GenerateService(Metadata, "generated.go")
}
