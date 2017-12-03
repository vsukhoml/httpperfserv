# httpperfserv

This is simple Origin server for performance testing of caching proxies and CDNs

## Use

go run

## Requests

web server understands requests in the format:

http://server:8080/SIZE/NAME

where:
* SIZE - size of file to produce
* NAME - any name. it is used as a seed value for content of file. Same NAME will result in same content

## Example

curl http://localhost:8080/1024/test

will produce a 1024 byte file

