  package main                                                                                                                                                                                                                             
   import (
   "net"
   "net/url"
   "os"
   )
   type Dialer interface {
   Dial (network, addr string) ( net.Conn ,  error)
   }
  type Auth struct {
  User, Password string
  }
  func main() {
  }
  func FromEnvironmentUsing(forward Dialer) Dialer {
  proxy := os.Getenv("HTTP_PROXY")
  if len(proxy)== 0 {
  return forward
  }
  proxyURL, err := url.Parse(proxy)
  if err != nil {
  return forward
  }
  _, err = FromURL(proxyURL,forward)
  if err != nil {
  return forward
  }
  return forward
  }
  func FromURL(u *url.URL,forward Dialer) (Dialer,error) {
  var auth *Auth
  if u.User != nil {
  auth = new(Auth)
  auth.User = u.User.Username ()
  if p , ok := u.User.Password (); ok {
  auth.Password = p
  }
  }
  switch u.Scheme {
  case  "socks5", "socks5h" :
  addr := u.Hostname()
  port := u.Port()
  if port == "" {
  port = "1080"
  }
  _ = net.JoinHostPort(addr,port)
  _=auth
  return forward,nil
  }
  return forward,nil
  }
