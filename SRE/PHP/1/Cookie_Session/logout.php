<?php
require_once 'config.php';

// 清除session
session_destroy();

// 清除cookie
setcookie('user_id', '', time() - 3600);
setcookie('username', '', time() - 3600);

header("Location: login.php");
exit();
?>