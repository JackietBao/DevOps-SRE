# Redis

# Memcached&Redis构建缓存数据库



## **前言**



许多Web应用都将数据保存到关系型数据库( RDBMS)中，应用服务器从中读取数据并在浏览器中显示。但随着数据量的增大、访问的集中，就会出现RDBMS的负担加重、数据库响应恶化、 网站显示延迟等重大影响。**Memcached/redis是高性能的分布式内存缓存服务器,通过缓存数据库查询结果，减少对关系型数据库访问次数，以提高动态Web等应用的速度、 提高可靠性。**



RDBMS即关系数据库管理系统(Relational Database Management System)



#### **1、简介**