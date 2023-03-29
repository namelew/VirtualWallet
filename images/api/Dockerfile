FROM golang:1.20.2-alpine3.17 as base
RUN apk update
RUN apk add git
RUN git clone https://github.com/namelew/VirtualWallet /app
WORKDIR /app
RUN go mod tidy; go build -o ./api ./cmd/api/main.go

FROM alpine:3.17 as binary

WORKDIR /app
COPY --from=base /app/api .
COPY --from=base /app/.env .
COPY --from=base /app/migrations ./migrations

ENV PORT=8080
ENV DBHOST=localhost
ENV DBPORT=5432
ENV DBUSER=postgres
ENV DBPASS=postgres
ENV DBNAME=postgres
ENV DBSSL=disable

EXPOSE ${PORT}

ENTRYPOINT [ "./api" ]