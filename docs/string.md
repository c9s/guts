String
================

Declaration
-------------------

```
str = "Hello World"
```

This compiles to:


String Concatenation
--------------------

str = "Hello" ++ "World"


Since PHP does not use "+" for string concatenation, and we have use "." as the
object method operator, we can't use the same operator to express string
concatenation, which is ambiguous for the parser and generator.  This is a
design decision.


String Interpolation
------------------

```
str = "Hello #{ name }"
```


Virtual String Functions
--------------------------

Although you could write function call to call the PHP string functions, but
you might already know the design of PHP string functions are inconsistent.
hence we provide a runtime package for you to call the same function but with
more consistent API.

```
import "strings"
out = strings.replace "needle", "something", "new"
```

Which is the same as:

```php
$out = str_replace("needle", "something", "new")
```




