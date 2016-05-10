Guts
==================
[![Build Status](https://secure.travis-ci.org/c9s/guts.svg)](http://travis-ci.org/c9s/guts)

Guts is a language makes your life easier,
it allows you to write less code to produce equivalent PHP code.

Guts compiler is written in Go, Go is faster than PHP, thus the compilation is
fast, and of course you can compile sources concurrently.

For more details, please check the `docs` for the language synopsis.

Guts aims to be
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
    # Print the name
    say :: (name) -> "Hello #{name}, Good morning"

    getPhone :: -> "12345678"

    setName :: (string name) -> @name = name

if str =~ /[a-z]/
    say "matched!"
```
        
The above code compiles to:

```php
class Person {
    /**
     * Print the name
     * 
     * @param mixed $name
     */
    function say($name) {
        return "Hello " . $name . ', Good morning';
    }

    function getPhone() {
        return "12345678";
    }

    /**
     * @param string $name
     */
    function setName($name) {
        $this->name = $name;
    }
}
if ( preg_match('[a-z]',$str) ) {
    echo "matched!";
}
```

File Extension
--------------------
The file extension is named with "\*.gs", the compiler compiles your .gs
files to .php files.





Guts Language
-------------------

To make guts compiler compatible with PHP, there are some syntax 
are simplified but the derived syntax shouldn't break the compiler.

(Move to separated doc file later)

### Constant

Defining const can be done by the same syntax of PHP.

```
const VAR = 10;
```

### Variable Definition

```
$foo = 'string'; // explicitly defined without type
```

```
var $foo = 'string'; // explicitly defined and  implicitly typed
```

```
int $foo = 10; // explicitly defined and explicitly typed
```

```
string $foo = "foo bar"; // explicitly defined and explicitly typed
```

### Array

```
string[] $list = []; // define an array of string.
```

### Array operations - map, reduce, filter

...

### Map

```
string[string] $contacts = []; // define a map for string -> string
$contacts["foo"] = "bar";
$contacts[0] = "bar"; // compilation error
```


### Control Structure

#### If

#### Foreach

#### While




### Function

The keyword `function` is optional, not required.

```
foo(Type $foo, Type $bar) : ReturnType {

}
```

### Class

The code below works

```
<?php
class Foo extends Bar implements DoorLocker {
}
```

But you can also write:

```
<?php
class Foo is Bar does DoorLocker {

}
```

The inheritance syntax was derived from Perl6, see
<https://doc.perl6.org/language/objects#Inheritance> for more details.

### Class Property

### Class Method

PHP method syntax still works in guts:

```
class Buffer {
    public function write($string) { ... }
}
```

However you can also do:

```
class Buffer {
    public write($string) : bool { ... }
}
```

The `function` keyword can be ignored.

To override a parent method, you can add a keyword "override" to explicitly 
define the override method.

```
class EncodedBuffer is Buffer {
    override public write($string) : bool { ... }
}
```

When `-Wmethod-implicitly-override` option is enabled, method overriding without 
`override` keyword will throw out the warning messages:


### Traits

...



### Namespace

Defining namespace is as the same as PHP:

```
namespace foo\bar;
```






Roadmap
-------------------

- [ ] Implement a parser that parses PHP5.3
- [ ] Getter generator











Implementation
---------------
Guts uses Go yacc parser generator to generate a LALR(1) parser. 

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
Not ready for production. this project is still in alpha stage. but I'd like
to recevie comments, patches and feature requests.


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
 x  | empty class
 x  | class with properties
 x  | class with methods
 x  | function call as expr
 o  | string concatenation
 o  | list structure
 o  | hash structure
 o  | for block
 o  | dot range
 o  | static method call
 o  | method call as expr.
 o  | assignment from method call
 o  | namespace design
 o	| map



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
