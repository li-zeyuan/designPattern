# 外观模式
### 定义
- 提供一个统一的接口供客户访问，然后统一接口再调用子系统完成各种任务
### 成员
- client：通过一个外观对象访问系统
- facade: 为多个子系统对外提供一个统一接口
- sub_system: 实现系统的部分功能，客户可以通过外观对象调用