# Understanding Go Generic Linked Lists: Walkthrough of `Push` and `AllElements`

---

## 1. The `Push` Method: Step-by-Step Pointer Tracing

The `Push` method appends a new generic value `v` to the end of the linked list:

```go
func (lst *List[T]) Push(v T) {
    fmt.Println("list:", lst.tail, "\n", "head:", lst.head)
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

```

### Execution Trace

#### Step 1: `Push(10)` (First Element)

* **Condition:** `lst.tail == nil` is `true`.
* **Execution:**
* `lst.head = &element[T]{val: 10}` ---> Creates node `{10, nil}`.
* `lst.tail = lst.head` -----> Assigns `lst.head` to `lst.tail`. Both head and tail now point to the exact same node: `{10, nil}`.



#### Step 2: `Push(13)` (Second Element)

* **Condition:** `lst.tail == nil` is `false` (tail points to `{10, nil}`).
* **Execution:**
* `lst.tail.next = &element[T]{val: 13}` ----> Appends the new node `{13, nil}` to the `next` pointer of the current tail. Because head and tail shared the same node reference, `lst.head.next` is also updated to `{13, nil}`.
* `lst.tail = lst.tail.next` -----> Moves `lst.tail` forward to point to the new last node (`{13, nil}`).


* **Current State:**
* `lst.head`: `{10, addr(node 13)}`
* `lst.tail`: `{13, nil}`



#### Step 3: `Push(23)` (Third Element)

* **Condition:** `lst.tail == nil` is `false`.
* **Execution:**
* `lst.tail.next = &element[T]{val: 23}` ------> Links the current tail to the new node `{23, nil}`.
* `lst.tail = lst.tail.next` ------> Shifts `lst.tail` forward to point to `{23, nil}`.



---

### Conceptual Memory Layout

Each pushed element acts as a node containing its own value and the memory address (`next` pointer) of the subsequent element:

```text
[10, addr(node 13)] ---> [13, addr(node 23)] ---> [23, nil]

```

---

## 2. The `AllElements` Method: Traversal Mechanics

The `AllElements` method traverses the linked list from head to tail, extracting each value into a standard Go slice:

```go
func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

```

### Traversal Breakdown

The pointer `e` starts from `lst.head` (`[10, addr(node 13)]`) and loops until `e == nil`:

1. **Iteration 1:**
* `e` = `[10, addr(node 13)]`
* `elems = append(elems, e.val)` -----> `elems = [10]`
* `e` advances to the next memory address: `[13, addr(node 23)]`.


2. **Iteration 2:**
* `e` = `[13, addr(node 23)]`
* `elems = append(elems, e.val)` --------> `elems = [10, 13]`
* `e` advances to the next memory address: `[23, nil]`.


3. **Iteration 3:**
* `e` = `[23, nil]`
* `elems = append(elems, e.val)` ------> `elems = [10, 13, 23]`
* `e` advances to `nil`. The condition `e != nil` evaluates to `false`, and the loop terminates.



---
