FROM golang:1.18.1-alpine as builder

ENV APP_PATH=github.com/Vladimir77715/otus/k8s_base

ARG TAG
ARG BUILD_TIME

WORKDIR ${GOPATH}/src/${APP_PATH}

COPY . .

RUN go mod download

RUN go build -ldflags "-X '${APP_PATH}/internal/version.Version=${TAG}' -X '${APP_PATH}/internal/version.BuildTime=${BUILD_TIME}'" \
    -o /bin/app ./cmd

EXPOSE 8000

ENTRYPOINT ["/bin/app"]

