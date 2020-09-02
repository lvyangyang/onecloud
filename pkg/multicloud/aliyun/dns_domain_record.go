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

package aliyun

import (
	"strconv"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SDomainRecords struct {
	//RequestID     string         `json:"RequestId"`
	TotalCount    int            `json:"TotalCount"`
	PageNumber    int            `json:"PageNumber"`
	PageSize      int            `json:"PageSize"`
	DomainRecords sDomainRecords `json:"DomainRecords"`
}

// https://help.aliyun.com/document_detail/29777.html?spm=a2c4g.11186623.6.666.aa4832307YdopF
type SDomainRecord struct {
	DomainId   string `json:"DomainId"`
	GroupId    string `json:"GroupId"`
	GroupName  string `json:"GroupName"`
	PunyCode   string `json:"PunyCode"`
	RR         string `json:"RR"`
	Status     string `json:"Status"`
	Value      string `json:"Value"`
	RecordID   string `json:"RecordId"`
	Type       string `json:"Type"`
	RequestID  string `json:"RequestId"`
	DomainName string `json:"DomainName"`
	Locked     bool   `json:"Locked"`
	Line       string `json:"Line"`
	TTL        int    `json:"TTL"`
	Priority   int    `json:"Priority"`
}
type sDomainRecords struct {
	Record []SDomainRecord `json:"Record"`
}

func (client *SAliyunClient) DescribeDomainRecords(domainName string, offset int, limit int) (SDomainRecords, error) {
	srecords := SDomainRecords{}
	params := map[string]string{}
	if limit < 1 {
		limit = 20
	}
	params["Action"] = "DescribeDomainRecords"
	params["DomainName"] = domainName
	params["PageNumber"] = strconv.Itoa(offset/limit + 1)
	params["PageSize"] = strconv.Itoa(limit)
	resp, err := client.alidnsRequest("DescribeDomainRecords", params)
	if err != nil {
		return srecords, errors.Wrap(err, "DescribeDomainRecords")
	}
	err = resp.Unmarshal(&srecords)
	if err != nil {
		return srecords, errors.Wrap(err, "resp.Unmarshal")
	}
	return srecords, nil
}

func (client *SAliyunClient) GetAllDomainRecords(domainName string) ([]SDomainRecord, error) {
	count := 0
	srecords := []SDomainRecord{}
	for {
		records, err := client.DescribeDomainRecords(domainName, count, 20)
		if err != nil {
			return nil, errors.Wrapf(err, "client.DescribeDomainRecords(%d, 20)", count)
		}
		count += len(records.DomainRecords.Record)
		srecords = append(srecords, records.DomainRecords.Record...)
		if count >= records.TotalCount {
			break
		}
	}
	return srecords, nil
}

func (client *SAliyunClient) DescribeDomainRecordInfo(recordId string) (*SDomainRecord, error) {
	srecord := SDomainRecord{}
	params := map[string]string{}
	params["Action"] = "DescribeDomainRecordInfo"
	params["RecordId"] = recordId

	resp, err := client.alidnsRequest("DescribeDomainRecordInfo", params)
	if err != nil {
		return nil, errors.Wrap(err, "DescribeDomainRecordInfo")
	}
	err = resp.Unmarshal(&srecord)
	if err != nil {
		return nil, errors.Wrap(err, "resp.Unmarshal")
	}
	return &srecord, nil
}

func (client *SAliyunClient) AddDomainRecord(domainName string, opts cloudprovider.DnsRecordSet) (string, error) {
	params := map[string]string{}
	params["Action"] = "AddDomainRecord"
	params["RR"] = opts.DnsName
	params["Type"] = string(opts.DnsType)
	params["Value"] = opts.DnsValue
	params["DomainName"] = domainName
	params["TTL"] = strconv.FormatInt(opts.Ttl, 10)
	ret, err := client.alidnsRequest("AddDomainRecord", params)
	if err != nil {
		return "", errors.Wrap(err, "AddDomainRecord")
	}
	recordId := ""

	return recordId, ret.Unmarshal(&recordId, "RecordId")
}

// line
func (client *SAliyunClient) UpdateDomainRecord(opts cloudprovider.DnsRecordSet) error {
	params := map[string]string{}
	params["Action"] = "UpdateDomainRecord"
	params["RR"] = opts.DnsName
	params["RecordId"] = opts.ExternalId
	params["Type"] = string(opts.DnsType)
	params["Value"] = opts.DnsValue
	params["TTL"] = strconv.FormatInt(opts.Ttl, 10)
	_, err := client.alidnsRequest("UpdateDomainRecord", params)
	if err != nil {
		return errors.Wrap(err, "UpdateDomainRecord")
	}
	return nil
}

// Enable: 启用解析 Disable: 暂停解析
func (client *SAliyunClient) SetDomainRecordStatus(recordId, status string) error {
	params := map[string]string{}
	params["Action"] = "SetDomainRecordStatus"
	params["RecordId"] = recordId
	params["Status"] = status
	_, err := client.alidnsRequest("SetDomainRecordStatus", params)
	if err != nil {
		return errors.Wrap(err, "SetDomainRecordStatus")
	}
	return nil
}

func (client *SAliyunClient) DeleteDomainRecord(recordId string) error {
	params := map[string]string{}
	params["Action"] = "DeleteDomainRecord"
	params["RecordId"] = recordId
	_, err := client.alidnsRequest("DeleteDomainRecord", params)
	if err != nil {
		return errors.Wrap(err, "DeleteDomainRecord")
	}
	return nil
}

func (self *SDomainRecord) GetGlobalId() string {
	return self.RecordID
}

func (self *SDomainRecord) GetDnsName() string {
	return self.RR
}

func (self *SDomainRecord) GetStatus() string {
	return api.DNS_ZONE_STATUS_AVAILABLE
}

func (self *SDomainRecord) GetEnabled() bool {
	return self.Status == "ENABLE"
}

func (self *SDomainRecord) GetDnsType() cloudprovider.TDnsType {
	return cloudprovider.TDnsType(self.Type)
}

func (self *SDomainRecord) GetDnsValue() string {
	return self.Value
}

func (self *SDomainRecord) GetTTL() int64 {
	return int64(self.TTL)
}

func (self *SDomainRecord) GetPolicyType() cloudprovider.TDnsPolicyType {
	return cloudprovider.DnsPolicyTypeSimple
}

func (self *SDomainRecord) GetPolicyValue() cloudprovider.TDnsPolicyValue {
	return ""
}

func (self *SDomainRecord) GetPolicyOptions() *jsonutils.JSONDict {
	return nil
}
