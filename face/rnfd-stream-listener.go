/* YaNFD - Yet another NDN Forwarding Daemon
 *
 * Copyright (C) 2020-2021 Eric Newberry.
 *
 * This file is licensed under the terms of the MIT License, as found in LICENSE.md.
 */

package face

import (
	"errors"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/named-data/YaNFD/core"
	"github.com/named-data/YaNFD/ndn"
	"github.com/named-data/YaNFD/ndn/tlv"
	"github.com/named-data/YaNFD/rnfd"
)

// RNFDStreamListener listens for incoming Unix stream connections.
type RNFDStreamListener struct {
	conn         net.Listener
	localURI     *ndn.URI
	HasQuit      chan bool
	transportMap map[string]*RNFDStreamTransport
}

// MakeRNFDStreamListener constructs a RNFDStreamListener.
func MakeRNFDStreamListener(localURI *ndn.URI) (*RNFDStreamListener, error) {
	localURI.Canonize()
	if !localURI.IsCanonical() || localURI.Scheme() != "unix" {
		return nil, core.ErrNotCanonical
	}

	l := new(RNFDStreamListener)
	l.localURI = localURI
	l.HasQuit = make(chan bool, 1)
	l.transportMap = make(map[string]*RNFDStreamTransport)
	rnfd.RnfdMgmtChan = make(chan interface{}, 100)
	return l, nil
}

func (l *RNFDStreamListener) String() string {
	return "RNFDStreamListener, " + l.localURI.String()
}

// Run starts the rNFD stream listener.
func (l *RNFDStreamListener) Run() {
	// Delete any existing socket
	os.Remove(l.localURI.Path())

	// Create listener
	var err error
	if l.conn, err = net.Listen(l.localURI.Scheme(), l.localURI.Path()); err != nil {
		core.LogFatal(l, "Unable to start rNFD stream listener: ", err)
	}

	// Set permissions to allow all local apps to communicate with us
	if err := os.Chmod(l.localURI.Path(), 0777); err != nil {
		core.LogFatal(l, "Unable to change permissions on rNFD stream listener: ", err)
	}

	core.LogInfo(l, "Listening")

	// Run accept loop
	for {
		newConn, err := l.conn.Accept()
		if err != nil {
			if err.Error() == "EOF" {
				// Must have failed due to being closed, so quit quietly
			} else {
				core.LogWarn(l, "Unable to accept connection: ", err)
			}
			break
		}

		go l.runReceive(newConn.(*net.UnixConn))
	}

	l.HasQuit <- true
}

// Close closes the RNFDStreamListener.
func (l *RNFDStreamListener) Close() {
	core.LogInfo(l, "Stopping listener")
	l.conn.Close()
}

func (l *RNFDStreamListener) processMgmtFrame(obj interface{}) ([]byte, error) {
	if obj == nil {
		return nil, errors.New("nil object")
	}

	switch obj.(type) {
	case *rnfd.InsertNextHopMsg:
		msg := obj.(*rnfd.InsertNextHopMsg)
		uri := FaceTable.Faces[msg.FaceID].RemoteURI()
		if uri.Scheme() != "udp4" {
			return nil, errors.New("invalid scheme")
		}
		// remove udp4:// from start
		remoteAddr := uri.String()[7:]

		b := tlv.NewEmptyBlock(1)                     // InsertNextHop
		b.Append(msg.Name.Encode())                   // Name
		b.Append(tlv.NewBlock(4, []byte(remoteAddr))) // UDP address
		b.Append(tlv.EncodeNNIBlock(5, msg.Cost))     // Cost
		b.Encode()
		return b.Wire()
	}

	return nil, errors.New("unknown message type")
}

// Accept connection and run receive thread
func (l *RNFDStreamListener) runReceive(conn *net.UnixConn) {
	core.LogTrace(l, "Starting receive thread")

	// Channel to quit writer
	quitChan := make(chan bool, 1)
	mutex := sync.Mutex{}

	// Read rnfdMgmtChan
	// WARN: only one connection can be active at a time
	go func() {
		core.LogInfo(l, "Starting rnfdMgmtChan thread")
		for {
			select {
			case frame := <-rnfd.RnfdMgmtChan:
				mutex.Lock()

				// Proess frame
				encodedFrame, err := l.processMgmtFrame(frame)
				if err != nil {
					core.LogWarn(l, "Unable to process frame: ", err)
					continue
				}
				b := tlv.NewBlock(3, encodedFrame)
				b.Encode()
				w, err := b.Wire()
				if err != nil {
					core.LogWarn(l, "Unable to encode TLV block: ", err)
					continue
				}
				_, err = conn.Write(w)
				if err != nil {
					core.LogWarn(l, "Unable to write to rNFD: ", err)
					return
				}
				core.LogInfo(l, "Wrote ", len(w), " mgmt bytes to rNFD")
				mutex.Unlock()
			case <-quitChan:
				return
			}
		}
	}()

	recvBuf := make([]byte, tlv.MaxNDNPacketSize*2)
	startPos := 0
	for {
		core.LogTrace(l, "Reading from socket")
		readSize, err := conn.Read(recvBuf[startPos:])
		startPos += readSize
		if err != nil {
			if err.Error() == "EOF" {
				core.LogWarn(l, "EOF - Face DOWN")
			} else {
				core.LogWarn(l, "Unable to read from rNFD socket (", err, ") - DROP and Face DOWN")
			}
			break
		}

		core.LogTrace(l, "Receive of size ", readSize)

		// Determine whether valid packet received
		tlvPos := 0
		for {
			if tlvPos >= startPos {
				startPos = 0
				break
			}

			_, _, tlvSize, err := tlv.DecodeTypeLength(recvBuf[tlvPos:])
			if err != nil {
				core.LogInfo(l, "Unable to process received packet: ", err)
				startPos = 0
				break
			} else if startPos >= tlvPos+tlvSize {
				// Packet was successfully received, send up to link service
				frame := recvBuf[tlvPos : tlvPos+tlvSize]
				l.processFrame(frame, conn, &mutex)
				tlvPos += tlvSize
			} else {
				if tlvPos > 0 {
					if startPos > tlvPos {
						// Move remaining data to beginning of buffer
						copy(recvBuf, recvBuf[tlvPos:startPos])
						startPos -= tlvPos
					} else {
						startPos = 0
					}
				}
				core.LogTrace(l, "Received packet is incomplete")
				break
			}
		}
	}

	// Close connection
	quitChan <- true
	conn.Close()
}

// Process frame received from rNFD
func (l *RNFDStreamListener) processFrame(frame []byte, conn *net.UnixConn, mutex *sync.Mutex) {
	// Decode TLV block
	tlvBlock, _, err := tlv.DecodeBlock(frame)
	if err != nil {
		core.LogWarn(l, "Unable to decode TLV block: ", err)
		return
	}

	err = tlvBlock.Parse()
	if err != nil {
		core.LogWarn(l, "Unable to parse TLV block: ", err)
		return
	}

	// iterate subelements
	udpAddr := ""
	for _, subElem := range tlvBlock.Subelements() {
		switch subElem.Type() {
		case 4:
			// UDP address
			udpAddr = string(subElem.Value())
			l.createTransport(udpAddr, conn, mutex)
		case 21:
			// NDN packet
			packet := subElem.Value()
			l.passPacket(packet, udpAddr)
		default:
			core.LogWarn(l, "Unknown TLV type: ", subElem.Type())
		}
	}
}

// Create transport for UDP address
func (l *RNFDStreamListener) createTransport(remoteAddr string, conn *net.UnixConn, mutex *sync.Mutex) {
	// Check if transport already exists
	if _, ok := l.transportMap[remoteAddr]; ok {
		return
	}

	// Create remote URI
	var remoteURI *ndn.URI
	host, port, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		core.LogWarn(l, "Unable to create face from ", remoteAddr, ": could not split host from port")
		return
	}
	portInt, _ := strconv.ParseUint(port, 10, 16)
	if err != nil {
		core.LogWarn(l, "Unable to create face from ", remoteAddr, ": could not split host from port")
		return
	}
	remoteURI = ndn.MakeUDPFaceURI(4, host, uint16(portInt))
	remoteURI.Canonize()
	if !remoteURI.IsCanonical() {
		core.LogWarn(l, "Unable to create face from ", remoteURI, ": remote URI is not canonical")
		return
	}

	// Create transport
	transport, err := MakeRNFDStreamTransport(remoteURI, l.localURI, conn, mutex)
	if err != nil {
		core.LogWarn(l, "Unable to create transport: ", err)
		return
	}
	transport.remoteAddr = remoteAddr

	newLinkService := MakeNDNLPLinkService(transport, MakeNDNLPLinkServiceOptions())
	if err != nil {
		core.LogError(l, "Failed to create new NDNLPv2 transport: ", err)
		return
	}

	core.LogInfo(l, "Accepting new rNFD stream face ", remoteURI)

	// Add face to table and start its thread
	FaceTable.Add(newLinkService)
	go newLinkService.Run(nil)

	// Sleep for a bit to allow face to be added to table
	time.Sleep(10 * time.Millisecond)

	// Add to local transport map
	l.transportMap[remoteAddr] = transport
}

// Pass packet to transport
func (l *RNFDStreamListener) passPacket(packet []byte, remoteAddr string) {
	// Find transport
	transport, ok := l.transportMap[remoteAddr]
	if !ok {
		core.LogWarn(l, "Unable to find transport for ", remoteAddr)
		return
	}

	// Pass packet to transport
	transport.receiveFrame(packet)
}
