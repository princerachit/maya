/*
Copyright 2018 The OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cstor

import (
	"context"
	"fmt"
	"log"

	"github.com/openebs/maya/cmd/cstor-volume-grpc/api"
	"google.golang.org/grpc"
)

var (
	cmdName = "cstor-volume-grpc"
	usage   = fmt.Sprintf("%s", cmdName)
)

// CmdSnaphotOptions holds the options for snapshot
// create command
type CmdSnaphotOptions struct {
	volName  string
	snapName string
}

//createSnapshot creates snapshots
func createSnapshot(volName, snapName, ip string) (*api.VolumeCommand, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, api.VolumeGrpcListenPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewRunCommandClient(conn)
	response, err := c.RunVolumeCommand(context.Background(),
		&api.VolumeCommand{
			Command:  api.CmdSnapCreate,
			Volume:   volName,
			Snapname: snapName,
			Status:   "requesting",
		})

	if err != nil {
		log.Fatalf("Error when calling RunVolumeCommand: %s", err)
	}
	log.Printf("Response from server: %s, %s, %s, %s",
		response.Command, response.Volume, response.Snapname, response.Status)
	return response, err
}

//destroySnapshot destroys snapshots
func destroySnapshot(volName, snapName, ip string) (*api.VolumeCommand, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, api.VolumeGrpcListenPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewRunCommandClient(conn)
	response, err := c.RunVolumeCommand(context.Background(),
		&api.VolumeCommand{
			Command:  api.CmdSnapDestroy,
			Volume:   volName,
			Snapname: snapName,
			Status:   "requesting",
		})

	if err != nil {
		log.Fatalf("Error when calling RunVolumeCommand: %s", err)
	}
	log.Printf("Response from server: %s, %s, %s, %s",
		response.Command, response.Volume, response.Snapname, response.Status)
	return response, err
}
