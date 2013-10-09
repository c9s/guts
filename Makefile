
all: yacc
	go build -x

yacc:
	for y in *.y ; \
	do \
		go tool yacc -o $${y%.y}.go -p Coffee $$y ; \
	done

install:
	cp -v coffeephp /usr/bin/cphp

