Map
=========================
The concept of 'map' is from Perl, although there are `array_filter` and `array_map` in PHP,
you should already know that it has a bad design on the function prototype.

`array_map(function, list)`
`array_filter(list, function)`

People usually get confused with these two functions, because the prototypes are different.


why not use map?
----------------------
In perl, we can do:

```perl
@list = map { $_ * $_ } @some;
```

So you can also do `map` in the gutscript:

```
list = map (x) -> { x * x } some
```

And this will be compiled to:

```php
$list = array();
for ( $_i = 0 ; $_i < count($some) ; $_i++ ) {
    $list[] = $some[$_i] * $some[$_i];
}
```





