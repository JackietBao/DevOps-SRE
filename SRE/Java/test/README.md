# Java博客系统

这是一个使用Spring Boot框架开发的简单博客系统，适合Java初学者学习和实践。

## 功能特点

- 用户注册和登录
- 发布、编辑和删除博客文章
- 查看所有文章和文章详情
- 响应式界面设计

## 技术栈

- Spring Boot 2.7.14
- Spring Data JPA
- Spring Security
- Thymeleaf模板引擎
- H2内存数据库
- Bootstrap 5

## 快速开始

### 环境要求

- JDK 1.8+
- Maven 3.6+

### 运行步骤

1. 克隆项目到本地
   ```
   git clone [项目GitLab地址]
   cd blog
   ```

2. 使用Maven构建项目
   ```
   mvn clean package
   ```

3. 运行应用
   ```
   java -jar target/blog-0.0.1-SNAPSHOT.jar
   ```

4. 访问应用
   打开浏览器，访问 http://localhost:8080

### 默认账号

应用启动时会自动创建两个测试账号：

- 管理员账号：
  - 用户名：admin
  - 密码：password

- 普通用户账号：
  - 用户名：user
  - 密码：password

## 项目结构

```
src/main/java/com/example/blog/
├── BlogApplication.java        # 应用程序入口
├── config/                     # 配置类
│   ├── SecurityConfig.java     # Spring Security配置
│   └── DataInitializer.java    # 数据初始化
├── controller/                 # 控制器
│   ├── HomeController.java     # 首页控制器
│   ├── UserController.java     # 用户控制器
│   └── PostController.java     # 文章控制器
├── model/                      # 实体类
│   ├── User.java               # 用户实体
│   └── Post.java               # 文章实体
├── repository/                 # 数据访问层
│   ├── UserRepository.java     # 用户仓库
│   └── PostRepository.java     # 文章仓库
└── service/                    # 服务层
    ├── UserService.java        # 用户服务接口
    ├── UserServiceImpl.java    # 用户服务实现
    ├── PostService.java        # 文章服务接口
    ├── PostServiceImpl.java    # 文章服务实现
    └── CustomUserDetailsService.java # 自定义用户详情服务
```

## 发布到GitLab

1. 在GitLab上创建一个新的项目

2. 将本地项目推送到GitLab
   ```
   git init
   git add .
   git commit -m "初始提交"
   git remote add origin [GitLab项目URL]
   git push -u origin master
   ```

## 学习资源

- [Spring Boot官方文档](https://spring.io/projects/spring-boot)
- [Spring Data JPA文档](https://spring.io/projects/spring-data-jpa)
- [Thymeleaf官方文档](https://www.thymeleaf.org/documentation.html)
- [Bootstrap文档](https://getbootstrap.com/docs/5.1/getting-started/introduction/) 