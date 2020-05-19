FROM golang:1.14.3-alpine3.11
COPY ./ /build
WORKDIR /build
RUN go mod verify && go build -o twitter-migrate main.go

FROM alpine:3.11
RUN mkdir -p /app/scripts
COPY ./scripts/ /app/scripts
WORKDIR /app
COPY --from=0 /build/twitter-migrate .
CMD [ "./twitter-migrate" ]