package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestGenerate(t *testing.T) {
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

	if err = Generate(in, out, pkgName, true); err != nil {
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
