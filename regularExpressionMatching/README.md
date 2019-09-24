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

We can build the solution and try to find the rule in there. For example. `s="aa", p="a."`. We can define the dp array:

```
s: "aa"
p: "a."

+-----+---+---+---+
| s\p |   | a | . |
+-----+---+---+---+
|     | T | F | F |
+-----+---+---+---+
|  a  | F |   |   |
+-----+---+---+---+
|  a  | F |   |   |
+-----+---+---+---+

- dp[0][0] is `true` because `s=""` and `p=""` are matched.
- dp[0][1] is `false` because `s=""` and `p="a"` aren’t matched.
- dp[0][2] is `false` because `s=""` and `p="a."` aren’t matched.
```

Now, try to extend the case. We can get the dp:

```
s: "aa"
p: "a."

+-----+---+---+---+
| s\p |   | a | . |
+-----+---+---+---+
|     | T | F | F |
+-----+---+---+---+
|  a  | F | T | F |
+-----+---+---+---+
|  a  | F | F | T |
+-----+---+---+---+
```

We know if `s[:i]` matched `p[:j]` or `p[j]=="."` then `s[:i-1]` matched `p[:j-1]` (See the parentheses in table).

```
s: "aa"
p: "a."

+-----+-----+-----+-----+
| s\p |     |  a  |  .  |
+-----+-----+-----+-----+
|     | (T) |  F  |  F  |
+-----+-----+-----+-----+
|  a  |  F  | (T) |  F  |
+-----+-----+-----+-----+
|  a  |  F  |  F  | (T) |
+-----+-----+-----+-----+
```

It is easy to understand. If `s="aa"` and `p="a."` are matched. Substring the last character and do match again. In this case would be:

```
s="aa"  matched  s="a"  matched  s=""
p="a."  ------>  p="a"  ------>  p=""
```

We find the first rule: If `s[:i]` and `p[:j]` matched or `p[j]` is `"."` string. `dp[i+1][j+1] = dp[i][j]`.

> Note: Because the dp contain the empty case. `dp[0][0]` means `s=""` and `p=""`. So the if `i` represents index of string. `j` represents index of pattern. The current match of dp would be `dp[i+1][j+1]`.

How about `"*"` character? We can also build the case to find the rule:

```
s="xaa"
p="xa*"

+-----+---+---+---+---+
| s\p |   | x | a | * |
+-----+---+---+---+---+
|     | T | F | F | F |
+-----+---+---+---+---+
|  x  | F | T | F | T |
+-----+---+---+---+---+
|  a  | F | F | T | T |
+-----+---+---+---+---+
|  a  | F | F | F | T |
+-----+---+---+---+---+
```

The result is `true` in `dp[1][3]`. Which means `s="x"` and `p="xa*"` are matched. So we know if `p[j]="*"`. We can just find the `p[:j-2]` and `s[:i]` are matched or not (See the parentheses in table).

```
s="xaa"
p="xa*"

+-----+---+-----+---+-----+
| s\p |   |  x  | a |  *  |
+-----+---+-----+---+-----+
|     | T |  F  | F |  F  |
+-----+---+-----+---+-----+
|  x  | F | (T) | F | (T) |
+-----+---+-----+---+-----+
|  a  | F |  F  | T |  T  |
+-----+---+-----+---+-----+
|  a  | F |  F  | F |  T  |
+-----+---+-----+---+-----+
```

What if this is not a case? If `p[:j-2]` and `s[:i]` are not matched. We should consider the string is one or more match. So we need to check the string and the pattern are equal (See the curly brackets in the table).

```
s="xaa"
p="xa*"

+-----+---+---+-----+---+
| s\p |   | x | [a] | * |
+-----+---+---+-----+---+
|     | T | F |  F  | F |
+-----+---+---+-----+---+
|  x  | F | T |  F  | T |
+-----+---+---+-----+---+
| [a] | F | F |  T  | T |
+-----+---+---+-----+---+
|  a  | F | F |  F  | T |
+-----+---+---+-----+---+
```

If `s[i]` equal `p[j]`. We can check `s[i-1]` and `p[j]` are matched. In this case. `i=1` and `j=2`. Check `s="xa"` and `p="xa*"` are matched (See the parentheses in table).

```
s="xaa"
p="xa*"

+-----+---+---+---+-----+
| s\p |   | x | a |  *  |
+-----+---+---+---+-----+
|     | T | F | F |  F  |
+-----+---+---+---+-----+
|  x  | F | T | F | (T) |
+-----+---+---+---+-----+
|  a  | F | F | T | (T) |
+-----+---+---+---+-----+
|  a  | F | F | F |  T  |
+-----+---+---+---+-----+
```

So we find the second rule. If `p[j]="*"` then check the zero match first. If the result is `false` then check one ore more match.

The third rule. If all of the above is not true. Then `s` and `p` are not matched.
