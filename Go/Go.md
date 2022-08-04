

日志库

### go.uber.org/zap
- etcd (client/pkg/logutil/zap.go:25)
- kubernetes (staging/src/k8s.io/component-base/logs/json/json.go:41)


## 代码片段


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




// 多久之后 const.go
time.Now().Add(duration365d).UTC()



### 生成文件夹

```
import (
    "github.com/pkg/errors"
)

dirName := "dirName"
timestampDirName := fmt.Sprintf("%s-%s", dirName, time.Now().Format("2006-01-02-15-04-05"))

tempDir := "tempDir"
timestampDir := path.Join(tempDir, timestampDirName)
if err := os.Mkdir(timestampDir, 0700); err != nil {
    return "", errors.Wrap(err, "could not create timestamp directory")
}
```