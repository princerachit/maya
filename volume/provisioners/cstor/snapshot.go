package cstor

import (
	"github.com/golang/glog"
	"github.com/openebs/maya/cmd/cstor-volume-grpc/api"
	"github.com/openebs/maya/cmd/cstor-volume-grpc/app/command"
)

func CreateSnapshot(volName, snapName, targetIP string) (*api.VolumeCommand, error) {
	glog.Infof("cStor.CreateSnapshot called volName:%s,snapName:%s, ip:%s", volName, snapName,targetIP)
	return command.CreateSnapshot(volName, snapName, targetIP)
}

func DeleteSnapshot(volName, snapName, targetIP string) (*api.VolumeCommand, error) {
	glog.Infof("cStor.DeleteSnapshot volName:%s,snapName:%s", volName, snapName)
	return command.DestroySnapshot(volName, snapName, targetIP)
}
