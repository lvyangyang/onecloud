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
	"yunion.io/x/onecloud/pkg/cloudcommon/notifyclient"
	"yunion.io/x/onecloud/pkg/compute/models"
	"yunion.io/x/onecloud/pkg/util/logclient"
)

type ElasticcacheRemoteUpdateTask struct {
	SGuestBaseTask
}

func init() {
	taskman.RegisterTask(ElasticcacheRemoteUpdateTask{})
}

func (self *ElasticcacheRemoteUpdateTask) taskFail(ctx context.Context, elasticcache *models.SElasticcache, reason jsonutils.JSONObject) {
	logclient.AddActionLogWithStartable(self, elasticcache, logclient.ACT_ENABLE, reason, self.UserCred, false)
	notifyclient.NotifySystemErrorWithCtx(ctx, elasticcache.Id, elasticcache.Name, api.LB_STATUS_DISABLED, reason.String())
	self.SetStageFailed(ctx, reason)
}

func (self *ElasticcacheRemoteUpdateTask) OnInit(ctx context.Context, obj db.IStandaloneModel, data jsonutils.JSONObject) {
	ec := obj.(*models.SElasticcache)
	region := ec.GetRegion()
	if region == nil {
		self.taskFail(ctx, ec, jsonutils.NewString(fmt.Sprintf("failed to find region for elastic cache %s", ec.GetName())))
		return
	}
	self.SetStage("OnRemoteUpdateComplete", nil)
	replaceTags := jsonutils.QueryBoolean(self.Params, "replace_tags", false)

	if err := region.GetDriver().RequestRemoteUpdateElasticcache(ctx, self.GetUserCred(), ec, replaceTags, self); err != nil {
		self.taskFail(ctx, ec, jsonutils.Marshal(err))
	}
}

func (self *ElasticcacheRemoteUpdateTask) OnRemoteUpdateComplete(ctx context.Context, obj db.IStandaloneModel, data jsonutils.JSONObject) {
	self.SetStageComplete(ctx, nil)
}
