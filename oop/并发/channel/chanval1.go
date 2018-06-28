package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)
