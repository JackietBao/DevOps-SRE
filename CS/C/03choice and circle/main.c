

/*
#include <stdio.h>

//练习if与else
int main() {
    int i;
    while(scanf("%d",&i))
    {
        if(i>0)//if下面加一个大括号
        {
            printf("i is bigger than 0\n");
        }else{
            printf("i is not bigger than 0\n");
        }
    }

    return 0;
}



#include <stdio.h>
//计算从1到100奇数的和

int main() {
    int i=1,total=0;
    while(i<=100)//在这里加分号会造成死循环
    {
        if(i%2==0)
        {
            i++;
            continue;//continue下面的代码均不会得到执行
        }
        total=total+i;//把i加到total上
        i++;//i++等价于 i+=1; 在循环体内没有让while判断表达式趋近于假的操作，死循环
    }
    printf("total=%d\n",total);
    return 0;
}

#include <stdio.h>
//for 循环实现从1加到100
int main() {
    int i,total;
    for(i=1,total=0;i<=100;i++)//for小括号后不要加分号
    {
        total+=i;
    }
    printf("total=%d\n",total);
    return 0;
}

 #include <stdio.h>
//for 循环实现从1加到100
//使用continue
int main() {
    int i,total;
    for(i=1,total=0;i<=100;i++)//for小括号后不要加分号
    {
        if(i%2==0)
        {
            continue;//continue下面的代码均不会得到执行
        }
        total+=i;
    }
    printf("total=%d\n",total);
    return 0;
}


#include <stdio.h>
//for 循环实现从1加到100
int main() {
    int i,total;
    for(i=1,total=0;i<=100;i++)//for小括号后不要加分号
    {
        if(total>2000)
        {
            break;//当和大于2000时，循环结束
        }
        total+=i;
    }
    printf("total=%d,i=%d\n",total,i);
    return 0;
}
 */


#include <stdio.h>

int main() {
    int a=12321;
    int b=0,backup_a;
    backup_a=a;
    while(a)
    {
        printf("%d\n",a%10);//1232    2
        a=a/10;
    }
    //判断b和backup_a相等，就是对称数，如果不等，就不是对称数
    return 0;
}


































































