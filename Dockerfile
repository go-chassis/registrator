FROM golang:1.11 as builder
WORKDIR /go/src/github.com/go-mesh/registrator/
COPY .    /go/src/github.com/go-mesh/registrator/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o registrator .


FROM alpine:latest
WORKDIR /opt/registrator
ENV CHASSIS_CONF_DIR=/etc/registrator/conf \
    SCHEMA_ROOT=/etc/registrator/schema
COPY --from=builder /go/src/github.com/go-mesh/registrator/registrator .
COPY ./conf/*.yaml $CHASSIS_CONF_DIR/
RUN mkdir -p $SCHEMA_ROOT
ENTRYPOINT ["/opt/registrator/registrator", "-c", "/etc/registrator/conf/reg.yaml"]