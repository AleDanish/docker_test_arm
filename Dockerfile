FROM golang:latest

ADD sorter.go /
ADD unsorted.txt /
WORKDIR /
RUN ["go","build","-o","/sorter_native","/sorter.go"].
CMD ["/sorter_native"]
