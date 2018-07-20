package main

import (
	"fmt"
	"net/http"

	"github.com/devilmustcry/gosuki/greeting"
	"github.com/devilmustcry/gosuki/subject"
)

func main() {
	fmt.Print(subject.GetSubject())
	http.ListenAndServe(":5000", http.HandlerFunc(greeting.MyGreeterHandler))
}
