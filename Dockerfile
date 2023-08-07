FROM golang:1.20 AS build

WORKDIR /go/src/github.com/codeinuit/fizzfuzz-api
COPY cmd cmd
COPY pkg pkg
COPY go.* ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api


FROM alpine:3.18

ARG USER=user
RUN apk --no-cache add curl

RUN adduser -D $USER
WORKDIR /run
COPY --from=build /go/src/github.com/codeinuit/fizzfuzz-api .

RUN chown $USER /run
RUN chmod 0755 /run/app

HEALTHCHECK --interval=1m --timeout=3s \
  CMD curl -X GET -f http://0.0.0.0:8080/health || exit 1

EXPOSE 8080
USER $USER
ENTRYPOINT ["./app"]
