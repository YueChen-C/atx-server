package proto

import (
	"encoding/json"
	"time"

	"github.com/openatx/androidutils"
)

type MessageType int

const (
	DeviceInfoMessage = MessageType(0)
	PingMessage       = MessageType(1)
)

type CommonMessage struct {
	Type MessageType
	Data interface{}
}

func (m *CommonMessage) MarshalJSON() []byte {
	data, _ := json.Marshal(m)
	return data
}

type CpuInfo struct {
	Cores    int    `json:"cores"`
	Hardware string `json:"hardware"`
}

type MemoryInfo struct {
	Total  int    `json:"total"` // unit kB
	Around string `json:"around,omitempty"`
}

type DeviceInfo struct {
	Udid         string                `json:"udid,omitempty"`       // Unique device identifier
	PropertyId   string                `json:"propertyId,omitempty"` // For device managerment, eg: HIH-PHO-1122
	Version      string                `json:"version,omitempty"`    // ro.build.version.release
	Serial       string                `json:"serial,omitempty"`     // ro.serialno
	Brand        string                `json:"brand,omitempty"`      // ro.product.brand
	Model        string                `json:"model,omitempty"`      // ro.product.model
	HWAddr       string                `json:"hwaddr,omitempty"`     // persist.sys.wifi.mac
	IP           string                `json:"ip,omitempty"`
	Sdk          int                   `json:"sdk,omitempty"`
	AgentVersion string                `json:"agentVersion,omitempty"`
	Display      *androidutils.Display `json:"display,omitempty"`
	Battery      *androidutils.Battery `json:"battery,omitempty"`
	Memory       *MemoryInfo           `json:"memory,omitempty"` // proc/meminfo
	Cpu          *CpuInfo              `json:"cpu,omitempty"`    // proc/cpuinfo

	ConnectionCount   int       `json:"-"` // > 1 happended when phone redial server
	Reserved          string    `json:"reserved,omitempty"`
	CreatedAt         time.Time `json:"-" gorethink:"createdAt,omitempty"`
	PresenceChangedAt time.Time `json:"presenceChangedAt,omitempty"`

	Ready   *bool `json:"ready,omitempty"`
	Present *bool `json:"present,omitempty"`
	Using   *bool `json:"using,omitempty"`

	Product *Product `json:"product,omitempty" gorethink:"product_id,reference,omitempty" gorethink_ref:"id"`
}

type Product struct {
	Id       string  `json:"id" gorethink:"id,omitempty"`
	Brand    string  `json:"brand" gorethink:"brand,omitempty"`
	Model    string  `json:"model" gorethink:"model,omitempty"`
	Name     string  `json:"name" gorethink:"name,omitempty"`
	Coverage float32 `json:"coverage" gorethink:"coverage,omitempty"`
	Gpu      string  `json:"gpu,omitempty"`
	Cpu      string  `json:"cpu,omitempty"`
	// AntutuScore int     `json:"antutuScore,omitempty"`
	// Alias: 高配 or 低配
}
