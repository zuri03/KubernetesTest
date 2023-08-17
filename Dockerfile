FROM golang:1.20.6 as builder

WORKDIR /go/src

COPY . .

RUN CGO_ENABLED=0 go build -o build/api ./main.go

FROM alpine:3.18.3
RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/zuri03/KubernetesTest/

COPY --from=builder /go/src/build/api ./main
EXPOSE 9000
ENTRYPOINT [ "./main" ]
