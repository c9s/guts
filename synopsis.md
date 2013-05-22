Synopsis
=========


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

## Array

indexed array:

```
a = [ "foo", "bar", "jack" ]
```

which is compiled to:

```php
$a = array( "foo", "bar", "jack" );
```

associative array:

```
a = [ foo: bar, key: value ]
```

which is compiled to:

```php
$a = array( 
    "foo": $bar,
    "key": $value,
);
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
```

which is compiled to:

```php
if ( $a > 3 ) {
    $a = 3;
}
```

## Class

```
class Wheel
    @params = []
    @name = "foo"

    def __constructor(@params):
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




