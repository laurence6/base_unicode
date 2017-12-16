build:
	go build -o main main.go decoder.go encoder.go common.go table.go

test:
	go test
