package wuid

import (
	"testing"
)

func Test_GenSonyflakeID(t *testing.T) {
	var machineID uint16 = 1
	//if err := InitSonyflake(&machineID); err != nil {
	//	t.Error(err)
	//}
	//if err := InitSonyflake(nil); err != nil {
	//	t.Error(err)
	//}

	id, err := GetSonyflakeID(&machineID)
	if err != nil {
		t.Error(err)
	}
	t.Log("Generated Sonyflake ID:", id)
	idHex, err := GetSonyflakeIDHex(&machineID)
	if err != nil {
		t.Error(err)
	}
	t.Log("Generated Sonyflake ID (hex):", idHex)
}
