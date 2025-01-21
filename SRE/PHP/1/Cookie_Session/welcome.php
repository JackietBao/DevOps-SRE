<?php
require_once 'config.php';

// 检查是否登录
if (!isset($_SESSION['user_id']) && !isset($_COOKIE['user_id'])) {
    header("Location: login.php");
    exit();
}

// 如果session过期但cookie存在，恢复session
if (!isset($_SESSION['user_id']) && isset($_COOKIE['user_id'])) {
    $_SESSION['user_id'] = $_COOKIE['user_id'];
    $_SESSION['username'] = $_COOKIE['username'];
}

$username = $_SESSION['username'];
?>

<!DOCTYPE html>
<html>
<head>
    <title>欢迎页面</title>
    <meta charset="UTF-8">
</head>
<body>
    <h2>欢迎回来，<?php echo htmlspecialchars($username); ?>！</h2>
    <a href="logout.php">退出登录</a>
</body>
</html>