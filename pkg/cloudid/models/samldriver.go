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

	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/util/samlutils"
	"yunion.io/x/onecloud/pkg/util/samlutils/idp"
)

type SamlInstance func() *idp.SSAMLIdpInstance

var (
	SamlIdpInstance SamlInstance = nil
)

type ICloudSAMLLoginDriver interface {
	GetEntityID() string

	GetMetadataFilename() string
	GetMetadataUrl() string

	GetIdpInitiatedLoginData(ctx context.Context, userCred mcclient.TokenCredential, cloudAccountId string, sp *idp.SSAMLServiceProvider) (samlutils.SSAMLIdpInitiatedLoginData, error)
	GetSpInitiatedLoginData(ctx context.Context, userCred mcclient.TokenCredential, cloudAccoutId string, sp *idp.SSAMLServiceProvider) (samlutils.SSAMLSpInitiatedLoginData, error)
}

var (
	driverTable = make(map[string]ICloudSAMLLoginDriver)
)

func Register(driver ICloudSAMLLoginDriver) {
	driverTable[driver.GetEntityID()] = driver
}

func FindDriver(entityId string) ICloudSAMLLoginDriver {
	if driver, ok := driverTable[entityId]; ok {
		return driver
	}
	return nil
}

func AllDrivers() map[string]ICloudSAMLLoginDriver {
	return driverTable
}
