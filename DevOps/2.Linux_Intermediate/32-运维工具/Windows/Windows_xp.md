## 1、准备工作

废话不多说，首先我们在VMware里面安装Windows XP需要做的准备工作。

①VMware Workstation软件

（演示版本：vmware-workstation-full-16.2.0-18760230.exe）

```text
【下载链接】http://ai95.microsoft-cloud.cn/d/9289114-49833935-3c06d9?p=ai95
（访问密码：ai95）

## 支持OneDrive网盘、城通网盘、阿里网盘、百度网盘等下载
```

②Windows XP系统原版镜像文件包

（演示版本：zh-hans_windows_xp_professional_with_service_pack_3_x86_cd_x14-80404.iso）

```text
【下载链接】http://ai95.microsoft-cloud.cn/d/9289114-49837613-071042?p=ai95
（访问密码：ai95）

## 支持OneDrive网盘、城通网盘、阿里网盘、百度网盘等下载
```

③序列号

```text
【产品密钥】http://www.xitongcheng.com/jiaocheng/xtazjc_article_33955.html
```

## 2、设置虚拟机

我个人喜欢使用”自定义“设置。

选择【自定义】

![img](assets/Windows_xp/v2-1c412f367ea4a4decb2df953e7c0e82f_1440w.webp)

【2-1】选择向导类型

选择最高的【硬件兼容性】

![img](assets/Windows_xp/v2-935ae38b5c038a20b459f20e73eb2fd7_1440w.webp)

【2-2】配置硬件兼容性

这里我们先选【稍后安装操作系统】，待会儿还有一些地方需要设置

![img](assets/Windows_xp/v2-ef91c598643c8e07b2fcb341334b46a3_1440w.webp)

【2-3】选择安装的操作系统

客户机操作系统选择【Microsoft Windows】，版本我们选【Windows XP Professional】

![img](assets/Windows_xp/v2-37ba6a28b0ff1a63176d6e79f8c91d81_1440w.webp)

【2-4】选择操作系统及版本

这里我们可以”重命名虚拟机“，设置好【虚拟机名称】和【安装位置】

![img](assets/Windows_xp/v2-3d6b9587f774de48814a0bd8e406d0b5_1440w.webp)

【2-5】设置“虚拟机名称”和“安装位置”

处理器配置可以【按需配置】

![img](assets/Windows_xp/v2-4eb3f5343e34ef6525a7616b40f1ef98_1440w.webp)

【2-6】配置处理器

内存配置可以【按需配置】

![img](assets/Windows_xp/v2-b11a3e5bc9df59c64ae793671b1b6b5f_1440w.webp)

【2-7】配置内存

网络类型选择【使用桥接网络】

![img](assets/Windows_xp/v2-7bbe439a7096ab5968fc9fb01a63f6cd_1440w.webp)

【2-8】配置网络类型

控制器类型选择【默认推荐】的就好

![img](assets/Windows_xp/v2-68eb27e1a1108ee6fd79939503b6f375_1440w.webp)

【2-9】选择I/O控制器类型

虚拟磁盘类型也选择【默认推荐】的

![img](assets/Windows_xp/v2-89b34ae250fc526e6bba3fb869f8f933_1440w.webp)

【2-10】选择磁盘类型

选择【创建新虚拟机硬盘】

![img](assets/Windows_xp/v2-d52fccdf8faa2aa1e9c96c869b9ef495_1440w.webp)

【2-11】创建新虚拟机硬盘

【按需配置】磁盘大小

![img](assets/Windows_xp/v2-e26902be1540d1d2e6c167b9d2b76397_1440w.webp)

【2-12】配置磁盘大小

确认无误后，点击”下一步“

![img](assets/Windows_xp/v2-8a98e9dd1c0f7e2cba5d0bb94ffa2d36_1440w.webp)

【2-13】确认磁盘文件位置

点击【完成】

![img](assets/Windows_xp/v2-c0a89082fd87dd026e337df9165b4417_1440w.webp)

【2-14】虚拟机创建完成

点击【CD/DVD（IDE）】或【编辑虚拟机设置】

![img](assets/Windows_xp/v2-4bf2f02754befd85abdfe99f76184f66_1440w.webp)

【2-15】编辑虚拟机设置

插入【ISO镜像文件】，选择我们之前下载好的【WindowsXP的镜像】

![img](assets/Windows_xp/v2-d151a8ca63378db909a43e4552e95a29_1440w.webp)

【2-16】添加IOS镜像文件

到这一步，虚拟机的设置就完成了，接下来我们来安装WindowsXP

## 3、安装WindowsXP

成功加载WindowsXP镜像文件

![img](assets/Windows_xp/v2-44be39bc87cb0e19eec9b77cdbef87c8_1440w.webp)

【3-1】镜像加载中

按【Enter】（回车键）进入XP安装程序

![img](assets/Windows_xp/v2-9eb36e22ba00a269e2a88a9229961e09_1440w.webp)

【3-2】安装程序选择

按【F8】同意许可条款

![img](assets/Windows_xp/v2-b69e4c5989168dfdbed68852affb8aee_1440w.webp)

【3-3】协议条款

按【Enter】（回车键）确认系统分区

![img](assets/Windows_xp/v2-4e3dc16bbb7d005d4666044ab657353b_1440w.webp)

【3-4】安装位置

选择【用NTFS文件系统格式化磁盘分区（快）】按【Enter】（回车键）确认

![img](assets/Windows_xp/v2-f4c28ca22f66b0314fffb96fd254ad9c_1440w.webp)

【3-5】新建分区

正在自动格式化

![img](assets/Windows_xp/v2-f6334553b690eb8d8ade767bb18c61ab_1440w.webp)

【3-6】格式化执行中

XP系统安装中

![img](assets/Windows_xp/v2-4608f463863253226781a2b7d20e7020_1440w.webp)

【3-7】系统安装①

![img](assets/Windows_xp/v2-6db8c2983ccb59de24e4470f0cc2bc25_1440w.webp)

【3-8】系统安装②

## 4、配置WindowsXP

选择区域和语言，确认无误后点击【下一步】

![img](assets/Windows_xp/v2-ede40fe7a66350e2be4e33535b7eb15b_1440w.webp)

【4-1】区域语言

填写用户名并确认

![img](assets/Windows_xp/v2-ddaffbba192006a613f45c7306c4155d_1440w.webp)

【4-2】填写用户名

填写产品密钥，网上有一大堆，大家都可以搜索得到

> 【产品密钥】[http://www.xitongcheng.com/jiaocheng/xtazjc_article_33955.html](https://link.zhihu.com/?target=http%3A//www.xitongcheng.com/jiaocheng/xtazjc_article_33955.html)

![img](assets/Windows_xp/v2-a2b10e5e6474a98de2f7d5ff02bd454b_1440w.webp)

【4-3】填写产品密钥

为计算机命名和设置密码（若不想设置密码，直接点击下一页，这样登录就不需要密码了）

![img](assets/Windows_xp/v2-7d8cce36b7734ed8a48e2344340929bb_1440w.webp)

【4-4】设置计算机名和密码

时区和时间设置

![img](assets/Windows_xp/v2-a14d37f1764bd4f688d4e51b7dce7f44_1440w.webp)

【4-5】时间和时区

网络设置选择【经典设置】

![img](assets/Windows_xp/v2-816261d9bdacd31e7b9f4dea940af7b0_1440w.webp)

【4-6】网络设置

设置工作组和计算机域，选择默认即可，直接下一步

![img](assets/Windows_xp/v2-a4ccc41391472dbdb7fcb319e1e70914_1440w.webp)

【4-7】组和域

等待组件自动配置

![img](assets/Windows_xp/v2-8be742b003ba29e07094d235e6ca328c_1440w.webp)

【4-8】组件配置

全部配置完成后会自动重启

![img](assets/Windows_xp/v2-f78ec2770b1f5592f8f2e36315071c8a_1440w.webp)

【4-9】XP启动界面

接下来我们设置一下系统功能

![img](assets/Windows_xp/v2-cfbf3b83c0313ce9bb0b0919865c33ff_1440w.webp)

【4-10】欢迎界面

是否开启自动更新系统，我们可以选择关闭掉【现在不启用】

![img](assets/Windows_xp/v2-7b0237d84a67c1a0051385da885708fc_1440w.webp)

【4-11】系统自动更新

网络连接设置，选择【否，次计算机将直接连接到Internet】

![img](assets/Windows_xp/v2-72cdbc26377f6a8f5d3c4ee361843569_1440w.webp)

【4-12】网络连接方式

产品激活，可以先选择【否，请每隔几天提醒我】

![img](assets/Windows_xp/v2-c0f0e70f743ddcb6440cb3c05cd94914_1440w.webp)

【4-13】激活Windows

创建用户名

![img](assets/Windows_xp/v2-47adafb172cc1ea3ddce0b342b19ea19_1440w.webp)

【4-14】创建用户名

设置完成，WindowsXP安装成功

![img](assets/Windows_xp/v2-1e9c60633143d31da515f8245ca787da_1440w.webp)

【4-15】安装完成①

![img](assets/Windows_xp/v2-23aba67885990fa39c2098ec3abff8b3_1440w.webp)

【4-16】安装完成②

熟悉的蓝天白云界面

![img](assets/Windows_xp/v2-05e6e76464077a45058ca3c23214b052_1440w.webp)

【4-17】WindowsXP桌面

在这里我们就安装完了WindowsXP，若大家按照这安装教材遇到了其它问题也可以随时和我说，我及时做修改，谢谢大家。