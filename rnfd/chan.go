package rnfd

import (
	"github.com/named-data/YaNFD/core"
	"github.com/named-data/YaNFD/ndn"
)

// global channel for rnfd mgmt
var RnfdMgmtChan chan interface{}

// RNFD mgmt messages
type InsertNextHopMsg struct {
	Name   *ndn.Name
	FaceID uint64
	Cost   uint64
}

// Send to rnfd
func SendToRnfd(msg interface{}) {
	select {
	case RnfdMgmtChan <- msg:
	default:
		core.LogError("Mgmt dropped due to full queue")
	}
}

// Add FIB entry
func InsertNextHop(name *ndn.Name, faceID uint64, cost uint64) {
	SendToRnfd(&InsertNextHopMsg{name, faceID, cost})
}
