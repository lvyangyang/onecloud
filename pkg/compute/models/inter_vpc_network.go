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

package models

import (
	"context"

	"yunion.io/x/jsonutils"
	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/util/stringutils2"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/sqlchemy"
)

type SInterVpcNetworkManager struct {
	db.SEnabledStatusInfrasResourceBaseManager
	db.SExternalizedResourceBaseManager
}

var InterVpcNetworkManager *SInterVpcNetworkManager

func init() {
	InterVpcNetworkManager = &SInterVpcNetworkManager{
		SEnabledStatusInfrasResourceBaseManager: db.NewEnabledStatusInfrasResourceBaseManager(
			SInterVpcNetwork{},
			"inter_vpc_networks_tbl",
			"inter_vpc_network",
			"inter_vpc_networks",
		),
	}
	InterVpcNetworkManager.SetVirtualObject(InterVpcNetworkManager)
}

type SInterVpcNetwork struct {
	db.SEnabledStatusInfrasResourceBase
	db.SExternalizedResourceBase
	CloudaccountId string `width:"36" charset:"ascii" nullable:"false" list:"user" create:"required" json:"cloudaccount_id"`
}

func (manager *SInterVpcNetworkManager) GetContextManagers() [][]db.IModelManager {
	return [][]db.IModelManager{
		{VpcManager},
	}
}

// 列表
func (manager *SInterVpcNetworkManager) ListItemFilter(
	ctx context.Context,
	q *sqlchemy.SQuery,
	userCred mcclient.TokenCredential,
	query api.InterVpcNetworkListInput,
) (*sqlchemy.SQuery, error) {
	var err error
	q, err = manager.SEnabledStatusInfrasResourceBaseManager.ListItemFilter(ctx, q, userCred, query.EnabledStatusInfrasResourceBaseListInput)
	if err != nil {
		return nil, errors.Wrap(err, "SEnabledStatusInfrasResourceBaseManager.ListItemFilter")
	}
	return q, nil
}

func (manager *SInterVpcNetworkManager) ValidateCreateData(
	ctx context.Context,
	userCred mcclient.TokenCredential,
	ownerId mcclient.IIdentityProvider,
	query jsonutils.JSONObject,
	input api.InterVpcNetworkCreateInput,
) (api.InterVpcNetworkCreateInput, error) {
	return input, nil
}

func (self *SInterVpcNetworkManager) PostCreate(ctx context.Context, userCred mcclient.TokenCredential, ownerId mcclient.IIdentityProvider, query jsonutils.JSONObject, data jsonutils.JSONObject) {
}

func (self *SInterVpcNetworkManager) GetExtraDetails(
	ctx context.Context,
	userCred mcclient.TokenCredential,
	query jsonutils.JSONObject,
	isList bool,
) (api.InterVpcNetworkDetails, error) {
	return api.InterVpcNetworkDetails{}, nil
}

func (manager *SInterVpcNetworkManager) FetchCustomizeColumns(
	ctx context.Context,
	userCred mcclient.TokenCredential,
	query jsonutils.JSONObject,
	objs []interface{},
	fields stringutils2.SSortedStrings,
	isList bool,
) []api.InterVpcNetworkDetails {
	rows := make([]api.InterVpcNetworkDetails, len(objs))
	stdRows := manager.SEnabledStatusInfrasResourceBaseManager.FetchCustomizeColumns(ctx, userCred, query, objs, fields, isList)
	for i := range rows {
		rows[i] = api.InterVpcNetworkDetails{
			EnabledStatusInfrasResourceBaseDetails: stdRows[i],
		}
	}
	return rows
}

func (self *SInterVpcNetworkManager) SyncWithCloudInterVpcNetwork(ctx context.Context, userCred mcclient.TokenCredential, ext cloudprovider.ICloudInterVpcNetwork) error{
	return nil
}

func (self *SInterVpcNetworkManager) newFromCloudInterVpcNetwork(ctx context.Context, userCred mcclient.TokenCredential, ext cloudprovider.ICloudInterVpcNetwork, account *SCloudaccount) (*SInterVpcNetwork, bool, error){
	return nil,false,nil
}

func (self *SInterVpcNetwork) CustomizeDelete(ctx context.Context, userCred mcclient.TokenCredential, query jsonutils.JSONObject, data jsonutils.JSONObject) error {
	return nil
}

func (self *SInterVpcNetwork) RealDelete(ctx context.Context, userCred mcclient.TokenCredential) error {
	return self.SEnabledStatusInfrasResourceBase.Delete(ctx, userCred)
}

func (self *SInterVpcNetwork) AllowPerformSyncstatus(ctx context.Context, userCred mcclient.TokenCredential, query jsonutils.JSONObject, data jsonutils.JSONObject) bool {
	return db.IsAdminAllowPerform(userCred, self, "syncstatus")
}

func (self *SInterVpcNetwork) PerformSync(ctx context.Context, userCred mcclient.TokenCredential, query jsonutils.JSONObject, input api.VpcSyncstatusInput) (jsonutils.JSONObject, error) {
	return nil, nil
}

func (self *SInterVpcNetwork) syncRemove(ctx context.Context, userCred mcclient.TokenCredential) error {
	return self.RealDelete(ctx, userCred)
}
