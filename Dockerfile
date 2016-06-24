FROM golang:latest

ADD . /
RUN ["go","build","-o","/sorter_native","/sorter.go"].
CMD ["/sorter_native"]
