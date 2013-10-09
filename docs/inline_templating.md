Inline Templating
==================
In the current PHP solutions, there are many template engines. these engines use
a lot of OO techniqs and Caching skills to improve the performance.

By using CoffeePHP, you can compile these template into static PHP code, so you 
event don't need a template engine to use nice template syntax and you could gain 
a nice performance.

To define a template block, use `%%%` keyword:

```
template :: (name, phone) ->
    # do other initialization... 
    %%%
    <div>name: {{ name }}</div>
    <div>phone: {{ phone }}</div>
    {% for i in 1..5 %}
        {{ i }}
    {% endfor %}
    %%%
```

which is compiled to:

```php
function template($name, $phone) {
    # do other initialization... 

    ?><div>name: <?= $vars["name"] ?><?
    ?><div>phone: <?= $vars["phone"] ?><?
    for ( $i = 0 ; $i < 5 ; $i++ ) { ?> 
        <?= $i ?>
    <?php endfor
    <?
};
```




