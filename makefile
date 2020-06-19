PROJECT="porter"
BINARY="porter"
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	@go build -o ${BINARY} -tags=jsoniter

fmt:
	@gofmt -s -w ${GOFILES}

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY}; fi

.PHONY: default fmt clean