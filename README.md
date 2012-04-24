CoffeePHP
=========
CoffeePHP is a little language that compiles into PHP.

CoffeePHP aims to be *like* CoffeeScript, you can compile CoffeePHP into PHP or
JavaScript.

**currently alpha stage**

<table>
<tr>
<td valign="top"><pre>

number = 42;
opposite = true;
number = -42 if opposite

singers = {Jagger: "Rock", Elvis: "Roll"}

array = 1..10;
array.map(  (e) -> e * e );

obj.method( );

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
array_map(function($e) { return $e * $e; },$array);

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
