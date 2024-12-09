FROM golang:1.23-alpine

WORKDIR /code/w_gen

COPY ./app .

RUN go mod download

CMD [ "api" ]
