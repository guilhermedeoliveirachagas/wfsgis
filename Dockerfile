FROM golang:1.9.2-alpine3.7

ADD ./ /go/src/github.com/boundlessgeo/wfs3

# Build the server command inside the container.
RUN apk add --no-cache git build-base bash && \
    go get -u github.com/golang/dep/cmd/dep && \
    cd /go/src/github.com/boundlessgeo/wfs3; dep ensure; \
		go install github.com/boundlessgeo/wfs3

ENTRYPOINT /go/bin/wfs3

EXPOSE 8080
