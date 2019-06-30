# Reverse Integer

Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

```
Input: 123
Output: 321
```

Example 2:

```
Input: -123
Output: -321
```

Example 3:

```
Input: 120
Output: 21
```

Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

## Solution 1

Convert integer to string and given two pointers in the middle of conveted string.

```
1 2 3 4 5
    r
    l
```

Starting loop until pointer `r` reach the end of the string or pointer `l` reach the start of the string. While looping, compose the characters to a new string.

```
1 2 3 4 5
    r
    l

result: 3
```

```
1 2 3 4 5
      r
  l

result: 4 3 2
```

If argument is a negative number. Add variable to catch the negative state, make number to positive number then convert to string.

```
input: -4321
base: -1

converted string:
4 3 2 1
    r
  l

result: 2 3
```

```
4 3 2 1
      r
l

result: 1 2 3 4
```

```
result * base = -1234
```

## Solution 2

We want to repeatedly "pop" the last digit off of `x` and "push" it to the back of the `rev`. In the end, `rev` will be the reverse of the `x`.

To "pop" and "push" digits without the help of some auxiliary stack/array, we can use math.

```
// pop operation:
pop = x % 10;
x /= 10;

// push operation:
temp = rev * 10 + pop;
rev = temp;
```
