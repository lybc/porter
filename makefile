PROJECT="porter"
BINARY="porter"
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default: help
## compile: Compile the binary.
compile:
	@go build -o ${BINARY} -tags=jsoniter
## fmt: Run gofmt for whole project
fmt:
	@gofmt -s -w ${GOFILES}
## clean: Delete old binary
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY}; fi
## help: Display make file command
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: compile fmt clean