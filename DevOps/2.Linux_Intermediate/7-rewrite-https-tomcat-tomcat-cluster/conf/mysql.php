<?php
	$servername = "localhost";
	$username = "root";
	$password = "oldxu.com123";
	
	
	$conn = mysqli_connect($servername, $username, $password);

	if (!$conn) {
		die("连接失败:" . mysqli_connect_error());
	}

	echo "php connection mysql successfuly";
?>
