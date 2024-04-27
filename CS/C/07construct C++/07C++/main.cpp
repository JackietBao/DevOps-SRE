/*
#include <stdio.h>

//当你在子函数中要修改主函数中变量的值，就用引用，不需要修改，就不用
void modify_num(int &r)//形参中写&，要称为引用
{
    r=r+1;
}
//C++的引用的讲解
//在子函数内修改主函数的普通变量的值
int main() {
    int a=10;
    modify_num(a);
    printf("after modify_num a=%d\n",a);
    return 0;
}



#include <stdio.h>

void modify_pointer(int *&p,int *q)//引用必须和变量名紧邻
{
    p=q;
}
//子函数内修改主函数的一级指针变量
int main() {
    int *p=NULL;
    int i=10;
    int *q=&i;
    printf("%d\n",*q);
    modify_pointer(p,q);
    printf("after modify_pointer *p=%d\n",*p);
    return 0;//进程已结束，退出代码为 -1073741819 ，不为0，那么代表进程异常结束
}

 #include <stdio.h>

int main() {
    bool a= true;
    bool b= false;
    printf("a=%d,b=%d\n",a,b);
    return 0;
}

#include <stdio.h>
//课时8的作业1 练习结构体
typedef struct student{
    int num;//学号
    char name[20];//姓名
    char sex;//性别
}stu;//声明一个结构体类型
int main() {
    stu s;
    scanf("%d%s %c",&s.num,s.name,&s.sex);
    printf("%d %s %c\n",s.num,s.name,s.sex);
    return 0;
}

 */


#include <stdio.h>
#include <stdlib.h>

//当子函数要修改主函数中的p，那么就加引用。引用如何实现的，完全不需要去关心
void modify_pointer(char *&p)
{
    p=(char*)malloc(100);//申请100个字节大小的空间
    fgets(p,100,stdin);//stdin代表标准输入，fgets是安全的
}

//课时8作业2 练习C++的引用的使用
int main() {
    char *p=NULL;
    modify_pointer(p);
    puts(p);
    free(p);//申请的空间不使用后，记得free，避免扣分
    return 0;
}


























