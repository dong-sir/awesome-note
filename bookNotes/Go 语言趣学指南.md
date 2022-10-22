# Go 语言趣学指南

很基础的一本书，入门的时候可以看，照着敲一遍代码

## 第 6 章 实数

6.3 浮点精确性 (舍入错误)
- (float64) 0.1 + 0.2 = 0.30000000000000004
- (float32) 0.1 + 0.2 = 0.3

## 第 7 章 整数

7.1.1 整数类型
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

## 第 10 章 类型转换

10.1 类型不能混合使用
10.2 数字类型转换 (显式)

``` golang
age := 41
marsAge := float64(age)
// 41 (float64)

earthDays := 365.2425
fmt.Println(int(earthDays))
// 365
```

比较不同类型会报错

10.3 类型转换的危险之处

注意类型的取值范围

``` golang
var bh float64 = 32768
var h = int16(bh)
fmt.Println(h)
// -32768
```
超出 int16 的范围(-32768 至 32767)，出现回绕行为，并称为 int16 能够表示的最小值 32768。

10.4 字符串转换
静态类型: 变量一但被声明之后就无法改变其类型。

## 第 12 章 函数

声明一个函数，它接受单个形参并返回单个值

``` golang
package main
import fmt

// kelvinToCelsius 函数会将开氏度转换为摄氏度
func kelvinToCelsius(k float64) float64 {
    k -= 273.15
    return k
}

func main()  {
    kelvin := 294.0
    celsius := kelvinToCelsius(kelvin)
    fmt.Print(kelvin, " 开氏度 = ", celsius, " 摄氏度")
    // 294 开氏度 = 20.850000000000023 摄氏度
}
```
## 第 13 章 方法

13.1 声明新类型
不同类型之间无法混用

## 第 16 章 数组

数组的长度是固定的

16.4 迭代数组