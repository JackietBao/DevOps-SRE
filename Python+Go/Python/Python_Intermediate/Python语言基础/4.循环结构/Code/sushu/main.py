from math import sqrt
#这一行从 Python 的 math 模块中导入 sqrt 函数，用于计算平方根。
num = int(input('请输入一个正整数: '))
#提示用户输入一个正整数，并将输入的字符串转换为整数类型，并将其赋值给变量 num
end = int(sqrt(num))
#计算出 num 的平方根，并将结果转换为整数类型，并将其赋值给变量 end。这个值将用于优化素数判断的范围
is_prime = True
for x in range(2, end + 1):
    #从2开始到 end（包括 end），用于检查 num 是否可以被除以这些数整除
    if num % x == 0:
        #判断 num 是否可以被 x 整除，即判断 num 是否有除了1和它本身以外的其他因子
        is_prime = False
        break
if is_prime and num != 1:
    print('%d是素数' % num)
else:
    print('%d不是素数' % num)
