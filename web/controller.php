<?php

$go_websocket_api = "localhost/ws/pub/a";

$data = $_POST['data'];

// send a POST request to the go websocket server
$ch = curl_init($go_websocket_api);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($ch);
curl_close($ch);

echo $response;

?>