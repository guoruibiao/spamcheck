<?php
// 101上没有代码环境，只能简单做下

class RedisHelper {
    public static $_instance = null;
    private function __construct(){
        $this->redis = new Redis();
        $this->redis->connect("192.168.32.103", 6379);
    }
    public static function getInstance() {
        if(self::$_instance == null) {
            self::$_instance = new RedisHelper();
        }
        return self::$_instance;
    }
}


$rs = RedisHelper::getInstance();
$key = "publicchatmsgqueue";

$ret = $rs->zrevrange($key, 0, 1, true);
var_dump($ret);
