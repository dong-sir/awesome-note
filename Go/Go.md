# 代码片段


每隔多少秒处理一次

``` golang
const (
    nodeListTimeout        = 2 * time.Minute
)

for start := time.Now(); time.Since(start) < nodeListTimeout; time.Sleep(2 * time.Second) {
    // xx
}

```

```go
    pollingPeriod := 2 * time.Second
    timeout := 10 * time.Second

    startTime := time.Now()
    for startTime.Add(timeout).After(time.Now()) {
        time.Sleep(pollingPeriod)
        // xx
    }
```
tag: ['test','time']