.PHONY: default run build tests docs clean

APP_NAME=gooportunities

default: run 

run:
	@go run main.go 

build:
	@go build -o $(APP_NAME) main.go 

test: 
	@go test ./ ...

docs: 
	@swag init 

clean:
	@rm -f $(APP_NAME)
