// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transit_realtime.proto

package transit_realtime

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The direction the train is moving.
type NyctTripDescriptor_Direction int32

const (
	NyctTripDescriptor_NORTH NyctTripDescriptor_Direction = 1
	NyctTripDescriptor_EAST  NyctTripDescriptor_Direction = 2
	NyctTripDescriptor_SOUTH NyctTripDescriptor_Direction = 3
	NyctTripDescriptor_WEST  NyctTripDescriptor_Direction = 4
)

var NyctTripDescriptor_Direction_name = map[int32]string{
	1: "NORTH",
	2: "EAST",
	3: "SOUTH",
	4: "WEST",
}
var NyctTripDescriptor_Direction_value = map[string]int32{
	"NORTH": 1,
	"EAST":  2,
	"SOUTH": 3,
	"WEST":  4,
}

func (x NyctTripDescriptor_Direction) Enum() *NyctTripDescriptor_Direction {
	p := new(NyctTripDescriptor_Direction)
	*p = x
	return p
}
func (x NyctTripDescriptor_Direction) String() string {
	return proto.EnumName(NyctTripDescriptor_Direction_name, int32(x))
}
func (x *NyctTripDescriptor_Direction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NyctTripDescriptor_Direction_value, data, "NyctTripDescriptor_Direction")
	if err != nil {
		return err
	}
	*x = NyctTripDescriptor_Direction(value)
	return nil
}
func (NyctTripDescriptor_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{2, 0}
}

type TripReplacementPeriod struct {
	// The replacement period is for this route
	RouteId *string `protobuf:"bytes,1,opt,name=route_id,json=routeId" json:"route_id,omitempty"`
	// The start time is omitted, the end time is currently now + 30 minutes for
	// all routes of the A division
	ReplacementPeriod *TimeRange `protobuf:"bytes,2,opt,name=replacement_period,json=replacementPeriod" json:"replacement_period,omitempty"`
	XXX_unrecognized  []byte     `json:"-"`
}

func (m *TripReplacementPeriod) Reset()                    { *m = TripReplacementPeriod{} }
func (m *TripReplacementPeriod) String() string            { return proto.CompactTextString(m) }
func (*TripReplacementPeriod) ProtoMessage()               {}
func (*TripReplacementPeriod) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TripReplacementPeriod) GetRouteId() string {
	if m != nil && m.RouteId != nil {
		return *m.RouteId
	}
	return ""
}

func (m *TripReplacementPeriod) GetReplacementPeriod() *TimeRange {
	if m != nil {
		return m.ReplacementPeriod
	}
	return nil
}

// NYCT Subway extensions for the feed header
type NyctFeedHeader struct {
	// Version of the NYCT Subway extensions
	// The current version is 1.0
	NyctSubwayVersion *string `protobuf:"bytes,1,req,name=nyct_subway_version,json=nyctSubwayVersion" json:"nyct_subway_version,omitempty"`
	// For the NYCT Subway, the GTFS-realtime feed replaces any scheduled
	// trip within the trip_replacement_period.
	// This feed is a full dataset, it contains all trips starting
	// in the trip_replacement_period. If a trip from the static GTFS is not
	// found in the GTFS-realtime feed, it should be considered as cancelled.
	// The replacement period can be different for each route, so here is
	// a list of the routes where the trips in the feed replace all
	// scheduled trips within the replacement period.
	TripReplacementPeriod []*TripReplacementPeriod `protobuf:"bytes,2,rep,name=trip_replacement_period,json=tripReplacementPeriod" json:"trip_replacement_period,omitempty"`
	XXX_unrecognized      []byte                   `json:"-"`
}

func (m *NyctFeedHeader) Reset()                    { *m = NyctFeedHeader{} }
func (m *NyctFeedHeader) String() string            { return proto.CompactTextString(m) }
func (*NyctFeedHeader) ProtoMessage()               {}
func (*NyctFeedHeader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *NyctFeedHeader) GetNyctSubwayVersion() string {
	if m != nil && m.NyctSubwayVersion != nil {
		return *m.NyctSubwayVersion
	}
	return ""
}

func (m *NyctFeedHeader) GetTripReplacementPeriod() []*TripReplacementPeriod {
	if m != nil {
		return m.TripReplacementPeriod
	}
	return nil
}

// NYCT Subway extensions for the trip descriptor
type NyctTripDescriptor struct {
	// The nyct_train_id is meant for internal use only. It provides an
	// easy way to associated GTFS-realtime trip identifiers with NYCT rail
	// operations identifier
	//
	// The ATS office system assigns unique train identification (Train ID) to
	// each train operating within or ready to enter the mainline of the
	// monitored territory. An example of this is 06 0123+ PEL/BBR and is decoded
	// as follows:
	//
	// The first character represents the trip type designator. 0 identifies a
	// scheduled revenue trip. Other revenue trip values that are a result of a
	// change to the base schedule include; [= reroute], [/ skip stop], [$ turn
	// train] also known as shortly lined service.
	//
	// The second character 6 represents the trip line i.e. number 6 train The
	// third set of characters identify the decoded origin time. The last
	// character may be blank "on the whole minute" or + "30 seconds"
	//
	// Note: Origin times will not change when there is a trip type change.  This
	// is followed by a three character "Origin Location" / "Destination
	// Location"
	TrainId *string `protobuf:"bytes,1,opt,name=train_id,json=trainId" json:"train_id,omitempty"`
	// This trip has been assigned to a physical train. If true, this trip is
	// already underway or most likely will depart shortly.
	//
	// Train Assignment is a function of the Automatic Train Supervision (ATS)
	// office system used by NYCT Rail Operations to monitor and track train
	// movements. ATS provides the ability to "assign" the nyct_train_id
	// attribute when a physical train is at its origin terminal. These assigned
	// trips have the is_assigned field set in the TripDescriptor.
	//
	// When a train is at a terminal but has not been given a work program it is
	// declared unassigned and is tagged as such. Unassigned trains can be moved
	// to a storage location or assigned a nyct_train_id when a determination for
	// service is made.
	IsAssigned *bool `protobuf:"varint,2,opt,name=is_assigned,json=isAssigned" json:"is_assigned,omitempty"`
	// Uptown and Bronx-bound trains are moving NORTH.
	// Times Square Shuttle to Grand Central is also northbound.
	//
	// Downtown and Brooklyn-bound trains are moving SOUTH.
	// Times Square Shuttle to Times Square is also southbound.
	//
	// EAST and WEST are not used currently.
	Direction        *NyctTripDescriptor_Direction `protobuf:"varint,3,opt,name=direction,enum=NyctTripDescriptor_Direction" json:"direction,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *NyctTripDescriptor) Reset()                    { *m = NyctTripDescriptor{} }
func (m *NyctTripDescriptor) String() string            { return proto.CompactTextString(m) }
func (*NyctTripDescriptor) ProtoMessage()               {}
func (*NyctTripDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *NyctTripDescriptor) GetTrainId() string {
	if m != nil && m.TrainId != nil {
		return *m.TrainId
	}
	return ""
}

func (m *NyctTripDescriptor) GetIsAssigned() bool {
	if m != nil && m.IsAssigned != nil {
		return *m.IsAssigned
	}
	return false
}

func (m *NyctTripDescriptor) GetDirection() NyctTripDescriptor_Direction {
	if m != nil && m.Direction != nil {
		return *m.Direction
	}
	return NyctTripDescriptor_NORTH
}

// NYCT Subway extensions for the stop time update
type NyctStopTimeUpdate struct {
	// Provides the planned station arrival track. The following is the Manhattan
	// track configurations:
	// 1: southbound local
	// 2: southbound express
	// 3: northbound express
	// 4: northbound local
	//
	// In the Bronx (except Dyre Ave line)
	// M: bi-directional express (in the AM express to Manhattan, in the PM
	// express away).
	//
	// The Dyre Ave line is configured:
	// 1: southbound
	// 2: northbound
	// 3: bi-directional
	ScheduledTrack *string `protobuf:"bytes,1,opt,name=scheduled_track,json=scheduledTrack" json:"scheduled_track,omitempty"`
	// This is the actual track that the train is operating on and can be used to
	// determine if a train is operating according to its current schedule
	// (plan).
	//
	// The actual track is known only shortly before the train reaches a station,
	// typically not before it leaves the previous station. Therefore, the NYCT
	// feed sets this field only for the first station of the remaining trip.
	//
	// Different actual and scheduled track is the result of manually rerouting a
	// train off it scheduled path.  When this occurs, prediction data may become
	// unreliable since the train is no longer operating in accordance to its
	// schedule.  The rules engine for the 'countdown' clocks will remove this
	// train from all schedule stations.
	ActualTrack      *string `protobuf:"bytes,2,opt,name=actual_track,json=actualTrack" json:"actual_track,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NyctStopTimeUpdate) Reset()                    { *m = NyctStopTimeUpdate{} }
func (m *NyctStopTimeUpdate) String() string            { return proto.CompactTextString(m) }
func (*NyctStopTimeUpdate) ProtoMessage()               {}
func (*NyctStopTimeUpdate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *NyctStopTimeUpdate) GetScheduledTrack() string {
	if m != nil && m.ScheduledTrack != nil {
		return *m.ScheduledTrack
	}
	return ""
}

func (m *NyctStopTimeUpdate) GetActualTrack() string {
	if m != nil && m.ActualTrack != nil {
		return *m.ActualTrack
	}
	return ""
}

var E_NyctFeedHeader = &proto.ExtensionDesc{
	ExtendedType:  (*FeedHeader)(nil),
	ExtensionType: (*NyctFeedHeader)(nil),
	Field:         1001,
	Name:          "nyct_feed_header",
	Tag:           "bytes,1001,opt,name=nyct_feed_header,json=nyctFeedHeader",
	Filename:      "transit_realtime.proto",
}

var E_NyctTripDescriptor = &proto.ExtensionDesc{
	ExtendedType:  (*TripDescriptor)(nil),
	ExtensionType: (*NyctTripDescriptor)(nil),
	Field:         1001,
	Name:          "nyct_trip_descriptor",
	Tag:           "bytes,1001,opt,name=nyct_trip_descriptor,json=nyctTripDescriptor",
	Filename:      "transit_realtime.proto",
}

var E_NyctStopTimeUpdate = &proto.ExtensionDesc{
	ExtendedType:  (*TripUpdate_StopTimeUpdate)(nil),
	ExtensionType: (*NyctStopTimeUpdate)(nil),
	Field:         1001,
	Name:          "nyct_stop_time_update",
	Tag:           "bytes,1001,opt,name=nyct_stop_time_update,json=nyctStopTimeUpdate",
	Filename:      "transit_realtime.proto",
}

func init() {
	proto.RegisterType((*TripReplacementPeriod)(nil), "TripReplacementPeriod")
	proto.RegisterType((*NyctFeedHeader)(nil), "NyctFeedHeader")
	proto.RegisterType((*NyctTripDescriptor)(nil), "NyctTripDescriptor")
	proto.RegisterType((*NyctStopTimeUpdate)(nil), "NyctStopTimeUpdate")
	proto.RegisterEnum("NyctTripDescriptor_Direction", NyctTripDescriptor_Direction_name, NyctTripDescriptor_Direction_value)
	proto.RegisterExtension(E_NyctFeedHeader)
	proto.RegisterExtension(E_NyctTripDescriptor)
	proto.RegisterExtension(E_NyctStopTimeUpdate)
}

func init() { proto.RegisterFile("transit_realtime.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb5, 0x49, 0x51, 0x93, 0x09, 0x4a, 0x53, 0x87, 0x94, 0x40, 0xa9, 0x08, 0xb9, 0x10,
	0x09, 0xb1, 0x87, 0x48, 0x5c, 0xc2, 0xa9, 0xa8, 0x45, 0x29, 0x87, 0x14, 0x39, 0x5b, 0xe0, 0x66,
	0xcc, 0x7a, 0x9a, 0x58, 0x24, 0xeb, 0x95, 0x77, 0x02, 0xca, 0x85, 0x67, 0xe0, 0xa1, 0x78, 0x19,
	0xde, 0x02, 0xd9, 0x9b, 0x34, 0x6c, 0xb3, 0x47, 0x7f, 0x33, 0xf6, 0xef, 0xf9, 0xff, 0x81, 0x13,
	0xb2, 0x32, 0xc9, 0x34, 0x09, 0x8b, 0x72, 0x41, 0x7a, 0x89, 0x61, 0x6a, 0x0d, 0x99, 0xa7, 0xed,
	0x19, 0xdd, 0x66, 0xaf, 0x8b, 0xb0, 0xff, 0x0b, 0x3a, 0x91, 0xd5, 0x29, 0xc7, 0x74, 0x21, 0x63,
	0x5c, 0x62, 0x42, 0x1f, 0xd1, 0x6a, 0xa3, 0xd8, 0x13, 0xa8, 0x59, 0xb3, 0x22, 0x14, 0x5a, 0x75,
	0x83, 0x5e, 0x30, 0xa8, 0xf3, 0x43, 0x7f, 0xbe, 0x52, 0xec, 0x03, 0x30, 0xbb, 0xeb, 0x17, 0xa9,
	0xbf, 0xd0, 0xad, 0xf4, 0x82, 0x41, 0x63, 0x78, 0x1a, 0xee, 0xa9, 0x47, 0x7a, 0x89, 0x5c, 0x26,
	0x33, 0xe4, 0xc7, 0xf6, 0xbe, 0x4c, 0xff, 0x77, 0x00, 0xcd, 0xc9, 0x3a, 0xa6, 0xf7, 0x88, 0x6a,
	0x8c, 0x52, 0xa1, 0x65, 0x21, 0xb4, 0x93, 0x75, 0x4c, 0x22, 0x5b, 0x7d, 0xfb, 0x29, 0xd7, 0xe2,
	0x07, 0xda, 0x4c, 0x9b, 0xa4, 0x1b, 0xf4, 0x2a, 0x83, 0x3a, 0x3f, 0x76, 0xa5, 0xa9, 0xaf, 0x7c,
	0xca, 0x0b, 0x6c, 0x02, 0x8f, 0xc9, 0xea, 0x54, 0x94, 0xfe, 0xa9, 0x3a, 0x68, 0x0c, 0x4f, 0xc2,
	0xd2, 0x11, 0x79, 0x87, 0xca, 0x70, 0xff, 0x4f, 0x00, 0xcc, 0x7d, 0xc9, 0x5d, 0xba, 0xc0, 0x2c,
	0xb6, 0x3a, 0x25, 0x63, 0x9d, 0x21, 0x64, 0xa5, 0x4e, 0xfe, 0x33, 0xc4, 0x9f, 0xaf, 0x14, 0x7b,
	0x0e, 0x0d, 0x9d, 0x09, 0x99, 0x65, 0x7a, 0x96, 0x60, 0xee, 0x44, 0x8d, 0x83, 0xce, 0xce, 0x37,
	0x84, 0xbd, 0x85, 0xba, 0xd2, 0x16, 0x63, 0x72, 0x83, 0x54, 0x7b, 0xc1, 0xa0, 0x39, 0x3c, 0x0b,
	0xf7, 0x35, 0xc2, 0x8b, 0x6d, 0x13, 0xdf, 0xf5, 0xf7, 0xdf, 0x40, 0xfd, 0x8e, 0xb3, 0x3a, 0x3c,
	0x98, 0x5c, 0xf3, 0x68, 0xdc, 0x0a, 0x58, 0x0d, 0x0e, 0x2e, 0xcf, 0xa7, 0x51, 0xab, 0xe2, 0xe0,
	0xf4, 0xfa, 0x26, 0x1a, 0xb7, 0xaa, 0x0e, 0x7e, 0xbe, 0x9c, 0x46, 0xad, 0x83, 0xfe, 0xd7, 0x7c,
	0x8a, 0x29, 0x99, 0xd4, 0x25, 0x70, 0x93, 0x2a, 0x49, 0xc8, 0x5e, 0xc2, 0x51, 0x16, 0xcf, 0x51,
	0xad, 0x16, 0xa8, 0x04, 0x59, 0x19, 0x7f, 0xdf, 0x0c, 0xd3, 0xbc, 0xc3, 0x91, 0xa3, 0xec, 0x05,
	0x3c, 0x94, 0x31, 0xad, 0xe4, 0x62, 0xd3, 0x55, 0xf1, 0x5d, 0x8d, 0x9c, 0xf9, 0x96, 0xd1, 0x17,
	0x68, 0xf9, 0xa0, 0x6e, 0x11, 0x95, 0x98, 0xe7, 0xe1, 0x3d, 0xdb, 0xcf, 0x7f, 0x17, 0x6d, 0xf7,
	0xef, 0xa1, 0x5f, 0x92, 0xa3, 0xb0, 0x18, 0x39, 0x6f, 0x26, 0x85, 0xf3, 0x68, 0x0e, 0x8f, 0xfc,
	0xcb, 0x3e, 0x57, 0xb5, 0xcb, 0xa0, 0x57, 0xb2, 0x5d, 0x05, 0x07, 0xb7, 0x0a, 0xed, 0x12, 0x77,
	0x39, 0x4b, 0xf6, 0xd8, 0x88, 0xa0, 0x93, 0x2f, 0x1b, 0x99, 0x54, 0xb8, 0x07, 0xc5, 0x2a, 0x37,
	0xea, 0x55, 0xb9, 0x54, 0x6e, 0x63, 0x58, 0x74, 0xb5, 0xa8, 0x5a, 0xac, 0xe5, 0xaa, 0x45, 0xf6,
	0xee, 0x0c, 0x4e, 0x63, 0xb3, 0x0c, 0x67, 0xc6, 0xcc, 0x16, 0xb8, 0x15, 0x0a, 0xb7, 0x42, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x28, 0x63, 0x1a, 0xd5, 0xc2, 0x03, 0x00, 0x00,
}