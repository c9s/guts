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

Inline Function
-------------------
Inline function will be expanded in place to replace the real function call:

    inline foo :: (x,y) -> x * y
    z = foo(1,3)

compiles to:

    z = 1 * 3


Calling Functions
------------------
You can call functions like this:

    r = rand(0.1)
    call_user_func(obj, "method")

parenthesis of function call is optional:

    r = rand 0.1
    call_user_func obj, "method"




Type Constraint
---------------

You can define type constraint to the function declaration:

    foo :: (int x) int ->
        if x > 3
            return 1
        else
            return nil

This makes your function consistent, but we can't do type constraint for runtime.

