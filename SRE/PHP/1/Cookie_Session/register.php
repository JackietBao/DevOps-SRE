<?php
require_once 'config.php';

if ($_SERVER['REQUEST_METHOD'] == 'POST') {
    // 净化输入数据
    $username = htmlspecialchars(trim($_POST['username']));
    $encrypted_password = trim($_POST['encrypted_password']);
    $email = filter_var(trim($_POST['email']), FILTER_SANITIZE_EMAIL);
    
    // 解密前端加密的密码
    $decrypted = openssl_decrypt(
        base64_decode($encrypted_password),
        'AES-256-CBC',
        'your-secret-key',
        OPENSSL_RAW_DATA,
        // 如果使用 IV，需要添加相应的 IV
    );
    
    // 输入验证
    $errors = [];
    if (strlen($username) < 3) {
        $errors[] = "用户名至少需要3个字符";
    }
    if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
        $errors[] = "邮箱格式不正确";
    }
    
    if (empty($errors)) {
        // 检查用户名是否已存在
        $stmt = $pdo->prepare("SELECT * FROM users WHERE username = ?");
        $stmt->execute([$username]);
        
        if ($stmt->rowCount() > 0) {
            $errors[] = "用户名已存在！";
        } else {
            // 对解密后的原始密码进行 password_hash
            $password_hash = password_hash($decrypted, PASSWORD_DEFAULT);
            
            // 插入新用户
            $stmt = $pdo->prepare("INSERT INTO users (username, password, email) VALUES (?, ?, ?)");
            if ($stmt->execute([$username, $password_hash, $email])) {
                header("Location: login.php");
                exit();
            } else {
                $errors[] = "注册失败，请稍后重试";
            }
        }
    }
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>用户注册</title>
    <meta charset="UTF-8">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/4.1.1/crypto-js.min.js"></script>
</head>
<body>
    <h2>用户注册</h2>
    <?php 
    if (!empty($errors)) {
        foreach ($errors as $error) {
            echo "<p style='color: red'>" . htmlspecialchars($error) . "</p>";
        }
    }
    ?>
    <form method="post" onsubmit="return encryptFormData()">
        用户名：<input type="text" name="username" minlength="3" required><br>
        密码：<input type="password" name="password" id="password" minlength="6" required><br>
        <input type="hidden" name="encrypted_password" id="encrypted_password">
        邮箱：<input type="email" name="email" required><br>
        <input type="submit" value="注册">
    </form>

    <script>
    function encryptFormData() {
        var password = document.getElementById('password').value;
        // 使用AES加密密码
        var encrypted = CryptoJS.AES.encrypt(password, 'your-secret-key').toString();
        document.getElementById('encrypted_password').value = encrypted;
        document.getElementById('password').value = '';
        return true;
    }
    </script>
</body>
</html> 