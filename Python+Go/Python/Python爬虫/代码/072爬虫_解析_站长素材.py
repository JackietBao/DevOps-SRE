

# (1) 请求对象的定制
# （2）获取网页的源码
# （3）下载


# 需求 下载的前十页的图片
# https://sc.chinaz.com/tupian/qinglvtupian.html   1
# https://sc.chinaz.com/tupian/qinglvtupian_page.html

import urllib.request
from lxml import etree

def create_request(page):
    if(page == 1):
        url = 'https://sc.chinaz.com/tupian/qinglvtupian.html'
    else:
        url = 'https://sc.chinaz.com/tupian/qinglvtupian_' + str(page) + '.html'

    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36',
    }
    request = urllib.request.Request(url=url, headers=headers)
    #用于构建 HTTP 请求
    return request

def get_content(request):
    response = urllib.request.urlopen(request)
    #向指定的 URL 发送 HTTP 请求，并返回一个响应对象
    content = response.read().decode('utf-8')
    #读取 HTTP 响应的内容，并将其解码为 UTF-8 编码的字符串
    return content


def down_load(content):
#     下载图片
    # urllib.request.urlretrieve('图片地址','文件的名字')
    tree = etree.HTML(content)
    #将 HTML 格式的字符串 content 解析为一个文档树对象 tree，以便后续对网页内容进行操作和提取信息
    #print(etree.tostring(tree, encoding='utf-8').decode('utf-8'))
    name_list = tree.xpath('//div[@class="item"]//a[@class="name"]/text()')


    # 一般设计图片的网站都会进行懒加载
    src_list = tree.xpath('//div[@class="item"]//img[@class="lazy"]/@data-original')

    for i in range(len(name_list)):
        name = name_list[i]
        src = src_list[i]
        url = 'https:' + src

        urllib.request.urlretrieve(url=url,filename='./loveImg/' + name + '.jpg')




if __name__ == '__main__':
    start_page = int(input('请输入起始页码'))
    end_page = int(input('请输入结束页码'))

    for page in range(start_page,end_page+1):
        # (1) 请求对象的定制
        request = create_request(page)
        # （2）获取网页的源码
        content = get_content(request)
        # （3）下载
        down_load(content)
