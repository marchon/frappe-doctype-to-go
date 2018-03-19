frappe-doctype-to-go: generator.go
	go build

test: generator.go
	go test

docker:
	docker build -t frappe-doctype-to-go .