from selenium import webdriver
#导入Selenium库中的webdriver模块，它包含了用于控制各种浏览器的类和方法
options = webdriver.ChromeOptions()
#创建一个ChromeOptions对象，它用于配置Chrome浏览器的选项，比如设置代理、浏览器窗口大小等
path = 'chromedriver.exe'
#指定Chrome浏览器驱动的路径，这个驱动程序是用来与Chrome浏览器进行通信的桥梁
service = webdriver.chrome.service.Service(path)
#创建一个ChromeDriver服务对象，指定了ChromeDriver的路径
browser = webdriver.Chrome(service=service, options=options)
#使用上述配置创建一个Chrome浏览器对象。通过service参数指定了ChromeDriver服务，通过options参数指定了Chrome浏览器的选项
url = 'https://www.JD.com/'

browser.get(url)
#打开浏览器并访问指定的网页
content = browser.page_source
#这是Selenium中WebDriver对象的一个属性，用于获取当前页面的HTML源代码
print(content)
#打印content变量的值，即当前页面的HTML源代码

browser.quit()
# 关闭浏览器