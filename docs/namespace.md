Namespace
===========

PHP Compatible Solution
---------------------------

While we need to support the below code:

    namespace Foo\Bar;
    $class = "Foo\\Bar\\Zoo";
    $obj = new $class;
    $obj = new \Foo\Bar\Zoo;

We might consider use the same separator of namespace:

    namespace Foo\Bar
    use Foo\Bar
    use Foo\Bar;   // semicolon is optional.

    obj = new Bar
    className = "Foo\\Bar\\Zoo"
    obj = new className
    obj = new \Foo\Bar\Zoo

SemiColon As Separator
-----------------------

    namespace Foo::Bar
    use Classname
    use Foo::Bar

    obj = new Bar
    className = "Foo\\Bar\\Zoo"   // we can't detemine the class name from a string
    obj = new className
    obj = new ::Foo::Bar::Zoo


