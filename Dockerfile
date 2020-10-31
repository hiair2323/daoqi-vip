FROM daocloud.io/arron2/golang as builder
WORKDIR $GOPATH/src/github.com/EDDYCJY/email
COPY . $GOPATH/src/github.com/EDDYCJY/email
RUN go build -o /go/email

FROM alpine:latest
COPY --from=builder /go/email /go/email
EXPOSE 8000
ENTRYPOINT ["./go/email"]
