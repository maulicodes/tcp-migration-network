   package main                                                                                                                                                                                                                             
   import (
   "fmt"
   "os"
   "bufio"
   "net"
   "strconv"
   "strings"
   "math/rand"
   "time"
  )
  func main () {
  arguements  := os.Args
  if len (arguements) != 2  {
  fmt.Println ("Provide a port number")
  return 
  }
  port := ":" + arguements[1]
  l, err := net.Listen ( "tcp", port )
  if err != nil {
  fmt.Println (err)
  return 
  }
  
  defer l.Close()
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Listening port",port)
  for {
  c,err := l.Accept()
  if err != nil {
  fmt.Println (err) 
  continue 
  }
  go handleConnection (c)
  }
  }
  
  func handleConnection(c net.Conn) {
  fmt.Printf("Serving %s\n", c.RemoteAddr().String())
  reader := bufio.NewReader(c)
  
         for {
                  netData, err := reader.ReadString('\n')
                  if err != nil {
                         fmt.Println(err)
                          return
                }
                temp := strings.TrimSpace(netData)
                  if temp == "STOP" {
                          fmt.Println("closing mechanism")
                         return
                }
                  result := strconv.Itoa(rand.Intn(1000)) + "\n"
                  c.Write([]byte(result))
          }
       }
