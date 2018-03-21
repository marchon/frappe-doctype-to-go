package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

const (
	commonOutputFileName   = "common.go"
	linkTypeOutputFileName = "link.go"
)

func main() {
	// TODO: Replace this with a CLI library
	// e.g. https://github.com/urfave/cli
	//			https://github.com/spf13/cobra
	defaultPkgName := "erpnext_api"
	pkgName := flag.String("p", defaultPkgName, "package name")
	output := flag.String("o", ".", "the outh path for common.go and link.go")
	includeCommon := flag.Bool("c", false, "include the Common file in the output")
	includeLink := flag.Bool("l", false, "include the Link structure in the output")
	flag.Parse()

	if *includeCommon {
		if out, err := os.Create(path.Join(*output, commonOutputFileName)); err == nil {
			if err := GenerateCommon(out, *pkgName); err != nil {
				fmt.Printf("Error generating 'common.go' file:\n%s", err)
			}
		}
	}
	if *includeLink {
		if out, err := os.Create(path.Join(*output, linkTypeOutputFileName)); err == nil {
			if err := GenerateLinkType(out, *pkgName); err != nil {
				fmt.Printf("Error generating 'link.go' file:\n%s", err)
			}
		}
	}
	if err := GenerateDocTypes(os.Stdin, os.Stdout, *pkgName); err != nil {
		fmt.Printf("Error generating doctype file:\n%s", err)
	}
}
