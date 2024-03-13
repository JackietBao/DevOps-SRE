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



d = difflib.HtmlDiff()
#difflib 模块中的 HtmlDiff 类来比较两个文本文件，并生成一个 HTML 格式的比较结果
print d.make_file(text1_lines, text2_lines)
#text1_lines 和 text2_lines 分别是两个文本文件的内容，通过 make_file 方法生成一个 HTML 文件，该文件显示了两个文本文件的差异之处
