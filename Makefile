## install dependencies
install:
	@echo "Installing dependencies"
	@go install go.k6.io/xk6/cmd/xk6@latest
	@go mod download

## build binary file
build:
	@echo "Building binary file"
	@xk6 build --with github.com/Juandavi1/xk6-prompt=. --output bin/

## run binary file example
run:
	@echo "Running binary file"
	@chmod +x ./bin/k6
	./bin/k6 run ./examples/sample.js