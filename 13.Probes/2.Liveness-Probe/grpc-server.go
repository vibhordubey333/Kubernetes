package main
import (
 "flag"
 "fmt"
 "log"
 "net"

 "google.golang.org/grpc"
 "google.golang.org/grpc/health"
 "google.golang.org/grpc/health/grpc_health_v1"
 "google.golang.org/grpc/reflection"
)

func main() {
 port := flag.Int("port", 8000, "The server port")
 host := flag.String("host", "", "The server host, perhaps 'localhost'")
 flag.Parse()

 s := grpc.NewServer()
 defer s.Stop()
 hs := health.NewServer() // will default to respond with SERVING
 grpc_health_v1.RegisterHealthServer(s, hs) // registration
 // register your own services
 reflection.Register(s)
 address := fmt.Sprintf("%s:%d", *host, *port)
 log.Default().Printf("starting service on  %s", address)
 listener, err := net.Listen("tcp", address)
 if err != nil {
   panic(err)
 }
 err = s.Serve(listener)
 if err != nil {
   panic(err)
 }
}
