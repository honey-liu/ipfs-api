package main

import (
	"fmt"
	"ipfs_api/driver"
)

func main() {
	fmt.Println("ipfs-api server starting...")
	driver.InitIpfs()
	InitApp()
}
