<?php
if(isset($_POST["submit"])) {
    // 设置上传文件的目标目录
    $target_dir = "uploads/";
    // 获取上传文件的完整路径
    $target_file = $target_dir . basename($_FILES["fileToUpload"]["name"]);
    // 获取文件类型
    $imageFileType = strtolower(pathinfo($target_file,PATHINFO_EXTENSION));
    // 标记是否可以上传
    $uploadOk = 1;

    // 检查文件是否已经存在
    if (file_exists($target_file)) {
        echo "抱歉，文件已经存在。";
        $uploadOk = 0;
    }

    // 检查文件大小（这里限制为5MB）
    if ($_FILES["fileToUpload"]["size"] > 5000000) {
        echo "抱歉，文件太大。";
        $uploadOk = 0;
    }

    // 限制文件类型（这里以图片为例）
    $allowed_types = array("jpg", "jpeg", "png", "gif");
    if (!in_array($imageFileType, $allowed_types)) {
        echo "抱歉，只允许上传 JPG, JPEG, PNG 和 GIF 文件。";
        $uploadOk = 0;
    }

    // 如果所有检查都通过，尝试上传文件
    if ($uploadOk == 1) {
        if (move_uploaded_file($_FILES["fileToUpload"]["tmp_name"], $target_file)) {
            echo "文件 ". basename( $_FILES["fileToUpload"]["name"]). " 上传成功。";
        } else {
            echo "抱歉，上传文件时发生错误。";
        }
    }
}
?> 