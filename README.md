# The Worra Programming Language

Worra is an open source programming language that was built purely for fun and educational purposes.

Worra is arrow spelt backwards.

Unless otherwise noted, the Worra source files are distributed under the
MIT license found in the LICENSE file.


## Language Specification
Must Have:
- Functions `function calculate() {}`
- Statically typed variables. `string name = "";`
- Variables are by default immutable, add the keyword ` mut` to make a variable mutable. `string mut name = "name";`
- Function recursion
```
function fibonacci(number n) {
    if (n < 2) {
        return 1;
    } else {
        return fibonacci(n - 2) + fibonacci(n - 1);
    }
}
```
-  
