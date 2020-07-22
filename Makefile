build: 
	go build -o dist/zoox main.go

install: build
	mv dist/zoox /usr/local/bin