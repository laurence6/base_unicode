build:
	go build -o base_unicode main.go decoder.go encoder.go common.go table.go

test:
	go test
