all: parser
	go build coffeephp

parser:
	for y in src/coffeephp/*.y ; \
	do \
		go tool yacc -o $${y%.y}.go -p Coffee $$y ; \
	done

install:
	cp -v bin/coffeephp /usr/bin/cphp

test:
	go test coffeephp -v

clean:
	rm -rfv bin/coffeephp pkg/darwin_amd64/
