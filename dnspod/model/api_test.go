package model

import (
	"fmt"
	"testing"
)

func TestDnspodAddRecordArg_ToQuery(t *testing.T) {
	args := DnspodAddRecordArg{
		DnspodCommonArg: DnspodCommonArg{
			LoginToken:   "token",
			Format:       "json",
			Lang:         "zh",
			ErrorOnEmpty: "yes",
			UserId:       "rex",
		},
		Domain:     "yfcloud.io",
		SubDomain:  "www",
		RecordType: "A",
		RecordLine: "default",
		Value:      "127.0.0.1",
		MX:         2,
		TTL:        1,
	}

	c := args.ToQuery()
	fmt.Println(c)
}
