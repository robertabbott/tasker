.PHONY: install src

default: src

src:
	cd src && make

install:
	cp bin/tasker-daemon /bin/ &&\
	cp upstart/tasker-daemon.conf /etc/init/
