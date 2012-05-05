CoffeePHP
=========
CoffeePHP is a little language that compiles into PHP.

CoffeePHP aims to be *like* CoffeeScript, you can compile CoffeePHP into PHP or
JavaScript.

**Currently Alpha Stage**


## Synopsis

    class Man is Person does SomeInterface
        has name
        has private address
        has private phone

        new: ->

        move: (x,y) ->

    man = new Man

## Transformation

<table>
<tr>
<td valign="top"><pre>

number = 42;
opposite = true;
number = -42 if opposite

singers = {Jagger: "Rock", Elvis: "Roll"}

array = 1..10
array2 = array.map (e) -> e * e

hash = { key: 1, foo: "bar" }

if hash.foo?
    echo hash.foo

obj.method key: value, key2: value

</pre></td>

<td valign="top"><pre>

$number = 42;
$opposite = true;
if( $opposite ) {
    $number = -42;
}

$singers = array(
    "Jagger" => "Rock", 
    "Elvis" => "Roll"
);

$array = range(1,10);
$array2 = array_map(function($e) { return $e * $e; },$array);

$obj-&gt;method(array( "key" =&gt; value, "key2" =&gt; value ));

$hash = array(
    'key' => 1,
    'foo' => 'bar',
);

if( isset($hash['foo']) ) {
    echo $hash['foo'];
}

</pre></td>
</tr>
</table>

Hacking
-------

The current recommended way to build CoffeePHP is via cabal-install with the
Haskell Platform.  The steps are:

0. Download and install the Haskell Platform:

    http://hackage.haskell.org/platform/

Alternately, one can also install only the GHC compiler, and manually set up
the following packages (in this order):

    http://hackage.haskell.org/cgi-bin/hackage-scripts/package/Cabal
    http://hackage.haskell.org/cgi-bin/hackage-scripts/package/HTTP
    http://hackage.haskell.org/cgi-bin/hackage-scripts/package/zlib
    http://hackage.haskell.org/cgi-bin/hackage-scripts/package/cabal-install

The instructions contained in the README file of Cabal contain more detailed
setup instructions, which should work for all the packages above.

(But then again, please consider simply installing the Haskell Platform. :-))

1. Install dependencies

    cabal update
    cabal install

2. Hack Hack

3. Build

    make 
