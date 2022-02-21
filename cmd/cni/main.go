package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/utils/buildversion"
	"github.com/the-prophet1/cni-example/pkg/log"
)

func Add(args *skel.CmdArgs) error {
	log.Write("进入到 cmdAdd")
	log.Write(
		"这里的 CmdArgs 是: ", "ContainerID: ", args.ContainerID,
		"Netns: ", args.Netns,
		"IfName: ", args.IfName,
		"Args: ", args.Args,
		"Path: ", args.Path,
		"StdinData: ", string(args.StdinData))

	return nil
}

func Check(args *skel.CmdArgs) error {
	log.Write("进入到 cmdCheck")
	log.Write(
		"这里的 CmdArgs 是: ", "ContainerID: ", args.ContainerID,
		"Netns: ", args.Netns,
		"IfName: ", args.IfName,
		"Args: ", args.Args,
		"Path: ", args.Path,
		"StdinData: ", string(args.StdinData))
	return nil
}
func Del(args *skel.CmdArgs) error {
	log.Write("进入到 cmdDel")
	log.Write(
		"这里的 CmdArgs 是: ", "ContainerID: ", args.ContainerID,
		"Netns: ", args.Netns,
		"IfName: ", args.IfName,
		"Args: ", args.Args,
		"Path: ", args.Path,
		"StdinData: ", string(args.StdinData))
	return nil
}

func main() {
	log.Write("hello cni")
	skel.PluginMain(Add, Check, Del, version.All, buildversion.BuildString("test-cni"))
}
