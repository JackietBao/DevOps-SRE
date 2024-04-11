

from selenium import webdriver

path = 'chromedriver.exe'
browser = webdriver.Chrome(path)


url = 'http://www.baidu.com'
browser.get(url)


input = browser.find_element_by_id('su')

# 获取标签的属性
print(input.get_attribute('class'))
# 获取标签的名字
print(input.tag_name)

# 获取元素文本
a = browser.find_element_by_link_text('新闻')
print(a.text)










from selenium import webdriver
from selenium.webdriver.common.by import By

options = webdriver.ChromeOptions()
path = 'chromedriver.exe'
service = webdriver.chrome.service.Service(path)
browser = webdriver.Chrome(service=service, options=options)
url = 'https://www.baidu.com'
browser.get(url)

# 使用find_element方法，并传入By.ID作为参数来查找元素
input = browser.find_element(By.ID, 'su')
print(input.get_attribute('class'))
print(input.tag_name)

a = browser.find_element(By.LINK_TEXT, '新闻')
print(a.text)

browser.quit()
