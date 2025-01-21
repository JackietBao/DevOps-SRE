<?php
// 数据库连接配置
$host = 'localhost';
$dbname = 'user_system';
$username = 'root';
$password = 'root';

try {
    $pdo = new PDO("mysql:host=$host;dbname=$dbname", $username, $password);
    $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
} catch(PDOException $e) {
    echo "连接失败: " . $e->getMessage();
}

// 开启session
session_start();
?> 