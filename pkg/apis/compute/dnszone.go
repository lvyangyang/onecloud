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

package compute

import (
	"yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/apis"
)

/*

Architecture For DnsZone


                                                   +-----------+                                                   +-----------+
                                                   | RecordSet |                  +---------------+                | RecordSet |                                         +--------+
                                                 __| (A)       |_________         | TrafficPolicy |                | (TXT)     |__                                      _|  Vpc   |
                                               _/  +-----------+         \________| (Google)      |___             +-----------+  \___                                _/ +--------+
                                            __/                                   +---------------+   \____                           \__                            /
                  +-----------------------+/       +-----------+                                           \____   +-----------+         \___  +-----------------+ _/    +--------+
    API           |  DnsZone  example.com |        | RecordSet |                                                \__| RecordSet |             \_| DnsZone abc.app |-------|  Vpc   |
                  |  (Public)             |--------| (AAAA)    |                  +---------------+                | (CAA)     |---------------| (Private)       |       +--------+
                  +-----------------------+        +-----------+                  | TrafficPolicy |                +-----------+             __+-----------------+_
                          ^                \_                                _____| (Telecom)     |____                                  ___/          ^           \_
                          │                  \_    +-----------+     _______/     +---------------+    \_______    +-----------+      __/              |             \_  +--------+
                          │                    \_  | RecordSet |____/                                          \___| RecordSet |  ___/                 |               \_|  Vpc   |
                          │                      \_| (NS)      |                                                   | (PTR)     |_/                     |                 +--------+
                          │                        +-----------+                                                   +-----------+                       |
                          │                                                                                                                            |
                          │                                                                                                                            |
               ___________│____________________________________________________________________________________________________________________________|__________________________________
                          │                                                                                                                            |
                          v                                                                                                                            |
                  +-----------------+                                                                                                                  |
                  |                 |                                                                                                                  v
                  |                 |            +----------+
                  |  example.com <──|──────────> | Account1 |                                                          +----------+           +---------------+
                  |                 |            | (阿里云) |                                                          | Account3 | <-------> |     abc.app   |
                  |                 |            +----------+                              +------------+              | (Aws)    |           +---------------+
                  |                 |                                                      | Account2   |              +----------+
                  |  example.com <──|────────────────────────────────────────────────────> | (腾讯云)   |
   Cache          |                 |                                                      +------------+
                  |                 |
                  |                 |            +----------+
                  |  example.com <──|──────────> | Account4 |
                  |                 |            | (阿里云) |
                  |                 |            +----------+
                  +-----------------+
               ___________________________________________________________________________________________________________________________________________________________________________




                                                *************                           ***************                  *************
                                            ****             ****                   ****               ****          ****             ****
   公有云                                    **     阿里云    **                     **      腾讯云     **            **	  Aws      **
                                            ****             ****                   ****               ****          ****             ****
                                                *************                           ***************                  *************








*/

const (
	DNS_ZONE_STATUS_AVAILABLE               = "available"               // 可用
	DNS_ZONE_STATUS_CREATING                = "creating"                // 创建中
	DNS_ZONE_STATUS_CREATE_FAILE            = "create_failed"           // 创建失败
	DNS_ZONE_STATUS_UNCACHING               = "uncaching"               // 云上资源删除中
	DNS_ZONE_STATUS_UNCACHE_FAILED          = "uncache_failed"          // 云上资源删除失败
	DNS_ZONE_STATUS_CACHING                 = "caching"                 // 云上资源创建中
	DNS_ZONE_STATUS_CACHE_FAILED            = "cache_failed"            // 云上资源创建失败
	DNS_ZONE_STATUS_ADD_VPCS                = "add_vpcs"                // 绑定VPC中
	DNS_ZONE_STATUS_ADD_VPCS_FAILED         = "add_vpcs_failed"         // 绑定VPC失败
	DNS_ZONE_STATUS_REMOVE_VPCS             = "remove_vpcs"             // 解绑VPC中
	DNS_ZONE_STATUS_REMOVE_VPCS_FAILED      = "remove_vpcs_failed"      // 解绑VPC失败
	DNS_ZONE_STATUS_SYNC_RECORD_SETS        = "sync_record_sets"        // 同步解析列表中
	DNS_ZONE_STATUS_SYNC_RECORD_SETS_FAILED = "sync_record_sets_failed" // 同步解析列表失败
	DNS_ZONE_STATUS_DELETING                = "deleting"                // 删除中
	DNS_ZONE_STATUS_DELETE_FAILED           = "delete_failed"           // 删除失败
	DNS_ZONE_STATUS_UNKNOWN                 = "unknown"                 // 未知
)

type DnsZoneFilterListBase struct {
	DnsZoneId string `json:"dns_zone_id"`
}

type DnsZoneCreateInput struct {
	apis.EnabledStatusInfrasResourceBaseCreateInput

	// 区域类型
	//
	//
	// | 类型			| 说明    |
	// |----------		|---------|
	// | PublicZone		| 公有    |
	// | PrivateZone	| 私有    |
	ZoneType string `json:"zone_type"`
	// 额外参数

	// VPC id列表, 仅在zone_type为PrivateZone时生效, vpc列表必须属于同一个账号
	VpcIds []string `json:"vpc_ids"`

	// 云账号Id, 仅在zone_type为PublicZone时生效, 若为空则不会在云上创建
	CloudaccountId string `json:"cloudaccount_id"`

	// 额外信息
	Options *jsonutils.JSONDict `json:"options"`
}

type DnsZoneDetails struct {
	apis.EnabledStatusInfrasResourceBaseDetails
	SDnsZone

	// Dns记录数量
	DnsRecordsetCount int `json:"dns_recordset_count"`
	// 关联vpc数量
	VpcCount int `json:"vpc_count"`
}

type DnsZoneListInput struct {
	apis.EnabledStatusInfrasResourceBaseListInput

	// 区域类型
	//
	//
	// | 类型			| 说明    |
	// |----------		|---------|
	// | PublicZone		| 公有    |
	// | PrivateZone	| 私有    |
	ZoneType string `json:"zone_type"`

	// Filter dns zone By vpc
	VpcId string `json:"vpc_id"`
}

type DnsZoneSyncStatusInput struct {
}

type DnsZoneCacheInput struct {
	// 云账号Id
	//
	//
	// | 要求								|
	// |----------							|
	// | 1. dns zone 状态必须为available		|
	// | 2. dns zone zone_type 必须为PublicZone |
	// | 3. 指定云账号未在云上创建相应的 dns zone |
	CloudaccountId string
}

type DnsZoneUnacheInput struct {
	// 云账号Id
	CloudaccountId string
}

type DnsZoneAddVpcsInput struct {
	// VPC id列表
	VpcIds []string `json:"vpc_ids"`
}

type DnsZoneRemoveVpcsInput struct {
	// VPC id列表
	VpcIds []string `json:"vpc_ids"`
}

type DnsZoneSyncRecordSetsInput struct {
}

type DnsZonePurgeInput struct {
}
