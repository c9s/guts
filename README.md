CoffeePHP
==================

CoffeePHP is a language that let you compile less code into PHP and makes your life easier.

CoffeePHP compiler is written in Go, Go is faster 10 times than PHP, so the compilation
will be fast enough , and of course you can compile sources concurrently.

CoffeePHP takes the good stuffs from Haskell, Go and CoffeeScript, so you could write
shorter code to generate the expected PHP code. e.g. To define a function, we use:

    getName :: -> "John"

Which is shorter than writing:

    function getName() {
        return "John";
    }

CoffeePHP aims to provide a simple optimizer to do optimizations like "dead
code elimination".. or so on.

For more details, please check the `docs` for the language synopsis.



Build
---------
Simply run make to produce the parser:

    make


