build: 
	go build -o dist/zoom main.go

install: build
	mv dist/zoom /usr/local/bin