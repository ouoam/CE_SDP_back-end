FROM golang:latest

# RUN apk update && apk add gcc git

WORKDIR /app
COPY dep.go .

RUN go get -d -v ./...

COPY . .

RUN go get -d -v ./...
RUN go build ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/app .

EXPOSE 3000

CMD ["./app"]