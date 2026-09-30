# How a Normal and Closure Go Function Works Under the Hood

This document breaks down exactly how the Go runtime and CPU execute a standard, non-closure and closure function at the hardware level using the following example and diagram:

<img width="1162" height="837" alt="Blank diagram" src="https://github.com/user-attachments/assets/81c50eed-3cea-4a51-af45-03af71f82edc" />


```go
func SampleFunc() int {
    var num1, num2 int = 4, 5
    var sum int = num1 + num2
    return sum
}
```

---

## Memory Layout (The Stack Frame)
Before execution, the Go compiler calculates the exact size needed for the function's variables. It reserves a dedicated block of memory on the thread stack called a **Stack Frame**:

## Step-by-Step Execution Lifecycle

### Step 1: Function Prologue (Allocation)
When `SampleFunc` is called, the CPU adjusts its **Stack Pointer (SP)** register downward to reserve the exact number of bytes required for this function's stack frame. 
* *Performance:* This is near-instantaneous, requiring only a single CPU subtraction instruction.

### Step 2: Variable Initialization
```go
var num1, num2 int = 4, 5
```
The CPU copies the literal values `4` and `5` from the compiled binary's data section directly into the stack memory locations reserved for `num1` and `num2`.

### Step 3: Arithmetic Operations
```go
var sum int = num1 + num2
```
1. The CPU moves the value of `num1` (`4`) from the stack into a CPU register (e.g., `R1`).
2. The CPU moves the value of `num2` (`5`) into a second CPU register (e.g., `R2`).
3. The Arithmetic Logic Unit (ALU) executes an `ADD` instruction: `R1 = R1 + R2`.
4. The CPU writes the resulting value (`9`) from the register back to the stack address designated for `sum`.

### Step 4: Return Value Assignment
```go
return sum
```
Before exiting, the CPU reads the value of `sum` (`9`) and copies it into the **Return Value Slot** allocated by the caller function (or passes it via a hardware register depending on Go's internal ABI design).

### Step 5: Function Epilogue (Deallocation)
The function completes its execution. The CPU increments the **Stack Pointer (SP)** register back to its original position prior to the function call. 
* *Performance:* This instantly deallocates the memory frame. The values `4`, `5`, and `9` are not erased, but their addresses are marked as available to be overwritten by the next function call The CPU then reads the **Return PC** to jump back to the caller's code.

---

# How a Closure Function Works Under the Hood in Go

<img width="1052" height="837" alt="Blank diagram(1)" src="https://github.com/user-attachments/assets/8a6b13c2-d9a0-4c8f-ac6c-bb81fdd09f19" />

```go
func ClosureFunc() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

---

## The Compiler's Hidden Struct
When Go compiles a closure, it realizes the inner anonymous function must access a variable (`count`) belonging to an outer scope. To handle this, the compiler generates a hidden internal structure to represent the closure value:

```go
// Hidden struct created behind the scenes by the compiler
type closureStruct struct {
    F     uintptr // Pointer to the underlying anonymous function machine code
    count *int    // A pointer to the captured 'count' variable on the heap
}
```

---

## Step-by-Step Execution Lifecycle

### Step 1: Escape Analysis & Heap Promotion
When `ClosureFunc` begins execution, the compiler's **Escape Analysis** engine intervenes.
* **The Problem:** If `count` were allocated on the standard stack frame, it would be instantly wiped out when `ClosureFunc` returns. The inner function would then point to corrupted or missing memory.
* **The Solution:** Go skips the stack and calls `runtime.newobject` to allocate 8 bytes for `count` directly on the **Heap**. `count` is initialized to its zero value (`0`).

### Step 2: Constructing the Closure Object
```go
return func() int { ... }
```
`ClosureFunc` creates an instance of the hidden `closureStruct`:
1. It assigns the `F` field to the static memory address of the anonymous function's compiled machine code.
2. It assigns the `count` field to the **heap memory address** generated in Step 1.
3. The function returns this 2-word pointer interface (Code Pointer + Context Pointer) to the caller. `ClosureFunc` exits and its local stack frame is deallocated, but `count` safely survives on the heap.

### Step 3: Invoking the Closure (The Hardware Bridge)
When the caller executes the returned closure (e.g., `myFunc()`), a new stack frame is created for the anonymous function. 

Because the anonymous function's machine code is static, it cannot hardcode a dynamic heap memory address. Go solves this using a strict hardware **Calling Convention (ABI)**:
1. The Go runtime automatically loads the memory address of the `closureStruct` into a specific hardware CPU register (typically `DX` on amd64/x86_64 architectures).
2. The anonymous function uses this `DX` register as a physical map to locate the closure context.

### Step 4: Dereferencing and State Mutation
```go
count++
return count
```
The CPU cannot execute arithmetic operations directly inside main memory (RAM/Heap). It must route everything through its internal registers:

1. **Read Context:** The CPU reads the address of the `closureStruct` out of the `DX` register.
2. **Dereference Heap Address:** The CPU looks inside the struct, extracts the heap pointer to `count`, and loads that pointer into register `AX`.
3. **Load Value:** The CPU reads the actual integer value (`0`) from the heap address stored in `AX` and copies it into math register `BX`.
4. **Increment:** The CPU's Arithmetic Logic Unit (ALU) increments `BX` from `0` to `1`.
5. **Write Back:** The CPU copies the new value `1` from register `BX` back down through the system memory bus into the physical heap address stored in `AX`.

---



