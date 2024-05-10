/*
#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;
typedef struct LNode{//结构体 LNode
    ElemType data;//数据域
    struct LNode *next;
}LNode,*LinkList;
//LNode*是结构体指针，和LinkList完全等价的
//输入3,4,5,6,7,9999
void list_head_insert(LNode *&L)
{
    L= (LinkList)malloc(sizeof(LNode));//申请头结点空间，头指针指向头结点
    L->next=NULL;
    ElemType x;
    scanf("%d",&x);
    LNode *s;//用来指向申请的新结点
    while(x!=9999)
    {
        s=(LinkList)malloc(sizeof(LNode));
        s->data=x;
        s->next=L->next;//s的next指向原本链表的第一个结点
        L->next=s;//头结点的next，指向新结点
        scanf("%d",&x);
    }
}

void print_list(LinkList L)
{
    L=L->next;
    while(L!=NULL)
    {
        printf("%3d",L->data);
        L=L->next;
    }
    printf("\n");
}

//头插法来新建链表
int main() {
    LinkList L;//L是链表头指针，是结构体指针类型
    list_head_insert(L);//输入数据可以为3 4 5 6 7 9999,头插法新建链表
    print_list(L);//链表打印
    return 0;
}


 头插法创建链表的逻辑图：
用户输入: 3 4 5 6 7 9999

初始状态: L -> NULL (L是链表的头指针，指向NULL，表示链表为空)

插入3:
        |
        V
L -> [3] -> NULL

插入4:
        |
        V
L -> [4] -> [3] -> NULL

插入5:
        |
        V
L -> [5] -> [4] -> [3] -> NULL

插入6:
        |
        V
L -> [6] -> [5] -> [4] -> [3] -> NULL

插入7:
        |
        V
L -> [7] -> [6] -> [5] -> [4] -> [3] -> NULL
当用户输入9999时，插入结束。

打印链表的逻辑图：
打印顺序: 从L开始，按箭头方向遍历

L -> [7] -> [6] -> [5] -> [4] -> [3] -> NULL

打印输出: 7 6 5 4 3
 */

#include <stdio.h>
#include <stdlib.h>

typedef int ElemType;
typedef struct LNode{
    ElemType data;//数据域
    struct LNode *next;
}LNode,*LinkList;
//LNode*是结构体指针，和LinkList完全等价的
//输入3,4,5,6,7,9999
void list_head_insert(LNode* &L)
{
    L= (LinkList)malloc(sizeof(LNode));//申请头结点空间，头指针指向头结点
    L->next=NULL;
    ElemType x;
    scanf("%d",&x);
    LNode *s;//用来指向申请的新结点
    while(x!=9999)
    {
        s=(LinkList)malloc(sizeof(LNode));
        s->data=x;
        s->next=L->next;//s的next指向原本链表的第一个结点
        L->next=s;//头结点的next，指向新结点
        scanf("%d",&x);
    }
}
//尾插法新建链表
void list_tail_insert(LNode* &L)
{
    L= (LinkList)malloc(sizeof(LNode));//申请头结点空间，头指针指向头结点
    L->next=NULL;
    ElemType x;
    scanf("%d",&x);
    LNode *s,*r=L;//s是用来指向申请的新结点，r始终指向链表尾部
    while(x!=9999)
    {
        s=(LinkList)malloc(sizeof(LNode));//为新结点申请空间
        s->data=x;
        r->next=s;//新结点给尾结点的next指针
        r=s;//r要指向新的尾部
        scanf("%d",&x);
    }
    r->next=NULL;//让尾结点的next为NULL
}

void print_list(LinkList L)
{
    L=L->next;
    while(L!=NULL)
    {
        printf("%3d",L->data);
        L=L->next;
    }
    printf("\n");
}

//头插法，尾插法来新建链表
int main() {
    LinkList L;//L是链表头指针，是结构体指针类型
//    list_head_insert(L);//输入数据可以为3 4 5 6 7 9999,头插法新建链表
    list_tail_insert(L);
    print_list(L);//链表打印
    return 0;
}


