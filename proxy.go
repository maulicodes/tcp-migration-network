package main 
import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)
func main () {
listener,_ := net.Listen ("tcp", ":9988") // bind and listen
if err != nil {
	log.Fatalf("Cannot bind listener: %v",err)
}
fmt.Println ("Proxy Active on :9998..connectmthe client now")
migrationSignal := make (chan string)
for { //accept portion
clientConn,_ := listener.Accept()
go handleMigration(clientConn,"localhost:8080",migrationSignal )
go func(){
fmt.Println("migration starting")
migrationSignal <- "localhost:8080"
}()
}
}
func handleMigration (clientConn net.Conn, initialTarget string , sig chan string) {
var mu sync.Mutex
go func(){
	_,err := io.Copy(backendConn, clientConn)
	if err != nil {
		log.Println("Copy error C to S",err)
}
}()
currentTarget := initialTarget
backendConn, _ := net.Dial ("tcp", currentTarget)
if err != nil {
	log.Println("Cannot connect to backend",err)
	return
}
defer backendConn.close() 
newAddr := <-sig
fmt.Printf(">>> PROXY: SWITCHING TARFFIC TO %s\n", newAddr)
mu.Lock()
backendConn.Close()
backendConn,_ = net.Dial("tcp",newAddr)
if err != nil {
	log.Println ("Cannot switch backend",err)
	mu.Unlock()
	return
}
mu.Unlock()
fmt.Println ("Proxy:Handover complete.")
