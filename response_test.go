package icapclient

import (
	"testing"
)

func TestResponse(t *testing.T) {

	// t.Run("ReadResponse", func(t *testing.T) {
	// 	respStr := `ICAP/1.0 200 OK
	//  Date: Mon, 10 Jan 2000  09:55:21 GMT
	//  Server: ICAP-Server-Software/1.0
	//  Connection: close
	//  ISTag: "W3E4R7U9-L2E4-2"
	//  Encapsulated: res-hdr=0, res-body=213
	//
	//  HTTP/1.1 403 Forbidden
	//  Date: Wed, 08 Nov 2000 16:02:10 GMT
	//  Server: Apache/1.3.12 (Unix)
	//  Last-Modified: Thu, 02 Nov 2000 13:51:37 GMT
	//  ETag: "63600-1989-3a017169"
	//  Content-Length: 58
	//  Content-Type: text/html
	//
	//  3a
	//  Sorry, you are not allowed to access that naughty content.
	//  0`
	//
	// 	resp, err := ReadResponse(bufio.NewReader(strings.NewReader(respStr)))
	//
	// 	if err != nil {
	// 		t.Fatal(err.Error())
	// 	}
	//
	// 	spew.Dump(resp)
	//
	// })

}
