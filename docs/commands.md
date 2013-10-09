Commands
==================

This section describes the command-line tool usage.


Watch and Compile
-----------------
Watch `src` directory and compiles to `lib`:

    cphp -wc src:lib


Optimization
----------------
Optimization flag is defined by `-O`, e.g.:

    cphp -O -wc src:lib

We only support dead code elimination.  Conditions like:

    if 0 and 1
        # ... code

Will be eliminated.
