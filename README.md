# Simple http server

A simple http server by go

## Example

    $ curl -v http://localhost:8080/foo/bar
    *   Trying ::1...
    * Connected to localhost (::1) port 8080 (#0)
    > GET /foo/bar HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.47.0
    > Accept: */*
    >
    < HTTP/1.1 200 OK
    < Date: Mon, 13 Nov 2017 07:46:56 GMT
    < Content-Length: 15
    < Content-Type: text/plain; charset=utf-8
    <
    Hello foo bar!
    * Connection #0 to host localhost left intact

## Implementation

Framework: `net/http` package of Golang

Since this server is too simple, I choose the built-in package of golang.

## ab Test

    $ ab -n 1000000 -c 1000 http://localhost:8080/foo/bar
    This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 100000 requests
    Completed 200000 requests
    Completed 300000 requests
    Completed 400000 requests
    Completed 500000 requests
    Completed 600000 requests
    Completed 700000 requests
    Completed 800000 requests
    Completed 900000 requests
    Completed 1000000 requests
    Finished 1000000 requests


    Server Software:
    Server Hostname:        localhost
    Server Port:            8080

    Document Path:          /foo/bar
    Document Length:        15 bytes

    Concurrency Level:      1000
    Time taken for tests:   32.112 seconds
    Complete requests:      1000000
    Failed requests:        0
    Total transferred:      132000000 bytes
    HTML transferred:       15000000 bytes
    Requests per second:    31140.73 [#/sec] (mean)
    Time per request:       32.112 [ms] (mean)
    Time per request:       0.032 [ms] (mean, across all concurrent requests)
    Transfer rate:          4014.23 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0   19  87.6     11    1054
    Processing:     0   13  15.1     12     417
    Waiting:        0    9  14.8      8     415
    Total:          0   32  89.9     23    1437

    Percentage of the requests served within a certain time (ms)
      50%     23
      66%     27
      75%     30
      80%     31
      90%     32
      95%     36
      98%     43
      99%    225
     100%   1437 (longest request)
