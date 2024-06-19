build:
	go build -o jp main.go

test:
	go build -o jp main.go && ./test.sh

run:
	go run main.go demo.json
