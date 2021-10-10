<?PHP
$url = "http://127.0.0.1:8080/";
for ($i = 1; $i <= 10; $i++) {
  $ch = curl_init();
  curl_setopt($ch,CURLOPT_URL,$url);

  if(curl_exec($ch)){ // ?? - if request and data are completely received
    continue; // ?? - go to the next loop
  }
  // DONT go to the next loop until the above data is complete or returns true
}
?>