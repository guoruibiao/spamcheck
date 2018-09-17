<?php
// get机上没有代码环境，只能简单做下

class RedisHelper {
    public static $_instance = null;
    private function __construct(){
        $this->redis = new Redis();
        $this->redis->connect("192.168.xx.xx", 6379);
    }
    public static function getInstance() {
        if(self::$_instance == null) {
            self::$_instance = new RedisHelper();
        }
        return self::$_instance;
    }

    public function getRedisInstance() {
        return $this->redis;
    }
}

$key = "key";
$redis = RedisHelper::getInstance()->getRedisInstance();
$contents = $redis->zrevrange($key, 0, 7, true);
$sockfile = "/tmp/spamcheck.sock";

foreach($contents as $data=>$timestamp) {
    $data = json_decode($data, true);
    $content = $data["msgbody"];
    $resp = getClassifyResult($sockfile, $content);
    echo "{$content}      {$resp}.\n";
}

function getClassifyResult($sockfile, $content) {
    $socket = socket_create(AF_UNIX, SOCK_STREAM, 0);
    socket_connect($socket, $sockfile);
    socket_send($socket, $content, strlen($content), 0);
    $response = socket_read($socket, 1024);
    socket_close($socket);
    return $response;
}

function sendMsgToDingDing($msg="", $webhook="", $phones=array(), $isAtAll=false) {
    $payload = array(
        "msgtype" => "text",
        "text" => array(
            "content" => $msg,
        ),
        "at" => array(
            "atMobiles" => $phones,
            "isAtAll" => $isAtAll,
        ),
    );
    $result = httpPost($webhook, json_encode($payload));
}

function httpPost( $url, $post = '', $timeout = 5   ){
    if(empty($url)){
        return ;
    }
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_HEADER, 0);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);

    if( $post != '' && !empty( $post   )   ){
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $post);
        curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type: application/json', 'Content-Length: ' . strlen($post)));
    }
    curl_setopt($ch, CURLOPT_TIMEOUT, $timeout);
    $result = curl_exec($ch);
    curl_close($ch);
    return $result;

}
