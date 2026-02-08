package main

import (
	"fmt"
	"time"
)

func main() {
	// ensure payment service stages happen in order
	cfgLoaded := make(chan struct{})
	go stage("load payment config", cfgLoaded)
	<-cfgLoaded

	gatewayReady := make(chan struct{})
	go func() {
		<-cfgLoaded
		stage("connect to gateway", gatewayReady)
	}()
	<-gatewayReady

	cacheWarmed := make(chan struct{})
	go func() {
		<-gatewayReady
		stage("warm risk cache", cacheWarmed)
	}()
	<-cacheWarmed

	fmt.Println("payment service startup complete")
}

// stage simulates a startup step and signals completion on a channel.
func stage(name string, done chan<- struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println(name, "done")
	close(done)
}
