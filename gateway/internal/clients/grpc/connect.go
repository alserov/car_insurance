package grpc

import "google.golang.org/grpc"

func Dial(addr string) *grpc.ClientConn {
	cc, err := grpc.NewClient(addr)
	if err != nil {
		panic("failed to dial: " + err.Error())
	}

	return cc
}
