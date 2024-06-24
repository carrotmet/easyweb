# 项目总结

## 路由设计
- 设计独立的路由
- 封装到命令行
- 使用postman测试

## 配置设计
- 设计了 .env 文件  和  config包init加载两套环境变量
- 利用viper进行配置文件匹配、配置加载、配置读取（根据数据类型不同调不同的viper方法读配置）

## 日志设计
- 利用zap重写了gin日志，切换为zap日志。
- 在需要错误处理的环境下利用 状态码+日志等级 进行 统一日志返回（gin日志、数据库日志）
- 为日志提供了登记参数方便日志归档以及日志大小管理。

## ORM设计
- 使用gorm，封装数据库连接作为main函数的启动自加载（包括mysql和redis）
- 使用hook实现密码加密
- 配置了数据库迁移并将数据库日志统一为zap自定义日志
- 配置了redis缓存管理用户验证的图片验证码与短信验证码

## 访问控制
- 使用jwt实现用户访问的统一token颁发（方便跨项目验证）
- 使用guest与auth中间件实现页面访问的鉴权
- 使用limier实现服务端页面的细粒度流量管理（对不同界面的访问流量进行限制）

## 命令行模式
- 使用cobra开发了serve --env命令
- 了解了Command作为cobra工作的基本对象。


