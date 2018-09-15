<?php

$msg = "充：38万星币=100元《需要购买》加QQ：７０９２７７ＩＩ４《货到付款》";
//$msg = "你好哇主播，今天没抢到星币红包诶";
//$msg = "好听好听";
$SOCKET_FILE = "/tmp/spamcheck.sock";
$socket = socket_create(AF_UNIX, SOCK_STREAM, 0);
socket_connect($socket, $SOCKET_FILE);
socket_send($socket, $msg, strlen($msg), 0);
$response = socket_read($socket, 1024);
socket_close($socket);
var_dump($response);


