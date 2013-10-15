For Block
===========

```
for i in [ 1..10 ]
	say i
```

compiles to:

```php
foreach( array(1,2,3,4,5,6,7,8,9,10) as $i ) {
	echo $i;
}
```

When using large number in dot range, e.g.,

```
for i in [ 1 .. 10000 ]
	say i
```

should automatically use `range` function to create the list:

```php
foreach( range(1,10000) as $i ) {
	echo $i;
}
```






