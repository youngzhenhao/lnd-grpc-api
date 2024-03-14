package main

import (
	"fmt"
	"lnd-grpc-api/api"
)

func main() {
	//fmt.Printf("%v\n", api.StructToJsonString(api.OpenChannel(0, "", 0, 0, 0, 0)))

	fmt.Printf("%v\n", api.StructToJsonString(api.GetInfo()))

}
