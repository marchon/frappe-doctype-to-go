frappe-doctype-to-go: main.go
	go build

test: main.go
	go test

docker:
	docker build -t frappe-doctype-to-go .