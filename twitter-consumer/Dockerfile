FROM golang:1.14.3-alpine3.11
COPY ./ /build
WORKDIR /build
RUN go mod verify && go build -o twitter-consumer main.go

FROM alpine:3.11
RUN mkdir /app
WORKDIR /app
COPY --from=0 /build/twitter-consumer .
CMD [ "./twitter-consumer" ]