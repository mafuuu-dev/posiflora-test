FROM golang:1.26.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD sh -c "air -c .air.toml"
