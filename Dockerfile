FROM golang:latest

#RUN mkdir /app
ADD . /
#WORKDIR /app
RUN go build -o requests .
CMD ["/requests"]
