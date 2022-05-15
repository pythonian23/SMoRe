package main

import (
	"fmt"
	"io/ioutil"
	"os"

	smore "github.com/pythonian23/SMoRe"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(smore.Render(string(b)))
}
