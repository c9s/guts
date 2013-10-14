Regular Expression
========================

Here comes the syntax surger:

```
if a =~ qr/[a-z]/
    # ... do something
```

compiles to PHP:

```php
if ( preg_match("[a-z]", a) ) {
    // ... do something
}
```

