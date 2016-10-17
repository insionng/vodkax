// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket_test

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

// Vodka the data received on the WebSocket.
func VodkaServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// This example demonstrates a trivial vodka server.
func ExampleHandler() {
	http.Handle("/vodka", websocket.Handler(VodkaServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
