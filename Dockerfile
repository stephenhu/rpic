FROM golang AS builder
WORKDIR /go/rpic
COPY . .
RUN GOOS=linux GOARCH=amd64 go get && go build

FROM ubuntu
WORKDIR /usr/local/rpic
RUN apt-get update -y
COPY --from=builder /go/rpic/rpic .
EXPOSE 12000
CMD ["/usr/local/rpic/rpic"]