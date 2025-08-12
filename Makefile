BINARY_NAME=main

.PHONY: run build test clean

build:
	docker build -t meu-app-go .

run:
	docker run -p 3000:3000 --rm meu-app-go

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)