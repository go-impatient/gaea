# 架构目录结构

- service 服务层/应用层
- biz 业务逻辑层
- data 存储层


## 项目大体结构
1.1 internal 使用，biz、data、service 等目录，可以携带myapp应用名或者不带（比如单项目），internal/pkg 为项目里共用
1.2 cmd/myapp，cmd下需要带上app名字
1.3 api/serivice/v1，按照gRPC 服务名，以及版本号
1.4 configs，放配置文件
2、对象初始化，biz、data、service，依赖的对象必须作为参数传入，在main里使用wire构建（IoC使用google wire 实现）和消费资源
3、biz中定义了 repository 的接口，实现在 data 目录中，biz中包含了 DomainObject的定义
4、main.go 中使用wire 进行对象后，有lifecycle进行服务的注册和启动


## internal目录
- 私有程序、库代码, 只允许本项目引入和使用
- 针对每个项目都应该新建一个对应的目录
- 如果需要调用不暴露的公共函数, 可以在internal目录下添加pkg目录
### xxxservice/data目录
- 类似DDD的repo, repo接口在这里定义, 使用依赖倒置原则
- 业务数据访问层, 包括cache
- 实现了biz定义的持久化接口逻辑
- 事务暂时在这里实现
- po(persistent Object) - 持久化对象, 与data层的数据结构一一对应
### xxxservice/biz目录
- 业务逻辑层, 类似DDD的domain
- 定义了业务逻辑实体, 业务实体应该在业务逻辑层, 定义了持久化接口
- 在写业务逻辑的时候才知道数据应该如何被持久化, 持久化的interface应该被定义在业务逻辑层
### xxxservice/service目录
- 实现了api定义的服务层, 类似DDD的application层
- 实现dto -> do, 贫血模型
- IOC 控制反转、依赖注入 - 1、方便测试 2、单次初始化和复用
- [https://github.com/google/wire](https://github.com/google/wire)
- 这里只应该有编排逻辑, 不应该有业务逻辑
