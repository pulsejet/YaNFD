/* YaNFD - Yet another NDN Forwarding Daemon
 *
 * Copyright (C) 2020-2021 Eric Newberry.
 *
 * This file is licensed under the terms of the MIT License, as found in LICENSE.md.
 */

package face

import (
	"encoding/binary"
	"strconv"

	"github.com/eric135/YaNFD/core"
	"github.com/eric135/YaNFD/dispatch"
	"github.com/eric135/YaNFD/fw"
	"github.com/eric135/YaNFD/ndn"
	"github.com/eric135/YaNFD/ndn/tlv"
)

// LinkService is an interface for link service implementations
type LinkService interface {
	String() string
	SetFaceID(faceID int)

	FaceID() int
	LocalURI() ndn.URI
	RemoteURI() ndn.URI
	Scope() ndn.Scope
	MTU() int

	State() ndn.State

	// Main entry point for running face thread
	Run()

	// SendPacket Add a packet to the send queue for this link service
	SendPacket(packet *ndn.PendingPacket)
	handleIncomingFrame(frame []byte)

	Close()
	tellTransportQuit()
	GetHasQuit() chan bool
}

// linkServiceBase is the type upon which all link service implementations should be built
type linkServiceBase struct {
	faceID           int
	transport        transport
	HasQuit          chan bool
	hasImplQuit      chan bool
	hasTransportQuit chan bool
	sendQueue        chan *ndn.PendingPacket
}

func (l *linkServiceBase) SetFaceID(faceID int) {
	l.faceID = faceID
	if l.transport != nil {
		l.transport.setFaceID(faceID)
	}
}

func (l *linkServiceBase) tellTransportQuit() {
	l.hasTransportQuit <- true
}

// GetHasQuit returns the channel that indicates when the face has quit.
func (l *linkServiceBase) GetHasQuit() chan bool {
	return l.HasQuit
}

//
// "Constructors" and threading
//

func (l *linkServiceBase) makeLinkServiceBase() {
	l.HasQuit = make(chan bool)
	l.hasImplQuit = make(chan bool)
	l.hasTransportQuit = make(chan bool)
	l.sendQueue = make(chan *ndn.PendingPacket, core.FaceQueueSize)
}

//
// Getters
//

// FaceID returns the ID of the face
func (l *linkServiceBase) FaceID() int {
	return l.faceID
}

// LocalURI returns the local URI of the underlying transport
func (l *linkServiceBase) LocalURI() ndn.URI {
	return l.transport.LocalURI()
}

// RemoteURI returns the remote URI of the underlying transport
func (l *linkServiceBase) RemoteURI() ndn.URI {
	return l.transport.RemoteURI()
}

// Scope returns the scope of the underlying transport
func (l *linkServiceBase) Scope() ndn.Scope {
	return l.transport.Scope()
}

// MTU returns the MTU of the underlying transport
func (l *linkServiceBase) MTU() int {
	return l.transport.MTU()
}

// State returns the state of the underlying transport
func (l *linkServiceBase) State() ndn.State {
	return l.transport.State()
}

//
// Forwarding pipeline
//

// SendPacket adds a packet to the send queue for this link service
func (l *linkServiceBase) SendPacket(packet *ndn.PendingPacket) {
	/*if l.State() != Up {
		core.LogWarn(l, "Cannot send packet on down face - DROP")
	}*/

	select {
	case l.sendQueue <- packet:
		// Packet queued successfully
		core.LogTrace(l, "Queued packet for Link Service")
	default:
		// Drop packet due to congestion
		core.LogWarn(l, "Dropped packet due to congestion")

		// TODO: Signal congestion
	}
}

func (l *linkServiceBase) dispatchIncomingPacket(netPacket *ndn.PendingPacket) {
	// Hand off to network layer by dispatching to appropriate forwarding thread(s)
	switch netPacket.Wire.Type() {
	case tlv.Interest:
		interest, err := ndn.DecodeInterest(netPacket.Wire)
		if err != nil {
			core.LogError(l, "Unable to decode Interest ("+err.Error()+") - DROP")
			break
		}
		thread := fw.HashNameToFwThread(interest.Name())
		core.LogTrace(l, "Dispatched Interest to thread "+strconv.Itoa(thread))
		fw.Threads[thread].QueueInterest(netPacket)
	case tlv.Data:
		if len(netPacket.PitToken) == 2 {
			// Decode PitToken. If it's for us, it's a uint16.
			pitToken := binary.BigEndian.Uint16(netPacket.PitToken)
			fwThread := dispatch.GetFWThread(int(pitToken))
			if fwThread == nil {
				// If invalid PIT token present, drop.
				core.LogError(l, "Invalid PIT token attached to Data packet - DROP")
				break
			}
			// If valid PIT token present, dispatch to that thread.
			core.LogTrace(l, "Dispatched Interest to thread "+strconv.FormatUint(uint64(pitToken), 10))
			fwThread.QueueData(netPacket)
		} else {
			// Otherwise, dispatch to threads matching every prefix.

			data, err := ndn.DecodeData(netPacket.Wire, false)
			if err != nil {
				core.LogError(l, "Unable to decode Data ("+err.Error()+") - DROP")
				break
			}

			core.LogDebug(l, "Missing PIT token from Data packet - performing prefix dispatching")
			for _, thread := range fw.HashNameToAllPrefixFwThreads(data.Name()) {
				core.LogTrace(l, "Prefix dispatched Data packet to thread "+strconv.Itoa(thread))
				fw.Threads[thread].QueueData(netPacket)
			}
		}
	default:
		core.LogError(l, "Cannot dispatch packet of unknown type "+strconv.FormatUint(uint64(netPacket.Wire.Type()), 10))
	}
}

func (l *linkServiceBase) Close() {
	l.transport.changeState(ndn.Down)
}
