FROM golang:alpine3.11

RUN apk update && apk add bash

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o patreon-statistics ./app

CMD ["./wait-for-it.sh", "db:5432", "--", "./patreon-statistics"]
