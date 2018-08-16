package cstor

import (
	"github.com/golang/glog"
	"github.com/openebs/maya/cmd/cstor-volume-grpc/api"
)

func CreateSnapshot(volName, snapName, targetIP string) (*api.VolumeCommand, error) {
	glog.Infof("cStor.CreateSnapshot called volName:%s,snapName:%s, ip:%s", volName, snapName, targetIP)
	return createSnapshot(volName, snapName, targetIP)
}

func DeleteSnapshot(volName, snapName, targetIP string) (*api.VolumeCommand, error) {
	glog.Infof("cStor.DreateSnapshot called volName:%s,snapName:%s, ip:%s", volName, snapName, targetIP)
	return destroySnapshot(volName, snapName, targetIP)
}
