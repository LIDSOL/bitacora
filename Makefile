all: clean build run

clean:
	rm -rf ./bin/*

build:
	go build -o bin/bitacora
	chmod +x bin/bitacora

run:
	./bin/bitacora