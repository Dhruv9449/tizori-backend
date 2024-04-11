# BUILD IMAGE
FROM golang:1.20.4-alpine3.18 AS builder

WORKDIR /usr/src/app

RUN apk update \
    && apk --no-cache --update add build-base git

COPY ./go.mod ./go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o bin/tizori

# RUNNER IMAGE
FROM alpine:3.18 AS runner

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/bin/tizori ./bin/tizori

COPY --from=builder /usr/src/app/credentials ./credentials

RUN chmod +x ./bin/tizori

CMD ["./bin/tizori", "run"]