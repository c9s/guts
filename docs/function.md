Function
=========

To define a function without parameters:

    foo :: ->
        return 1

Or even shorter:

    foo :: -> 1

the function will return the last expression, and the above code compiles to:

```php
function foo() {
    return 1;
}
```

To define a function with parameters:

    foo :: (x,y) -> x * y

The above code compiles to:

    function foo($x, $y) {
        return $x * $y
    }


Type Constraint
---------------

You can define type constraint to the function declaration:

    foo :: (int x) int ->
        if x > 3
            return 1
        else
            return nil

This makes your function consistent, but we can't do type constraint for runtime.

