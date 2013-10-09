Variable
========

Integer, String, Floating Number
---------------------------------

To define a variable:

    a = 1

    b = "string"

    floatingNumber = 3.1415926

compiles to:

```php
$a = 1;
$b = "string"
$floatingNumber = 3.1415926;
```

Hash
----

To define a hash:

    varKey = "key3"
    dict = {
        key1: value1
        key2: value2
        varKey: value3
        "varKey": value4
    }

compiles to:

```php
    $dict = array(
        "key1" => $value1,
        "key2" => $value2,
        $varKey => $value3,
        "varKey" => $value4,
    );
```

Please note that if you defined a variable and use the variable name in the hash table as a key,
the value of the key will be the variable.

Array
------

To define an array:

    list = [1,2,3,4]
    list = [1..10]

compiles to:

```php
$list = array(1,2,3,4);
$list = range(1,10);
```

We compile array with `array()` for the backward compatibility, but it would be nice to have 
an option to configure the generated code version for PHP.



