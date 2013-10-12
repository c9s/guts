Condition
============


If statement
------------
```
if a > 3 and b < 2
    a = 0
else if a == 0
    a = 1
else
    a = 2
```

which is compiled to:

```php
if ( $a > 3 && $b < 2) {
    $a = 0;
} else if ( $a == 0 ) {
    $a = 1;
} else {
    $a = 2;
}
```
