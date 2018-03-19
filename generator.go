package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	tt "text/template"
)

const (
	templateName = "generator.tmpl"
	templatePath = "./generator.tmpl"
)

type docTypeField struct {
	Type       string `json:"fieldtype"`
	MappedType string
	Required   int    `json:"reqd"`
	JSONName   string `json:"fieldname"`
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
func Generate(in io.Reader, out io.Writer, pkgName string, withLinkStruct bool) (err error) {
	t, err := getTemplate(in)
	if err != nil {
		return
	}

	t.PkgName = pkgName
	t.WithLinkStruct = withLinkStruct
	t.TypeFields = mapTypeFields(t.TypeFields)

	err = applyTemplate(templateName, templatePath, t, out)

	return
}

func getTemplate(in io.Reader) (t *template, err error) {
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

func mapTypeFields(tf []docTypeField) (r []docTypeField) {
	hasLinks := false
	typeMapping := map[string]string{
		"Data":     "string",
		"Select":   "string",
		"Check":    "bool",
		"Currency": "float32",
		"Text":     "string",
		"Float":    "float32",
	}

	for _, field := range tf {
		field.MappedType = typeMapping[field.Type]

		if !hasLinks && field.Type == "Link" {
			hasLinks = true
		}

		if len(field.MappedType) > 0 {
			r = append(r, field)
		}
	}

	if hasLinks {
		link := docTypeField{
			JSONName:   "links",
			MappedType: "[]Link",
		}
		r = append(r, link)
	}
	return
}

func applyTemplate(templateName string, templatePath string, data *template, out io.Writer) (err error) {
	t, err := tt.New(templateName).ParseFiles(templatePath)
	if err != nil {
		return
	}

	err = t.Execute(out, data)
	return
}
