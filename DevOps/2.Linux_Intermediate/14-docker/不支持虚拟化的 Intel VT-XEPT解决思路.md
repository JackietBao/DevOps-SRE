# 不支持虚拟化的 Intel VT-XEPT解决思路

![img](assets/不支持虚拟化的 Intel VT-XEPT解决思路/1680145939288-e875be86-de39-4929-a6f8-8c63fa8e61de.png)

**解决思路**

以管理员权限运行cmd

![img](assets/不支持虚拟化的 Intel VT-XEPT解决思路/1680145939650-94967977-bca5-4171-a989-68f460c5cf3b.png)

关闭 hypervisorlaunchtype

bcdedit /set hypervisorlaunchtype off

![img](assets/不支持虚拟化的 Intel VT-XEPT解决思路/1680145940101-68e89e2d-fc79-4950-b1d0-170359ef36b0.png)



![img](assets/不支持虚拟化的 Intel VT-XEPT解决思路/1680145940591-38349f0e-2d49-4e79-82ba-60d5d0cd2889.png)

![img](assets/不支持虚拟化的 Intel VT-XEPT解决思路/1680145941004-d6938471-fcc7-4f29-86c4-b626d221e61b.png)

**最后重启系统**