package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/utils/buildversion"
	"github.com/the-prophet1/cni-example/pkg/log"
)

func Add(args *skel.CmdArgs) error {
	log.Write("进入到 cmdCheck")
	log.Write("这里的 CmdArgs 是: ", "容器ID: ", args.ContainerID)
	log.Write("网络命名空间: ", args.Netns)
	log.Write("网卡名称: ", args.IfName)
	log.Write("环境变量: ", args.Args)
	log.Write("路径: ", args.Path)
	log.Write("输入数据: ", string(args.StdinData))

	return nil
}

func Check(args *skel.CmdArgs) error {
	log.Write("进入到 cmdCheck")
	log.Write("这里的 CmdArgs 是: ", "容器ID: ", args.ContainerID)
	log.Write("网络命名空间: ", args.Netns)
	log.Write("网卡名称: ", args.IfName)
	log.Write("环境变量: ", args.Args)
	log.Write("路径: ", args.Path)
	log.Write("输入数据: ", string(args.StdinData))

	return nil
}
func Del(args *skel.CmdArgs) error {
	log.Write("进入到 cmdCheck")
	log.Write("这里的 CmdArgs 是: ", "容器ID: ", args.ContainerID)
	log.Write("网络命名空间: ", args.Netns)
	log.Write("网卡名称: ", args.IfName)
	log.Write("环境变量: ", args.Args)
	log.Write("路径: ", args.Path)
	log.Write("输入数据: ", string(args.StdinData))
	return nil
}

func main() {
	log.Write("hello cni")
	skel.PluginMain(Add, Check, Del, version.All, buildversion.BuildString("test-cni"))
}
