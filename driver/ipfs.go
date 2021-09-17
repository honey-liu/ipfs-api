package driver

import (
	"bytes"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)


var Ipfs *shell.Shell

func InitIpfs()  {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	shId, err:= sh.ID()
	if err != nil {
		panic(err)
	}else {
		fmt.Printf("connect ipfs %s success!", shId)
	}
	Ipfs = sh
}

func UploadIPFS(data []byte) (string,error) {
	hash, err := Ipfs.Add(bytes.NewReader(data))
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
		return "", err
	}
	return hash,nil
}

func CatIPFS(hash string) string {
	read, err := Ipfs.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)
	return string(body)
}
