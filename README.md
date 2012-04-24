CoffeePHP
=========
CoffeePHP is a little language that compiles into PHP.

CoffeePHP aims to be *like* CoffeeScript, you can compile CoffeePHP into PHP or
JavaScript.

<table>
<tr>
<td><pre>

number = 42;
opposite = true;
number = -42 if opposite

</pre></td>

<td><pre>

$number = 42;
$opposite = true;
if( $opposite ) {
    $number = -42;
}

</pre></td>
</tr>
</table>

Reference
=========
- http://phpjs.org/
- https://github.com/kvz/phpjs
