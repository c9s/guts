all: parser build test

parser:
	for y in src/coffeephp/*.y ; \
	do \
		go tool yacc -o $${y%.y}.go -p Coffee $$y ; \
	done

install:
	cp -v bin/coffeephp /usr/bin/cphp

build:
	go build coffeephp

test:
	go test coffeephp -v

clean:
	rm -rfv bin/coffeephp pkg/darwin_amd64/
