package doctype2go

import (
	"os"
)

func main() {
	pkgName := "erpnext_api"
	Generate(os.Stdin, os.Stdout, pkgName)
}
