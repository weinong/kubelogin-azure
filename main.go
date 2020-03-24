package main

import(
	"flag"
	"os"
	"github.com/spf13/pflag"
	"k8s.io/klog"
	"github.com/weinong/kubelogin-azure/pkg/cmd"
)

func main() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlag(flag.CommandLine.Lookup("v"))
	pflag.CommandLine.AddGoFlag(flag.CommandLine.Lookup("logtostderr"))
	pflag.CommandLine.Set("logtostderr", "true")
	root := cmd.NewRootCmd(v.String())
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
