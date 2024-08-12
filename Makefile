build:
	@go build -o ./bin/imagick .

run: build
	@./bin/imagick

deps:
	sudo apt-get install libvips libvips-dev libvips-tools

test:
	@go test -cover ./...