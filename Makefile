SOURCES=$(shell find src -name *.hs)

dist/setup-config:
	cabal configure

dist/build/coffeephp/coffeephp: dist/setup-config $(SOURCES)
	cabal build

coffeephp: dist/build/coffeephp/coffeephp
	cp $< $@

clean:
	rm -rf dist coffeephp
