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

Inheritance
-------------
You can simply write `is` keyword to define an inheritance:

    class Person is Object do ArrayIterator
        getName :: () -> "name"

The above code compiles to:

```php
class Person extends Object implements ArrayIterator
{
    function getName() {
        return "name";
    }
}
```

Implements Interface
--------------------
We use simpler keyword to define interface implementation:

    class LockableDoor is Object does DoorInterface
        openDoor :: () -> ..

Which compiles to

```php
class LockableDoor extends Object implements DoorInterface {
    function openDoor() {
    }
}
```

Class Members
-------------
To define class members:

    class LockableDoor is Object does DoorInterface
        @password = "password"

And here is the way to access the member variable:

    class LockableDoor is Object does DoorInterface
        @password = "password"

        public @username = "username"
        private @password = "password"

        openDoor :: (password) ->
            if @password == password
                # do something



Method Scope
------------

Methods are default to public scope, to define private method, please use
underscore as the prefix.

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


Magic Methods
--------------

Just like what you did in PHP 

    class Person
        __get :: (key) ->
        __set :: (key, val) ->
        __call :: (methodName, args) ->


Calling Object Method
---------------------

    obj = new Person
    name = obj.getName()

