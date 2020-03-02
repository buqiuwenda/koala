package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

type GrpcGenerator struct {
}

func (d *GrpcGenerator) Run(opt *Option, metaData *ServiceMetaData) (err error) {

	//protoc --go_out=plugins=grpc:. hello.proto
	dir := path.Join(opt.Output, "generate", path.Join(metaData.serviceNameParts...))
	os.MkdirAll(dir, 0755)
	outputParams := fmt.Sprintf("plugins=grpc:%s", dir)

	cmd := exec.Command("protoc", "--gofast_out", outputParams, opt.Proto3Filename)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf("grpc generator failed, err:%v\n", err)
		return

	}
	return
}

func init() {
	gc := &GrpcGenerator{}

	RegisterServerGenerator("grpc generator", gc)
	RegisterClientGenerator("grpc generator", gc)
}
