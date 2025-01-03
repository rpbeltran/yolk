# Yolk

A stack based Virtual Machine for executing eggshell

## Implemented Instructions

|  Instruction  |         Argument(s)        |
| ------------- | -------------------------- |
| ASSIGN_NAME   | name: *Name*               |
| BINOP         | operation: *[add]*         |
| DECLARE_NAME  | name: *Name*               |
| EXEC          | arg_count: *uint*          |
| JUMP          | label: *uint*              |
| JUMP_IF_TRUE  | label: *uint*              |
| JUMP_IF_FALSE | label: *uint*              |
| .LABEL        | label: *uint*              |
| LOAD_NAME     | name: *Name*               |
| PIPELINE      | mode: *[begin, next, end]* |
| PRINT         |                            |
| PUSH_BOOL     | value: *[true, false]*     |
| PUSH_NUM      | value: *Number*            |
| PUSH_STR      | value: *Quoted*            |


*EXEC is only implemented in mock execution mode

### Todo Instructions

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