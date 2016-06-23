FROM golang:latest

ADD . /
RUN ["go","build","-o","/requests_native","/requests.go"].
CMD ["/requests_native"]
