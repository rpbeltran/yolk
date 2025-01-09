# Yolk Instruction Specifications


## List of Implemented Instructions

|  Instruction  |             First Argument         |     Second Argument     |
| ------------- | ---------------------------------- | ----------------------- |
| ASSIGN        | name: *Name*                       |                         |
| BINOP         | operation: *[add, and, ...]*       |                         |
| BINOP_INPLACE | operation: Operation               | name: *Name*            |
| COMPARE       | test_mode: *[equal, less, ...]*    |                         |
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

## `ASSIGN ${name}`

Pops the value from the top of the stack and stores it in an existing variable.
If a variable with the given name does not exist in any of the current scopes,
or if the stack is empty execution will terminate with an error state.

If the variable was declared with a type constraint and the value popped from the stack does not
have the specified type, the value will be implicitly cast to the type if possible.
If the implicit cast is not possible, execution will terminate with
an error state.

Arguments:
* name: name of a variable to update the value of (angle quoted)

Example:

```
# egg: (foo)
# -- vm.stack:[7]
# -- vm.globals: {foo=1,}
DECLARE <foo>
# -- vm.stack:[]
# -- vm.globals: {foo=7,}
```

## `BINOP ${operation}`

Pops two values off the stack and attempts to perform an operation with both values, and then push
the resulting value onto the stack. The first popped value is the right operand, and the second
popped value becomes the left operand. If the operations fails, or if the stack has less than two
elements execution will terminate with an error state.

Arguments:
* operation: enum, specifies a binary operation to perform (see ./Operations.md for options)

Example:

```
# egg: (10 / 5)
# -- vm.stack:[]
PUSH_NUM 5
# -- vm.stack:[5]
PUSH_NUM 10
# -- vm.stack:[5 10]
BINOP divide
# -- vm.stack:[2]
```

## `BINOP_INPLACE ${operation} ${name}`

Attempts to update the value of a variable in memory based on the result of a binary operation with
its current value as the left hand side. Pops a value from the stack to use as the right hand side.
The same operations are supported as for BINOP.

If the operations fails or there is no variable with the given name, or if the stack is empty
execution will terminate with an error state.

Arguments:
* operation: enum, specifies a binary operation to perform (see ./Operations.md for options)
* name: name of the variable to update (angle quoted)

Example:

```
# egg: foo /= 3
# -- vm.stack:[]
# -- vm.globals{foo=15,}
PUSH_NUM 3
# -- vm.stack:[3]
# -- vm.globals{foo=15,}
BINOP_INPLACE divide <foo>
# -- vm.stack:[]
# -- vm.globals{foo=3,}
```

## `COMPARE ${test_mode}`

Pops two values off the stack and attempts to perform a comparison of the two values then push
the resulting value onto the stack. The first popped value is the right operand, and the second
popped value becomes the left operand. If the operations fails, execution will terminate with an
error state.

The resulting value depends on the test_mode provided, which can be any of the following:
* equal: check if left == right
* unequal: check if left != right
* less: check if left < right
* lte: check if left <= right
* greater: check if left > right
* gte: check if left >= right

Arguments:
* test_mode: enum, specifies the comparison operation to use.
Must be one of the above mentioned options.

Example:

```
# egg: (10 == 5)
# -- vm.stack:[]
PUSH_NUM 5
# -- vm.stack:[10]
PUSH_NUM 10
# -- vm.stack:[10 5]
COMPARE equal
# -- vm.stack:[false]
```

## `DECLARE ${name} ${type???}`

Pops the value from the top of the stack and stores it in a new variable in the current scope.
If a variable with the given name already exists in the current scope, or if the stack is empty
execution will terminate with an error state.

An optional second argument specifies a type constraint for the new variable. If this argument is
provided, the variable created will be required to have the given type for the life of the variable.

If the initial value popped from the stack does not have the specified type, the value will be
implicitly cast to the type if possible. If the implicit cast is not possible, execution will
terminate with an error state.

Arguments:
* name: name of a variable to create (angle quoted)
* type (optional): name of the type that this variable is restricted to (angle quoted)

Example:

```
# egg: (foo)
# -- vm.stack:[10 7]
# -- vm.globals{}
DECLARE <foo>
# -- vm.stack:[10]
# -- vm.globals{foo=7,}
DECLARE <foo2> <int>
# -- vm.stack:[]
# -- vm.globals{foo=7, foo2=10}
```

## `.DEFINE ${name} ${type???}`

Designates that instructions between this point and the next `.DEFINE_END` instruction are part of a
defined function with the given name.

Arguments:
 * name: name of the function being defined
 * type (optional): name of a type to return (angle quoted)

```
.DEFINE <say_hello>
PUSH_STRING "hello world!!!"
PRINT
.DEFINE_END
```

## `.DEFINE_END`

Designates an end to a defined function and return to the global scope.

Arguments: None


## `EXEC ${arg_count}`

Calls an executable in a new process. Pops the stack for the path to the executable then pops the
stack again *arg_count* times to get additional arguments. Arguments will be added in the order
they are popped. The path to execute must be a string, other arguments be serialized by their
ToString() methods.

If execution fails, a "recoverable error" will be pushed to the top of the stack, which will cause
a panic if popped, otherwise, a Result object will be pushed to the top of the stack.

If vm.MockExecutions is enabled, EXEC will instead return JSON blobs formatted as follows:
```
{
    "command\": "foo",
    "args": [a b c],
    "stdin":\"Hello World!\"
}
```
The stdin line will be omitted if not inside of a pipeline.

Arguments:
* arg_count: uint, amount of additional arguments to pop, not including the path to the executable

Example:

```
# egg: `echo a b c` 
# -- vm.stack:[]
PUSH_STR "c"
# -- vm.stack:["c"]
PUSH_STR "b"
# -- vm.stack:["c", "b"]
PUSH_STR "a"
# -- vm.stack:["c", "b", "a"]
PUSH_STR "foo"
# -- vm.stack:["c", "b", "a", "foo"]
EXEC 3
# -- vm.stack:["a b c"]
```


## `JUMP ${label_id}`

Unconditionally sets the instruction pointer to the label with the given id.

If label id is not in the vm's list of known labels, execution will terminate with an error state.

Arguments:
 * label_id: uint, a unique id associated with this label

```
# Example Egg:
# loop {
#   say("hello")
# }

.LABEL 123
PUSH_STR "hello"
PRINT
JUMP 123
```

## `JUMP_IF_FALSE ${label_id}`

Pops the top value from the stack, and if that value is false, sets the instruction pointer to the
label with the given id. If the data stack is empty, or if the top value if not Boolean, execution
will terminate with an error state.

If the data stack is empty, or if the top value if not Boolean, execution will terminate with an
error state.
If label id is not in the vm's list of known labels, execution will terminate with an error state.

Arguments:
 * label_id: uint, a unique id associated with this label

```
# Example Egg:
# if (a) {
#   say("hello")
# }
# say("goodbye")


LOAD  a
JUMP_IF_FALSE  123
PUSH_STR "hello"
PRINT
.LABEL 123
PUSH_STR "goodbye"
PRINT
```

## `JUMP_IF_TRUE ${label_id}`

Pops the top value from the stack, and if that value is true, sets the instruction pointer to the
label with the given id. 

If the data stack is empty, or if the top value if not Boolean, execution will terminate with an
error state.
If label id is not in the vm's list of known labels, execution will terminate with an error state.

Arguments:
 * label_id: uint, a unique id associated with this label

```
# Example Egg: (a or b)
LOAD  a
DUPLICATE
JUMP_IF_TRUE  123
LOAD  b
.LABEL 123
```

## `.LABEL ${label_id}`

Designates a point in program execution for JUMP statements to go to.
Executing a label instruction is a no-op.
Label locations become known to the VM when new instructions are added to the program.

Arguments:
 * label_id: uint, a unique id associated with this label


## `LOAD ${name}`

Loads a variable from memory and pushes it onto the top of the stack.
The variable will be searched for in local scope first and then in global scope if it cannot be
found in the local scope.
If a variable with the given name cannot be found in either the local scope or the global scope
execution will terminate with an error state.

Arguments:
* name: name of a variable to load onto the stack (unquoted)

Example:

```
# egg: (foo)
# -- vm.stack:[]
# -- vm.globals{foo=7,}
LOAD foo
# -- vm.stack:[7]
# -- vm.globals{foo=7,}
```

## `NEGATE`

Pops a value off the stack and attempts to push its negated value to the stack.
If the value is not a number or if the stack is empty, execution will terminate with an error state.

Arguments: None

Example:

```
# egg: -(10)
# -- vm.stack:[]
PUSH_NUM 10
# -- vm.stack:[10]
NEGATE
# -- vm.stack:[-10]
```

## `NOT`

Pops a value `val` off the stack and attempts to push `not(val)` to the stack.
If the value is not a number or if the stack is empty, execution will terminate with an error state.

Arguments: None

Example:

```
# egg: !(true)
# -- vm.stack:[]
PUSH_BOOL true
# -- vm.stack:[true]
NOT
# -- vm.stack:[false]
```

## `PIPELINE ${mode}`

Modifies the VMs pipeline_state stack to facilitate data pipelines. Behavior depends on *mode*:
* begin: Adds a new entry to the pipeline_state stack. The entry will have a value of nil.
* next: Pops the data stack and replaces the top of the pipeline_state stack with that value if that
    value is not a callable or if the top pipeline_state was previously nil.
    If the top of the data stack is a callable and the pipeline state is not nil, the callable will
    be invoked with the top of the pipeline_state and the pipeline_state will be replaced with the
    result of that invocation. If the pipeline_state stack or the data stack are empty, execution
    will terminate with an error state.

* end: If the top of the data stack is a callable, invoke the callable with the top value of the
    pipeline_state stack, and pop the pipeline_state stack. Otherwise, only pop the top value of
    the pipeline state stack. If the pipeline_state stack or the data stack are empty, execution
    will terminate with an error state.

Arguments:
* mode: enum, one of begin, next, or end. 

Example:

```
# egg: 1 | f | g
# -- vm.stack:[]
# -- vm.pipeline_states:[]
PIPELINE begin
# -- vm.stack:[]
# -- vm.pipeline_states:[nil]
PUSH_NUM 3
# -- vm.stack:[1]
# -- vm.pipeline_states:[nil]
PIPELINE next
# -- vm.stack:[]
# -- vm.pipeline_states:[3]
LOAD add3
# -- vm.stack:[add3]
# -- vm.pipeline_states:[3]
PIPELINE next
# -- vm.stack:[]
# -- vm.pipeline_states:[f(3)]
LOAD g
# -- vm.stack:[g]
# -- vm.pipeline_states:[f(3)]
PIPELINE end
# -- vm.stack:[f(g(3))]
# -- vm.pipeline_states:[]
```

## `PRINT`

Pops the top value and writes it to the buffer specified by vm.stdout.
If the data stack is empty, execution will terminate with an error state.

Arguments: None

Example:

```
# egg: do {say("Hello World!")}
# -- vm.stack:[]
# -- stdout:""
PUSH_STRING "Hello World!"
# -- vm.stack:[10]
PUSH_NUM 5
# -- vm.stack:[10 5]
BINOP divide
# -- vm.stack:[2]
```

## `PUSH_BOOL ${value}`

Pushes a boolean value to the top of the stack.

Arguments:
* value: boolean value to push, either 'true' or 'false' (unquoted)

Example:

```
# egg: (true)
# -- vm.stack:[]
PUSH_BOOL true
# -- vm.stack:[true]
```

## `PUSH_INT ${value}`

Pushes an integer value to the top of the stack.

Arguments:
* value: int, specifies an int to be pushed to the stack

Example:

```
# egg: (10)
# -- vm.stack:[]
PUSH_INT 10
# -- vm.stack:[10]
```

## `PUSH_NUM ${value}`

Pushes a numeric value to the top of the stack.

Arguments:
* value: number, specifies a number in decimal format with optional fractional 
    components (example 1.5) to be be pushed to the stack

Example:

```
# egg: (10)
# -- vm.stack:[]
PUSH_NUM 10
# -- vm.stack:[10]
```

## `PUSH_STR ${value}`

Pushes a string value to the top of the stack. The value should be double quoted and UTF-8.

Arguments:
* value: a double-quoted string to push to the top of the stack

Example:

```
# egg: (10)
# -- vm.stack:[]
PUSH_NUM 10
# -- vm.stack:[10]
```