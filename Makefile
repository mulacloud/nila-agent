build:
	@go build -o dist/nila-agent-dev
release:
	@GOOS=linux GOARCH=amd64 go build -o dist/nila-agent-linux
	@GOOS=freebsd GOARCH=amd64 go build -o dist/nila-agent-freebsd
	@GOOS=solaris GOARCH=amd64 go build -o dist/nila-agent-solaris
