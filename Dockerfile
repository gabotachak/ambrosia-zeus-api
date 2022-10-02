FROM golang:1.19-alpine

WORKDIR /go/ambrosia-zeus-api
COPY . /go/ambrosia-zeus-api

RUN go mod download
RUN go build -o /ambrosia-zeus-api

EXPOSE 8080

CMD [ "/ambrosia-zeus-api" ]
