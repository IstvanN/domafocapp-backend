
FROM golang:alpine AS builder
WORKDIR /go/src/domafocapp-backend
COPY . .
RUN go get -d -v ./... &&\
    go install -v ./...

FROM alpine:latest
COPY --from=builder /go/bin/domafocapp-backend /domafocapp-backend
LABEL Name=domafocapp-backend Version=0.0.1
EXPOSE 8080
CMD ["/domafocapp-backend"]
