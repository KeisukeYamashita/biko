# Build Go Server Binary
FROM golang:1.16
LABEL MAINTAINER=KeisukeYamashita

ENV GO111MODULE on

ARG SERVICE_NAME
ARG VERSION

WORKDIR /project

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go install -v \
            -ldflags="-w -s -X main.version=${VERSION} -X main.serviceName=${SERVICE_NAME}" \
            ./cmd/biko

FROM alpine:latest
COPY --from=0 /go/bin/biko /bin/biko
ENTRYPOINT ["/bin/biko"]