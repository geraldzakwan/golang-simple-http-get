# golang-simple-http-get

## How to Run

- To start the web server, run the following command:

`go run main.go`

- If it boots succesfully, you should see the below prompt:

```
2021/04/28 17:16:31 Data is succesfully loaded
2021/04/28 17:16:31 ID:  1 Name:  A
2021/04/28 17:16:31 ID:  2 Name:  B
2021/04/28 17:16:31 ID:  3 Name:  C
2021/04/28 17:16:31 Starting server at port 8080
```

## How to Run Automated Test

- To do an automated test, run the following command:

`go test`

- If all test pass, you should see the result below:

```
2021/04/28 17:23:36 Data is succesfully loaded
2021/04/28 17:23:36 ID:  1 Name:  A
2021/04/28 17:23:36 ID:  2 Name:  B
2021/04/28 17:23:36 ID:  3 Name:  C
2021/04/28 17:23:36 Data is succesfully loaded
2021/04/28 17:23:36 ID:  1 Name:  A
2021/04/28 17:23:36 ID:  2 Name:  B
2021/04/28 17:23:36 ID:  3 Name:  C
2021/04/28 17:23:36 Data is succesfully loaded
2021/04/28 17:23:36 ID:  1 Name:  A
2021/04/28 17:23:36 ID:  2 Name:  B
2021/04/28 17:23:36 ID:  3 Name:  C
2021/04/28 17:23:36 Data is succesfully loaded
2021/04/28 17:23:36 ID:  1 Name:  A
2021/04/28 17:23:36 ID:  2 Name:  B
2021/04/28 17:23:36 ID:  3 Name:  C
2021/04/28 17:23:36 Data is succesfully loaded
2021/04/28 17:23:36 ID:  1 Name:  A
2021/04/28 17:23:36 ID:  2 Name:  B
2021/04/28 17:23:36 ID:  3 Name:  C
PASS
ok  	golang-simple-get-http	0.523s
```

## How to Run Manual Test

- To test manually, below are five examples using `curl` (and the response) for each test case:

1. `curl -i localhost:8080`

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 28 Apr 2021 10:29:33 GMT
Content-Length: 81

{"code":200,"data":[{"id":1,"name":"A"},{"id":2,"name":"B"},{"id":3,"name":"C"}]}
```

2. `curl -i localhost:8080?id=2`

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 28 Apr 2021 10:29:56 GMT
Content-Length: 41

{"code":200,"data":[{"id":2,"name":"B"}]}
```

3. `curl -i localhost:8080?id=1,3,4`

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 28 Apr 2021 10:30:34 GMT
Content-Length: 61

{"code":200,"data":[{"id":1,"name":"A"},{"id":3,"name":"C"}]}
```

4. `curl -i localhost:8080?id=xxx`

```
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 28 Apr 2021 10:31:13 GMT
Content-Length: 53

{"code":400,"message":"Invalid or empty ID: \"xxx\""}
```

5. `curl -i localhost:8080?id=4`

```
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 28 Apr 2021 10:31:46 GMT
Content-Length: 58

{"code":404,"message":"Resource with ID: 4 doesn't exist"}
```
