FROM golang:1.12.3 AS BUILD

RUN mkdir /wfs3
WORKDIR /wfs3

ADD go.mod .
ADD go.sum .
RUN go mod download

#now build source code
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/wfs3 .

FROM golang:1.12.3

ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_DBNAME=postgres
ENV POSTGRES_USERNAME=postgres
ENV POSTGRES_PASSWORD=postgres

COPY --from=BUILD /go/bin/* /bin/
ENTRYPOINT /bin/wfs3

EXPOSE 8080