Sorting
===============

A basic sorting

```
list = sort (a,b) -> { a <=> b } list
```

compiles to:

```php
function __sort1($a, $b) {
    if ( $a == $b ) {
        return 0;
    }
    return ( $a < $b ) ? -1 : 1;
}
```
