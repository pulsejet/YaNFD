/* YaNFD - Yet another NDN Forwarding Daemon
 *
 * Copyright (C) 2020-2021 Eric Newberry.
 *
 * This file is licensed under the terms of the MIT License, as found in LICENSE.md.
 */

package fw

import (
	"reflect"

	"github.com/named-data/YaNFD/core"
	"github.com/named-data/YaNFD/ndn_defn"
	"github.com/named-data/YaNFD/table"
	enc "github.com/zjkmxy/go-ndn/pkg/encoding"
)

// Multicast is a forwarding strategy that forwards Interests to all nexthop faces.
type Multicast struct {
	StrategyBase
}

func init() {
	strategyTypes = append(strategyTypes, reflect.TypeOf(new(Multicast)))
	StrategyVersions["multicast"] = []uint64{1}
}

func (s *Multicast) Instantiate(fwThread *Thread) {
	s.NewStrategyBase(fwThread, enc.Component{
		Typ: enc.TypeGenericNameComponent, Val: []byte("multicast"),
	}, 1, "Multicast")
}

func (s *Multicast) AfterContentStoreHit(
	packet *ndn_defn.PendingPacket,
	pitEntry table.PitEntry,
	inFace uint64,
) {
	core.LogTrace(s, "AfterContentStoreHit: Forwarding content store hit Data=", packet.NameCache, " to FaceID=", inFace)
	s.SendData(packet, pitEntry, inFace, 0) // 0 indicates ContentStore is source
}

func (s *Multicast) AfterReceiveData(
	packet *ndn_defn.PendingPacket,
	pitEntry table.PitEntry,
	inFace uint64,
) {
	core.LogTrace(s, "AfterReceiveData: Data=", packet.NameCache, ", ", len(pitEntry.InRecords()), " In-Records")
	for faceID := range pitEntry.InRecords() {
		core.LogTrace(s, "AfterReceiveData: Forwarding Data=", packet.NameCache, " to FaceID=", faceID)
		s.SendData(packet, pitEntry, faceID, inFace)
	}
}

func (s *Multicast) AfterReceiveInterest(
	packet *ndn_defn.PendingPacket,
	pitEntry table.PitEntry,
	inFace uint64,
	nexthops []*table.FibNextHopEntry,
) {
	if len(nexthops) == 0 {
		core.LogDebug(s, "AfterReceiveInterest: No nexthop for Interest=", packet.NameCache, " - DROP")
		return
	}

	for _, nexthop := range nexthops {
		core.LogTrace(s, "AfterReceiveInterest: Forwarding Interest=", packet.NameCache, " to FaceID=", nexthop.Nexthop)
		s.SendInterest(packet, pitEntry, nexthop.Nexthop, inFace)
	}
}

func (s *Multicast) BeforeSatisfyInterest(pitEntry table.PitEntry, inFace uint64) {
	// This does nothing in Multicast
}
