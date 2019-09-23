# Regular Expression Matching

Given an input string (`s`) and a pattern (`p`), implement regular expression matching with support for `'.'` and `'*'`.

```
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
```

The matching should cover the **entire** input string (not partial).

Note:

- `s` could be empty and contains only lowercase letters `a-z`.
- `p` could be empty and contains only lowercase letters `a-z`, and characters like `.` or `*`.

Example 1:

```
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

Example 2:

```
Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

Example 3:

```
Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

Example 4:

```
Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
```

Example 5:

```
Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
```

# Solution 1

We can breakdown the issue. First, matches string without `*` character. It only needs to check the character is the same as the pattern or the character is equal `.`.

```
func isMatch(s string, p string) bool {
    if p == "" {
        return s == ""
    }

    firstMatch := s != "" && (s[0] == p[0] || string(p[0]) == ".")

    return firstMatch && isMatch(s[1:], p[1:])
}
```

Then handle `*` character. The length of the pattern should bigger than 2. Also, the second character should equal `*`. If the condition is matched then split string and pattern to two-part and continue matching.

The first part is to match `*`. So the first character of the string should match the first character of the pattern. And remove the first character continues the match with `*`. The second part matches the whole without `*` character. So remove the first two character in the pattern and continue matching.

```
func isMatch(s string, p string) bool {
    if p == "" {
        return s == ""
    }

    firstMatch := s != "" && (s[0] == p[0] || string(p[0]) == ".")

    // firstMatch && isMatch(s[1:], p) continue match with *
    // isMatch(s, p[2:]) skip * pattern and match again
    if len(p) >= 2 && string(p[1]) == "*" {
        return firstMatch && isMatch(s[1:], p) || isMatch(s, p[2:])
    }

    return firstMatch && isMatch(s[1:], p[1:])
}
```

# Solution 2

s: aacba
p: c*a*cb.
result: true

If pattern is match string or pattern is "." then match the previous string and pattern.

Example:

```
s: aa
p: a.

Base case

+-----+---+---+---+
| s\p |   | a | . |
+-----+---+---+---+
|     | T | F | F |
+-----+---+---+---+
|  a  | F |   |   |
+-----+---+---+---+
|  a  | F |   |   |
+-----+---+---+---+

Expend case, s[1] == p[1] then match s[0] and p[0]. s[0] == p[0] so is matched.

+-----+-----+-----+---+
| s\p |     |  a  | . |
+-----+-----+-----+---+
|     | (T) |  F  | F |
+-----+-----+-----+---+
|  a  |  F  | (T) |   |
+-----+-----+-----+---+
|  a  |  F  |     |   |
+-----+-----+-----+---+

s[1] == p[2] or p[2] == "." then match s[0] and p[1]. s[0] != p[1] so is not matched.

+-----+---+-----+-----+
| s\p |   |  a  |  .  |
+-----+---+-----+-----+
|     | T | (F) |  F  |
+-----+---+-----+-----+
|  a  | F |  T  | (F) |
+-----+---+-----+-----+
|  a  | F |     |     |
+-----+---+-----+-----+

s[2] == p[2] then match s[1] and p[1]. s[1] == p[1] so is matched.

+-----+---+-----+-----+
| s\p |   |  a  |  .  |
+-----+---+-----+-----+
|     | T |  F  |  F  |
+-----+---+-----+-----+
|  a  | F | (T) |  F  |
+-----+---+-----+-----+
|  a  | F |  F  | (T) |
+-----+---+-----+-----+

So we know if s[j] == p[j] or p[j] == '.' then try to match s[i-1] and p[j-1].
```

If pattern is "*" then do zero match first. If zero match is not matched then match the previous string.

Example

```
s: xaa
p: xa*

Base case

+-----+---+---+---+---+
| s\p |   | x | a | * |
+-----+---+---+---+---+
|     | T | F | F | F |
+-----+---+---+---+---+
|  x  | F |   |   |   |
+-----+---+---+---+---+
|  a  | F |   |   |   |
+-----+---+---+---+---+
|  a  | F |   |   |   |
+-----+---+---+---+---+

Expend case. Current index of s is 1. if p[3] == "*" then match p[1] and s[1]. p[1] == s[1] so is matched.

+-----+---+-----+---+-----+
| s\p |   |  x  | a |  *  |
+-----+---+-----+---+-----+
|     | T |  F  | F |  F  |
+-----+---+-----+---+-----+
|  x  | F | (T) | F | (T) |
+-----+---+-----+---+-----+
|  a  | F |     |   |     |
+-----+---+-----+---+-----+
|  a  | F |     |   |     |
+-----+---+-----+---+-----+

Current index of s is 2. p[3] == "*", match p[1] and s[2]. p[1] and s[2] is not matched. Try to match p[3] and s[1]. p[3] and s[1] is matched. So, p[3] and s[2] is matched.

+-----+---+-----+---+-----+
| s\p |   |  x  | a |  *  |
+-----+---+-----+---+-----+
|     | T |  F  | F |  F  |
+-----+---+-----+---+-----+
|  x  | F |  T  | F | (T) |
+-----+---+-----+---+-----+
|  a  | F | [F] | T | (T) |
+-----+---+-----+---+-----+
|  a  | F |     |   |     |
+-----+---+-----+---+-----+

Current index of s is 3. p[3] == "*", match p[1] and s[3]. p[1] and s[3] is not matched. Try to match p[3] and s[2]. p[3] and s[2] is matched. So, p[3] and s[3] is matched.

+-----+---+-----+---+-----+
| s\p |   |  x  | a |  *  |
+-----+---+-----+---+-----+
|     | T |  F  | F |  F  |
+-----+---+-----+---+-----+
|  x  | F |  T  | F |  T  |
+-----+---+-----+---+-----+
|  a  | F |  F  | T | (T) |
+-----+---+-----+---+-----+
|  a  | F | [F] | F | (T) |
+-----+---+-----+---+-----+

We know if p[i] == "*". Match p[i-2] and s[j]. If not matched then match p[i] and s[j-1].
```
