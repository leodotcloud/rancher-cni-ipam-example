package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/rancher/rancher-cni-ipam/ipfinder/metadata"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		logrus.Errorf("Invalid arguments, need just one argument which is container id")
		os.Exit(1)
	}
	cid := os.Args[1]

	logrus.Infof("cid: %v", cid)

	ipf, err := metadata.NewIPFinderFromMetadata()
	if err != nil {
		logrus.Errorf("err: %v", err)
		os.Exit(1)
	}

	ip := ipf.GetIP(cid)

	logrus.Infof("ip: %v", ip)
}
