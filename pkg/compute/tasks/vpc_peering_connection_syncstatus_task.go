// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tasks

import (
	"context"
	"fmt"

	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/cloudcommon/db/taskman"
	"yunion.io/x/onecloud/pkg/compute/models"
	"yunion.io/x/onecloud/pkg/util/logclient"
)

type VpcPeeringConnectionSyncstatusTask struct {
	taskman.STask
}

func init() {
	taskman.RegisterTask(VpcPeeringConnectionSyncstatusTask{})
}

func (self *VpcPeeringConnectionSyncstatusTask) taskFail(ctx context.Context, peer *models.SVpcPeeringConnection, msg jsonutils.JSONObject) {
	peer.SetStatus(self.UserCred, api.VPC_PEERING_CONNECTION_STATUS_UNKNOWN, msg.String())
	db.OpsLog.LogEvent(peer, db.ACT_SYNC_STATUS, msg, self.GetUserCred())
	logclient.AddActionLogWithStartable(self, peer, logclient.ACT_SYNC_STATUS, msg, self.UserCred, false)
	self.SetStageFailed(ctx, msg)
}

func (self *VpcPeeringConnectionSyncstatusTask) OnInit(ctx context.Context, obj db.IStandaloneModel, data jsonutils.JSONObject) {
	peer := obj.(*models.SVpcPeeringConnection)

	svpc, err := peer.GetVpc()
	if err != nil {
		msg := fmt.Sprintf("fail to find vpc for VpcPeeringConnection %s", err)
		self.taskFail(ctx, peer, jsonutils.NewString(msg))
		return
	}

	extVpc, err := svpc.GetIVpc()
	if err != nil {
		msg := fmt.Sprintf("fail to find ICloudVpc for vpc %s", err)
		self.taskFail(ctx, peer, jsonutils.NewString(msg))
		return
	}

	ipeer, err := extVpc.GetICloudVpcPeeringConnectionById(peer.GetExternalId())
	if err != nil {
		msg := fmt.Sprintf("fail to find ICloudVpcPeeringConnection for VpcPeeringConnection %s", err)
		self.taskFail(ctx, peer, jsonutils.NewString(msg))
		return
	}

	err = ipeer.Refresh()
	if err != nil {
		msg := fmt.Sprintf("fail to refresh ICloudVpcPeeringConnection status %s", err)
		self.taskFail(ctx, peer, jsonutils.NewString(msg))
		return
	}

	err = peer.SyncWithCloudPeerConnection(ctx, self.UserCred, ipeer, nil)
	if err != nil {
		msg := fmt.Sprintf("fail to sync vpc status %s", err)
		self.taskFail(ctx, peer, jsonutils.NewString(msg))
		return
	}

	logclient.AddActionLogWithStartable(self, peer, logclient.ACT_SYNC_STATUS, nil, self.UserCred, true)
	self.SetStageComplete(ctx, nil)
}
