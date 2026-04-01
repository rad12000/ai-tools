.PHONY: executable

executable:
	go build -ldflags="-s -w" -trimpath -o schema-template main.go

build: executable
	cp schema-template ~/.ai-tools/bin
