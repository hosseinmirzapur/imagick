build:
	go build -o ./bin/imagick .

run: build
	./bin/imagick

deps:
	sudo apt-get install libvips