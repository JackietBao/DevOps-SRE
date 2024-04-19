/*
#include <stdio.h>

int main() {
    int i=5;
    //定义了一个指针变量，i_pointer就是指针变量名
    //指针变量的初始化是某个变量取地址来赋值，不能随机写个数
    int *i_pointer;
    i_pointer=&i;
    printf("i=%d\n",i);//直接访问
    printf("*i_pointer=%d\n",*i_pointer);//间接访问
    return 0;
}


#include <stdio.h>
//在子函数内去改变主函数中某个变量的值
void change(int k)//j是形参
{
    k=5;
}

int main() {
    int i=10;
    printf("before change i=%d\n",i);
    change(i);//C语言的函数调用是值传递，实参赋值给形参，j=i
    printf("after change i=%d\n",i);
    return 0;
}

 #include <stdio.h>
//在子函数内去改变主函数中某个变量的值
void change(int *j)//j是形参
{
    *j=5;//*j等价于变量i，只是间接访问
}

int main() {
    int i=10;
    printf("before change i=%d\n",i);
    change(&i);//C语言的函数调用是值传递，实参赋值给形参，j=&i
    printf("after change i=%d\n",i);
    return 0;
}


 #include <stdio.h>
//指针的偏移使用场景，也就是对指针进行加和减

#define N 5
int main() {
    int a[N]={1,3,2,4,5};//数组名内存储了数组的起始地址，a中存储的就是一个地址值
    int *p;//定义指针变量p
    p=a;
    int i;
    for(i=0;i<N;i++)
    {
        printf("%3d",*(p+i));//这里写a[i]等价的
    }
    printf("\n-----------------\n");
    p=&a[4];//指针变量p指向了数组的最后一个元素
    for(i=0;i<N;i++)
    {
        printf("%3d",*(p-i));//这里写a[i]等价的
    }
    return 0;
}


 #include <stdio.h>
//指针与一维数组的传递
//数组名作为实参传递给子函数时，是弱化为指针的
//练习传递与偏移
void change(char *d)
{
    *d='H';
    d[1]='E';//*(d+1)='E'与其等价
    *(d+2)='E';
}

int main() {
    char c[10]="hello";
    change(c);
    puts(c);
    return 0;
}

 #include <stdio.h>
#include <stdlib.h> //malloc使用的头文件
#include <string.h>
int main() {
    int size;//size代表我们要申请多大字节的空间
    char *p;//void*类型的指针不能偏移的，因此不会定义无类型指针
    scanf("%d",&size);//输入要申请的空间大小
    //malloc返回的void*代表无类型指针
    p=(char*)malloc(size);//使用 malloc() 函数动态地分配了 size 字节的内存空间，并将其地址赋给 p 指针。
    // malloc() 返回的是 void* 类型的指针，因此需要进行类型转换为 char* 类型的指针
    strcpy(p,"malloc success");
    puts(p);
    free(p);//释放申请的空间时，给的地址，必须是最初malloc返回给我们的地址
    printf("free success\n");
    return 0;
}


 #include <stdio.h>
#include <stdlib.h>
#include <string.h>
//堆和栈的差异

char* print_stack()
{
    char c[100]="I am print_stack func";
    char *p;
    p=c;
    puts(p);
    return p;
}

char* print_malloc()
{
    char *p=(char*)malloc(100);//堆空间在整个进程中一直有效，不因为函数结束，而消亡
    strcpy(p,"I am print malloc func");
    puts(p);
    return p;
}

int main() {
    char *p;
    p=print_stack();
    puts(p);
    p=print_malloc();
    puts(p);
    free(p);//只有free时，堆空间才会释放
    return 0;
}

 #include <stdio.h>

void change(int  *j)
{
    *j=*j/2;
}

int main() {
    int i;
    int *p=&i;
    scanf("%d",p);
    change(p);
    printf("%d\n",*p);
    return 0;
}
 */





#include <stdio.h>
#include <stdlib.h>

int main() {
    int i;
    scanf("%d",&i);
    char c;
    scanf("%d",&c);
    char *p;
    p=(char*) malloc(i);
    gets(p);
    puts(p);
    return 0;
}

































































