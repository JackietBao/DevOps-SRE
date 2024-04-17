/*
#include <stdio.h>
//练习算术运算符
int main() {
    int result=4+5*2-6/3+11%4;
    printf("result=%d\n",result);
    return 0;
}


 #include <stdio.h>

//关系运算符，优先级小于算术运算符
int main() {
    int a;
    while(scanf("%d",&a))
    {
        if(3<a && a<10)//a大于3同时a小于10要这样写
        {
            printf("a is between 3 and 10\n");
        }else{
            printf("a is not between 3 and 10\n");
        }
    }
    return 0;
}

 #include <stdio.h>
//记住优先级目的，不能够加多余的括号
int main() {
    int year,i,j=6;
    while(scanf("%d",&year))
    {
        if(year%4==0&&year%100!=0 || year%400==0)
        {
            printf("%d is leap year\n",year);
        }else{
            printf("%d is not leap year\n",year);
        }
    }
    i=!!j;
    printf("i value=%d\n",i);
    return 0;
}

 #include <stdio.h>

//逻辑与和逻辑或 短路运算
int main() {
    int i=0;
    i&&printf("you can't see me !\n");//当i为假时，不会执行逻辑与后的表达式，称为短路运算
    i=1;
    i||printf("you can't see me !\n");
    return 0;
}

 #include <stdio.h>
//赋值的练习
int main() {
    int a=1,b=2;
    a+=3;
    b*=5;
    printf("a=%d\n",a);
    printf("b=%d\n",b);
    return 0;
}
 */

#include <stdio.h>
//sizeof运算符
int main() {
    int i=0;
    printf("i size is %d\n",sizeof(i));
    return 0;
}


























