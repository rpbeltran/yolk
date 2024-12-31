# Yolk

A stack based Virtual Machine for executing eggshell

## Implemented Instructions

|  Instruction  |         Argument(s)        |
| ------------- | -------------------------- |
| BINOP         | operation: *[add]*         |
| EXEC          | arg_count: *uint*          |
| JUMP          | label: *uint*              |
| JUMP_IF_TRUE  | label: *uint*              |
| JUMP_IF_FALSE | label: *uint*              |
| .LABEL        | label: *uint*              |
| PIPELINE      | mode: *[begin, next, end]* |
| PRINT         |                            |
| PUSH_NUM      | value: *Number*            |
| PUSH_STR      | value: *Quoted*            |


*EXEC is only implemented in mock execution mode

### Todo Instructions

* LOAD_LOCAL
* LOAD_GLOBAL
* STORE_LOCAL
* STORE_GLOBAL
* BINOP_INPLACE
* UNOP
* UNOP_INPLACE
* SLICE
* BUILD_LIST
* BUILD_MAP
* BUILD_SET
* BUILD_CALLABLE
* DEFINE_STRUCT
* ASSERT
* ...

## Implemented Binary Operators

* add
* subtract
* multiply
* divide
* int_divide
* power
* modulus
* concat
* and
* or

### Todo Operators

* compare_eq
* compare_lt
* compare_lt
* curry
* ...

## Implemented Types

* Number: *Arbitrary precision rational numbers*
* String: *UTF-8 encoded strings*

### TODO Types

* Boolean
* Integer
* Struct
* List
* Map
* Set
* Callable