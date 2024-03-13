
"""

#类
#电脑（类）：
#   特征（属性）：主机，显示器，键盘，鼠标…
host = "4c8G"
displayer = "27寸"
keyboard = "机械键盘"
mouse = "无线鼠标"

#功能（方法）：办公，上网，看片…
def office():
    办公
def internet():
    上网
def movies():
    玩游戏
"""




"""

class Computer():
    #电脑类
    #属性(特征)
    def __init__(self, name):
        self.host = "4C8G"
        self.displayer = "27寸"
        self.keyboard = "机械键盘"
        self.mouse = "无线鼠标"
        self.name = name
    #方法(功能)
    def office(self):
        print("办公")
    def internet(self):
        print(f"{self.name}在上网。。。")
    def movies(self):
        print(f"{self.name}看片")

pc1 = Computer("刘惠")
pc1.internet()

pc2 = Computer("田豹")
pc2.movies()


class Computer():
    pass

pc1 = Computer()
pc2 = Computer()

print(pc1)
print(pc2)

pc1.host = "4C8G"
pc1.displayer = "27寸"
pc2.host = "2C4G"
pc2.displayer = "24寸"

print(pc1.host)
print(pc2.host)
"""


class Calc():
    '''计算器类'''
    def __init__(self, num1, num2):
        #定义了类的初始化方法（构造函数）
        #self 是类实例的一个引用，而 num1 和 num2 是进行运算的两个数字
        self.num1 = num1
        self.num2 = num2
        #两行代码将传入的参数 num1 和 num2 分别赋值给实例的属性
    def jia(self):
        return self.num1 + self.num2
    def jian(self):
        return self.num1 - self.num2
    def cheng(self):
        return self.num1 * self.num2
    def chu(self):
        return self.num1 / self.num2

calc1 = Calc(6,6)
#创建了一个 Calc 类的实例 calc1，并将两个运算数字都设置为 6
print(calc1.cheng())
#调用 calc1 实例的 cheng 方法，打印出 6 乘以 6 的结果，即 36
calc2 = Calc(6,6)
print(calc2.chu())













