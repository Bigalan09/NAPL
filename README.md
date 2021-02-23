# The Hoora Programming Language

Hoora is an open source programming language that was built purely for fun and educational purposes.

Unless otherwise noted, the Hoora source files are distributed under the
MIT license found in the LICENSE file.


## Language Specification
Must Have:
- Comments `//` or `/*  */`
- Statically typed variables. `var name = "";`
- Data types: `string`, `number`, `boolean`, `null`
- Variables can be defined as arrays. `var digits = [1, 2, 3]`
- Variables are by default immutable, add the keyword ` mut` to make a variable mutable. `var mut name = "name";`
- Precedence and grouping `var average = (min + max) / 2;`
- A basic standard library containing `print()` and `clock()`.
- Control Flow: `if` statements, `while` loop, `for` loop
- Functions
```
function addPair(a, b) {
  return a + b;
}
```
- Closures
```
function addPair(a, b) {
  return a + b;
}

function do(a) {
  return a;
}

print(do(addPair)(1, 2)); // Prints "3".
```
- Function recursion.
```
function fibonacci(n) {
    if (n < 2) {
        return 1;
    } else {
        return fibonacci(n - 2) + fibonacci(n - 1);
    }
}

print(fibonacci(5));
```

