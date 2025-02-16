```code
for _, num := range nums {
  curMax := max(curMax+num, num)
  res := max(res, curMax)
}
```
