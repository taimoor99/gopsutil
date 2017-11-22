// +build darwin
// +build !cgo

package host

import "github.com/taimoor99/gopsutil/internal/common"

func SensorsTemperatures() ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
