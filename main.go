package main // import "go-depmgmt-testrepo"

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println(errors.New("hello world!"))
}
