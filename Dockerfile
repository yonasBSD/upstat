FROM golang:1.21.8-alpine AS build
WORKDIR /app
COPY . .
RUN apk -U upgrade \
    && apk add --no-cache yarn
WORKDIR web
RUN yarn install && yarn build
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN chmod +x ./startup.sh

FROM alpine:latest as release
WORKDIR /app
COPY --from=build /app/main .
RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && chmod +x /app/main
EXPOSE 8000
CMD ["./main"]
