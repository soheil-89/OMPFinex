FROM golang:1.20

WORKDIR /usr/src/go
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/apis ./
CMD ["apis"]

