FROM golang

WORKDIR /source

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

ADD main.go main.go
COPY auth auth
COPY config config
COPY models models

RUN go build

CMD ["go", "run", "main.go"]