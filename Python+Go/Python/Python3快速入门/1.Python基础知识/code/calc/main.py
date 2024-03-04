choice = input("请输入编号：")
n1 = int(input("请输入第一个数字："))
n2 = int(input("请输入第二个数字："))
if choice == '1':
    print(n1+n2)
elif choice == '2':
    print(n1-n2)
elif choice == '3':
    print(n1*n2)
elif choice == '4':
    print(n1/n2)
else:
    print("输入的编号不对！")