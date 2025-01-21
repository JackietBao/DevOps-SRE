<?php
require_once 'config.php';

if ($_SERVER['REQUEST_METHOD'] == 'POST') {
    // 净化输入数据
    $username = htmlspecialchars(trim($_POST['username']));
    $encrypted_password = trim($_POST['encrypted_password']);
    
    // 解密前端加密的密码
    $decrypted = openssl_decrypt(
        base64_decode($encrypted_password),
        'AES-256-CBC',
        'your-secret-key',
        OPENSSL_RAW_DATA,
        // 如果使用 IV，需要添加相应的 IV
    );
    
    // 验证用户
    $stmt = $pdo->prepare("SELECT * FROM users WHERE username = ?");
    $stmt->execute([$username]);
    $user = $stmt->fetch();
    
    if ($user && password_verify($decrypted, $user['password'])) {
        // 登录成功，设置session
        session_regenerate_id(true); // 防止会话固定攻击
        $_SESSION['user_id'] = $user['id'];
        $_SESSION['username'] = $user['username'];
        
        // 设置cookie，保持登录状态30天
        if (isset($_POST['remember'])) {
            // 生成安全的记住我令牌
            $token = bin2hex(random_bytes(32));
            $token_hash = hash('sha256', $token);
            
            // 存储令牌到数据库（需要在users表添加remember_token字段）
            $stmt = $pdo->prepare("UPDATE users SET remember_token = ? WHERE id = ?");
            $stmt->execute([$token_hash, $user['id']]);
            
            // 设置安全的cookie
            setcookie('remember_token', $token, [
                'expires' => time() + 30 * 24 * 60 * 60,
                'path' => '/',
                'httponly' => true,
                'secure' => true,
                'samesite' => 'Strict'
            ]);
        }
        
        header("Location: welcome.php");
        exit();
    } else {
        $error = "用户名或密码错误！";
    }
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>用户登录</title>
    <meta charset="UTF-8">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/4.1.1/crypto-js.min.js"></script>
</head>
<body>
    <h2>用户登录</h2>
    <?php if (isset($error)) echo "<p style='color: red'>" . htmlspecialchars($error) . "</p>"; ?>
    <form method="post" onsubmit="return encryptFormData()">
        用户名：<input type="text" name="username" required><br>
        密码：<input type="password" name="encrypted_password" id="encrypted_password" required><br>
        <input type="checkbox" name="remember"> 记住我<br>
        <input type="submit" value="登录">
    </form>

    <script>
    function encryptFormData() {
        var password = document.getElementById('encrypted_password').value;
        // 使用AES加密密码
        var encrypted = CryptoJS.AES.encrypt(password, 'your-secret-key').toString();
        document.getElementById('encrypted_password').value = encrypted;
        document.getElementById('password').value = '';
        return true;
    }
    </script>
</body>
</html>