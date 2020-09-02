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

package cloudprovider

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/gotypes"
)

type TDnsZoneType string
type TDnsPolicyType string

type TDnsType string

type TDnsPolicyTypeValue jsonutils.JSONObject

const (
	PublicZone  = TDnsZoneType("PublicZone")
	PrivateZone = TDnsZoneType("PrivateZone")
)

const (
	DnsPolicyTypeSimple           = TDnsPolicyType("Simple")           //简单
	DnsPolicyTypeByCarrier        = TDnsPolicyType("ByCarrier")        //运营商
	DnsPolicyTypeByGeoLocation    = TDnsPolicyType("ByGeoLocation")    //地理区域
	DnsPolicyTypeBySearchEngine   = TDnsPolicyType("BySearchEngine")   //搜索引擎
	DnsPolicyTypeIpRange          = TDnsPolicyType("IpRange")          //自定义IP范围
	DnsPolicyTypeWeighted         = TDnsPolicyType("Weighted")         //加权
	DnsPolicyTypeFailover         = TDnsPolicyType("Failover")         //故障转移
	DnsPolicyTypeMultiValueAnswer = TDnsPolicyType("MultiValueAnswer") //多值应答
	DnsPolicyTypeLatency          = TDnsPolicyType("Latency")          //延迟
)

const (
	DnsTypeA            = TDnsType("A")
	DnsTypeAAAA         = TDnsType("AAAA")
	DnsTypeCAA          = TDnsType("CAA")
	DnsTypeCNAME        = TDnsType("CNAME")
	DnsTypeMX           = TDnsType("MX")
	DnsTypeNS           = TDnsType("NS")
	DnsTypeSRV          = TDnsType("SRV")
	DnsTypeSOA          = TDnsType("SOA")
	DnsTypeTXT          = TDnsType("TXT")
	DnsTypePTR          = TDnsType("PTR")
	DnsTypeDS           = TDnsType("DS")
	DnsTypeDNSKEY       = TDnsType("DNSKEY")
	DnsTypeIPSECKEY     = TDnsType("IPSECKEY")
	DnsTypeNAPTR        = TDnsType("NAPTR")
	DnsTypeSPF          = TDnsType("SPF")
	DnsTypeSSHFP        = TDnsType("SSHFP")
	DnsTypeTLSA         = TDnsType("TLSA")
	DnsTypeREDIRECT_URL = TDnsType("REDIRECT_URL") //显性URL转发
	DnsTypeFORWARD_URL  = TDnsType("FORWARD_URL")  //隐性URL转发
)

var (
	SUPPORTED_DNS_TYPES = []TDnsType{
		DnsTypeA,
		DnsTypeAAAA,
		DnsTypeCAA,
		DnsTypeCNAME,
		DnsTypeMX,
		DnsTypeNS,
		DnsTypeSRV,
		DnsTypeSOA,
		DnsTypeTXT,
		DnsTypePTR,
		DnsTypeDS,
		DnsTypeDNSKEY,
		DnsTypeIPSECKEY,
		DnsTypeNAPTR,
		DnsTypeSPF,
		DnsTypeSSHFP,
		DnsTypeTLSA,
		DnsTypeREDIRECT_URL,
		DnsTypeFORWARD_URL,
	}
)

var (
	DnsPolicyValueNil TDnsPolicyTypeValue = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{}))

	DnsPolicyTypeByCarrierUnicom      = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"carrier": "unicom"}))
	DnsPolicyTypeByCarrierTelecom     = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"carrier": "telecom"}))
	DnsPolicyTypeByCarrierChinaMobile = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"carrier": "chinamobile"}))
	DnsPolicyTypeByCarrierCernet      = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"carrier": "cernet"}))

	DnsPolicyTypeByGeoLocationOversea  = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"location": "oversea"}))
	DnsPolicyTypeByGeoLocationMainland = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"location": "mainland"}))

	DnsPolicyTypeBySearchEngineBaidu   = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"searchengine": "baidu"}))
	DnsPolicyTypeBySearchEngineGoogle  = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"searchengine": "google"}))
	DnsPolicyTypeBySearchEngineBing    = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"searchengine": "bing"}))
	DnsPolicyTypeBySearchEngineYoudao  = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"searchengine": "youdao"}))
	DnsPolicyTypeBySearchEngineSousou  = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"searchengine": "sousou"}))
	DnsPolicyTypeBySearchEngineSougou  = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"searchengine": "sougou"}))
	DnsPolicyTypeBySearchEngineQihu360 = TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"searchengine": "qihu360"}))
)

var AwsGeoLocations = []TDnsPolicyTypeValue{
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "*", "CountryName": "Default"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AD", "CountryName": "Andorra"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AE", "CountryName": "United Arab Emirates"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AF", "CountryName": "Afghanistan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AG", "CountryName": "Antigua and Barbuda"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AI", "CountryName": "Anguilla"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AL", "CountryName": "Albania"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AM", "CountryName": "Armenia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AO", "CountryName": "Angola"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AQ", "CountryName": "Antarctica"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AR", "CountryName": "Argentina"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AS", "CountryName": "American Samoa"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AT", "CountryName": "Austria"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AU", "CountryName": "Australia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AW", "CountryName": "Aruba"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AX", "CountryName": "Åland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "AZ", "CountryName": "Azerbaijan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BA", "CountryName": "Bosnia and Herzegovina"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BB", "CountryName": "Barbados"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BD", "CountryName": "Bangladesh"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BE", "CountryName": "Belgium"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BF", "CountryName": "Burkina Faso"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BG", "CountryName": "Bulgaria"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BH", "CountryName": "Bahrain"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BI", "CountryName": "Burundi"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BJ", "CountryName": "Benin"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BL", "CountryName": "Saint Barthélemy"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BM", "CountryName": "Bermuda"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BN", "CountryName": "Brunei"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BO", "CountryName": "Bolivia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BQ", "CountryName": "Bonaire"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BR", "CountryName": "Brazil"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BS", "CountryName": "Bahamas"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BT", "CountryName": "Bhutan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BW", "CountryName": "Botswana"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BY", "CountryName": "Belarus"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "BZ", "CountryName": "Belize"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CA", "CountryName": "Canada"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CC", "CountryName": "Cocos [Keeling] Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CD", "CountryName": "Congo"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CF", "CountryName": "Central African Republic"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CG", "CountryName": "Republic of the Congo"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CH", "CountryName": "Switzerland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CI", "CountryName": "Ivory Coast"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CK", "CountryName": "Cook Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CL", "CountryName": "Chile"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CM", "CountryName": "Cameroon"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CN", "CountryName": "China"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CO", "CountryName": "Colombia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CR", "CountryName": "Costa Rica"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CU", "CountryName": "Cuba"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CV", "CountryName": "Cape Verde"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CW", "CountryName": "Curaçao"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CY", "CountryName": "Cyprus"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "CZ", "CountryName": "Czech Republic"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "DE", "CountryName": "Germany"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "DJ", "CountryName": "Djibouti"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "DK", "CountryName": "Denmark"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "DM", "CountryName": "Dominica"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "DO", "CountryName": "Dominican Republic"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "DZ", "CountryName": "Algeria"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "EC", "CountryName": "Ecuador"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "EE", "CountryName": "Estonia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "EG", "CountryName": "Egypt"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ER", "CountryName": "Eritrea"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ES", "CountryName": "Spain"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ET", "CountryName": "Ethiopia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "FI", "CountryName": "Finland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "FJ", "CountryName": "Fiji"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "FK", "CountryName": "Falkland Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "FM", "CountryName": "Federated States of Micronesia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "FO", "CountryName": "Faroe Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "FR", "CountryName": "France"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GA", "CountryName": "Gabon"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GB", "CountryName": "United Kingdom"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GD", "CountryName": "Grenada"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GE", "CountryName": "Georgia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GF", "CountryName": "French Guiana"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GG", "CountryName": "Guernsey"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GH", "CountryName": "Ghana"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GI", "CountryName": "Gibraltar"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GL", "CountryName": "Greenland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GM", "CountryName": "Gambia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GN", "CountryName": "Guinea"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GP", "CountryName": "Guadeloupe"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GQ", "CountryName": "Equatorial Guinea"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GR", "CountryName": "Greece"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GS", "CountryName": "South Georgia and the South Sandwich Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GT", "CountryName": "Guatemala"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GU", "CountryName": "Guam"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GW", "CountryName": "Guinea-Bissau"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "GY", "CountryName": "Guyana"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "HK", "CountryName": "Hong Kong"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "HN", "CountryName": "Honduras"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "HR", "CountryName": "Croatia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "HT", "CountryName": "Haiti"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "HU", "CountryName": "Hungary"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ID", "CountryName": "Indonesia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IE", "CountryName": "Ireland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IL", "CountryName": "Israel"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IM", "CountryName": "Isle of Man"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IN", "CountryName": "India"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IO", "CountryName": "British Indian Ocean Territory"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IQ", "CountryName": "Iraq"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IR", "CountryName": "Iran"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IS", "CountryName": "Iceland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "IT", "CountryName": "Italy"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "JE", "CountryName": "Jersey"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "JM", "CountryName": "Jamaica"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "JO", "CountryName": "Hashemite Kingdom of Jordan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "JP", "CountryName": "Japan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KE", "CountryName": "Kenya"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KG", "CountryName": "Kyrgyzstan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KH", "CountryName": "Cambodia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KI", "CountryName": "Kiribati"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KM", "CountryName": "Comoros"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KN", "CountryName": "Saint Kitts and Nevis"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KP", "CountryName": "North Korea"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KR", "CountryName": "Republic of Korea"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KW", "CountryName": "Kuwait"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KY", "CountryName": "Cayman Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "KZ", "CountryName": "Kazakhstan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LA", "CountryName": "Laos"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LB", "CountryName": "Lebanon"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LC", "CountryName": "Saint Lucia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LI", "CountryName": "Liechtenstein"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LK", "CountryName": "Sri Lanka"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LR", "CountryName": "Liberia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LS", "CountryName": "Lesotho"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LT", "CountryName": "Lithuania"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LU", "CountryName": "Luxembourg"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LV", "CountryName": "Latvia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "LY", "CountryName": "Libya"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MA", "CountryName": "Morocco"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MC", "CountryName": "Monaco"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MD", "CountryName": "Republic of Moldova"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ME", "CountryName": "Montenegro"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MF", "CountryName": "Saint Martin"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MG", "CountryName": "Madagascar"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MH", "CountryName": "Marshall Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MK", "CountryName": "Macedonia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ML", "CountryName": "Mali"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MM", "CountryName": "Myanmar [Burma]"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MN", "CountryName": "Mongolia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MO", "CountryName": "Macao"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MP", "CountryName": "Northern Mariana Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MQ", "CountryName": "Martinique"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MR", "CountryName": "Mauritania"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MS", "CountryName": "Montserrat"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MT", "CountryName": "Malta"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MU", "CountryName": "Mauritius"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MV", "CountryName": "Maldives"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MW", "CountryName": "Malawi"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MX", "CountryName": "Mexico"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MY", "CountryName": "Malaysia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "MZ", "CountryName": "Mozambique"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NA", "CountryName": "Namibia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NC", "CountryName": "New Caledonia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NE", "CountryName": "Niger"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NF", "CountryName": "Norfolk Island"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NG", "CountryName": "Nigeria"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NI", "CountryName": "Nicaragua"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NL", "CountryName": "Netherlands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NO", "CountryName": "Norway"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NP", "CountryName": "Nepal"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NR", "CountryName": "Nauru"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NU", "CountryName": "Niue"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "NZ", "CountryName": "New Zealand"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "OM", "CountryName": "Oman"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PA", "CountryName": "Panama"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PE", "CountryName": "Peru"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PF", "CountryName": "French Polynesia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PG", "CountryName": "Papua New Guinea"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PH", "CountryName": "Philippines"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PK", "CountryName": "Pakistan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PL", "CountryName": "Poland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PM", "CountryName": "Saint Pierre and Miquelon"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PN", "CountryName": "Pitcairn Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PR", "CountryName": "Puerto Rico"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PS", "CountryName": "Palestine"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PT", "CountryName": "Portugal"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PW", "CountryName": "Palau"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "PY", "CountryName": "Paraguay"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "QA", "CountryName": "Qatar"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "RE", "CountryName": "Réunion"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "RO", "CountryName": "Romania"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "RS", "CountryName": "Serbia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "RU", "CountryName": "Russia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "RW", "CountryName": "Rwanda"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SA", "CountryName": "Saudi Arabia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SB", "CountryName": "Solomon Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SC", "CountryName": "Seychelles"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SD", "CountryName": "Sudan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SE", "CountryName": "Sweden"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SG", "CountryName": "Singapore"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SH", "CountryName": "Saint Helena"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SI", "CountryName": "Slovenia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SJ", "CountryName": "Svalbard and Jan Mayen"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SK", "CountryName": "Slovakia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SL", "CountryName": "Sierra Leone"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SM", "CountryName": "San Marino"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SN", "CountryName": "Senegal"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SO", "CountryName": "Somalia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SR", "CountryName": "Suriname"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SS", "CountryName": "South Sudan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ST", "CountryName": "São Tomé and Príncipe"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SV", "CountryName": "El Salvador"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SX", "CountryName": "Sint Maarten"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SY", "CountryName": "Syria"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "SZ", "CountryName": "Swaziland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TC", "CountryName": "Turks and Caicos Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TD", "CountryName": "Chad"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TF", "CountryName": "French Southern Territories"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TG", "CountryName": "Togo"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TH", "CountryName": "Thailand"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TJ", "CountryName": "Tajikistan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TK", "CountryName": "Tokelau"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TL", "CountryName": "East Timor"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TM", "CountryName": "Turkmenistan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TN", "CountryName": "Tunisia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TO", "CountryName": "Tonga"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TR", "CountryName": "Turkey"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TT", "CountryName": "Trinidad and Tobago"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TV", "CountryName": "Tuvalu"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TW", "CountryName": "Taiwan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "TZ", "CountryName": "Tanzania"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "UA", "CountryName": "Ukraine"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "UG", "CountryName": "Uganda"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "UM", "CountryName": "U.S. Minor Outlying Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "AK", "SubdivisionName": "Alaska"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "AL", "SubdivisionName": "Alabama"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "AR", "SubdivisionName": "Arkansas"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "AZ", "SubdivisionName": "Arizona"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "CA", "SubdivisionName": "California"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "CO", "SubdivisionName": "Colorado"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "CT", "SubdivisionName": "Connecticut"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "DC", "SubdivisionName": "District of Columbia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "DE", "SubdivisionName": "Delaware"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "FL", "SubdivisionName": "Florida"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "GA", "SubdivisionName": "Georgia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "HI", "SubdivisionName": "Hawaii"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "IA", "SubdivisionName": "Iowa"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "ID", "SubdivisionName": "Idaho"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "IL", "SubdivisionName": "Illinois"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "IN", "SubdivisionName": "Indiana"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "KS", "SubdivisionName": "Kansas"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "KY", "SubdivisionName": "Kentucky"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "LA", "SubdivisionName": "Louisiana"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "MA", "SubdivisionName": "Massachusetts"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "MD", "SubdivisionName": "Maryland"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "ME", "SubdivisionName": "Maine"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "MI", "SubdivisionName": "Michigan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "MN", "SubdivisionName": "Minnesota"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "MO", "SubdivisionName": "Missouri"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "MS", "SubdivisionName": "Mississippi"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "MT", "SubdivisionName": "Montana"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "NC", "SubdivisionName": "North Carolina"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "ND", "SubdivisionName": "North Dakota"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "NE", "SubdivisionName": "Nebraska"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "NH", "SubdivisionName": "New Hampshire"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "NJ", "SubdivisionName": "New Jersey"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "NM", "SubdivisionName": "New Mexico"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "NV", "SubdivisionName": "Nevada"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "NY", "SubdivisionName": "New York"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "OH", "SubdivisionName": "Ohio"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "OK", "SubdivisionName": "Oklahoma"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "OR", "SubdivisionName": "Oregon"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "PA", "SubdivisionName": "Pennsylvania"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "RI", "SubdivisionName": "Rhode Island"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "SC", "SubdivisionName": "South Carolina"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "SD", "SubdivisionName": "South Dakota"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "TN", "SubdivisionName": "Tennessee"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "TX", "SubdivisionName": "Texas"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "UT", "SubdivisionName": "Utah"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "VA", "SubdivisionName": "Virginia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "VT", "SubdivisionName": "Vermont"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "WA", "SubdivisionName": "Washington"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "WI", "SubdivisionName": "Wisconsin"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "WV", "SubdivisionName": "West Virginia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "US", "CountryName": "United States", "SubdivisionCode": "WY", "SubdivisionName": "Wyoming"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "UY", "CountryName": "Uruguay"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "UZ", "CountryName": "Uzbekistan"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "VA", "CountryName": "Vatican City"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "VC", "CountryName": "Saint Vincent and the Grenadines"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "VE", "CountryName": "Venezuela"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "VG", "CountryName": "British Virgin Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "VI", "CountryName": "U.S. Virgin Islands"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "VN", "CountryName": "Vietnam"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "VU", "CountryName": "Vanuatu"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "WF", "CountryName": "Wallis and Futuna"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "WS", "CountryName": "Samoa"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "XK", "CountryName": "Kosovo"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "YE", "CountryName": "Yemen"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "YT", "CountryName": "Mayotte"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ZA", "CountryName": "South Africa"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ZM", "CountryName": "Zambia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"CountryCode": "ZW", "CountryName": "Zimbabwe"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"ContinentCode": "AF", "ContinentName": "Africa"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"ContinentCode": "AN", "ContinentName": "Antarctica"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"ContinentCode": "AS", "ContinentName": "Asia"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"ContinentCode": "EU", "ContinentName": "Europe"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"ContinentCode": "NA", "ContinentName": "North America"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"ContinentCode": "OC", "ContinentName": "Oceania"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"ContinentCode": "SA", "ContinentName": "South America"})),
}

var AwsRegions = []TDnsPolicyTypeValue{
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "us-east-2"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "us-east-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "us-west-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "us-west-2"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "af-south-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "ap-east-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "ap-south-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "ap-northeast-3"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "ap-northeast-2"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "ap-southeast-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "ap-southeast-2"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "ap-northeast-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "ca-central-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "cn-north-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "cn-northwest-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "eu-central-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "eu-west-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "eu-west-2"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "eu-south-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "eu-west-3"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "eu-north-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "me-south-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "sa-east-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "us-gov-east-1"})),
	TDnsPolicyTypeValue(jsonutils.Marshal(map[string]string{"region": "us-gov-west-1"})),
}

type SPrivateZoneVpc struct {
	Id       string
	RegionId string
}

type SDnsZoneCreateOptions struct {
	Name     string
	Desc     string
	ZoneType TDnsZoneType
	Vpcs     []SPrivateZoneVpc
	Options  *jsonutils.JSONDict
}

func IsSupportPolicyValue(v1 TDnsPolicyTypeValue, arr []TDnsPolicyTypeValue) bool {
	for i := range arr {
		if IsPolicyValueEqual(v1, arr[i]) {
			return true
		}
	}
	return false
}

func IsPolicyValueEqual(v1, v2 TDnsPolicyTypeValue) bool {
	if gotypes.IsNil(v1) {
		v1 = DnsPolicyValueNil
	}
	if gotypes.IsNil(v2) {
		v2 = DnsPolicyValueNil
	}
	return jsonutils.Marshal(v1).Equals(jsonutils.Marshal(v2))
}

type DnsRecordSet struct {
	Id         string
	ExternalId string

	Enabled      bool
	DnsName      string
	DnsType      TDnsType
	DnsValue     string
	Status       string
	Ttl          int64
	PolicyType   TDnsPolicyType
	PolicyParams TDnsPolicyTypeValue
}

func (r DnsRecordSet) GetGlobalId() string {
	return r.ExternalId
}

func (r DnsRecordSet) GetName() string {
	return r.DnsName
}

func (r DnsRecordSet) GetDnsName() string {
	return r.DnsName
}

func (r DnsRecordSet) GetDnsValue() string {
	return r.DnsValue
}

func (r DnsRecordSet) GetPolicyType() TDnsPolicyType {
	return r.PolicyType
}

func (r DnsRecordSet) GetPolicyParams() TDnsPolicyTypeValue {
	return r.PolicyParams
}

func (r DnsRecordSet) GetStatus() string {
	return r.Status
}

func (r DnsRecordSet) GetTTL() int64 {
	return r.Ttl
}

func (r DnsRecordSet) GetDnsType() TDnsType {
	return r.DnsType
}

func (r DnsRecordSet) GetEnabled() bool {
	return r.Enabled
}

func (record DnsRecordSet) Equals(r DnsRecordSet) bool {
	if record.DnsName != r.DnsName {
		return false
	}
	if record.DnsType != r.DnsType {
		return false
	}
	if record.DnsValue != r.DnsValue {
		return false
	}
	if record.PolicyType != r.PolicyType {
		return false
	}
	if !IsPolicyValueEqual(record.PolicyParams, r.PolicyParams) {
		return false
	}
	return true
}

func (record DnsRecordSet) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", record.DnsType, record.DnsName, record.DnsValue, record.PolicyType, jsonutils.Marshal(record.PolicyParams).String())
}

func CompareDnsRecordSet(iRecords []ICloudDnsRecordSet, local []DnsRecordSet, debug bool) ([]DnsRecordSet, []DnsRecordSet, []DnsRecordSet, []DnsRecordSet) {
	common, added, removed, updated := []DnsRecordSet{}, []DnsRecordSet{}, []DnsRecordSet{}, []DnsRecordSet{}

	localMaps := map[string]DnsRecordSet{}
	remoteMaps := map[string]DnsRecordSet{}
	for i := range iRecords {
		record := DnsRecordSet{
			ExternalId: iRecords[i].GetGlobalId(),

			DnsName:      iRecords[i].GetDnsName(),
			DnsType:      iRecords[i].GetDnsType(),
			DnsValue:     iRecords[i].GetDnsValue(),
			Status:       iRecords[i].GetStatus(),
			Enabled:      iRecords[i].GetEnabled(),
			Ttl:          iRecords[i].GetTTL(),
			PolicyType:   iRecords[i].GetPolicyType(),
			PolicyParams: iRecords[i].GetPolicyParams(),
		}
		remoteMaps[record.String()] = record
	}
	for i := range local {
		localMaps[local[i].String()] = local[i]
	}

	for key, record := range localMaps {
		remoteRecord, ok := remoteMaps[key]
		if ok {
			record.ExternalId = remoteRecord.ExternalId
			if remoteRecord.Ttl != record.Ttl || remoteRecord.Enabled != record.Enabled {
				updated = append(updated, record)
			} else {
				common = append(common, record)
			}
		} else {
			added = append(added, record)
		}
	}

	for key, record := range remoteMaps {
		_, ok := localMaps[key]
		if !ok {
			removed = append(removed, record)
		}
	}

	return common, added, removed, updated
}
