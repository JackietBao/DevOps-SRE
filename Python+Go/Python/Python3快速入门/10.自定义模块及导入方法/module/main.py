import sys
print(sys.path)
#印了模块搜索路径 sys.path 的内容
import os
current_dir = os.getcwd()
#os.getcwd() 函数返回一个字符串，表示当前工作目录的路径
module_dir = os.path.join(current_dir, "m")
#os.path.join() 方法将当前目录和子目录 "m" 组合成了一个新的路径 module_dir
sys.path.append(module_dir)
#将这个新的路径添加到了模块搜索路径 sys.path 中，以便 Python 解释器可以搜索到这个目录中的模块

import mymodule

print(mymodule.count(6,6))
calc = mymodule.Calc(6,6)
print(calc.count())
print(sys.version)