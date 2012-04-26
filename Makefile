SOURCES=$(shell find src -name *.hs)

coffeephp: dist/build/coffeephp/coffeephp
	cp -v $< $@

dist/build/coffeephp/coffeephp: dist/setup-config $(SOURCES)
	cabal build

dist/setup-config:
	cabal configure


clean:
	rm -rf dist coffeephp Setup
