#!/usr/bin/python
# -*- coding: utf-8 -*-
import difflib
#difflib 模块，这个模块提供了用于比较序列的类和函数，包括生成差异的功能
text1 = """text1:
This module provides classes and functions for comparing sequences.
including HTML and context and unified diffs.
difflib document v7.4
add string
"""
#这段文本是我们要进行比较的第一个文本

text1_lines = text1.splitlines()
#使用 splitlines() 方法将 text1 字符串分割成行，并将结果存储在 text1_lines 变量中

text2 = """text2:
This module provides classes and functions for Comparing sequences.
including HTML and context and unified diffs.
difflib document v7.5"""

text2_lines = text2.splitlines()

d = difflib.Differ()
#创建了一个 Differ 对象，用于执行两个文本之间的差异比较
diff = d.compare(text1_lines, text2_lines)
#使用 Differ 对象的 compare() 方法比较两个文本的行，并将结果存储在 diff 变量中
print '\n'.join(list(diff))
#diff 变量转换为列表，然后使用换行符将列表中的元素连接起来，并通过 print 语句输出到控制台
