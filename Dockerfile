FROM golang:1.11 as builder
WORKDIR /go/src/github.com/go-mesh/registrator/
COPY .    /go/src/github.com/go-mesh/registrator/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o registrator .


FROM alpine:latest
WORKDIR /opt/registrator
COPY --from=builder /go/src/github.com/go-mesh/registrator/registrator .
ENV CHASSIS_CONF_DIR /etc/registrator/
ENTRYPOINT ["/opt/registrator/registrator", "-c", "/etc/registrator/reg.yaml"]