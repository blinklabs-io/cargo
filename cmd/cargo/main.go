package main

import (
	"fmt"
	"github.com/cloudstruct/cargo/version"
)

func main() {
	fmt.Printf("cargo %s\n", version.GetVersionString())
}
