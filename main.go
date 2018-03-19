package main

import (
	"flag"
	"os"
)

func main() {
	defaultPkgName := "erpnext_api"
	pkgName := flag.String("package", defaultPkgName, "package name")
	withLink := flag.Bool("withlink", true, "include the Link structure in the output")

	Generate(os.Stdin, os.Stdout, *pkgName, *withLink)
}
