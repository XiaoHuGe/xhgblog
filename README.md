# xhgblog
### 简介
- 基于golang + bootstrap实现的个人博客项目
- 示例地址：[zhangbaihu.com](http://zhangbaihu.com)
### 主要第三方组件：
- Gin: 轻量级Web框架
- GORM: ORM工具。配合Mysql使用
- Gin-Session: Gin框架提供的Session操作工具
- yaml.v2: 环境变量工具
- dep：依赖包管理
### 运行
1. 修改cong/app.yaml文件下database和oauth的内容
2. 安装dep ```brew install dep```
3. 项目根路径下同步依赖并运行
```
cd xhgblog
# 同步依赖
dep ensure
go run main.go
```
### 问题
dep同步出现```i/o timeout``` 错误, 解决方案查看 [Go依赖管理工具dep](http://zhangbaihu.com/article/2)

