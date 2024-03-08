computer = {"主机": 5000, "显示器": 100,"鼠标":150, "键盘":60}

#查询
print(computer["显示器"])
print(computer.get("显示器2"))
print(computer.keys())
print(computer.values())
print(computer.items())
#增加或更新
computer["音响"] = 200
print(computer)
c = {"音响2": 180}
computer.update(c)
print(computer)

#computer.setdefault("鼠标垫", "12")
computer.setdefault("刘慧", "哈哈")
#在字典中查找指定键，如果该键存在，则返回对应的值；如果键不存在，则将该键插入到字典中，并设置默认值，然后返回默认值
print(computer)

#删除
computer.pop("音响")
#从字典中移除指定键对应的键值对，并返回被移除的值
print(computer)
computer.popitem()
#随机地移除并返回字典中的一对键和值（即键值对）
#常会删除最后一个插入到字典中的项，但在某些情况下，可能会从字典中的其他位置删除项
print(computer)

#嵌套
computer = {"主机": {"CPU": 1500, "内存": 400, "硬盘": 500}, "显示器": 100,"鼠标":150, "键盘":["机械键盘","薄膜键盘"]}
print(computer)
print(computer["主机"]["内存"])
computer["主机"]["显卡"] = 1500
print(computer)
print("========================")
computer["键盘"].append("其他")
print(computer)
print(computer["键盘"][0])







