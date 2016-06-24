FROM golang:latest

ADD . /
WORKDIR /
RUN ["go","build","-o","/sorter_native","/sorter.go"].
CMD ["/sorter_native"]
