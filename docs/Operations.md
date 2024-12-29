# Binary Operators

These operators are supported by the `BINOP ${operator}` instruction.

| Operator | Left Operand | Right Allowed |
| ----------- | --------- | --------- |
| `add` | `Number` | `Number` |
| `subtract` | `Number` | `Number` |
| `multiply` | `Number` | `Number` |
| `divide` | `Number` | `Number` |
| `int_divide` | `Number` | `Number` |
| `modulus` | `Number` | `Number` |
| `power` | `Number` | `Number` |
| `concat` | `String` | any |
| `and` | `Bool` | any |
| `or` | `Bool` | any |


Left and Right Operand are first and second by popping order (as opposed to pushing order)
respectively so:
```
# right operand
PUSH_NUM 5
# left operand
PUSH_NUM 10 
# operation
BINOP divide
```
Will result in 2 being placed on the stack

For the descriptions below, `left` designates the first operand and `right` designates the right
operand.

## Math Operators

### add

Pushes `left + right` to the stack.

`left` and `right` must both be Numbers otherwise execution will terminate with an error.

### subtract

Pushes `left - right` to the stack. `left` and `right` must both be Numbers otherwise execution will
terminate with an error.

### multiply

Pushes `left * right` to the stack.

`left` and `right` must both be Numbers otherwise otherwise execution will terminate with an error.

### divide

Pushes `left / right` to the stack. `left` and `right` must both be Numbers otherwise execution will
terminate with an error. If `right` is zero, execution will terminate with an error.

### int_divide

Pushes `left / right` is computed, then the result is truncated towards zero before being pushed to
the stack. `left` and `right` must both be Numbers otherwise execution will terminate with an error.
If `right` is zero, execution will terminate with an error.

### modulus

Pushes `left % right` to the stack. 
`left` and `right` must both be Numbers otherwise execution will terminate with an error.
If `right` is zero, execution will terminate with an error.
If `left` is negative and `right` is not a whole number, execution will terminate with an error.

### power

Pushes `left ** right` to the stack. 
`left` and `right` must both be Numbers otherwise execution will terminate with an error.

## Math Operators

### concat

Pushes `left ++ right`, the result of concatenating left and right together as strings to the stack.

`left` must be a string otherwise a "recoverable error" will be pushed to the top of
the stack. If `right`  is not a string, it will be converted to a string with it's ToString() method
before concatenating it.

## Math Operators

### and

Pushes `left && right` to the stack.

This behavior is equivalent to the ternary expression:

    left ? right : left

`left` must both be Bools otherwise execution will terminate with an error. 
If `right` is not a Bool, it will be interpreted as a bool by testing if it is `truthy`.

**Note:** This method is only suitable for implementing `and`/`&&` between two identifiers and should
not be used in the case where one side is a generic expression which may include side effects such
as an execution since it would not be possible to implement shortcutting with this 

### or

Pushes `left || right` to the stack.

This behavior is equivalent to the ternary expression:

    left ? left : right

`left` must both be Bools otherwise execution will terminate with an error. 
If `right` is not a Bool, it will be interpreted as a bool by testing if it is `truthy`.

**Note:** This method is only suitable for implementing `or`/`||` between two identifiers and should
not be used in the case where one side is a generic expression which may include side effects such
as an execution since it would not be possible to implement shortcutting with this 

