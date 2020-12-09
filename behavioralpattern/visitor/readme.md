# 访问者模式
### 定义
- 一个集合中的元素，每个元素都有多中访问方式
### 成员
- 抽象访问接口：定义一个visit()接口，参数为抽象元素接口。
- 具体访问类：实现了抽象访问接口所定义的接口，根据角色的类型提供不同的访问操作
- 抽象元素接口：定义一个accept()接口，参数为抽象访问接口
- 具体元素类：实现了抽象元素接口，一般是visitor.visit(this) 
- 对象结构类：用有一个容器，用来装元素