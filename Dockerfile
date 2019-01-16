FROM golang:1.11 as builder
WORKDIR /go/src/github.com/go-mesh/registrator/
COPY .    /go/src/github.com/go-mesh/registrator/
RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o registrator .


FROM alpine:latest
WORKDIR /opt/registrator
COPY --from=builder /go/src/github.com/go-mesh/registrator/registrator .
COPY ./template /opt/registrator
ENTRYPOINT ["/opt/registrator/registrator", "-c", "/etc/registrator/reg.yaml"]