# Evaluate Nested Arithmetic Expression

## Problem

You are given an arithmetic expression represented as a nested array.

Each expression follows this structure:

```
[operator, operand1, operand2]
```

Where:

* `operator` is a string:

  * `"add"` → addition
  * `"mult"` → multiplication
  * `"sub"` → subtraction
  * `"div"` → division
* `operand1` and `operand2` can be:

  * an integer
  * another nested expression following the same format

Your task is to write a function that **evaluates the expression and returns the result as an integer**.

---

## Examples

### Example 1

Input:

```
["add", 1, 2]
```

Output:

```
3
```

---

### Example 2

Input:

```
["mult", 3, ["add", 2, 4]]
```

Output:

```
18
```

Explanation:

```
3 * (2 + 4) = 18
```

---

### Example 3

Input:

```
["add", ["mult", 2, 3], ["mult", 4, 5]]
```

Output:

```
26
```

Explanation:

```
(2 * 3) + (4 * 5) = 6 + 20 = 26
```

---

## Constraints

* Expressions are always valid
* Only the following operators will be used:

  * `"add"`
  * `"mult"`
  * `"sub"`
  * `"div"`
* Division is integer division
* Maximum nesting depth ≤ 100
* All numbers fit in 32-bit integer range


## Notes
<!-- Edge cases, thoughts -->
