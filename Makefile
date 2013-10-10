all: parser
	source goenv && go build -x

parser:
	for y in *.y ; \
	do \
		go tool yacc -o $${y%.y}.go -p Coffee $$y ; \
	done

install:
	cp -v coffeephp /usr/bin/cphp

test:
	go test -v
