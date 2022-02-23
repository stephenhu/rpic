FROM golang AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go get && go build

FROM migrate/migrate-sqlite3 AS data
WORKDIR /db
COPY db/migrations/ .
RUN migrate -source file://db/migrations -database sqlite3://rpic.db up 2

FROM ubuntu
WORKDIR /usr/local/rpic
RUN apt-get update -y
COPY --from=builder /app/rpic .
COPY --from=data rpic.db .
EXPOSE 9008
CMD ["/usr/local/rpic/rpic"]