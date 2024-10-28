FROM golang:1.23.2

WORKDIR /usr/src/llink

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/llink

CMD [ "llink" ]

