## Overview
This assignment was done in WSL 2 using Golang. The server's address is **http://localhost:8000/** The unit test cases have a 59.2% statement coverage.

There is a config.json which contains information regarding which API url to retrieve user information and the file directory messages should be stored in.

After making and running the executable, the user needs to submit a GET or POST request containing the username and password. If it is a POST request then the user can also send a message less than or equal to 256 characters long.

The server will then give a response back to the user via a HTTP response with a respective status code along with more information in the response body. For example if a user successfully submits a POST request to write a message, then the server will respond with a status code of 201 along with a JSON body with the attached message. If there are any unacceptable user requests, the body will contain additional information.

Overall this assignment was straightforward. One of the biggest challenges for me was getting used to how Golang parsed JSON from the response body. In Python, it is very simple with the help of libraries JSON responses can be directly translated into a dictionary and their values accessed by keys. In Golang its more cumbersome because I had to first make a struct as a data structure to fit the responses that I would get. I definitely wanted to use Golang since that is the focus of this position and wanted to be ready. However, after this assignment I remembered more things about Go and learned some new tricks, especially its unit testing feature. I know some things are not necessary good conventional implementations, such as iterating through every user to validate credentials rather than having a map or requiring the user to have a body in their GET request. These things can be fixed in little time, but I wanted to focus more on writing some test cases.
## Reading Messages
To read a message, the user submits a GET request with a body containing their credentials in JSON.
```
{"username": "*<username>*", "password": "*<password>*", "message" : ""}
```
In curl it would look something like this:
```
curl -v -X GET -H "Content-Type: application/json" \
    -d '{"username": "tracey.ramos@reqres.in", "password": "password", "message" : ""}' \
    localhost:8000
```

## Writing Messages
To write a message, the user submits a POST request with a body containing their credentials and their message in JSON
```
{"username": "*<username>*", "password": "*<password>*", "message" : "*<message>*"}
```
In curl it would look something like this:
```
curl -v -X POST -H "Content-Type: application/json" \
    -d '{"username": "emma.wong@reqres.in", "password": "password", "message" : "Hello, this is my message."}' \
    localhost:8000
```

## Building
To run:

```
make build
chmod +x ./server
./server
```

To just run unit tests:
```
make test
```

Cleanup:
```
make clean
```

## Example
First we start the server after building:
```
./server
```
Then we use curl to make a request to our local server:
```
curl -v -X POST -H "Content-Type: application/json" \
    -d '{"username": "tracey.ramos@reqres.in", "password": "password", "message" : "Hey, I got your message"}' \
    localhost:8000
```

we would get:
```
*   Trying 127.0.0.1:8000...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8000 (#0)
> POST / HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.68.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 82
>
* upload completely sent off: 82 out of 82 bytes
* Mark bundle as not supporting multiuse
< HTTP/1.1 201 Created
< Content-Type: application/json
< Date: Wed, 31 Mar 2021 20:14:47 GMT
< Content-Length: 19
<
* Connection #0 to host localhost left intact
{"message": "cool"}
```