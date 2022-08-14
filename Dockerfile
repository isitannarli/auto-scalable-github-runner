FROM golang:1.19-alpine

WORKDIR /workspace

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o app

CMD [ "/workspace/app" ]

