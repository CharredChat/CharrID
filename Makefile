BASE=charrid
OS := $(shell uname)

build:
	GOOS=linux GOARCH=amd64 go build -o build/${BASE}-linux-x86_64 charrid-test.go
	GOOS=windows GOARCH=amd64 go build -o build/${BASE}-windows-x86_64.exe charrid-test.go
	GOOS=darwin GOARCH=amd64 go build -o build/${BASE}-macos-x86_64 charrid-test.go

run: build
ifeq ($(OS),Darwin)
	./build/${BASE}-macos-x86_64	 
else
	./build/${BASE}-linux-x86_64	 
endif

clean:
	go clean
	rm -rf build