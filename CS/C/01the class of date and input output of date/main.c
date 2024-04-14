
//符号常量

/*
#include <stdio.h>
#define PI 3+2
int main() {
    int i=PI*2;
    printf("i=%d\n",i);
    return 0;
}


 #include <stdio.h>
int main() {
    char c='A';
    printf("%c\n",c+32);
    printf("%d\n",c);
    return 0;
}


 //强制类型转换
#include <stdio.h>
int main() {
    int i=5;
    float j=i/2;
    float k=(float)i/2;
    printf("%f\n",j);
    printf("%f\n",k);
    return 0;
}


 //强制类型转换
#include <stdio.h>
int main() {
    int i=10;
    float f=96.3;
    printf("student number=%3d score=%5.2f\n",i,f);
    //%3d指定一个整数字段宽度为3
    //%5.2f指定一个浮点数字段宽度为5，保留2位小数
    printf("student number=%-3d score=%5.2f\n",i,f);
    //-表示左对齐。因此，i的值将左对齐，并且字段宽度为3
    printf("%10s\n","hello");
    //%10s指定一个字符串字段宽度为10
    return 0;
}
 */


//scanf读取标准输入
//scanf %d %f 发现里边有\n 空格，忽略
//scanf %c 不忽略内容
#include <stdio.h>
int main() {
    int i=10;
    char c;
    scanf("%d",&i);
    printf("%d\n",i);
    fflush(stdin);
    scanf("%c",&c);
    printf("c=%c\n",c);
    return 0;
}















