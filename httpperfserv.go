package main

import (
	"log"
	"regexp"
	"strconv"
//	"math/rand"
	"hash/fnv"
	"net/http"
)

var validPath = regexp.MustCompile("^/([0-9]+)/([a-zA-Z0-9]+)$")

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var randomString = RandStringBytesRmndr(1048576*64)


func hash(s string) uint64 {
        h := fnv.New64a()
        h.Write([]byte(s))
        return h.Sum64()
}

func RandStringBytesRmndr(n int) string {
    b := make([]byte, n)
    var seed uint64 = 0
    for i := range b {
	seed = seed*16807+1
        b[i] = letterBytes[seed % uint64(len(letterBytes))]
    }
    return string(b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	m:=validPath.FindStringSubmatch(r.URL.Path)
	maxs:=uint64(len(randomString))
	if m==nil {
		http.NotFound(w,r)
		return 
	}
	size, err:=strconv.Atoi(m[1])
	if  err==nil && size>=0 {
	size64:=uint64(size)

//	log.Printf("Got request  for %s bytes interpreted as %v, name %s, seed value %v\n", m[1], size, m[2], hashv)
	if size64>maxs { size64=maxs }
	hashv:=hash(m[2]) % uint64(maxs)
	w.Header().Set("Content-Type", "text/plain")
	if (hashv+size64)<maxs {
		w.Write([]byte(randomString[hashv:hashv+size64]))	
	} else
	{
		w.Write([]byte(randomString[hashv:]))	
		w.Write([]byte(randomString[0:hashv+size64-maxs]))	

	}
	
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 8080. Go to http://localhost:8080/size/name, where size is a positive number")
//	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	err := http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}
