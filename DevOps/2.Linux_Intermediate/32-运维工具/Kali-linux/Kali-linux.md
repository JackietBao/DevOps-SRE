# **二、新建虚拟机向导配置**

Note:系统也可以安装到U盘，个人觉得意义不大，此处只安装到本地磁盘，需要安装的U盘的可参考其他博文！

**1、Vmware中打开文件 —> 新建虚拟机 —> 自定义（高级） —> \*下一步\* —> 硬件兼容性根据自己情况选择 —> \*下一步\***

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70.png)

**2、安装客户操作系统、选择客户机操作系统**

先选择“稍后安装系统” 　—>　***下一步\*** —> 　客户操作系统选择**“Linux(L)”,**版本本人选择的是Debian 9.x —> ***下一步\***；

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-16905121970318.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051221197511.png)

3、命名虚拟机、处理器配置、虚拟机内存

虚拟机名称自己随便想一个，名字最后会在Vmware中展示出来，位置根据自己的磁盘定义，建议新建一个文件夹保存，接着又到了下一步；

处理器数量和单个处理器内核数量根据自己的物理主机选吧，1 1或者2 2都行，我这里选择2 2；

内存建议（建议的话还是要听一下哈，不然容易造成卡顿或死机）不要超过“最大推荐内存”，我这里选择的是4G。再下一步。
![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051223481914.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051223863217.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051224331920.png)



**4、网络类型、Ｉ／ｏ控制器类型、磁盘类型**

使用桥接网络和使用网络地址转换（NAT）都行，我这里使用NAT，***下一步***　选择I／O控制类型和选择磁盘类型中按默认推荐来就行，***下一步\***。

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051225455623.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051226178826.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051226572629.png)

 **５、选择磁盘、指定磁盘容量、指定磁盘文件**

选择创建新虚拟机磁盘　—＞　***下一步\***　—＞　最大磁盘大小建议超过推荐大小２０G，不然使用时磁盘空间不足就比较麻烦，选择将虚拟机存储位单个文件或者多个文件都行　—＞　***下一步\***　—＞　磁盘文件选择默认就行。最后***下一步\***，再点击完成。

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051234197038.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051236573841.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051237890244.png)

# 三、系统安装

**１、安装前虚拟机设置**

如图

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051239703847.png)



**２、点击开启此虚拟机**

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051241061950.png)

**３、进入安装界面**

选择默认Graphical　Install　按Enter键确认。往上滑选择中文　—＞　Continue　—＞默认选择中国　—＞　***继续\***　—＞　汉语　—＞　***继续\***　—＞　主机名自己输入一个　—＞　***继续\***　—＞　域名不用管　—＞　***继续\***

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051242538153.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051243081856.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051243673959.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051244208662.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051245486565.png)



![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051245848868.png)

**４、设置用户名和密码**

自定义输入ｋａｌｉ系统的用户名，继续两下，设置密码，等待配置

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051247808571.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051248426174.png)

**５、磁盘分区**

这里我建议选择手动吧　***继续\***　选择刚刚在Vmware中新建的硬盘　***继续\***　这是一定要选择“是”　***继续\***，后面的如果是新手或者不确定的根据图中我的选择来吧，大神任意哈。

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051250051677.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051250528380.png)



![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051250959083.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051251337886.png)



![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051251853489.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051252645992.png)



![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051253188395.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-169051253884998.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-1690512543412101.png)

６、此处有个坑

在选择并安装软件中，会多次提示安装失败，我试验过多次都是安装失败三次，这里不要跳过，重复安装。安装成功后再继续，若直接跳过软件安装最后的系统将无图形操作界面

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-1690512566465104.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-1690512571539107.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-1690512576006110.png)

这三张图片的步骤多重复几次，直到安装成功。这里需要一些时间，请耐心等待吧。

**７、将GRUB安装至硬盘**

选择　是　和　／ｄｅｖ／ｓｄａ

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-1690512589444113.png)

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-1690512594786116.png)

**８、安装完成，自动重启进入系统**

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-1690512605080119.png)

**９、输入此教程中２.４设置的用户名和密码，接下来你会看到一个极度舒适的ｋａｌｉ界面。**

![img](assets/Kali-linux/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0ZlbmdmZW5nX195,size_16,color_FFFFFF,t_70-1690512617815122.png)

安装结束，请开始你的Kali　Linux之旅吧～～，有疑问可以留言或私信我

补充：为保护系统安全，Kali　系统默认无root用户，若需要使用时可用ｓｕｄｏ命令代替，或者参考

Kali２０２０.２以root权限执行

如果只是做一般测试可以使用，若涉及系统安全性问题则不建议更改。

系统中需要使用到root用户的可以参考下面这篇文章


https://blog.csdn.net/Fengfeng__y/article/details/107390413
