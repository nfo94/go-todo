FROM golang:1.22 as BUILDER

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/ ./cmd/
COPY main/ ./main/

RUN go build -o /go-todo ./cmd

EXPOSE 8080

CMD ["/go-todo"]
