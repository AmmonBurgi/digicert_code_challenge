FROM golang:1.16-alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /digicert_book_api

EXPOSE 8080

CMD [ "/digicert_book_api" ]