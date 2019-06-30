# Add Two Numbers

## Description

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

# Solution

The solution is simple: Init a linked for result and give a current pointer for it, then add two pointers in given two linked lists. Sum the values from the value of the parameters and the pointer go to next. If the value is bigger than 10, then catch the 1 for next value and sub 10 for current value.

```
Loop1:

2 -> 4 -> 3
5 -> 6 -> 4
i
j

i + 7 = 7

Current Value: 7
Next Value: 0
Result: 7
```

```
Loop2:

2 -> 4 -> 3
5 -> 6 -> 4
     i
     j

i + j = 10

Current Value: 0 (i+j bigger than 10. Current value -10 and Next Value +1)
Next Value: 1
Result: 7 -> 0
```

```
Loop3:

2 -> 4 -> 3
5 -> 6 -> 4
          i
          j

i + j = 7 + 1 (In privious next value is 1) = 8

Current Value: 8
Next Value: 0
Result: 7 -> 0 -> 8
```