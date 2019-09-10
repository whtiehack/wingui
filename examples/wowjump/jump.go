package main

import "sync"

var mux = sync.Mutex{}

var running = false
var logouts []*Logout

var clicking = false
