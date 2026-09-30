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
* *Performance:* This instantly deallocates the memory frame. The values `4`, `5`, and `9` are not erased, but their addresses are marked as available to be overwritten by the next function call. The CPU then reads the **Return PC** to jump back to the caller's code.

---


