<?php 
$redis = new Redis();

//Connecting to Redis
$redis->connect('redis');
 
// "url":"https://httpbin.org/get?title={mascot}&image={location}&foo={bar}"
$message = <<<JSON
{
  "endpoint":{
    "method":"GET", 
    "url": "https://postman-echo.com/delay/5"
  }, 
  "data":[
    {
      "mascot":"Gopher", 
      "location":"https://blog.golang.org/gopher/gopher.png"
    } 
  ]
}
JSON;


for ($i=0; $i<10; $i++) {
 
  
  $redis->publish(
    'postback-channel', $message
  );
  
}

if ($redis->ping()) {
  echo "OK";
}

