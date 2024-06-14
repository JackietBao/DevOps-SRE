# linux搭建ftp

Yum -y install vsftpd

![img](assets/linux搭建ftp/1693980321457-427592b6-e9c3-4eaf-a42d-980035aefffc.png)

创建ftp目录,创建ftp的管理用户,设置密码

![img](assets/linux搭建ftp/1693980321704-810f5af0-c04b-4234-a79a-f6d0246adaaf.png)

![img](assets/linux搭建ftp/1693980321948-77b95360-ed63-41f7-9ab3-e30a99434992.png)

把ftp的目录设置为拥有者: 给该目录755权限

![img](assets/linux搭建ftp/1693980322217-c87e4255-1c1f-4e0e-b43e-8de7e3fc01c7.png)

![img](assets/linux搭建ftp/1693980322485-bd482ac6-9674-4b5b-8164-cf041dd9e409.png)

配置vsftpd配置文件

![img](assets/linux搭建ftp/1693980322739-47f1260e-6c07-4bdc-9a02-35918c194fb3.png)

![img](assets/linux搭建ftp/1693980323007-6f40ff01-3899-4e37-bcbb-f2577ef8fe4c.png)



关闭防火墙,文件防火墙!启动vsftpd 然后进行测试

![img](assets/linux搭建ftp/1693980323260-c5f4a84c-fc0e-4831-8279-40a3917b281c.png)

![img](assets/linux搭建ftp/1693980323472-cc249178-b669-40d6-baf6-21131ffa1ad3.png)

linux版本的ftp；有管理员用户可下载上传删除，匿名用户只能下载