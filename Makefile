VERSION=1.0.0

name = "sui-hackathon"

globe = "command/main.go"

${name}:
	@go build \
	-o ${name} ${globe}

${name}-amd64-linux:
	@GOOS=linux GOARCH=amd64go build \
	-o ${name} ${globe}

${name}-x86-64-linux:
	@GOOS=linux GOARCH=386 go build \
	-o ${name} ${globe}

default:
	make ${name}

.PHONY: ${name}