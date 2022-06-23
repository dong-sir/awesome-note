# 深入PHP面向对象、模式与实践 (第 5 版)

## 第 4 章 高级特性

### trait

#### 去除重复代码

定义一个 `calulateTax()` 方法的 <b>trait</b>，然后在 ShopProduct 类和 UtilityService 类中引用了该 trait: