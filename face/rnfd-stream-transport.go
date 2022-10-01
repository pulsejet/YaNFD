/* YaNFD - Yet another NDN Forwarding Daemon
 *
 * Copyright (C) 2020-2021 Eric Newberry.
 *
 * This file is licensed under the terms of the MIT License, as found in LICENSE.md.
 */

package face

import (
	"net"
	"strconv"

	"github.com/named-data/YaNFD/core"
	"github.com/named-data/YaNFD/ndn"
	"github.com/named-data/YaNFD/ndn/tlv"
)

// RNFDStreamTransport is a fake UDP transport
type RNFDStreamTransport struct {
	remoteAddr string
	conn       *net.UnixConn
	transportBase
}

// MakeRNFDStreamTransport creates a rNFD stream transport.
func MakeRNFDStreamTransport(remoteURI *ndn.URI, localURI *ndn.URI, conn *net.UnixConn) (*RNFDStreamTransport, error) {
	t := new(RNFDStreamTransport)
	t.makeTransportBase(remoteURI, localURI, PersistencyPersistent, ndn.Local, ndn.PointToPoint, tlv.MaxNDNPacketSize)

	// Set connection
	t.conn = conn

	t.changeState(ndn.Up)

	return t, nil
}

func (t *RNFDStreamTransport) String() string {
	return "RNFDStreamTransport, FaceID=" + strconv.FormatUint(t.faceID, 10) + ", RemoteURI=" + t.remoteURI.String() + ", LocalURI=" + t.localURI.String()
}

// SetPersistency changes the persistency of the face.
func (t *RNFDStreamTransport) SetPersistency(persistency Persistency) bool {
	if persistency == t.persistency {
		return true
	}

	if persistency == PersistencyPersistent {
		t.persistency = persistency
		return true
	}

	return false
}

// GetSendQueueSize returns the current size of the send queue.
func (t *RNFDStreamTransport) GetSendQueueSize() uint64 {
	return 0
}

func (t *RNFDStreamTransport) sendFrame(frame []byte) {
	if len(frame) > t.MTU() {
		core.LogWarn(t, "Attempted to send frame larger than MTU - DROP")
		return
	}

	// Create new TLV block with frame
	wire := tlv.NewEmptyBlock(6)
	wire.Append(tlv.NewBlock(4, []byte(t.remoteAddr)))
	wire.Append(tlv.NewBlock(21, frame))
	wire.Encode()
	bytes, err := wire.Wire()
	if err != nil {
		core.LogError(t, "Error encoding TLV block for frame - DROP")
		return
	}

	core.LogDebug(t, "Sending frame of length ", len(bytes))
	_, err = t.conn.Write(bytes)
	if err != nil {
		core.LogWarn(t, "Unable to send on socket - DROP and Face DOWN")
		t.changeState(ndn.Down)
	}

	t.nOutBytes += uint64(len(frame))
}

func (t *RNFDStreamTransport) receiveFrame(frame []byte) {
	t.linkService.handleIncomingFrame(frame)
}

func (t *RNFDStreamTransport) runReceive() {

}

func (t *RNFDStreamTransport) changeState(new ndn.State) {
	if t.state == new {
		return
	}

	core.LogInfo(t, "state: ", t.state, " -> ", new)
	t.state = new

	if t.state != ndn.Up {
		t.hasQuit <- true
		// Stop link service
		t.linkService.tellTransportQuit()
		FaceTable.Remove(t.faceID)
	}
}
