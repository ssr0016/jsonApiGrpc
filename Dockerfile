FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /jsonApiGrpc

EXPOSE 3000

CMD [ "/jsonApiGrpc" ]