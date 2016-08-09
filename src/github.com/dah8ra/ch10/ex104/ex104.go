package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
)

type deps struct {
	Deps []string `json:"Deps"`
}

type imports struct {
	Imports []string `json:"Imports"`
}

var b = flag.String("p", "ex104", "Specified package to show the dependency.")

/////////////////////////////////////////////////////
// Sample command
// go run ex104.go -p github.com/dah8ra/ch10/ex104
/////////////////////////////////////////////////////
func main() {
	flag.Parse()

	importsOut, importsErr := exec.Command("go", "list", "-f", "{{.Imports}}", "-json", *b).Output()
	if importsErr != nil {
		fmt.Println(importsErr)
		return
	}

	depsOut, depsErr := exec.Command("go", "list", "-f", "{{.Deps}}", "-json").Output()
	if depsErr != nil {
		fmt.Println(depsErr)
		return
	}

	importsReader := bytes.NewReader(importsOut)
	depsReader := bytes.NewReader(depsOut)
	importsDec := json.NewDecoder(importsReader)
	depsDec := json.NewDecoder(depsReader)
	var i imports
	var d deps
	importsDec.Decode(&i)
	depsDec.Decode(&d)
	fmt.Printf("[ All packages ]\n%+v\n\n", d)
	fmt.Printf("[ Specified package: %s ]\n%+v\n",*b, i)
}
