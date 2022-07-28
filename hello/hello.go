package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
	"rsc.io/quote"
)

func main() {
	//fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(quote.Go())
	fmt.Println(stringutil.ToUpper("Hello"))
}
