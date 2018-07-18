package main

import (
	"fmt"

	"github.com/devilmustcry/gosuki/request"
	"github.com/devilmustcry/gosuki/subject"
)

func main() {
	fmt.Println(subject.GetSubject())
	request.Get()
}
