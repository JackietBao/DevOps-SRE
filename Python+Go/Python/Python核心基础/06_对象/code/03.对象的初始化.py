class Person :
    # 在类中可以定义一些特殊方法（魔术方法）
    # 特殊方法都是以__开头，__结尾的方法
    # 特殊方法不需要我们自己调用，不要尝试去调用特殊方法
    # 特殊方法将会在特殊的时刻自动调用
    # 学习特殊方法：
    #   1.特殊方法什么时候调用
    #   2.特殊方法有什么作用
    # 创建对象的流程
    # p1 = Person()的运行流程
    #   1.创建一个变量
    #   2.在内存中创建一个新对象
    #   3.__init__(self)方法执行
    #   4.将对象的id赋值给变量

    # init会在对象创建以后离开执行
    # init可以用来向新创建的对象中初始化属性
    # 调用类创建对象时，类后边的所有参数都会依次传递到init()中
    def __init__(self,name):
        # print(self)
        # 通过self向新建的对象中初始化属性
        self.name = name

    def say_hello(self):
        print('大家好，我是%s'%self.name)


# 目前来讲，对于Person类来说name是必须的，并且每一个对象中的name属性基本上都是不同
# 而我们现在是将name属性在定义为对象以后，手动添加到对象中，这种方式很容易出现错误
# 我们希望，在创建对象时，必须设置name属性，如果不设置对象将无法创建
#   并且属性的创建应该是自动完成的，而不是在创建对象以后手动完成
# p1 = Person()
# # 手动向对象添加name属性
# p1.name = '孙悟空'

# p2 = Person()
# p2.name = '猪八戒'

# p3 = Person()
# p3.name = '沙和尚'

# p3.say_hello()

p1 = Person('孙悟空')
p2 = Person('猪八戒')
p3 = Person('沙和尚')
p4 = Person('唐僧')
# p1.__init__() 不要这么做

# print(p1.name)
# print(p2.name)
# print(p3.name)
# print(p4.name)

p4.say_hello()



class Person :
    #声明了一个名为 Person 的类。在Python中，类是一种用来创建对象的模板，它定义了对象的属性和方法
    def __init__(self, name):
        #这是一个特殊的方法，称为构造函数或初始化方法。它在创建类的实例时被调用。self
        # 是一个指向对象自身的引用，name 是构造函数的参数。在这个方法内部，self.name 表示类的实例的 name 属性
        print(self)
        self.name = name
#这行代码将传递给构造函数的 name 参数赋值给当前实例的 name 属性
    def say_hello(self):
        # 方法定义，用于在类中定义行为。它接受 self 作为参数，以访问类的属性和其他方法
        print('大家好，我是%s' % self.name)

p1 = Person('孙悟空')
#创建了一个名为 p1 的 Person 类的实例，并将 '孙悟空' 作为参数传递给构造函数
print(p1.name)
#打印了 p1 实例的 name 属性，即 '孙悟空'
p1.say_hello()