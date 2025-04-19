# Yolk

A stack based Virtual Machine for executing eggshell

## Implemented Instructions

|  Instruction  |             First Argument         |     Second Argument     |
| ------------- | ---------------------------------- | ----------------------- |
| ASSIGN        | name: *Name*                       |                         |
| BINOP         | operation: *[add, and, ...]*       |                         |
| BINOP_INPLACE | operation: Operation               | name: *Name*            |
| COMPARE       | test_mode: *[equal, less, ...]*    |                         |
| COMPARE_CHAIN | test_mode: *[equal, less, ...]*    |                         |
| DECLARE       | name: *Name*                       | type: *Name* (optional) |
| .DEFINE       | name: *Name*                       | type: *Name* (optional) |
| .DEFINE_END   |                                    |                         |
| EXEC          | arg_count: *uint*                  |                         |
| JUMP          | label: *uint*                      |                         |
| JUMP_IF_TRUE  | label: *uint*                      |                         |
| JUMP_IF_FALSE | label: *uint*                      |                         |
| .LABEL        | label: *uint*                      |                         |
| LOAD          | name: *Name*                       |                         |
| NEGATE        |                                    |                         |
| NOT           |                                    |                         |
| PIPELINE      | mode: *[begin, next, end]*         |                         |
| PRINT         |                                    |                         |
| PUSH_BOOL     | value: *[true, false]*             |                         |
| PUSH_INT      | value: *int*                       |                         |
| PUSH_NUM      | value: *Number*                    |                         |
| PUSH_STR      | value: *Quoted*                    |                         |


*EXEC is only implemented in mock execution mode

### Todo Instructions

* .DEFINE
* SCOPE_BLOCK push/pop
* CALL
* BUILD_LIST
* SLICE
* BUILD_MAP
* SELECT
* BUILD_SET
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
* Int: *64-bit signed integers*
* String: *UTF-8 encoded strings*
* Bool: *boolean value; true or false*

### TODO Types

* Struct
* List
* Map
* Set
* Callable
