CoffeePHP
==================

CoffeePHP is a language makes your life easier,
it let you write less code to compile a workable PHP code.

CoffeePHP compiler is written in Go, Go is faster 10+ times than PHP, hence the
compilation is fast, and of course you can compile sources concurrently.

CoffeePHP takes the good stuffs from Haskell, Perl6, Go and CoffeeScript, you can write
shorter code to generate the expected PHP code. e.g. To define a function, we use:

    getName :: -> "John"

Which is shorter than writing:

    function getName() {
        return "John";
    }

CoffeePHP aims to provide a simple optimizer to do optimizations like "dead
code elimination"..

So if you have framework with development/production mode, this dead code
elimination could improve the overhead of mode checking for your production
site.

For more details, please check the `docs` for the language synopsis.


CoffeePHP aims to be
---------------------

* simple to learn.
* easy to use.
* fast compilation.
* shorter code.
* phpdoc friendly.


Build
---------
Simply run make to produce the parser:

    git clone http://github.com/c9s/coffeephp
    cd coffeephp
    source goenv
    make

Synopsis
---------
For more language details, please check the `docs` for the language synopsis.

    class Person
        say :: (name) -> "Hello #{name}, Good morning"
        getPhone :: () -> "12345678"
        
The above code compiles to:

```php
class Person {
    function say($name) {
        return "Hello " . $name . ', Good morning';
    }
    function getPhone() {
        return "12345678";
    }
}
```

File Extension
--------------------
The file extension is named with "\*.cphp", the compiler compiles your .cphp
files into .php file.


Implementation
---------------
CoffeePHP uses Go yacc parser generator to generate a LALR(1) parser. 

To add new syntax, please checkout the parser.y file, which is located in
`src/coffeephp/parser.y`

The lexer uses concurrent strategy to parse tokens, the state machine runs in
another goroutine, so the parser receives the tokens from lexer through the
channel concurrently.

Codegen is expected to be a separated tree structure traverser, and 
which is not implemented inside the ast node structure like the
coffee-script codegen. So the code generator can be PHP code generator, LLVM
bit-code generator or JavaScript generator.

Basic optimization is on the roadmap. Optimizer uses tree-pattern matching
strategy to traverse the IR structure. 2 basic optimizations are in the plan --
constant folding and dead code elimination.


Current State
-------------
Not ready for production. we're still in alpha stage. but we'd like to recevie
comments, patches and feature requests.


Lexer
------

    | Token Type          | Remark
--- | ------------------- | -------------
 x  | `T_IDENTIFIER     ` | 
 x  | `T_FLOATING       ` |
 x  | `T_NUMBER         ` |
 x  | `T_NEWLINE        ` |
 x  | `T_EOF            ` |
 x  | `T_ONELINE_COMMENT` |
 x  | `T_COMMENT        ` |
 x  | `T_IF             ` |
 x  | `T_ELSEIF         ` |
 x  | `T_CLASS          ` |
 x  | `T_FOR            ` |
 x  | `T_FOREACH        ` | 
 x  | `T_STRING         ` |  "String", and 'string'
 x  | `T_ECHO           ` |  echo "output"
 x  | `T_DOES           ` |  class Gate does DoorInterface

Parser
---------

    | Rule                
--- | -------------------
 x  | assignment statement
 x  | if statement
 x  | if else statement
 x  | if elseif else statement
 o  | foreach block
 o  | for block
 o  | dot range
 o  | class property
 o  | class method
 o  | function call, method call as expr.
 o  | assignment from function call, method call


Author
-------------

Yo-An Lin (yoanlin93@gmail.com)

