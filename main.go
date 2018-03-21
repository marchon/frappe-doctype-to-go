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
	includeCommon := flag.Bool("includeCommon", true, "include the Common file in the output")
	includeLink := flag.Bool("includeLink", true, "include the Link structure in the output")
	flag.Parse()

	outPath, _ := filepath.Abs(filepath.Dir(os.Stdout.Name()))

	if *includeCommon {
		if out, err := os.Create(fmt.Sprintf("%s/%s", outPath, commonOutputFileName)); err == nil {
			if err := GenerateCommon(out, *pkgName); err != nil {
				fmt.Printf("Error generating 'common.go' file:\n%s", err)
			}
		}
	}
	if *includeLink {
		if out, err := os.Create(fmt.Sprintf("%s/%s", outPath, linkTypeOutputFileName)); err == nil {
			if err := GenerateLinkType(out, *pkgName); err != nil {
				fmt.Printf("Error generating 'link.go' file:\n%s", err)
			}
		}
	}
	if err := GenerateDocTypes(os.Stdin, os.Stdout, *pkgName); err != nil {
		fmt.Printf("Error generating doctype file:\n%s", err)
	}
}
