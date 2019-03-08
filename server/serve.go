package server

import (
	"net"
	"net/http"
)

type Server interface {


}


type lzServer struct {

	*http.Server
	ln net.Listener


}


func hello(){



}