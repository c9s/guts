Synopsis
=========

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




