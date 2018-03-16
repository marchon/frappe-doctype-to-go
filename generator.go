package doctype2go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/jmoiron/jsonq"
)

type template struct {
	PkgName        string
	WithLinkStruct bool
	TypeName       string
	TypeFields     []struct {
		Name      string
		Type      string
		Decorator string
	}
}

// Generate generate the GO structure file from doctype
func Generate(in *os.File, out *os.File, pkgName string) (err error) {
	data := map[string]interface{}{}
	dec := json.NewDecoder(in)
	if err = dec.Decode(&data); err != nil {
		return
	}

	jq := jsonq.NewQuery(data)
	fields, err := jq.ArrayOfObjects("fields")

	fields["fieldname"]
	fields["fieldtype"]
	fields["reqd"]

	var contact Contact
	if err = json.Unmarshal([]byte(body), &contact); err != nil {
		contacts = append(contacts, contact)
	}

	return
}
