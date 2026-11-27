package urlParsing

import (
	"fmt"
	"net"
	"net/url"
)

func GoUrlPrasing(){
	fmt.Println("Go url parsing is running")
}

func UrlParsingCaller(){
  s := "postgres://user:pass@host.com:5432/path?k=v#f"
  u, err := url.Parse(s)
  if err != nil{
	panic(err)
  }
  fmt.Printf("Type %t\n",u)

  fmt.Println(u.Scheme) // returns the scheme of the url
  fmt.Println(u.User) // returns both username and password of url
  fmt.Println("Username:", u.User.Username()) // returns the username of the user
  p, isTrue := u.User.Password() // returns the password and boolean value
  fmt.Println("Password", p, isTrue) 
  fmt.Println(u.Host) // returns both hostname and port together
  host, port, err := net.SplitHostPort(u.Host) // returns host, port and err message(if any)
  if err != nil{
	panic(err)
  }
  fmt.Println("Host:", host, "\n", "Port:", port)
  fmt.Println("Path:", u.Path) // returns the path of the url
  fmt.Println("Fragment:", u.Fragment) // returns fragments after #
  fmt.Println("Query Params:", u.RawQuery)
  m, _ := url.ParseQuery(u.RawQuery) // retunrs a map containing key and value of query params
  fmt.Println(m)
  fmt.Println(m["k"][0])
}