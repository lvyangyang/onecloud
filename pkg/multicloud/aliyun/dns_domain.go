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
	"yunion.io/x/onecloud/pkg/multicloud"
)

type SDomains struct {
	PageNumber int `json:"PageNumber"`
	TotalCount int `json:"TotalCount"`
	PageSize   int `json:"PageSize"`
	// RequestID  string  `json:"RequestId"`
	Domains sDomains `json:"Domains"`
}
type DNSServers struct {
	DNSServer []string `json:"DnsServer"`
}
type SDomain struct {
	multicloud.SResourceBase
	client      *SAliyunClient
	PunyCode    string     `json:"PunyCode"`
	VersionCode string     `json:"VersionCode"`
	InstanceID  string     `json:"InstanceId"`
	AliDomain   bool       `json:"AliDomain"`
	DomainName  string     `json:"DomainName"`
	DomainID    string     `json:"DomainId"`
	DNSServers  DNSServers `json:"DnsServers"`
	GroupID     string     `json:"GroupId"`
}
type sDomains struct {
	Domain []SDomain `json:"Domain"`
}

// https://help.aliyun.com/document_detail/29751.html?spm=a2c4g.11186623.6.638.55563230d00kzJ
func (client *SAliyunClient) DescribeDomains(offset int, limit int) (SDomains, error) {
	sdomains := SDomains{}
	params := map[string]string{}
	if limit < 1 {
		limit = 20
	}
	params["Action"] = "DescribeDomains"
	params["PageNumber"] = strconv.Itoa(offset/limit + 1)
	params["PageSize"] = strconv.Itoa(limit)
	resp, err := client.alidnsRequest("DescribeDomains", params)
	if err != nil {
		return sdomains, errors.Wrap(err, "DescribeDomains")
	}
	err = resp.Unmarshal(&sdomains)
	if err != nil {
		return sdomains, errors.Wrap(err, "resp.Unmarshal")
	}
	return sdomains, nil
}

func (client *SAliyunClient) GetAllDomains() ([]SDomain, error) {
	count := 0
	sdomains := []SDomain{}
	for {
		domains, err := client.DescribeDomains(count, 20)
		if err != nil {
			return nil, errors.Wrapf(err, "client.DescribeDomains(%d, 20)", count)
		}
		count += len(domains.Domains.Domain)
		sdomains = append(sdomains, domains.Domains.Domain...)
		if count >= domains.TotalCount {
			break
		}
	}
	for i := 0; i < len(sdomains); i++ {
		sdomains[i].client = client
	}
	return sdomains, nil
}

func (client *SAliyunClient) DescribeDomainInfo(domainName string) (SDomain, error) {
	sdomain := SDomain{client: client}
	params := map[string]string{}
	params["Action"] = "DescribeDomainInfo"
	params["DomainName"] = domainName
	resp, err := client.alidnsRequest("DescribeDomainInfo", params)
	if err != nil {
		return sdomain, errors.Wrap(err, "DescribeDomainInfo")
	}
	err = resp.Unmarshal(&sdomain)
	if err != nil {
		return sdomain, errors.Wrap(err, "resp.Unmarshal")
	}
	return sdomain, nil
}

func (client *SAliyunClient) AddDomain(domainName string) (SDomain, error) {
	sdomain := SDomain{client: client}
	params := map[string]string{}
	params["Action"] = "AddDomain"
	params["DomainName"] = domainName
	resp, err := client.alidnsRequest("AddDomain", params)
	if err != nil {
		return sdomain, errors.Wrap(err, "AddDomain")
	}
	err = resp.Unmarshal(&sdomain)
	if err != nil {
		return sdomain, errors.Wrapf(err, "%s:resp.Unmarshal()", resp)
	}
	return sdomain, nil
}

func (client *SAliyunClient) DeleteDomain(domainName string) error {
	params := map[string]string{}
	params["Action"] = "DeleteDomain"
	params["DomainName"] = domainName
	_, err := client.alidnsRequest("DeleteDomain", params)
	if err != nil {
		return errors.Wrap(err, "DeleteDomain")
	}
	return nil
}

func (self *SDomain) GetId() string {
	return self.DomainID
}

func (self *SDomain) GetName() string {
	return self.DomainName
}

func (self *SDomain) GetGlobalId() string {
	return self.GetId()
}

func (self *SDomain) GetStatus() string {
	return api.DNS_ZONE_STATUS_AVAILABLE
}

func (self *SDomain) Refresh() error {
	sdomain, err := self.client.DescribeDomainInfo(self.DomainName)
	if err != nil {
		return errors.Wrapf(err, "self.client.DescribeDomainInfo(%s)", self.DomainName)
	}
	return jsonutils.Update(self, sdomain)
}

func (self *SDomain) GetZoneType() cloudprovider.TDnsZoneType {
	return cloudprovider.PublicZone
}
func (self *SDomain) GetOptions() *jsonutils.JSONDict {
	return nil
}

func (self *SDomain) GetICloudVpcIds() ([]string, error) {
	return nil, nil
}
func (self *SDomain) AddVpc(vpc *cloudprovider.SPrivateZoneVpc) error {
	return cloudprovider.ErrNotSupported
}
func (self *SDomain) RemoveVpc(vpc *cloudprovider.SPrivateZoneVpc) error {
	return cloudprovider.ErrNotSupported
}

func (self *SDomain) GetIDnsRecordSets() ([]cloudprovider.ICloudDnsRecordSet, error) {
	irecords := []cloudprovider.ICloudDnsRecordSet{}
	records, err := self.client.GetAllDomainRecords(self.DomainName)
	if err != nil {
		return nil, errors.Wrapf(err, "self.client.GetAllDomainRecords(%s)", self.DomainName)
	}
	for i := 0; i < len(records); i++ {
		irecords = append(irecords, &records[i])
	}
	return irecords, nil
}

func (self *SDomain) SyncDnsRecordSets(common, add, del, update []cloudprovider.DnsRecordSet) error {
	for i := 0; i < len(add); i++ {
		recordId, err := self.client.AddDomainRecord(self.DomainName, add[i])
		if err != nil {
			return errors.Wrapf(err, "self.client.AddDomainRecord(%s,%s)", self.DomainName, jsonutils.Marshal(add[i]).String())
		}
		if !add[i].Enabled {
			// Enable: 启用解析 Disable: 暂停解析
			err = self.client.SetDomainRecordStatus(recordId, "Disable")
			if err != nil {
				return errors.Wrapf(err, "self.client.SetDomainRecordStatus(%s,%t)", recordId, add[i].Enabled)
			}
		}
	}

	for i := 0; i < len(del); i++ {
		err := self.client.DeleteDomainRecord(del[i].ExternalId)
		if err != nil {
			return errors.Wrapf(err, "self.client.DeleteDomainRecord(%s)", del[i].ExternalId)
		}
	}

	for i := 0; i < len(update); i++ {
		// Enable: 启用解析 Disable: 暂停解析
		status := "Enable"
		if !update[i].Enabled {
			status = "Disable"
		}
		err := self.client.SetDomainRecordStatus(update[i].ExternalId, status)
		if err != nil {
			return errors.Wrapf(err, "self.client.SetDomainRecordStatus(%s,%t)", update[i].ExternalId, update[i].Enabled)
		}
		err = self.client.UpdateDomainRecord(update[i])
		if err != nil {
			return errors.Wrapf(err, "self.client.UpdateDomainRecord(%s)", jsonutils.Marshal(update[i]).String())
		}
	}

	return nil
}

func (self *SDomain) Delete() error {
	return self.client.DeleteDomain(self.DomainName)
}
