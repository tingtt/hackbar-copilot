FROM golang:1.24.0-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 80


FROM golang:1.24.0-alpine as builder

ENV ROOT=/go/src/app
ARG GO_ENTRYPOINT=main.go
WORKDIR ${ROOT}

RUN apk update
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/hackbar-copilot ${GO_ENTRYPOINT}

FROM alpine:latest as certs
RUN apk update && apk add ca-certificates

FROM busybox as prod

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=certs /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /bin/hackbar-copilot /bin/hackbar-copilot

EXPOSE 80
ENTRYPOINT ["/bin/hackbar-copilot", "--host", "0.0.0.0", "--port", "80"]
