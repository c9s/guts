Gutscript
==================

Gutscript is a language makes your life easier,
it allows you to write less code to produce equivalent PHP code.

Gutscript compiler is written in Go, Go is faster 10+ times than PHP, thus the
compilation is fast, and of course you can compile sources concurrently.

Gutscript takes the good stuffs from Haskell, Perl6, Go and CoffeeScript, you can write
shorter code to generate a equivalent PHP code. e.g. To define a function, we use:

    getName :: -> "John"

Which is shorter than writing:

    function getName() {
        return "John";
    }

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

    class Person
        say :: (name) -> "Hello #{name}, Good morning"
        getPhone :: -> "12345678"

    if str =~ /[a-z]/
        say "matched!"
        
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
 x  | function statement
 x  | function param list
 o  | list structure
 o  | hash structure
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



License
---------------------
MIT License Copyright (c) 2012 Yo-An Lin

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
