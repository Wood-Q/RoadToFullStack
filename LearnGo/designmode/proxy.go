package main

import "fmt"

//type Server interface {
//	HandleReq(url string) string
//}
//
//type RealServer struct {
//}
//
//func (r RealServer) HandleReq(url string) string {
//	handle := fmt.Sprintf("handle Req for %s\n", url)
//	return handle
//}
//
//type ProxyServer struct {
//	realServer Server
//}
//
//func (s ProxyServer) HandleReq(url string) string {
//	fmt.Println("using proxy server")
//	return s.realServer.HandleReq(url)
//}
//
//func main() {
//	proxyServer := ProxyServer{
//		realServer: &RealServer{},
//	}
//
//	response := proxyServer.HandleReq("http://google.com")
//	fmt.Println(response)
//}

type Server interface {
	HandleReq(req string) string
}

type RealServer struct {
}

func (server *RealServer) HandleReq(req string) string {
	return "handle：" + req
}

type proxyServer struct {
	realServer RealServer
}

func (s *proxyServer) HandleReq(req string) string {
	fmt.Println("服务启动！")
	return s.realServer.HandleReq(req)
}

func main() {
	server := &proxyServer{realServer: RealServer{}}
	req := "123"
	fmt.Println(server.HandleReq(req))
}
