all: yacc parser build test

yacc:
	go install -x yacc

parser:
	for y in src/gutscript/*.y ; \
	do \
		bin/yacc -o $${y%.y}.go -p Guts $$y ; \
	done

install:
	cp -v bin/gutscript /usr/bin/cphp

build:
	go build gutscript

executable:
	go build bin/gutscript/main.go

test:
	go test gutscript -v

testphp:
	for file in src/gutscript/tests/*.guts ; \
	do \
		echo "Testing " $$file; \
		go run src/gutscript/guts/main.go $$file | php ; \
	done

clean:
	rm -rfv pkg/darwin_amd64/
