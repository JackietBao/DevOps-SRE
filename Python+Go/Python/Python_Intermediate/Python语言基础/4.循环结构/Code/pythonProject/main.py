row = int(input('请输入行数: '))
for i in range(row):
    for _ in range(i + 1):
        #通常使用 _ 作为一个占位符或者哑变量（dummy variable）。当我们不关心循环中的索引或者循环的某些结果时，可以使用 _ 来表示这个值不需要被使用，只是为了保持语法的正确性。
        print('*', end='')
    print()


for i in range(row):
    for j in range(row):
        if j < row - i - 1:
            print(' ', end='')
        else:
            print('*', end='')
    print()

for i in range(row):
    for _ in range(row - i - 1):
        print(' ', end='')
    for _ in range(2 * i + 1):
        print('*', end='')
    print()
