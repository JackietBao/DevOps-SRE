/*
#include <stdio.h>

struct student{
    int num;
    char name[20];
    char sex;
    int age;
    float score;
    char addr[30];
};//结构体类型声明，注意最后一定要加分号

int main() {
    struct student s={1001,"lele",'M',20,85.4,"Shenzhen"};
    struct student sarr[3];//定义一个结构体数组变量
    int i;
    //结构体输出必须单独去访问内部的每个成员
    s.num=1003;
    printf("%d %s %c %d %f %s\n",s.num,s.name,s.sex,s.age,s.score,s.addr);
    printf("--------------------------------------\n");
    scanf("%d%s %c%d%f%s",&s.num,s.name,&s.sex,&s.age,&s.score,s.addr);
    for(i=0;i<3;i++)
    {
        scanf("%d%s %c%d%f%s",&sarr[i].num,sarr[i].name,&sarr[i].sex,&sarr[i].age,&sarr[i].score,sarr[i].addr);
    }
    for(i=0;i<3;i++)//结构体数组的输出
    {
        printf("%d %s %c %d %f %s\n",sarr[i].num,sarr[i].name,sarr[i].sex,sarr[i].age,sarr[i].score,sarr[i].addr);
    }

    return 0;
}



#include <stdio.h>

struct student_type1{
    double score;//double是一种浮点类型，8个字节，浮点分为float和double，记住有这两种即可
    short age;//short 是整型，占2个字节
};

struct student_type2{
    double score;
    int height;//如果两个小存储之和是小于最大长度8，那么它们就结合在一起
    short age;
};

struct student_type3{
    int height;
    char sex;
    short age;
};
//结构体对齐
int main() {
    struct student_type1 s1;
    struct student_type2 s2={1,2,3};
    struct student_type3 s3;
    printf("s1 size=%d\n",sizeof(s1));
    printf("s2 size=%d\n",sizeof(s2));
    printf("s3 size=%d\n",sizeof(s3));
    return 0;
}


#include <stdio.h>


struct student{
    int num;
    char name[20];
    char sex;
};
//结构体指针的练习
int main() {
    struct student s={1001,"wangle",'M'};
    struct student sarr[3]={1001,"lilei",'M',1005,"zhangsan",'M',1007,"lili",'F'};
    struct student *p;//定义了一个结构体指针变量
    p=&s;
    printf("%d %s %c\n",(*p).num,(*p).name,(*p).sex);//方式1访问通过结构体指针去访问成员
    printf("%d %s %c\n",p->num,p->name,p->sex);//方式2访问通过结构体指针去访问成员，用这种
    p=sarr;
    printf("%d %s %c\n",(*p).num,(*p).name,(*p).sex);//方式1访问通过结构体指针去访问成员
    printf("%d %s %c\n",p->num,p->name,p->sex);//方式2访问通过结构体指针去访问成员，用这种
    printf("------------------------------\n");
    p=p+1;
    printf("%d %s %c\n",(*p).num,(*p).name,(*p).sex);//方式1访问通过结构体指针去访问成员
    printf("%d %s %c\n",p->num,p->name,p->sex);//方式2访问通过结构体指针去访问成员，用这种
    return 0;
}

#include <stdio.h>

//stu 等价于 struct student，pstu等价于struct student*
typedef struct student{
    int num;
    char name[20];
    char sex;
}stu,*pstu;

typedef int INGETER;//特定地方使用
//typedef的使用，typedef起别名
int main() {
    stu s={1001,"wangle",'M'};
    stu *p=&s;//定义了一个结构体指针变量
    pstu p1=&s;//定义了一个结构体指针变量
    INGETER num=10;
    printf("num=%d,p->num=%d\n",num,p->num);
    printf("num=%d,p1->num=%d\n",num,p1->num);
    return 0;
}
 */






















