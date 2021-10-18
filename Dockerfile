FROM alpine:3.14.0

RUN apk add --no-cache git make musl-dev go protobuf protobuf-dev
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

CMD [ "protoc -v" ]