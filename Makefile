SOURCES=$(shell find src -name *.hs)

coffeephp: dist/build/coffeephp/coffeephp
	cp -v $< $@

dist/build/coffeephp/coffeephp: dist/setup-config $(SOURCES)
	cabal build

dist/setup-config:
	cabal configure


phpwiki.phar:
	curl -O https://raw.github.com/c9s/PHPWikiGen/master/phpwiki.phar

doc: phpwiki.phar
	php phpwiki.phar doc/synopsis/ html/synopsis/

clean:
	rm -rf dist coffeephp Setup

.PHONY: doc clean
