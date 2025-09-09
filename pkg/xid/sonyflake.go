package xid

import (
	"fmt"
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sf     *sonyflake.Sonyflake
	sfOnce sync.Once
)

// internalInit initializes Sonyflake only once
func internalInit(machineIDPtr *uint16) error {
	var initErr error
	sfOnce.Do(func() {
		st := sonyflake.Settings{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		if machineIDPtr != nil {
			st.MachineID = func() (uint16, error) {
				return *machineIDPtr, nil
			}
		}
		sf = sonyflake.NewSonyflake(st)
		if sf == nil {
			initErr = fmt.Errorf("failed to create Sonyflake")
		}
	})
	return initErr
}

// GetSonyflakeID generates a Sonyflake unique ID (uint64)
// machineIDPtr: optional machine ID (can pass nil)
// On first call, initializes Sonyflake
func GetSonyflakeID(machineIDPtr *uint16) (uint64, error) {
	if sf == nil {
		if err := internalInit(machineIDPtr); err != nil {
			return 0, err
		}
	}
	return sf.NextID()
}

// GetSonyflakeIDHex generates a Sonyflake unique ID (hex string)
func GetSonyflakeIDHex(machineIDPtr *uint16) (string, error) {
	id, err := GetSonyflakeID(machineIDPtr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%016x", id), nil
}
