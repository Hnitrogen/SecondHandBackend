## 技术栈

- 后端
  - 语言：Go
  - 框架：Kratos
  - 数据库：MySQL
  - 缓存：Redis
  - 架构：微服务
  - 服务发现：Consul
  - API网关: Nginx

- 前端
  - 语言：TypeScript
  - 框架：Vue
  - 状态管理：Pinia
  - 路由：Vue Router

- 项目结构
    - 后端
        - App 每个App的目录格式：
            - api 存放RPC接口定义
            - internal 存放业务逻辑
                - biz 存放业务逻辑
                - service 存放服务端代码
                - data 存放数据层代码
                - server 存放服务端代码
                - conf 存放配置文件
            - cmd 存放服务端代码
            - pkg 存放公共组件
            - server 存放服务端代码
            - third_party 存放第三方库

        - 服务包含 （每个服务是独立的go.mod，可以独立运行，存放在根目录）：
            - user 用户服务     Kratos 框架
            - stuff 物品服务    Kratos 框架
            - Gin-Media 图床服务 Gin 框架
            - Gin-Chat 聊天服务 Gin 框架 (未完成)
    
        - 微服务之间采用RPC通讯，通过Nginx暴露Http接口给前端访问
            

    - 前端
        - src 存放前端代码
            - api 存放Http接口定义
            - assets 存放静态资源
            - components 存放公共组件
            - router 存放路由
            - store 存放状态管理
            - utils 存放工具函数
            - views 存放页面

