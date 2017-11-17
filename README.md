"# OnlineTestGo" 


run following command to after cloning

$> go run main.go

you should be able to hit the web service using curl/ Rest client

Sample Request

URL : hitting http://<address>:<port>/registerUser

Header :
accept: application/json
accept-encoding: gzip, deflate
accept-language: en-US,en;q=0.8
content-type: application/json
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36


payload : 

{
  "fname":"yourfirstname",
  "lname":"yourlastname",
  "phone":"123456789",
  "email":"somone@something.com",
  "test":"Java"
}# OnlineTestgo
