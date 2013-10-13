BuiltIn
=======

GutsPHP builtin provides a consistent interface to solve the inconsistent PHP API,
for example, PHP string functions sometimes use "str_" as prefix, sometimes use "str" as prefix.

By using GutsPHP builtin package, you can use these functions with consistentcy:

    import "strings"

    strings.pad( )   => str_pad()
    strings.repeat( )  => str_repeat()
    strings.replace(  )
    strings.indexOf(  )
    strings.join(needle, sep)
    strings.explode(needle, sep)
    strings.cmp( )   => strcmp()
