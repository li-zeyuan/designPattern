# 享元模式
### 定义
- 在系统中共享细粒度的对象，节省内存空间
### 成员
- FlyWeight: 享元类，
- ConcreteFlyWeight：具体享元类，共享了享元类
- FlyweightFactory：享元工厂，可以用map存多个对应的具体享元类