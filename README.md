# frappe-doctype-to-go
Generates GO structures from a given Frappe DocType file

## Installation
```bash
go get -v github.com/deidle/frappe-doctype-to-go
```

## Usage
Generates go functions from a given Postman JSON collection as standard input.
```bash
frappe-doctype-to-go <postman-collection-file.json >output.go -o=./output -p=my-package-name -l -c
```
* _o_ optional: the output folder for link.go and commond.go
* _p_ optional: the package name for the generated file
* _c_ optional: indicates whether to generate the Common file in the output folder
* _l_ optional: indicates whether to generate the Link structure in the output folder

## Docker
Run the following shell command to build the docker image
```bash
make docker
```

### Docker based development
```bash
docker run --rm -it --name=frappe-doctype-to-go -v "$(pwd)":/go/src/github.com/user/frappe-doctype-to-go -w /go/src/github.com/user/frappe-doctype-to-go golang sh -c '/bin/bash'
```
