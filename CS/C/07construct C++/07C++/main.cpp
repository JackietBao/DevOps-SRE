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
 */





#include <stdio.h>

int main() {
    bool a= true;
    bool b= false;
    printf("a=%d,b=%d\n",a,b);
    return 0;
}






















