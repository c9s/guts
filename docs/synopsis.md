Synopsis
=========


## Overview

```
var foo = 1
bar = 1

multiply = (value) ->
    return value * 2
```

```php
$foo = 1;
$bar = 1;
```


## Variable

```
var success = true
var number = 1
var floatVariable = 1.23
var str = "string"
```

which is compiled to:

```php
$success = true;
$number = 1;
$floatVariable = 1.23;
$str = "string";
```

associative array with yaml format definitions:

```
a = 
    foo: bar
    key: value
```


## Conditions

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
if ( $a > 3 ) {
    $a = 3;
} else if ( $a == 0 ) {
    $a = 1;
}
```

## Class

```
class Wheel
    @params = []
    @name = "foo"

    def __constructor(name, @params)
        # logics here

```

which is compiled to:

```php
class Wheel {
    public $params = array();
    public $name = "foo";

    public function __constructor($params) {
        $this->params = $params;
    }
}
```


Extending class

```
class Human extends Animal
    @name = "foo"
    @phone = ""
```

## Built-In HTML Template

```
template = %%%
    <div>name: {{ name }}</div>
    <div>phone: {{ phone }}</div>
    {% for i in 1..5 %}
        {{ i }}
    {% endfor %}
    %%%
```

which is compiled to:

```php
$template = function($vars = array() ) {
    ?><div>name: <?= $vars["name"] ?><?
    ?><div>phone: <?= $vars["phone"] ?><?
    for ( $i = 0 ; $i < 5 ; $i++ ) { ?> 
        <?= $i ?>
    <?php endfor
    <?
};
```




