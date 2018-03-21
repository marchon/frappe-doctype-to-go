package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	tt "text/template"
)

const (
	commonTemplateName   = "common.tmpl"
	commonTemplatePath   = "./templates/common.tmpl"
	docTypesTemplateName = "doctypes.tmpl"
	docTypesTemplatePath = "./templates/doctypes.tmpl"
	linkTypeTemplateName = "link.tmpl"
	linkTypeTemplatePath = "./templates/link.tmpl"
)

type docTypeField struct {
	Type       string `json:"fieldtype"`
	MappedType string
	Required   int    `json:"reqd"`
	JSONName   string `json:"fieldname"`
}

type docTypeTemplate struct {
	PkgName    string
	TypeName   string         `json:"name"`
	TypeFields []docTypeField `json:"fields"`
}

type commonTemplate struct {
	PkgName string
}

type linkTypeTemplate struct {
	PkgName string
}

func (d *docTypeTemplate) Name() string {
	return strings.Replace(d.TypeName, " ", "", -1)
}

func (d *docTypeField) Name() (s string) {
	s = strings.Replace(d.JSONName, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// GenerateCommon generate the GO structure file for common types
func GenerateCommon(out io.Writer, pkgName string) (err error) {
	t := &commonTemplate{pkgName}
	err = applyTemplate(commonTemplateName, commonTemplatePath, t, out)
	return
}

// GenerateLinkType generate the GO structure file from Link doctype
func GenerateLinkType(out io.Writer, pkgName string) (err error) {
	t := &linkTypeTemplate{pkgName}
	err = applyTemplate(linkTypeTemplateName, linkTypeTemplatePath, t, out)
	return
}

// GenerateDocTypes generate the GO structure file from doctype
func GenerateDocTypes(in io.Reader, out io.Writer, pkgName string) (err error) {
	t, err := getDocTypeTemplate(in)
	if err != nil {
		return
	}

	t.PkgName = pkgName
	t.TypeFields = mapTypeFields(t.TypeFields)
	err = applyTemplate(docTypesTemplateName, docTypesTemplatePath, t, out)

	return
}

func getDocTypeTemplate(in io.Reader) (t *docTypeTemplate, err error) {
	raw, err := ioutil.ReadAll(in)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(raw), &t)
	return
}

func mapTypeFields(tf []docTypeField) (r []docTypeField) {
	hasLinks := false
	typeMapping := map[string]string{
		"Data":     "string",
		"Select":   "string",
		"Check":    "int",
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

func applyTemplate(name string, path string, template interface{}, out io.Writer) (err error) {
	t, err := tt.New(name).ParseFiles(path)
	if err != nil {
		return
	}

	err = t.Execute(out, template)
	return
}
