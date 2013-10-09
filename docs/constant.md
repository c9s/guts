Constant
==============

Using constant is just like the `define` in PHP, you can simply write:

	const foo = 1

to get:

	define('foo', 1)

You can also define multiple constants at one time:

	const (
		foo = 1
		bar = 2
	)

This produces:

	define('foo',1);
	define('bar',2);


Enumerating
------------

You can also use `itoa` to enumerate the constant number:

	const (
		foo = itoa    // get 1
		bar           // get 2
	)




