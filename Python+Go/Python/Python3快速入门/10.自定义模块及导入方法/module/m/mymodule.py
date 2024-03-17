
'''
这是一个统计的模块
提供了有以下几种使用方法：
count(6,6)
Calc.count(6,6)
'''

name = "aliang"

def count(a, b):
    return  a * b

class Calc():
    def __init__(self, a, b):
        self.a = a
        self.b = b
    def count(self):
        return self.a * self.b

if __name__ == "__main__":
    print("我在手动执行这个程序")
    print(count(6,6))
    calc = Calc(6,6)
    print(calc.count())

