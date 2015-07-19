.PHONY: all

default: all

all:
	cd cli &&\
	go build -o cli &&\
	mv cli ../../bin &&\
	cd ../daemon &&\
	go build -o daemon &&\
	mv daemon ../../bin
