FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o requests .
CMD ["/app/requests"]
