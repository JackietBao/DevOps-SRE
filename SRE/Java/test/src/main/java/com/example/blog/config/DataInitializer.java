package com.example.blog.config;

import com.example.blog.model.Post;
import com.example.blog.model.User;
import com.example.blog.service.PostService;
import com.example.blog.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import java.time.LocalDateTime;

@Component
public class DataInitializer implements CommandLineRunner {

    private final UserService userService;
    private final PostService postService;

    @Autowired
    public DataInitializer(UserService userService, PostService postService) {
        this.userService = userService;
        this.postService = postService;
    }

    @Override
    public void run(String... args) {
        // 创建示例用户
        User admin = new User();
        admin.setUsername("admin");
        admin.setPassword("password");
        admin.setEmail("admin@example.com");
        admin.setBio("管理员账号");
        admin = userService.registerNewUser(admin);

        User user = new User();
        user.setUsername("user");
        user.setPassword("password");
        user.setEmail("user@example.com");
        user.setBio("普通用户账号");
        user = userService.registerNewUser(user);

        // 创建示例文章
        Post post1 = new Post();
        post1.setTitle("欢迎使用博客系统");
        post1.setContent("这是一个简单的Java博客系统，使用Spring Boot框架开发。\n\n您可以在这里发布、编辑和删除文章。");
        post1.setAuthor(admin);
        post1.setCreatedAt(LocalDateTime.now());
        postService.createPost(post1);

        Post post2 = new Post();
        post2.setTitle("Spring Boot简介");
        post2.setContent("Spring Boot是由Pivotal团队提供的全新框架，其设计目的是用来简化Spring应用的初始搭建以及开发过程。\n\n该框架使用了特定的方式来进行配置，从而使开发人员不再需要定义样板化的配置。");
        post2.setAuthor(admin);
        post2.setCreatedAt(LocalDateTime.now().minusDays(1));
        postService.createPost(post2);

        Post post3 = new Post();
        post3.setTitle("Java编程技巧");
        post3.setContent("Java是一种广泛使用的计算机编程语言，拥有跨平台、面向对象、泛型编程的特性，广泛应用于企业级Web应用开发和移动应用开发。\n\n本文将分享一些Java编程的实用技巧。");
        post3.setAuthor(user);
        post3.setCreatedAt(LocalDateTime.now().minusDays(2));
        postService.createPost(post3);
    }
} 