package models

import (
	"fmt"
	"log"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

// ***************************** //
// ** ConnectionRequest Mgnmt ** //
// ***************************** //
func (req *ConnectionRequest) AttachGeoToRequest(geo *GeoIP) *ConnectionRequest {
	if req.Location.Latitude != 0 && req.Location.Longitude != 0 {
		req.IpLocation = geo.GetLocation()
		return req
	} else {
		req.Location = &Location{
			Latitude:  geo.Latitude,
			Longitude: geo.Longitude,
		}
		req.IpLocation = geo.GetLocation()
		return req
	}
}

func (req *ConnectionRequest) Latitude() float64 {
	loc := req.GetLocation()
	return loc.GetLatitude()
}

func (req *ConnectionRequest) Longitude() float64 {
	loc := req.GetLocation()
	return loc.GetLongitude()
}

func (g *GeoIP) GetLocation() *Location {
	return &Location{
		Name:        g.State,
		Continent:   g.Continent,
		CountryCode: g.Country,
		Latitude:    g.Latitude,
		Longitude:   g.Longitude,
		MinorOlc:    olc.Encode(float64(g.Latitude), float64(g.Longitude), 6),
		MajorOlc:    olc.Encode(float64(g.Latitude), float64(g.Longitude), 4),
	}
}

func (req *ConnectionRequest) IsDesktop() bool {
	return req.Device.Platform == Platform_MacOS || req.Device.Platform == Platform_Linux || req.Device.Platform == Platform_Windows
}

func (req *ConnectionRequest) IsMobile() bool {
	return req.Device.Platform == Platform_IOS || req.Device.Platform == Platform_Android
}

// ********************** //
// ** Lobby Management ** //
// ********************** //

// ^ Get Remote Point Info ^
func GetRemoteInfo(list []string) RemoteInfo {
	return RemoteInfo{
		Display: fmt.Sprintf("%s %s %s", list[0], list[1], list[2]),
		Topic:   fmt.Sprintf("%s-%s-%s", list[0], list[1], list[2]),
		Count:   int32(len(list)),
		IsJoin:  false,
		Words:   list,
	}
}

// ^ Returns as Lobby Buffer ^
func (l *Lobby) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes, nil
}

// ^ Add/Update Peer in Lobby ^
func (l *Lobby) Add(peer *Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Remove Peer from Lobby ^
func (l *Lobby) Delete(id peer.ID) {
	// Update Peer with new data
	delete(l.Peers, id.String())
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Sync Between Remote Peers Lobby ^
func (l *Lobby) Sync(ref *Lobby, remotePeer *Peer) {
	// Validate Lobbies are Different
	if l.Count != ref.Count {
		// Iterate Over List
		for id, peer := range ref.Peers {
			if l.User.IsNotPeerIDString(id) {
				l.Add(peer)
			}
		}
	}
	l.Add(remotePeer)
}
