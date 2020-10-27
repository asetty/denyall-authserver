FROM golang:1.15.3-alpine3.12 as build

RUN apk add --no-cache --update alpine-sdk

WORKDIR /go/src/denyall-authserver/

# Download dependencies
COPY . .
RUN go mod download

RUN make build

FROM alpine:3.12.0

RUN apk add --no-cache --update dumb-init
COPY --from=build /go/src/denyall-authserver/bin/denyall-authserver /usr/bin/denyall-authserver

ENTRYPOINT ["dumb-init", "--", "denyall-authserver"]
