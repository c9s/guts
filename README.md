CoffeePHP
=========
CoffeePHP is a little language that compiles into PHP.

CoffeePHP aims to be *like* CoffeeScript, you can compile CoffeePHP into PHP or
JavaScript.

**currently alpha stage**

<table>
<tr>
<td><pre>

number = 42;
opposite = true;
number = -42 if opposite

array = [ 1..10 ];
array.map(  (e) -> e * e );

</pre></td>

<td><pre>

$number = 42;
$opposite = true;
if( $opposite ) {
    $number = -42;
}

$array = range(1,10);
array_map(function($e) { return $e * $e; },$array);

</pre></td>
</tr>
</table>

Reference
=========
- http://phpjs.org/
- https://github.com/kvz/phpjs
