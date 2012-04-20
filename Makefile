SOURCES=$(shell find src -name *.hs)


dist/build/coffeephp/coffeephp: dist/setup-config $(SOURCES)
	cabal build

dist/setup-config:
	cabal configure

coffeephp: dist/build/coffeephp/coffeephp
	cp $< $@

clean:
	rm -rf dist coffeephp Setup
