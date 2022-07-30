# 代码片段


每隔多少秒处理一次

```
const (
    nodeListTimeout        = 2 * time.Minute
)

for start := time.Now(); time.Since(start) < nodeListTimeout; time.Sleep(2 * time.Second) {
    // xx
}

```

```
    pollingPeriod := 2 * time.Second
    timeout := 10 * time.Second

    startTime := time.Now()
    for startTime.Add(timeout).After(time.Now()) {
        time.Sleep(pollingPeriod)
        // xx
    }
```
tag: ['test','time']