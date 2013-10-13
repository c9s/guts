all: yacc parser build test

yacc:
	go install -x yacc

parser:
	for y in src/gutscript/*.y ; \
	do \
		bin/yacc -o $${y%.y}.go -p Coffee $$y ; \
	done

install:
	cp -v bin/gutscript /usr/bin/cphp

build:
	go build gutscript

executable:
	go build bin/gutscript/main.go

test:
	go test gutscript -v

clean:
	rm -rfv pkg/darwin_amd64/
