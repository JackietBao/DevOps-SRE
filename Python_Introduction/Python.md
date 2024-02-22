## 一、安装与运行

各个系统的 Python 安装教程请自行查阅资料，这里不再赘述。

检查 Python 版本，在命令行输入 `python` 即可，同时会进入命令行交互模式，可以在这里执行 python 命令。

如果电脑中安装了 python2.x 和 python3.x 两个版本，输入 `python` 运行的是 2.x 版本。想运行 3.x，则需输入 `python3`。

在命令行输入 `python` ：

```bash
复制代码Solo-mac:~ solo$ python
Python 2.7.10 (default, Aug 17 2018, 19:45:58)
[GCC 4.2.1 Compatible Apple LLVM 10.0.0 (clang-1000.0.42)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>>
```

在命令行输入 `python3` ：

```bash
复制代码Solo-mac:~ solo$ python3
Python 3.7.0 (v3.7.0:1bf9cc5093, Jun 26 2018, 23:26:24)
[Clang 6.0 (clang-600.0.57)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>>
```

输入 `exit()` 即可退出命令行模式。

### 命令行运行 python 文件

如果是写好了一个 python 文件，想通过命令行运行它，进入这个目录，在命令行输入 `python 文件名.py` 即可。

比如桌面上有个文件 `hello.py`，内容是打印一句话：

```bash
复制代码print("Hello, Python")
```

想运行它，先进入 Desktop 目录，再在命令行输入 `python hello.py` 就能运行：

```bash
复制代码Solo-mac:Desktop solo$ python hello.py
Hello, Python
```

## 二、变量和简单数据类型

### 2.1 变量命名规则

- 变量名只能包含字母、数字和下划线。变量名可以字母或下划线打头，但不能以数字打 头，例如，可将变量命名为message_1，但不能将其命名为1_message。
- 变量名不能包含空格，但可使用下划线来分隔其中的单词。例如，变量名greeting_message 可行，但变量名greeting message会引发错误。
- 不要将Python关键字和函数名用作变量名，即不要使用Python保留用于特殊用途的单词， 如print。
- 变量名应既简短又具有描述性。例如，name比n好，student_name比s_n好，name_length比length_of_persons_name好。
- 慎用小写字母l和大写字母O，因为它们可能被人错看成数字1和0。

变量名应该是小写的，虽然没有强制规定，但是约定俗称的规则。

### 2.2 字符串

字符串就是一系列字符。在Python中，用引号括起的都是字符串，其中的引号可以是单引号，也可以是双引号，还可以同时使用。如：

```python
复制代码"This is a string." 
'This is also a string.'
"I love 'python'"
```

#### 2.2.1 字符串的简单运算

下面介绍字符串的简单运算。

##### title()

title()以首字母大写的方式显示每个单词，即将每个单词的首字母都改为大写。

```bash
复制代码>>> name = 'solo coder'
>>> name.title()
'Solo Coder'
```

##### upper()、lower()

将字符串改为全部大写或全部小写。

```bash
复制代码>>> name='solo coder'
>>> name.upper()
'SOLO CODER'
>>> name.lower()
'solo coder'
```

注意：title()、upper()、lower() 均不改变原字符串，只是输出了一个新的字符串。

#### 2.2.2 合并（拼接）字符串

Python使用加号(+)来合并字符串。

```bash
复制代码>>> first = 'solo'
>>> last = 'coder'
>>> full = first + ' ' + last
>>> full
'solo coder'
```

#### 2.2.3 使用制表符或换行符来添加空白

在编程中，空白泛指任何非打印字符，如空格、制表符和换行符。

要在字符串中添加制表符，可使用字符组合 `\t`，要在字符串中添加换行符，可使用字符组合 `\n` 。

```bash
复制代码>>> print('\tPython')
	Python
>>> print('Hello,\nPython')
Hello,
Python
```

#### 2.2.4 删除空白

`rstrip()` 删除右侧空白，`lstrip()` 删除左侧空白，`strip()` 删除两端空白。

```bash
复制代码>>> msg = ' Python '
>>> msg
' Python '
>>> msg.rstrip()
' Python'
>>> msg.lstrip()
'Python '
>>> msg.strip()
'Python'
>>> msg
' Python '
```

注意执行完去空格命令后，再打印出 msg，还是原来的字符串，这说明 `strip()` 也不改变原来的字符串。

#### 2.2.5 Python 2 中的 print 语句

在Python 2中，print语句的语法稍有不同:

```bash
复制代码>>> python2.7
>>> print "Hello Python 2.7 world!" 
Hello Python 2.7 world!
```

在Python 2中，无需将要打印的内容放在括号内。从技术上说，Python 3中的print是一个函数，因此括号必不可少。有些Python 2 print语句也包含括号，但其行为与Python 3中稍有不同。简单地说，在Python 2代码中，有些print语句包含括号，有些不包含。

### 2.3 数字

#### 2.3.1 整数

在Python中，可对整数执行加(+)减(-)乘(*)除(/)运算。

```bash
复制代码>>> 2 + 3 
5
>>> 3 - 2 
1
>>> 2 * 3 
6
>>> 3 / 2 
1.5
```

Python还支持运算次序，因此你可在同一个表达式中使用多种运算。你还可以使用括号来修 改运算次序，让Python按你指定的次序执行运算，如下所示:

```bash
复制代码>>> 2 + 3*4
14
>>> (2 + 3)*4-20
```

#### 2.3.2 浮点数

Python将带小数点的数字都称为浮点数。大多数编程语言都使用了这个术语，它指出了这样一个事实:小数点可出现在数字的任何位置。

从很大程度上说，使用浮点数时都无需考虑其行为。你只需输入要使用的数字，Python通常都会按你期望的方式处理它们:

```bash
复制代码>>> 0.1 + 0.1
0.2
>>>2 * 0.1
0.2
>>>2 * 0.2
0.4
```

但需要注意的是，结果包含的小数位数可能是不确定的:

```bash
复制代码>>> 0.2 + 0.1 
0.30000000000000004 
>>> 3 * 0.1 
0.30000000000000004
```

所有语言都存在这种问题，没有什么可担心的。Python会尽力找到一种方式，以尽可能精确地表示结果，但鉴于计算机内部表示数字的方式，这在有些情况下很难。后面将会学习更多的处理方式。

#### 2.3.3 使用函数str()避免类型错误

如果用数字跟字符串拼接，就会出现类型错误。为避免这个问题，可以使用 `str()` 将数字转换为字符串再进行操作。

```bash
复制代码>>> age = 18
>>> print('my age is ' + age)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: can only concatenate str (not "int") to str
>>> print('my age is ' + str(age))
my age is 18
```

#### 2.3.4 Python 2 中的整数

在Python 2中，将两个整数相除得到的结果稍有不同:

```bash
复制代码>>> python2.7 
>>> 3 / 2
1
```

Python返回的结果为1，而不是1.5。在Python 2中，整数除法的结果只包含整数部分，小数部 分被删除。请注意，计算整数结果时，采取的方式不是四舍五入，而是将小数部分直接删除。

在Python 2中，若要避免这种情况，务必确保至少有一个操作数为浮点数，这样结果也将为 浮点数:

```bash
复制代码>>> 3 / 2
1
>>> 3.0 / 2
1.5
>>> 3 / 2.0 
1.5
>>> 3.0 / 2.0 
1.5
```

从Python 3转而用Python 2或从Python 2转而用Python 3时，这种除法行为常常会令人迷惑。使用或编写同时使用浮点数和整数的代码时，一定要注意这种异常行为。

#### 2.3.5 注释

在Python中，注释用井号(#)标识。井号后面的内容都会被Python解释器忽略。如

```bash
复制代码# 向大家问好
print("Hello Python people!")
```

## 三、列表

列表由一系列按特定顺序排列的元素组成。

在Python中，用方括号([])来表示列表，并用逗号来分隔其中的元素。

```bash
复制代码>>> list = []
>>> list.append('haha')
>>> list.append('heihei')
>>> list.append('hehe')
>>> list
['haha', 'heihei', 'hehe']
>>> list[0]
'haha'
```

> 获取最后一个元素可以用 -1，如 list[-1] 是获取最后一个元素，list[-2] 是获取倒数第二个元素。

### 3.1 列表的增删改查

#### 3.1.1 修改元素

修改元素直接用索引修改

```bash
复制代码>>> list[0] = 'nihao'
>>> list
['nihao', 'heihei', 'hehe']
```

#### 3.1.2 添加元素

可以在末尾添加，也可以在任意位置插入。

**在末尾添加：append**

```bash
复制代码>>> list.append('wa')
>>> list
['nihao', 'heihei', 'hehe', 'wa']
```

**插入：insert**

```bash
复制代码>>> list.insert(1, 'hello')
>>> list
['nihao', 'hello', 'heihei', 'hehe', 'wa']
```

#### 3.1.3 删除元素

删除有三种方式：

- del：按索引删除
- pop()：删除列表最后一个元素并返回最后一个元素的值。也可以传索引删除任意位置的值。
- remove()：按值删除

```bash
复制代码>>> list
['nihao', 'hello', 'heihei', 'hehe', 'wa']
>>> del list[1]
>>> list
['nihao', 'heihei', 'hehe', 'wa']
>>> list.pop()
'wa'
>>> list.remove('hehe')
>>> list
['nihao', 'heihei']
```

给 `pop()` 传索引删除其他位置的值

```bash
复制代码>>> list
['nihao', 'heihei']
>>> list.pop(0)
'nihao'
>>> list
['heihei']
```

> 注意：
>
> 方法remove()只删除第一个指定的值。如果要删除的值可能在列表中出现多次，就需要使用循环来判断是否删除了所有这样的值。
>
> 如果你不确定该使用del语句还是pop()方法，下面是一个简单的判断标准:如果你要从列表中删除一个元素，且不再以任何方式使用它，就使用del语句;如果你要在删除元素后还能继续使用它，就使用方法pop()。

### 3.2 组织列表

本节将介绍列表的排序、反转、计算长度等操作。

列表的排序主要有两种方式：

- 使用方法sort()对列表进行永久性排序
- 使用函数sorted()对列表进行临时排序

#### 3.2.1 使用方法sort()对列表进行永久性排序

使用 `sort()` 方法将改变原列表。如果要反转排序，只需向sort()方法传递参数 reverse=True。

```
复制代码>>> list
['zhangsan', 'lisi', 'bob', 'alex']
>>> list.sort()
>>> list
['alex', 'bob', 'lisi', 'zhangsan']
>>> list.sort(reverse=True)
>>> list
['zhangsan', 'lisi', 'bob', 'alex']
```

#### 3.2.2 使用函数sorted()对列表进行临时排序

函数 `sorted()` 让你能够按特定顺序显示列表元素，同时不影响它们在列表中的原始排列顺序。

如果要反转排序，只需向 `sorted()` 传递参数 reverse=True。

```
复制代码>>> list = ['douglas','alex','solo','super']
>>> sorted(list)
['alex', 'douglas', 'solo', 'super']
>>> list
['douglas', 'alex', 'solo', 'super']
>>> sorted(list, reverse=True)
['super', 'solo', 'douglas', 'alex']
>>> list
['douglas', 'alex', 'solo', 'super']
```

#### 3.2.3 反转列表

要反转列表元素的排列顺序，可使用方法 `reverse()`。 `reverse()` 也会改变原始列表。

`reverse()` 只会按原来的顺序反转，不会进行额外的按字母排序。

```
复制代码>>> list
['douglas', 'alex', 'solo', 'super']
>>> list.reverse()
>>> list
['super', 'solo', 'alex', 'douglas']
```

#### 3.2.4 确定列表的长度

使用函数len()可快速获悉列表的长度。

```
复制代码>>> list
['super', 'solo', 'alex', 'douglas']
>>> len(list)
4
```

### 3.3 操作列表

#### 3.3.1 循环

使用 `for…in` 循环。

python 以缩进来区分代码块，所以需要正确的缩进

```
复制代码>>> cats
['super', 'solo', 'alex', 'douglas']
>>> for cat in cats:
...     print(cat)
...
super
solo
alex
douglas
```

#### 3.3.2 range()

Python函数range()让你能够轻松地生成一系列的数字。

```
复制代码>>> for value in range(1,5):
...     print(value)
...
1
2
3
4
```

注意：range() 会产生包含第一个参数但不包含第二个参数的一系列数值。

使用 `range()` 创建列表

```
复制代码>>> numbers = list(range(1,6))
>>> numbers
[1, 2, 3, 4, 5]
```

`range()` 还可以指定步长。下面的例子生成了从0开始，到11的偶数：

```
复制代码>>> nums = list(range(0,11,2))
>>> nums
[0, 2, 4, 6, 8, 10]
```

#### 3.3.3 对列表简单的计算

有几个专门用于处理数字列表的Python函数。

- min()：计算最小值
- max()：计算最大值
- sum()：计算总和

```
复制代码>>> numbers
[1, 2, 3, 4, 5]
>>> min(numbers)
1
>>> max(numbers)
5
>>> sum(numbers)
15
```

#### 3.3.4 列表解析

列表解析将for循环和创建新元素的代码合并成一行，并自动附加新元素。

```
复制代码>>> squares = [value**2 for value in range(1,11)]
>>> squares
[1, 4, 9, 16, 25, 36, 49, 64, 81, 100]
```

要使用这种语法，首先指定一个描述性的列表名，如squares;然后，指定一个左方括号， 并定义一个表达式，用于生成你要存储到列表中的值。在这个示例中，表达式为 value ** 2，它计 算平方值。接下来，编写一个for循环，用于给表达式提供值，再加上右方括号。在这个示例中， for循环为for value in range(1,11)，它将值1~10提供给表达式 value ** 2。请注意，这里的for 语句末尾没有冒号。

### 3.4 切片

要创建切片，可指定要使用的第一个元素和最后一个元素的索引。与函数range()一样，Python在到达你指定的第二个索引前面的元素后停止。要输出列表中的前三个元素，需要指定索引0~3，这将输出分别为0、1和2的元素。

```
复制代码>>> names = ['aa','bb','cc','dd']
>>> print(names[1:4])
['bb', 'cc', 'dd']
```

如果你没有指定第一个索引，Python将自动从列表开头开始:

```
复制代码>>> print(names[:4])
['aa', 'bb', 'cc', 'dd']
```

如果没有指定终止索引，将自动取到列表末尾

```
复制代码>>> print(names[2:])
['cc', 'dd']
```

也可以使用负数索引，比如返回最后三个元素

```
复制代码>>> print(names[-3:])
['bb', 'cc', 'dd']
```

遍历切片

```
复制代码>>> for name in names[1:3]:
...     print(name)
...
bb
cc
```

### 3.5 复制列表

可以使用切片来快速复制列表，不指定开始索引和结束索引。

```
复制代码>>> names
['aa', 'bb', 'cc', 'dd']
>>> names2 = names[:]
>>> names2
['aa', 'bb', 'cc', 'dd']
```

用切片复制出来的新列表，跟原来的列表是完全不同的列表，改变其实一个不会影响另一个列表。

```
复制代码>>> names.append('ee')
>>> names
['aa', 'bb', 'cc', 'dd', 'ee']
>>> names2
['aa', 'bb', 'cc', 'dd']
```

而如果简单的通过赋值将 names 赋值给 names2，就不能得到两个列表，实际上它们都指向了同一个列表。如果改变其中一个，另一个也将被改变。

```
复制代码>>> names
['aa', 'bb', 'cc', 'dd']
>>> names2 = names
>>> names2
['aa', 'bb', 'cc', 'dd']
>>> names.append('ee')
>>> names
['aa', 'bb', 'cc', 'dd', 'ee']
>>> names2
['aa', 'bb', 'cc', 'dd', 'ee']
```

### 3.6 元组

Python将不能修改的值称为不可变的，而不可变的列表被称为元组。

元组看起来犹如列表，但使用圆括号而不是方括号来标识。定义元组后，就可以使用索引来访问其元素，就像访问列表元素一样。

```
复制代码>>> food = ('apple', 'orange')
>>> food[0]
'apple'
>>> food[1]
'orange'
>>> food[1] = 'banana'
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'tuple' object does not support item assignment
```

遍历用法跟列表一致。

## 四、条件判断

每条if语句的核心都是一个值为True或False的表达式，这种表达式被称为条件测试。

- 检查是否相等，用 `==`

- 检查是否不相等，用 `!=`

- 数字比较 `>`、 `<`、 `>=`、 `<=`

- 多个条件与 `and`

- 多个条件或 `or`

- 判断列表是否包含某元素 `in`

  ```
  复制代码>>> names
  ['aa', 'bb', 'cc', 'dd', 'ee']
  >>> 'bb' in names
  True
  ```

- 判断列表是否不包含某元素

  ```
  复制代码>>> names
  ['aa', 'bb', 'cc', 'dd', 'ee']
  >>> 'ff' not in names
  True
  ```

### if 语句

简单的 if-else

```
复制代码>>> a = 10
>>> if a > 10:
...     print('hello')
... else:
...     print('bye')
...
bye
```

if-elif-else

```
复制代码>>> if a<5:
...     print(a<5)
... elif 5<a<10:
...     print('5<a<10')
... else:
...     print('a>10')
...
a>10
```

## 五、字典

在Python中，字典是一系列键-值对。每个键都与一个值相关联，你可以使用键来访问与之相关联的值。与键相关联的值可以是数字、字符串、列表乃至字典。事实上，可将任何Python对象用作字典中的值。

### 5.1 字典的增删改查

#### 使用字典

在Python中，字典用放在花括号{}中的一系列键-值对表示。

```
复制代码>>> user = {'name':'bob', 'sex':'male', 'age':20}
>>> user
{'name': 'bob', 'sex': 'male', 'age': 20}
```

#### 访问字典中的值

要获取与键相关联的值，可依次指定字典名和放在方括号内的键。

```
复制代码>>> user
{'name': 'bob', 'sex': 'male', 'age': 20}
>>> user['name']
'bob'
>>> user['age']
20
```

#### 添加键值对

字典是一种动态结构，可随时在其中添加键—值对。

```
复制代码>>> user['city']='beijing'
>>> user
{'name': 'bob', 'sex': 'male', 'age': 20, 'city': 'beijing'}
```

#### 修改字典中的值

要修改字典中的值，可依次指定字典名、用方括号括起的键以及与该键相关联的新值。

```
复制代码>>> cat = {}
>>> cat['color'] = 'white'
>>> cat['age'] = 4
>>> cat
{'color': 'white', 'age': 4}
>>> cat['age'] = 6
>>> cat
{'color': 'white', 'age': 6}
```

#### 删除键值对

对于字典中不再需要的信息，可使用del语句将相应的键—值对彻底删除。使用del语句时，必须指定字典名和要删除的键。

```
复制代码>>> del cat['color']
>>> cat
{'age': 6}
```

### 5.2 遍历字典

字典可用于以各种方式存储信息，因此有多种遍历字典的方式:可遍历字典的所有键—值对、键或值。

#### 遍历所有键值对 `items()`

```
复制代码>>> cat
{'age': 6, 'color': 'white', 'city': 'beijing'}
>>> for k,v in cat.items():
...     print(k + '-' + str(v))
...
age-6
color-white
city-beijing
```

通过 `for k,v in cat.items()` 的方式遍历所有的键值对，`k` 代表键，`v` 代表值。

> 注意：即便遍历字典时，键—值对的返回顺序也与存储顺序不同。Python不关心键—值对的存储顺序，而只跟踪键和值之间的关联关系。

#### 遍历所有键 `keys()`

如果不需要用值，可以用 `keys()` 遍历出所有的键。

```
复制代码>>> cat
{'age': 6, 'color': 'white', 'city': 'beijing'}
>>> for k in cat.keys():
...     print(k.title())
...
Age
Color
City
```

上面的例子打印出了 `cat` 的所有键，用字符串的 `title()` 方法使每个单词的首字母大写。

遍历字典时会默认遍历所有的键，`for k in cat.keys()` 和 `for k in cat` 的效果一样。

按顺序遍历所有键，可用 `sorted()` 排序，这让Python列出字典中的所有键，并在遍历前对这个列表进行排序。

```
复制代码>>> for k in sorted(cat.keys()):
...     print(k.title())
...
Age
City
Color
```

#### 遍历所有值 `values()`

```
复制代码>>> for value in cat.values():
...     print(str(value))
...
6
white
beijing
```

如果需要剔除重复项，可以使用 `set()`

```
复制代码>>> cat
{'age': 6, 'color': 'white', 'city': 'beijing', 'city2': 'beijing'}
>>> for value in cat.values():
...     print(str(value))
...
6
white
beijing
beijing
>>> for value in set(cat.values()):
...     print(str(value))
...
beijing
white
6
```

### 5.3 嵌套

可以在列表中嵌套字典、在字典中嵌套列表以及在字典中嵌套字典。这里就不演示了。

## 六、用户输入和while循环

### 6.1 用户输入

函数input()让程序暂停运行，等待用户输入一些文本。获取用户输入后，Python将其存储在一个变量中，以方便你使用。

```
复制代码>>> msg = input('Please input your name: ')
Please input your name: solo
>>> msg
'solo'
```

如果你使用的是Python 2.7，应使用函数raw_input()来提示用户输入。这个函数与Python 3中的input()一样，也将输入解读为字符串。

Python 2.7也包含函数input()，但它将用户输入解读为Python代码，并尝试运行它们。如果你使用的是Python 2.7，请使用raw_input()而不是input()来获取输入。

如果想将输入的内容转换为数字，可以用 `int()` 来转换。

### 6.2 while 循环

for循环用于针对集合中的每个元素都一个代码块，而while循环不断地运行，直到指定的条件不满足为止。

```
复制代码>>> num = 1
>>> while num <= 5:
...     print(str(num))
...     num += 1
...
1
2
3
4
5
```

#### break

要立即退出while循环，不再运行循环中余下的代码，也不管条件测试的结果如何，可使用break语句。break语句用于控制程序流程，可使用它来控制哪些代码行将执行，哪些代码行不执行，从而让程序按你的要求执行你要执行的代码。

#### continue

要返回到循环开头，并根据条件测试结果决定是否继续执行循环，可使用continue语句，它不像 break 语句那样不再执行余下的代码并退出整个循环。

## 七、函数

Python 用关键字 `def` 来定义函数，函数名以冒号 `:` 结尾，冒号之后的缩进里的内容都是函数体。

```
复制代码>>> def greet():
...     print('Hello World!')
...
>>> greet()
Hello World!
```

### 7.1 函数参数

可以向函数传递参数。下面的例子向函数 `greet()` 传递了个参数 `name`。其中 `name` 是形参，`solo` 是实参。

```
复制代码>>> def greet(name):
...     print('Hello,' + name)
...
>>> greet('solo')
Hello,solo
```

向函数传递实参的方式很多，可使用位置实参，这要求实参的顺序与形参的顺序相同;也可使用关键字实参，其 中每个实参都由变量名和值组成;还可使用列表和字典。

#### 位置实参

你调用函数时，Python必须将函数调用中的每个实参都关联到函数定义中的一个形参。为此，最简单的关联方式是基于实参的顺序。这种关联方式被称为位置实参。

```
复制代码>>> def student(name, age):
...     print('Hello, My name is ' + name + ', I am ' + str(age) + ' years old')
...
>>> student('solo', 18)
Hello, My name is solo, I am 18 years old
```

按照形参定义的顺序传递的实参就称为位置实参。

#### 关键字实参

关键字实参是传递给函数的名称—值对。关键字实参让你无需考虑函数调用中的实参顺序，还清楚地指出了函数调用中各个值的用途。

```
复制代码>>> student(age=18, name='solo')
Hello, My name is solo, I am 18 years old
```

接着位置实参中的例子，`student(name, age)` 方法第一个参数是 `name`，第二个参数是 `age` 。我们用关键字实参指明传递的是哪一个，即使顺序写乱了得到的结果也不会乱。

#### 默认值

编写函数时，可给每个形参指定默认值。在调用函数中给形参提供了实参时，Python将使用指定的实参值;否则，将使用形参的默认值。因此，给形参指定默认值后，可在函数调用中省略相应的实参。使用默认值可简化函数调用，还可清楚地指出函数的典型用法。

```
复制代码>>> def student(name, age=18):
...     print('Hello, My name is ' + name + ', I am ' + str(age) + ' years old')
...
>>> student('bob')
Hello, My name is bob, I am 18 years old
>>> student('nicole')
Hello, My name is nicole, I am 18 years old
>>> student('bob', 20)
Hello, My name is bob, I am 20 years old
```

如上，给 `student()` 函数定义的第二个参数 `age` 设置了默认值 `18`，如果调用时只传一个参数，无论传的是什么 `age` 都是 18。当传两个参数时，传递的实参就会覆盖掉默认值。

> 注意：使用默认值时，在形参列表中必须先列出没有默认值的形参，再列出有默认值的实参。这让Python依然能够正确地解读位置实参。

### 7.2 返回值

函数并非总是直接显示输出，相反，它可以处理一些数据，并返回一个或一组值。函数返回 的值被称为返回值。在函数中，可使用return语句将值返回到调用函数的代码行。返回值让你能够将程序的大部分繁重工作移到函数中去完成，从而简化主程序。

```
复制代码>>> def student(name):
...     return name
...
>>> name = student('solo')
>>> name
'solo'
```

#### 返回字典

函数可返回任何类型的值，包括列表和字典等较复杂的数据结构。例如，下面的函数接受姓名和年龄，并返回一个表示人的字典:

```
复制代码>>> def build_person(name,age):
...     person = {'name':name, 'age':age}
...     return person
...
>>> p = build_person('solo',18)
>>> p
{'name': 'solo', 'age': 18}
```

### 7.3 传递任意数量的实参

有时候，你预先不知道函数需要接受多少个实参，好在Python允许函数从调用语句中收集任意数量的实参。

```
复制代码>>> def person(*args):
...     print(args)
...
>>> person('name','age','address')
('name', 'age', 'address')
```

上面定义了一个函数 `person()` ，只有一个形参 `*args` 。形参名 `*args` 中的星号让 Python 创建一个名为 `args` 的空元组，并将收到的所有值都封装到这个元组中。

#### 结合使用位置实参和任意数量实参

如果要让函数接受不同类型的实参，必须在函数定义中将接纳任意数量实参的形参放在最后。Python 先匹配位置实参和关键字实参，再将余下的实参都收集到最后一个形参中。

```
复制代码>>> def person(city, *args):
...     print('city: ' + city + ', other args:')
...     for value in args:
...             print(value)
...
>>> person('beijing', 'name', 'age', 'tel')
city: beijing, other args:
name
age
tel
```

函数 `person()` 有两个形参，第一个 `city` 是普通的位置实参，第二个 `*args` 是可变参数。

#### 使用任意数量的关键字实参

有时候，需要接受任意数量的实参，但预先不知道传递给函数的会是什么样的信息。在这种情况下，可将函数编写成能够接受任意数量的键—值对——调用语句提供了多少就接受多少。一个这样的示例是创建用户简介:你知道你将收到有关用户的信息，但不确定会是什么样的信息。

```
复制代码def build_profile(first, last, **user_info):
	profile = {}
	profile['first_name'] = first
	profile['last_name'] = last

	for key,value in user_info.items():
		profile[key] = value

	return profile

user = build_profile('steven', 'bob', city='beijing', age=18)

print(user)
```

执行代码，输出结果是：

```
复制代码{'first_name': 'steven', 'last_name': 'bob', 'city': 'beijing', 'age': 18}
```

### 7.4 导入导出

可以将函数存储在被称为模块的独立文件中，再将模块导入到主程序中。import语句允许在当前运行的程序文件中使用模块中的代码。

#### 7.4.1 导入整个模块

模块是扩展名为.py的文件，包含要导入到程序中的代码。

> [cat.py](https://link.juejin.cn/?target=http%3A%2F%2Fcat.py)

```
复制代码def eat(food):
    print('I am cat, I eat ' + food)
```

> [animal.py](https://link.juejin.cn/?target=http%3A%2F%2Fanimal.py)

```
复制代码import cat

cat.eat('fish')
```

控制台输出

```
复制代码I am cat, I eat fish
```

#### 7.4.2 导入特定的函数

你还可以导入模块中的特定函数，这种导入方法的语法如下:

```
复制代码from module_name import function_name 
```

通过用逗号分隔函数名，可根据需要从模块中导入任意数量的函数:

```
复制代码from module_name import function_0, function_1, function_2 
```

上面的例子只导入 `cat.py` 中的 `eat()` 方法

```
复制代码from cat import eat

eat('fish')
```

得到相同的结果。

#### 7.4.3 使用 as 给函数指定别名

如果要导入的函数的名称可能与程序中现有的名称冲突，或者函数的名称太长，可指定简短而独一无二的别名——函数的另一个名称，类似于外号。要给函数指定这种特殊外号，需要在导入它时这样做。

```
复制代码from cat import eat as cat_eat

cat_eat('fish')
```

将 `cat.py` 中的 `eat()` 方法导入并指定了别名 `cat_eat`，使用时可以直接用别名使用。

#### 7.4.4 使用 as 给模块指定别名

你还可以给模块指定别名。通过给模块指定简短的别名，让你 能够更轻松地调用模块中的函数。

通用语法：`import module_name as mn`

```
复制代码import cat as c

c.eat('fish')
```

#### 7.4.5 导入模块中的所有函数

使用星号(*)运算符可让Python导入模块中的所有函数:

> [cat.py](https://link.juejin.cn/?target=http%3A%2F%2Fcat.py)

```
复制代码def eat(food):
    print('I am cat, I eat ' + food)


def run():
    print('cat run')
```

> [animal.py](https://link.juejin.cn/?target=http%3A%2F%2Fanimal.py)

```
复制代码from cat import *

eat('fish')
run()
```

输出结果

```
复制代码I am cat, I eat fish
cat run
```

由于导入 了每个函数，可通过名称来调用每个函数，而无需使用句点表示法。然而，使用并非自己编写的 大型模块时，最好不要采用这种导入方法:如果模块中有函数的名称与你的项目中使用的名称相 同，可能导致意想不到的结果: Python 可能遇到多个名称相同的函数或变量，进而覆盖函数，而 不是分别导入所有的函数。

最佳的做法是，要么只导入你需要使用的函数，要么导入整个模块并使用句点表示法。这能 让代码更清晰，更容易阅读和理解。

### 7.5 函数编写指南

- 应给函数指定描述性名称

- 函数名应只包含小写字母和下划线

- 每个函数都应包含简要地阐述其功能的注释，该注释应紧跟在函数定义后面，并采用文档字符串格式。

- 给形参指定默认值时，等号两边不要有空格:

  ```
  复制代码def function_name(parameter_0, parameter_1='default value')
  ```

  对于函数调用中的关键字实参，也应遵循这种约定:

  ```
  复制代码function_name(value_0, parameter_1='value')
  ```

- 如果程序或模块包含多个函数，可使用两个空行将相邻的函数分开，这样将更容易知道前一个函数在什么地方结束，下一个函数从什么地方开始。

- 所有的import语句都应放在文件开头，唯一例外的情形是，在文件开头使用了注释来描述整个程序。

## 八、类

### 8.1 创建和使用类

```
复制代码class Cat():
    def __init__(self, name, color):
        self.name = name
        self.color = color

    def eat(self):
        print('cat ' + self.name + ' color ' + self.color + ', now eat')

    def run(self):
        print('cat ' + self.name + ' color ' + self.color + ', now run')


my_cat = Cat('Spring', 'white')
print(my_cat.name)
print(my_cat.color)
my_cat.eat()
my_cat.run()
```

上面创建了类 `Cat` ，并实例化了 `my_cat`，然后调用了类的方法 `eat()` 和 `run()`。输出结果：

```
复制代码Spring
white
cat Spring color white, now eat
cat Spring color white, now run
```

类中的函数称为方法。`__init__()` 是函数的构造方法，每档创建新实例时 Python 都会自动运行它。注意构造方法名字必须是这个，是规定好的。

上面的例子中`__init__(self, name, color)` 有三个形参，第一个形参 `self` 必不可少，还必须位于其他形参的前面。其他的形参可以根据需要调整。`self` 是一个指向实例本身的引用，让实例能够访问类中的属性和方法。

还可以通过实例直接访问属性：`my_cat.name`。但在其他语言中并不建议这样做。

#### 在Python 2.7中创建类

在Python 2.7中创建类时，需要做细微的修改——在括号内包含单词object:

```
复制代码class ClassName(object):
```

### 8.2 类的属性

#### 8.2.1 给属性设置默认值

类中的每个属性都必须有初始值，哪怕这个值是0或空字符串。在有些情况下，如设置默认值时，在方法`__init__()` 内指定这种初始值是可行的;如果你对某个属性这样做了，就无需包含为它提供初始值的形参。

重新定义 `Cat` ，在构造方法中给属性 `age` 设置默认值。

```
复制代码class Cat():
    def __init__(self, name, color):
        self.name = name
        self.color = color
        self.age = 3

    def eat(self):
        print('cat ' + self.name + ' color ' + self.color + ', now eat')

    def run(self):
        print('cat ' + self.name + ' color ' + self.color + ', now run')

    def print_age(self):
        print('cat`s age is ' + str(self.age))
```

#### 8.2.2 修改属性的值

可以以三种不同的方式修改属性的值：直接通过实例进行修改，通过方法进行设置。

##### 1. 直接修改属性的值

要修改属性的值，最简单的方式是通过实例直接访问它。

```
复制代码class Cat():
    def __init__(self, name, color):
        self.name = name
        self.color = color
        self.age = 3

    def eat(self):
        print('cat ' + self.name + ' color ' + self.color + ', now eat')

    def run(self):
        print('cat ' + self.name + ' color ' + self.color + ', now run')

    def print_age(self):
        print('cat`s age is ' + str(self.age))


my_cat = Cat('Spring', 'white')

my_cat.print_age()
my_cat.age = 4
my_cat.print_age()
```

输出结果为

```
复制代码cat`s age is 3
cat`s age is 4
```

上例直接通过 `my_cat.age = 4` 修改了 `age` 属性的值。

##### 2. 通过方法修改属性的值

再来更新代码，加入 `update_age()` 方法来修改 `age` 的属性。

```
复制代码class Cat():
    def __init__(self, name, color):
        self.name = name
        self.color = color
        self.age = 3

    def eat(self):
        print('cat ' + self.name + ' color ' + self.color + ', now eat')

    def run(self):
        print('cat ' + self.name + ' color ' + self.color + ', now run')

    def print_age(self):
        print('cat`s age is ' + str(self.age))

    def update_age(self, age):
        self.age = age


my_cat = Cat('Spring', 'white')
my_cat.print_age()
my_cat.update_age(10)
my_cat.print_age()
```

运行代码输出：

```
复制代码cat`s age is 3
cat`s age is 10
```

### 8.3 继承

一个类继承另一个类时，它将自动获得另一个类的所有属性和方法;原有的类称为父类，而新类称为子类。子类继承了其父类的所有属性和方法，同时还可以定义自己的属性和方法。

```
复制代码class Animal():
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def run(self):
        print('Animal ' + self.name + ' run')


class Cat(Animal):
    def __init__(self, name, age):
        super().__init__(name, age)


cat = Cat('Tony', 2)
cat.run()
```

运行程序，输出：

```
复制代码Animal Tony run
```

先定义了类 `Animal`，又定义了 `Cat` 继承自 `Animal`。 `Animal`称为父类， `Cat` 称为子类。通过输出可以验证，子类继承了父类的方法。

在子类的构造方法中要先实现父类的构造方法：`super().__init__(name, age)`。

还可以给子类定义自己的方法，或者重写父类的方法。

```
复制代码class Animal():
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def run(self):
        print('Animal ' + self.name + ' run')


class Cat(Animal):
    def __init__(self, name, age):
        super().__init__(name, age)

    def play(self):
        print('Cat ' + self.name + ' play')

    def run(self):
        print('Cat ' + self.name + ' run')


cat = Cat('Tony', 2)
cat.run()
cat.play()
```

我们来修改下程序，`Animal` 类不变，`Cat` 类还是继承了 `Animal` ，但定义了自己的方法 `play()` 并重写了父类方法 `run()` 。运行程序，得到输出：

```
复制代码Cat Tony run
Cat Tony play
```

#### Python2.7 中的继承

在Python 2.7中，继承语法稍有不同，ElectricCar类的定义类似于下面这样:

```
复制代码class Car(object):
	def __init__(self, make, model, year):
    	--snip--
    
class ElectricCar(Car):
	def __init__(self, make, model, year):
		super(ElectricCar, self).__init__(make, model, year)
		--snip--
```

函数super()需要两个实参:子类名和对象self。为帮助Python将父类和子类关联起来，这些实参必不可少。另外，在Python 2.7中使用继承时，务必在定义父类时在括号内指定object。

### 8.4 导入类

当一个文件过长时，可以将其中一部分代码抽离出去，然后导入到主文件中。

导入方式有多种：

- 导入单个类

  假如 `car.py` 里定义了类 `Car`

  ```
  复制代码from car import Car
  ```

- 从一个模块中导入多个类

  假如 `car.py` 包含了三个类 `Car` ， `Battery` 和 `ElectricCar` 。

  只导入一个类：

  ```
  复制代码from car import ElectricCar
  ```

  导入多个类，中间用逗号隔开：

  ```
  复制代码from car import Car, ElectricCar
  ```

- 导入整个模块

  还可以导入整个模块，再使用句点表示法访问需要的类。这种导入方法很简单，代码也易于阅读。由于创建类实例的代码都包含模块名，因此不会与当前文件使用的任何名称发生冲突。

  ```
  复制代码import car
  my_car = car.Car()
  ```

- 导入模块中的所有类

  要导入模块中的每个类，可使用下面的语法:

  ```
  复制代码from module_name import *
  ```

  > 不推荐使用这种导入方式，其原因有二。
  >
  > 首先，如果只要看一下文件开头的import语句，就 能清楚地知道程序使用了哪些类，将大有裨益;但这种导入方式没有明确地指出你使用了模块中 的哪些类。这种导入方式还可能引发名称方面的困惑。如果你不小心导入了一个与程序文件中其 他东西同名的类，将引发难以诊断的错误。这里之所以介绍这种导入方式，是因为虽然不推荐使 用这种方式，但你可能会在别人编写的代码中见到它。
  >
  > 需要从一个模块中导入很多类时，最好导入整个模块，并使用module_name.class_name语法 来访问类。这样做时，虽然文件开头并没有列出用到的所有类，但你清楚地知道在程序的哪些地 方使用了导入的模块;你还避免了导入模块中的每个类可能引发的名称冲突。

## 九、文件和异常

### 9.1 从文件中读取数据

要使用文本文件中的信息，首先需要将信息读取到内存中。为此，你可以一次性读取文件的全部内容，也可以以每次一行的方式逐步读取。

#### 9.1.1 读取整个文件

```
复制代码with open('test.txt') as file_obj:
    contents = file_obj.read()
    print(contents)
```

`open()` 用于打开一个文件，参数为文件的路径。

关键字 `with` 在不再需要访问文件后将其关闭。有了 `with` 你只管打开文件，并在需要时使用它，Python自会 在合适的时候自动将其关闭。

相比于原始文件，该输出唯一不同的地方是末尾多了一个空行。为何会多出这个空行呢?因为 `read()` 到达文件末尾时返回一个空字符串，而将这个空字符串显示出来时就是一个空行。要删除多出来的空行，可在print语句中使用 `rstrip()`。

文件路径可以是相对路径，也可以是绝对路径。

#### 9.1.2 逐行读取

```
复制代码with open('test.txt') as file_obj:
    for line in file_obj:
        print(line.rstrip())
```

要以每次一行的方式检查文件，可对文件对象使用for循环。

### 9.2 写入文件

```
复制代码with open('test.txt', 'w') as file_obj:
    file_obj.write("I love python")
```

在这个示例中，调用open()时提供了两个实参，第一个实参也是要打开的文件的名称；第二个实参('w')告诉Python，我们要以写入模式打开这个文件。

可选模式：

- `r` ：只读。
- `w` : 只写。如果文件不存在则创建，如果文件存在则先清空，再写入。
- `a` ：附加模式，写入的内容追加到原始文件后面。如果文件不存在则创建。
- `r+` ：可读可写。

如果你省略了模式实参，Python将以默认的只读模式打开文件。

### 9.3 异常

异常是使用try-except代码块处理的。try-except代码块让Python执行指定的操作，同时告诉Python发生异常时怎么办。

```
复制代码try:
    print(5/0)
except ZeroDivisionError:
    print("You can't divide by zero!")
```

#### else 代码块

```
复制代码try:
    print(5/0)
except ZeroDivisionError:
    print("You can't divide by zero!")
else:
	print("no exception")
```

如果 `try` 中的代码运行成功，没有出现异常，则执行 `else` 代码块中的代码。

### 9.4 用 json 存储数据

Python 中使用 `json.dump()` 和 `json.load()` 来存储和读取 json 文件。

```
复制代码import json

userInfo = {'username': 'jack', 'age': 18}

with open('test.txt', 'w') as obj:
    json.dump(userInfo, obj)


with open('test.txt') as obj:
    content = json.load(obj)
    print(content)
```

上例中用 `json.dump()` 把数据存入到了 `test.txt` 中，又用 `json.load()` 把数据从文件中取出并打印。

注意使用前先导入 json 模块。

## 十、单元测试

先定义一个拼接名字的函数 `name_function.py`

```
复制代码def get_formatted_name(first, last):
    full_name = first + ' ' + last
    return full_name.title()
```

再写测试类来测试这个函数

```
复制代码import unittest
from name_function import get_formatted_name


class NamesTestCase(unittest.TestCase):

    def test_name_function(self):
        full_name = get_formatted_name('david', 'alex')
        self.assertEqual(full_name, 'David Alex')


unittest.main()
```

测试类要继承 `unittest.TestCase` ，通过 `self.assertEqual` 断言是否得到的结果和预期相等。

#### 常见的断言方法

| 方法                    | 用途               |
| ----------------------- | ------------------ |
| assertEqual(a, b)       | 核实a == b         |
| assertNotEqual(a, b)    | 核实a != b         |
| assertTrue(x)           | 核实x为True        |
| assertFalse(x)          | 核实x为False       |
| assertIn(item, list)    | 核实item在list中   |
| assertNotIn(item, list) | 核实item不在list中 |



#### `setUp()` 方法

如果你在TestCase类中包含了方法 `setUp()` ，Python将先运行它，再运行各个以test_打头的方法。

通常用于做一些初始化操作。

------

## 全栈全平台开源项目 CodeRiver

CodeRiver 是一个免费的项目协作平台，愿景是打通 IT 产业上下游，无论你是产品经理、设计师、程序员或是测试，还是其他行业人员，只要有好的创意、想法，都可以来 CodeRiver 免费发布项目，召集志同道合的队友一起将梦想变为现实！

CodeRiver 本身还是一个大型开源项目，致力于打造全栈全平台企业级精品开源项目。涵盖了 React、Vue、Angular、小程序、ReactNative、Android、Flutter、Java、Node 等几乎所有主流技术栈，主打代码质量。

目前已经有近 `100` 名优秀开发者参与，github 上的 `star` 数量将近 `1000` 个。每个技术栈都有多位经验丰富的大佬坐镇，更有两位架构师指导项目架构。无论你想学什么语言处于什么技术水平，相信都能在这里学有所获。

通过 `高质量源码 + 博客 + 视频`，帮助每一位开发者快速成长。

项目地址：[github.com/cachecats/c…](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fcachecats%2Fcoderiver)