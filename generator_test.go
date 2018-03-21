package main

import (
	"github.com/stretchr/testify/assert"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestMain(m *testing.M) {
	gofakeit.Seed(0)
	os.Exit(m.Run())
}

func TestGenerateCommon(t *testing.T) {
	out := bytes.NewBufferString("")
	pkgName := strings.ToLower(gofakeit.FirstName())

	if err := GenerateCommon(out, pkgName); err != nil {
		t.Error(err)
	}
	
	generated := out.String()
	assert.NotEmpty(t, generated)
}

func TestGenerateLinkType(t *testing.T) {
	out := bytes.NewBufferString("")
	pkgName := strings.ToLower(gofakeit.FirstName())

	if err := GenerateLinkType(out, pkgName); err != nil {
		t.Error(err)
	}
		
	generated := out.String()
	assert.NotEmpty(t, generated)
}

func TestGenerateDocTypes(t *testing.T) {
	expected, err := ioutil.ReadFile("./samples/customer.go")
	if err != nil {
		t.Error(err)
	}
	pkgName := strings.ToLower(gofakeit.FirstName())
	in, err := os.Open("./samples/customer.json")
	if err != nil {
		t.Error(err)
	}
	out := bytes.NewBufferString("")

	if err = GenerateDocTypes(in, out, pkgName); err != nil {
		t.Error(err)
	}

	generated := out.String()
	if len(generated) == 0 {
		t.Errorf("No outout has been generated")
	}

	if generated != string(expected) {
		t.Errorf("Generated is different from expected")
	}
}
