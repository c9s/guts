Class
===========

    class Person
        getName :: () -> "John"
        getPhone :: () -> "12345678"
        
The above code compiles to:

```php
class Person {
    function getName() {
        return "John";
    }
    function getPhone() {
        return "12345678";
    }
}
```

Method Scope
------------

Methods are default to public scope, to define private method, please use underscore as the prefix.

Static methods are defined by a tilde `~`

    class Person
        getName :: () -> "John"

        ~getInstance :: () -> ....

Or you can declare scope keyword explicitly:

    class Person
        # public method
        public getName :: () -> "John"

        # private method
        private privateMethodName :: () ->

        # static method
        public static getInstance :: () -> ....


