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

executable:
	go build bin/coffeephp/main.go

test:
	go test coffeephp -v

clean:
	rm -rfv pkg/darwin_amd64/
