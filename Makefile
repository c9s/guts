

parser:
	for y in *.y ; \
	do \
		go tool yacc -o $${y%.y}.go -p Coffee $$y ; \
	done

all: parser
	go build -x

install:
	cp -v coffeephp /usr/bin/cphp

