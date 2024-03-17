"""

这是一个统计的模块
提供了有以下几种使用方法：
count(6,6)
Calc.count(6,6)



"""

name = "liuhui"

def count(a, b):
    #函数定义，接受两个参数 a 和 b，然后返回它们的乘积
    return a * b

class Calc():
    #类的定义，名为 Calc
    def __init__(self, a, b):
        #__init__ 方法用于初始化新创建的对象
        #受 self 和两个参数 a 和 b，并将它们分别赋给 self.a 和 self.b
        self.a = a
        self.b = b
    def count(self):
        #Calc 类中的另一个方法 count，它不接受其他参数（除了 self，这是类方法的标准）
        #返回 self.a 和 self.b 的乘积
        return self.a * self.b

if __name__ == "__main__":
    #检查当前脚本是否是被直接执行的主程序，而不是被导入为模块
    #如果当前脚本是主程序，则条件成立，下面的代码块将被执行
    print("我在手动执行这个程序")
    print(count(6,6))
    #调用全局函数 count，并传入参数 6 和 6。该函数将返回参数的乘积
    Calc = Calc(6,6)
    #创建了一个 Calc 类的对象，并传入参数 6 和 6
    print(Calc.count())
    #调用 Calc 对象的 count 方法，该方法返回对象属性 self.a 和 self.b 的乘积，并将结果打印出来












