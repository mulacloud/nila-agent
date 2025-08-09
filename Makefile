build:
	@go build -o dist/nila-agent-dev cmd/nila/main.go
run:
	@$(MAKE) build
	@./dist/nila-agent-dev -host 192.168.122.60 -port 8069
release:
	@GOOS=linux GOARCH=amd64 go build -o dist/nila-agent-linux cmd/nila/main.go
	@GOOS=freebsd GOARCH=amd64 go build -o dist/nila-agent-freebsd cmd/nila/main.go
	@GOOS=solaris GOARCH=amd64 go build -o dist/nila-agent-solaris cmd/nila/main.go
