#!/bin/bash

# Usage:
# ./golnet <IP> <PORT> <PASSWORD> '<COMMAND>'
# Example:
# ./golnet 127.0.0.1 8081 you_password 'say "[FF0000]Checking the server label[-] without panic."'

go clean -cache

# Linux build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o golnet -ldflags="-w -s" golnet.go

# Windows build
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o golnet.exe -ldflags="-w -s" golnet.go

