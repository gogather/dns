package model

import (
	"github.com/gogather/dns/utils"
)

// https://www.dnspod.cn/docs/records.html#record-create
type DnspodAddRecordArg struct {
	DnspodCommonArg

	Domain     string `query:"domain"`
	SubDomain  string `query:"sub_domain"`
	RecordType string `query:"record_type"`
	RecordLine string `query:"record_line"`
	Value      string `query:"value"`
	MX         int    `query:"mx"`
	TTL        int    `query:"ttl"`
}

func (dara *DnspodAddRecordArg) ToQuery() string {
	c := utils.Marshal(*dara)
	return string(c)
}
