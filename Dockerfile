FROM golang:1.26.2

WORKDIR /source

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

ADD main.go main.go
COPY auth auth
COPY config config
COPY models models
COPY chatroom chatroom
COPY migrations migrations
COPY routes routes
COPY eventbus eventbus
COPY bot bot

RUN go build

ENV FILE_STORAGE_PATH=/data/uploads
VOLUME ["/data/uploads"]

CMD ["go", "run", "main.go"]
