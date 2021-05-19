FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/franciscoruizar/quasar-fire
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/quasar-fire/cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/quasar-fire-api /go/bin/quasar-fire-api
ENTRYPOINT ["/go/bin/quasar-fire-api"]
