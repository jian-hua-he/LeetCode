# Longest Substring Without Repeating Characters

Given a string, find the length of the **longest substring** without repeating characters.

**Example 1:**

```
Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3.
```

**Example 2:**

```
Input: "bbbbb"
Output: 1
Explanation:** The answer is "b", with the length of 1.
```

**Example 3:**

```
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## Solution 1

Iterate the string and put it to map or set to finding longest substring.

```
longest = 0
temp = {}

abcabcbb
i
j
```

```
longest = 0
temp = {"a"}

abcabcbb
i
 j
```

```
longest = 0
temp = {"a", "b"}

abcabcbb
i
  j
```

```
longest = 3
temp = {"a", "b", "c"} 

abcabcbb
i
   j
```

`a` is repeated and the length of temp is bigger than variable `longest`. Assign the length of temp to variable `longest` and move `i` to next.

```
longest = 3
temp = {}

abcabcbb
   i
   j
```

```
longest = 3
temp = {"a"}
abcabcbb
 i
 j

// And so on...
```

## Solution 2

If a substring `s[i:j]` is already checked to have no duplicate characters. We only need to check if `s[j]` is already in the substring `s[i:j]`.
​   
```
longest = 0
temp = ""

abcabcbb
i
j
```

```
longest = 1
temp = "a"

abcabcbb
i
 j
```

```
longest = 2
temp = "ab"

abcabcbb
i
  j
```

```
longest = 3
temp = "abc"

abcabcbb
i
   j
```

Repeated, slice the element of head and move `i` to next.

```
longest = 3
temp = "bc"

abcabcbb
 i
   j
```

temp is not contained character `a`. Put `a` to temp and move `j` to next.

```
longest = 3
temp = "bca"
abcabcbb
 i
    j

// And so on...
```
