all: clean build install

clean:
	-rm -rf ./bin/
	rm -f bitacora.log
	rm -f bitacora.db

build:
	mkdir ./bin/
	go build -o ./bin/bitacora
	chmod +x ./bin/bitacora

install:
	cp ./bin/bitacora /usr/local/bin/bitacora
