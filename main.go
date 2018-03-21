package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	commonOutputFileName   = "common.go"
	linkTypeOutputFileName = "link.go"
)

func main() {
	defaultPkgName := "erpnext_api"
	pkgName := flag.String("p", defaultPkgName, "package name")
	output := flag.String("o", ".", "the outh path for common.go and link.go")
	includeCommon := flag.Bool("c", false, "include the Common file in the output")
	includeLink := flag.Bool("l", false, "include the Link structure in the output")
	flag.Parse()

	fmt.Println(os.Stdout.Name())

	if *includeCommon {
		if out, err := os.Create(fmt.Sprintf("%s/%s", *output, commonOutputFileName)); err == nil {
			if err := GenerateCommon(out, *pkgName); err != nil {
				fmt.Printf("Error generating 'common.go' file:\n%s", err)
			}
		}
	}
	if *includeLink {
		if out, err := os.Create(fmt.Sprintf("%s/%s", *output, linkTypeOutputFileName)); err == nil {
			if err := GenerateLinkType(out, *pkgName); err != nil {
				fmt.Printf("Error generating 'link.go' file:\n%s", err)
			}
		}
	}
	if err := GenerateDocTypes(os.Stdin, os.Stdout, *pkgName); err != nil {
		fmt.Printf("Error generating doctype file:\n%s", err)
	}
}
