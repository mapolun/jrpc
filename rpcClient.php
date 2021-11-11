<?php
$message = 'ggg';
$atMobiles = [
    15517955313
];
/*******************************************
 * 格式1，自定义关键词机器人
********************************************/
$url = "https://oapi.dingtalk.com/robot/send?access_token=7afdc8ea25abd65cf3b0be1591563af2025cdef0db525a0ed716f2852af7e2f0";
$params=[
    "id"=>"1",
    "jsonrpc"=>"2.0",
    "method"=>"/News/DingDing",
    "params"=>['host'=>['address'=>$url,'secret'=>''], "msg" => $message, 'atmobiles'=>[$atMobiles]],
];
$resArr = send($params);
var_dump($resArr);
/*******************************************
 * 格式2，自定义机器人 + 秘钥
********************************************/
$url = "https://oapi.dingtalk.com/robot/send?access_token=638672a2e216dcc7ac6b23acb172216a190b1a889b085b0a2dd235072e2c98c2";
$secret = "SEC05eaae4431ec45b0d9769418f355b8aa86d73fefde58ddce8b07b703288d9b38";
$params=[
    "id"=>"1",
    "jsonrpc"=>"2.0",
    "method"=>"/News/DingDing",
    "params"=>['host'=>['address'=>$url,'secret'=>$secret], "msg" => $message, 'atmobiles'=>[$atMobiles]],
];
$resArr = send($params);
var_dump($resArr);


function send($params){
    $ip_api_url = 'http://192.168.2.160:34712';
    $response = post($ip_api_url,json_encode($params));
    if (!$response) {
        return [];
    }
    return json_decode($response,true);
}

function post($url,$data){
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 5);
    curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Content-Type: application/json;charset=utf-8'
        ]
    );
    curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
    $data = curl_exec($ch);
    curl_close($ch);
    return $data;
}