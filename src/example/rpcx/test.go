package rpcx

import (
	"github.com/smallnest/rpcx/client"
	"flag"
)

var (
	consulAddr = flag.String("consulAddr","localhost:8500","consul service")
	basePath = flag.String("basePath","/test_rpcx","base path")
)
func main(){

	flag.Parse()

	d := client.NewConsulDiscovery(*basePath,"Arith",[]string{*consulAddr}, nil)
	client.NewXClient("Arith",client.Failtry,client.RandomSelect,d,client.DefaultOption)
}
