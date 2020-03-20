.DEFAULT_GOAL := build-all

build-all: clean build

build:
	@ echo "Building..."
	# @ go build -o caracal-pty cmd/main.go
	@ python bin/build.py
	@ echo "bin file:"
	@ ls caracal-pty

pack:
	@ echo "pack..."

clean:
	@ echo "Clean..."
	@ rm -f caracal-pty

fmt:
	@ echo "gofmt format..."
	go fmt ./

test:
	# TODO 
	@ echo "test..."