Gutscript
==================
[![Build Status](https://secure.travis-ci.org/c9s/gutscript.png)](http://travis-ci.org/c9s/gutscript)

Gutscript is a language makes your life easier,
it allows you to write less code to produce equivalent PHP code.

Gutscript compiler is written in Go, Go is faster 10+ times than PHP, thus the
compilation is fast, and of course you can compile sources concurrently.

Gutscript takes the good stuffs from Haskell, Perl6, Go and CoffeeScript, you can write
shorter code to generate a equivalent PHP code. e.g. To define a function, we use:

```hs
hello :: (name)-> "Hello #{name}!"
```

Which is shorter than writing:

```php
function hello($name) {
    return "Hello " . $name . "!";
}
```

Gutscript aims to provide a simple optimizer to do optimizations like "dead
code elimination"..

So if you have framework with development/production mode, this dead code
elimination could improve the overhead of mode checking for your production
site.

For more details, please check the `docs` for the language synopsis.


Gutscript aims to be
---------------------

* simple to learn.
* easy to use.
* fast compilation.
* shorter code.
* phpdoc friendly.


Build
---------
Simply run make to produce the parser:

    git clone http://github.com/c9s/gutscript
    cd gutscript
    source goenv
    make

Synopsis
---------
For more language details, please check the `docs` for the language synopsis.

```coffee
class Person
    say :: (name) -> "Hello #{name}, Good morning"
    getPhone :: -> "12345678"

if str =~ /[a-z]/
    say "matched!"
```
        
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
if ( preg_match('[a-z]',$str) ) {
    echo "matched!";
}
```

File Extension
--------------------
The file extension is named with "\*.guts", the compiler compiles your .guts
files to .php files.


Implementation
---------------
Gutscript uses Go yacc parser generator to generate a LALR(1) parser. 

To add new syntax, please checkout the parser.y file, which is located in
`src/gutscript/parser.y`

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


Contribution
------------------------------
Feature requests, pull requests, comments are welcomed, but please discuss first 
on our GitHub issue. Just feel free to drop a line there.



Development
------------------------------

First, go get a workable go compiler at <http://golang.org>

Run the below command to setup GOPATH:

	source goenv

Run make to build and test:

	make

If you modified the grammar, remember to update the parser:

    make parser

Run test cases:

    make test

To dump lex tokens:

    go run utils/lexdump/main.go src/gutscript/tests/09_class_properties.guts


Test the command tool with test case:

    ./bin/guts src/gutscript/tests/03_if_else_statement.guts


Grammar Rules
-------------

    | Rule                
--- | -------------------
 x  | assignment statement
 x  | if statement
 x  | if else statement
 x  | if elseif else statement
 x  | function statement
 x  | function param list
 o  | string concatenation
 o  | list structure
 o  | hash structure
 o  | foreach block
 o  | for block
 o  | dot range
 x  | empty class
 o  | class with properties
 o  | class with methods
 o  | static method call
 o  | function call, method call as expr.
 o  | assignment from function call, method call
 o  | namespace design



Author
-------------

Yo-An Lin (yoanlin93@gmail.com)


IRC Channel
----------------
Join us on irc.freenode.net #gutscript


License
---------------------
MIT License Copyright (c) 2012 Yo-An Lin

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
