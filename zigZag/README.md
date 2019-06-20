# ZigZag Conversion

## Description

The string `PAYPALISHIRING` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```
And then read line by line: `PAHNAPLSIIGYIR`

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = `PAYPALISHIRING`, numRows = `3`
Output: `PAHNAPLSIIGYIR`
Example 2:

Input: s = `PAYPALISHIRING`, numRows = `4`
Output: `PINALSIGYAHRPI`
Explanation:

```
P     I     N
A   L S   I G
Y A   H R
P     I
```

## Solution 1

In previous example (numRow = `4`). We can add a map on it:

```
0: P     I     N
1: A   L S   I G
2: Y A   H R
3: P     I
```

Decomposition issues:

```
P     | I     | N
A   L | S   I | G
Y A   | H R   |
P     | I     |
====================
PAYPAL | ISHIRI | NG
```

The length of the string is 6. We learn the length is `2n - 2`. For example:

- numRow = 3, length = 2 * 3 - 2 = 4
- numRow = 4, length = 2 * 4 - 2 = 6
- numRow = 5, length = 2 * 5 - 2 = 8

And if we focus on the smaller pieces:

```
P |    
A |   L 
Y | A   
P |
=========
PAYP | AL
```

So if the index of the string is less than numbers of the row, than put the string normally. Otherwise put string into `m[numRow - index - 2]`. For example:

string = `AL`, numRow = `4`
- index = 0, m[4-0-2] = A is equal m[2] = A
- index = 1, m[4-1-2] = L is equal m[1] = L

```
0: P |
1: A |     (L)
2: Y | (A)
3: P |
```

Finally, loop map to build result

```
0: P     I     N
1: A   L S   I G
2: Y A   H R
3: P     I
=======================
PIN + ALSIG + YAHR + PI
```
