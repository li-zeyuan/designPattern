# 访问者模式

### 定义
- 在不修改具体元素对象情况下，通过扩展访问者方式添加新的操作

### 成员
- 抽象元素接口：定义 Accept(Visitor) 接口，用来接受访问者；如 `Shape`
- 具体元素类：实现 Accept(Visitor) 接口，调用 Visitor 中对应自己的方法；如 `Circle`

- 抽象访问者接口：定义了操作对应【具体元素类】Visit_xx 接口；如 `Visitor`
- 具体访问者类：实现【访问者接口】；如 `AreaCalculator`

### 流程
- 1、【具体元素类】调用 Accept 方法，接受一个 Visitor
- 2、Accept 方法的实现：Visitor 调用【抽象访问者接口】中所定义对应自己的 Visit_xx 接口，并传入【具体访问者类】
- 3、最终是 1 步骤 Visitor 对应具体类 Visit_xx 方法，执行具体的操作

### 参考
- https://cloud.tencent.com/developer/article/2611194