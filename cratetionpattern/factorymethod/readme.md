# 工厂方法
### 思想
- 用于生成无关联的产品
### 步骤
- 1、抽出公共属性方法成基类
- 2、子类（PlusOperator）通过结构体嵌套继承基类
- 3、封装子类的所有方法成interfere（Operator）
- 4、工厂类定义Create方法，返回子类接口（Operator）