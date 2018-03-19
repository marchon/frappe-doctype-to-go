package doctype2go

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type docTypeField struct {
	Type       string `json:"fieldtype"`
	MappedType string
	Required   bool   `json:"reqd"`
	JSONName   string `json:"fieldname"`
	Generate   bool
}

type template struct {
	PkgName        string
	WithLinkStruct bool
	TypeName       string         `json:"name"`
	TypeFields     []docTypeField `json:"fields"`
}

func (d *docTypeField) Name() (s string) {
	s = strings.Replace(d.JSONName, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// Generate generate the GO structure file from doctype
func Generate(in *os.File, out *os.File, pkgName string, withLinkStruct bool) (err error) {
	t := getTemplate(in)

	t.PkgName = pkgName
	t.WithLinkStruct = withLinkStruct
	hasLinks := false
	typeMapping := map[string]string{
		"Data":     "string",
		"Select":   "string",
		"Check":    "bool",
		"Currency": "float32",
		"Text":     "string",
		"Float":    "float32",
	}

	for _, field := range t.TypeFields {
		field.MappedType = typeMapping[field.Type]
		if len(field.MappedType) > 0 {
			field.Generate = true
		}

		if !hasLinks && field.Type == "Link" {
			hasLinks = true
		}
	}

	if hasLinks {
		link := docTypeField{
			JSONName:   "links",
			MappedType: "[]Link",
			Generate:   true,
		}
		t.TypeFields = append(t.TypeFields, link)
	}

	return
}

func getTemplate(in *os.File) (t* template) {
	// data := map[string]interface{}{}
	// dec := json.NewDecoder(in)
	// if err = dec.Decode(&data); err != nil {
	// 	return
	// }
	raw, err := ioutil.ReadAll(in)
	if err != nil {
		return
	}

	// jq := jsonq.NewQuery(data)
	// f, err := jq.String("fields")
	// var fields []docTypeField
	err = json.Unmarshal([]byte(raw), &t)
	return
}