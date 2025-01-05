# Yolk

A stack based Virtual Machine for executing eggshell

## Implemented Instructions

|  Instruction  |             Argument(s)            |
| ------------- | ---------------------------------- |
| ASSIGN_NAME   | name: *Name*                       |
| BINOP         | operation: *[add, and, ...]*       |
| BINOP_INPLACE | operation: Operation, name: *Name* |
| COMPARE       | test_mode: *[equal, less, ...]*    |
| DECLARE_NAME  | name: *Name*                       |
| EXEC          | arg_count: *uint*                  |
| JUMP          | label: *uint*                      |
| JUMP_IF_TRUE  | label: *uint*                      |
| JUMP_IF_FALSE | label: *uint*                      |
| .LABEL        | label: *uint*                      |
| LOAD_NAME     | name: *Name*                       |
| NEGATE        |                                    |
| NOT           |                                    |
| PIPELINE      | mode: *[begin, next, end]*         |
| PRINT         |                                    |
| PUSH_BOOL     | value: *[true, false]*             |
| PUSH_NUM      | value: *Number*                    |
| PUSH_STR      | value: *Quoted*                    |


*EXEC is only implemented in mock execution mode

### Todo Instructions

* DECLARE_NAME_TYPED
* .DEFINE
* .DEFINE_TYPED
* SUB_SCOPE
* CALL
* CALL_NAME


* BUILD_LIST
* SLICE
* BUILD_MAP
* SELECT
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