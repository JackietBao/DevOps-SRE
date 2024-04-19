
/*
#include <stdio.h>

int main() {
    int a[]={1,2,3,4,5};
    return 0;
}

 #include <stdio.h>

int main() {
    int a[5]={1,2,3,4,5};
    int j=20;
    int i=10;
    a[5]=6;//访问越界
    a[6]=7;
    printf("i=%d\n",i);//i我们并没有赋值，但是值却变化了
    return 0;
}

 #include <stdio.h>
//子函数把某一个常用功能，封装起来的作用
//数组名传递到子函数后，子函数的形参接收到是数组的起始地址，因此不能把数组的长度传递给子函数
void print(int a[],int length)
{
    int i;
    for(i=0;i<length;i++)
    {
        printf("%3d",a[i]);
    }
    a[3]=20;
    printf("\n");
}

//main函数就是主函数
int main() {
    int a[5]={1,2,3,4,5};
    print(a,5);//数组在传递给子函数时，它的长度传递不过去
    printf("a[3]=%d\n",a[3]);
    return 0;
}

 #include <stdio.h>
//模拟printf("%s",c)操作
void print(char d[])
{
    int i=0;
    while(d[i])//当走到结束符时，循环结束
    {
        printf("%c",d[i]);
        i++;
    }
    printf("\n");
}
//如何初始化字符数组，字符串如何输出
//输出字符串乱码时，要去查看字符数组中是否存储了结束符'\0'
int main() {
    char c[6]= "hello";//使用这种方式初始化字符数组
    char d[5]="how";
    printf("%s\n",c);//使用%s来输出一个字符串，直接把字符数组名放到printf后面位置
    print(d);
    return 0;
}
 #include <stdio.h>
//scanf读取字符串操作，会自动往字符数组中放结束符
int main() {
    char c[10];
    char d[10];
//    scanf("%s",c);//字符数组名c中存储了数组的起始地址，因此不需要取地址
//    printf("%s\n",c);
    scanf("%s%s",c,d);
    printf("c=%s,d=%s\n",c,d);
    return 0;
}

 #include <stdio.h>

int main() {
    char c[20];
    gets(c);//gets中放入我们字符数组的数组名即可
    puts(c);//puts等价于printf("%s\n",c); puts内放的参数是字符数组名
    return 0;
}

#include <stdio.h>
#include <string.h>

int mystrlen(char c[])
{
    int i=0;
    while(c[i])//找到结束符后，循环结束，从而得出字符串长度
    {
        i++;
    }
    return i;
}

int main() {
    int len;
    char c[20];
    char d[100]="world";
    char e[100];
    gets(c);
    puts(c);
    len=strlen(c);//统计字符串的长度
    printf("len=%d\n",len);
    len= mystrlen(c);
    printf("my len=%d\n",len);
    strcat(c,d);//把d中的字符串拼接到c中
    puts(c);
    strcpy(e,c);//把c中的字符串复制到e中
    puts(e);
    //c大于“how"，返回值是正值，相等是0，c小于”how",返回负值
    printf("c?d=%d\n",strcmp(c,"how"));
    return 0;
}

#include <stdio.h>

int main() {
    int a;
    int arr[100];
    scanf("%d",&a);
    int i;
    for(i=0;i<a;i++)
    {
        scanf("%d",&arr[i]);
    }
    int count=0;
    for(i=0;i<a;i++)
    {
        if(arr[i]==2)
        {
            count++;
        }
    }
    printf("%d\n",count);
    return 0;
}
 */


#include <stdio.h>
#include <string.h>
int main() {
    char a[100];
    char b[100]={0};
    gets(a);
    int i,j;
    for(i=0,j=strlen(a)-1;i< strlen(a);i++,j--)
    {
        b[j]=a[i];
    }
//    puts(b);
    int result = strcmp(a,b);
    if(result < 0)
    {
        printf("%d\n",-1);
    } else  if(result > 0)
    {
        printf("%d\n",1);
    } else{
        printf("%d\n",0);
    }
    return 0;
}




















