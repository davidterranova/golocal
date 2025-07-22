FROM golang:alpine as builder
RUN apk update && apk add --no-cache git make

RUN mkdir /golocal
WORKDIR /golocal
COPY . .
RUN go get -d -v
RUN make build GOOS=linux GOARCH=arm64

FROM scratch
COPY --from=builder /golocal/target/golocal /go/bin/golocal
ENTRYPOINT [ "/go/bin/golocal" ]