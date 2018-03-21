package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	commonOutputFileName   = "common.go"
	linkTypeOutputFileName = "link.go"
)

func main() {
	defaultPkgName := "erpnext_api"
	pkgName := flag.String("package", defaultPkgName, "package name")
	includeCommon := flag.Bool("includeCommon", false, "include the Common file in the output")
	includeLink := flag.Bool("includeLink", false, "include the Link structure in the output")
	flag.Parse()

	outPath, _ := filepath.Abs(filepath.Dir(os.Stdout.Name()))

	if *includeCommon {
		if out, err := os.Create(outPath + commonOutputFileName); err == nil {
			if err := GenerateCommon(out, *pkgName); err != nil {
				fmt.Printf("Error generating 'common.go' file:\n%s", err)
			}
		}
	}
	if *includeLink {
		if out, err := os.Create(outPath + linkTypeOutputFileName); err == nil {
			if err := GenerateLinkType(out, *pkgName); err != nil {
				fmt.Printf("Error generating 'link.go' file:\n%s", err)
			}
		}
	}
	if err := GenerateDocTypes(os.Stdin, os.Stdout, *pkgName); err != nil {
		fmt.Printf("Error generating doctype file:\n%s", err)
	}
}
