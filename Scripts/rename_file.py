import os

def rename_files(directory):
    # 遍历指定目录下的所有文件
    for filename in os.listdir(directory):
        # 获取文件的完整路径
        filepath = os.path.join(directory, filename)
        # 如果是文件而不是目录
        if os.path.isfile(filepath):
            # 判断文件名中是否包含"_尚硅谷_"
            if "_尚硅谷_" in filename:
                # 将"_尚硅谷_"替换为空字符串
                new_filename = filename.replace("_尚硅谷_", "")
                # 构建新的文件路径
                new_filepath = os.path.join(directory, new_filename)
                # 重命名文件
                os.rename(filepath, new_filepath)
                print(f"文件 {filename} 重命名为 {new_filename}")

# 指定目录路径
directory_path = r"D:\GitHub\DevOps-SRE\Python+Go\Python\Python爬虫\代码"
# 调用函数进行重命名
rename_files(directory_path)