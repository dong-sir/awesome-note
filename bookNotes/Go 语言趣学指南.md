# Go 语言趣学指南

很基础的一本书，推荐入门的时候看，照着敲一遍代码

## 第 6 章 实数

6.3 浮点精确性 (舍入错误)
- (float64) 0.1 + 0.2 = 0.30000000000000004
- (float32) 0.1 + 0.2 = 0.3

## 第 7 章 整数

7.1.1 整数类型 补一张图

7.1.2 了解类型

打印变量的类型

```golang
year := 2022
fmt.Printf("Type %T for %v\n", year, year)
// Type int for 2022

days := 365.2425
fmt.Printf("Type %T for %[1]v\n", days, days)
// Type float64 for 365.2425
```

7.3 整数回绕 (有限的取值范围)

``` golang
var red uint8 = 255
red ++
fmt.Println(red)
// 0

var number int8 = 127
number ++
fmt.Println(number)
// -128
```

7.3.1 聚焦二进制位
7.3.2 避免时间回绕