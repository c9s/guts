Inline Templating
==================
In the current PHP solutions, there are many template engines. these engines use
a lot of OO techniqs, to use it, you need declare loader, extension to initialize the engine.
and you also need cache to improve the performance.

By using CoffeePHP, you can simply compile a template block into static PHP
code. it's so lightweight, so you don't even need a template engine to use the simple
template syntax and you could gain a nice performance.

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




