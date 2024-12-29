# Yolk Instruction Specifications


## List of Implemented Instructions
| Instruction | Argument(s) |
| ----------- | --------- |
| BINOP       | operation: *[add]* |
| EXEC*       | arg_count: *uint*   |
| PIPELINE    | mode: *[begin, next, end]* |
| PRINT       |  |
| PUSH_NUM    | value: *Number* |
| PUSH_STR    | value: *Quoted* |



## BINOP ${operation}

Pops two values off the stack and attempts to perform an operation with both values, and then push
the resulting value onto the stack. The first popped value is the right operand, and the second
popped value becomes the left operand. If the operations fails, execution will terminate with an
error state.

Arguments:
* operation: enum, specifies a binary operation to perform (see ./Operations.md for options)

Example:

```
# egg: (10 / 5)
# -- vm.stack:[]
PUSH_NUM 10
# -- vm.stack:[10]
PUSH_NUM 5
# -- vm.stack:[10 5]
BINOP divide
# -- vm.stack:[2]
```

## EXEC ${arg_count}

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


## PIPELINE ${mode}


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
LOAD_NAME add3
# -- vm.stack:[add3]
# -- vm.pipeline_states:[3]
PIPELINE next
# -- vm.stack:[]
# -- vm.pipeline_states:[f(3)]
LOAD_NAME g
# -- vm.stack:[g]
# -- vm.pipeline_states:[f(3)]
PIPELINE end
# -- vm.stack:[f(g(3))]
# -- vm.pipeline_states:[]
```

## PRINT

Pops the top value and writes it to the buffer specified by vm.stdout.
If the data stack is empty, execution will terminate with an error state.

Arguments: None

Example:

```
# egg: do {say("Hello World!")
# -- vm.stack:[]
# -- stdout:""
PUSH_STRING "Hello World!"
# -- vm.stack:[10]
PUSH_NUM 5
# -- vm.stack:[10 5]
BINOP divide
# -- vm.stack:[2]
```

## PUSH_NUM ${value}

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

## PUSH_STR ${value}

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