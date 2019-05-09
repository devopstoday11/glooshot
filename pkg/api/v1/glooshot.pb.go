// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/glooshot/api/v1/glooshot.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ExperimentResult_State int32

const (
	// Experiment has not started
	ExperimentResult_Pending ExperimentResult_State = 0
	// Experiment started but threshold not met
	ExperimentResult_Started ExperimentResult_State = 1
	// Experiment failed, threshold was exceeded
	ExperimentResult_Failed ExperimentResult_State = 2
	// Experiment succeeded, duration elapsed
	// If duration is not specified, the Experiment will never
	// be marked Succeeded
	ExperimentResult_Succeeded ExperimentResult_State = 3
)

var ExperimentResult_State_name = map[int32]string{
	0: "Pending",
	1: "Started",
	2: "Failed",
	3: "Succeeded",
}

var ExperimentResult_State_value = map[string]int32{
	"Pending":   0,
	"Started":   1,
	"Failed":    2,
	"Succeeded": 3,
}

func (x ExperimentResult_State) String() string {
	return proto.EnumName(ExperimentResult_State_name, int32(x))
}

func (ExperimentResult_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{1, 0}
}

//
//Describes an Experiment that GlooShot should run
type Experiment struct {
	// the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	// indicates whether or not the spec is valid
	// set by glooshot, intended to be read by clients
	Status core.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status"`
	// configuration for the Experiment
	Spec *ExperimentSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// the result of the experiment
	Result               ExperimentResult `protobuf:"bytes,4,opt,name=result,proto3" json:"result"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Experiment) Reset()         { *m = Experiment{} }
func (m *Experiment) String() string { return proto.CompactTextString(m) }
func (*Experiment) ProtoMessage()    {}
func (*Experiment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{0}
}
func (m *Experiment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Experiment.Unmarshal(m, b)
}
func (m *Experiment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Experiment.Marshal(b, m, deterministic)
}
func (m *Experiment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Experiment.Merge(m, src)
}
func (m *Experiment) XXX_Size() int {
	return xxx_messageInfo_Experiment.Size(m)
}
func (m *Experiment) XXX_DiscardUnknown() {
	xxx_messageInfo_Experiment.DiscardUnknown(m)
}

var xxx_messageInfo_Experiment proto.InternalMessageInfo

func (m *Experiment) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Experiment) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Experiment) GetSpec() *ExperimentSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Experiment) GetResult() ExperimentResult {
	if m != nil {
		return m.Result
	}
	return ExperimentResult{}
}

type ExperimentResult struct {
	// the current state of the experiment as reported by glooshot
	State ExperimentResult_State `protobuf:"varint,1,opt,name=state,proto3,enum=glooshot.solo.io.ExperimentResult_State" json:"state,omitempty"`
	// arbitrary data summarizing a failure in case one occurred
	FailureReport map[string]string `protobuf:"bytes,2,rep,name=failure_report,json=failureReport,proto3" json:"failure_report,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// time the experiment was started
	TimeStarted time.Time `protobuf:"bytes,3,opt,name=time_started,json=timeStarted,proto3,stdtime" json:"time_started"`
	// the time the experiment completed
	TimeFinished         time.Time `protobuf:"bytes,4,opt,name=time_finished,json=timeFinished,proto3,stdtime" json:"time_finished"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExperimentResult) Reset()         { *m = ExperimentResult{} }
func (m *ExperimentResult) String() string { return proto.CompactTextString(m) }
func (*ExperimentResult) ProtoMessage()    {}
func (*ExperimentResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{1}
}
func (m *ExperimentResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExperimentResult.Unmarshal(m, b)
}
func (m *ExperimentResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExperimentResult.Marshal(b, m, deterministic)
}
func (m *ExperimentResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExperimentResult.Merge(m, src)
}
func (m *ExperimentResult) XXX_Size() int {
	return xxx_messageInfo_ExperimentResult.Size(m)
}
func (m *ExperimentResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExperimentResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExperimentResult proto.InternalMessageInfo

func (m *ExperimentResult) GetState() ExperimentResult_State {
	if m != nil {
		return m.State
	}
	return ExperimentResult_Pending
}

func (m *ExperimentResult) GetFailureReport() map[string]string {
	if m != nil {
		return m.FailureReport
	}
	return nil
}

func (m *ExperimentResult) GetTimeStarted() time.Time {
	if m != nil {
		return m.TimeStarted
	}
	return time.Time{}
}

func (m *ExperimentResult) GetTimeFinished() time.Time {
	if m != nil {
		return m.TimeFinished
	}
	return time.Time{}
}

type ExperimentSpec struct {
	// the faults this experiment will inject
	// if empty, Glooshit will run a "control" experiment with no faults injected
	Faults []*ExperimentSpec_InjectedFault `protobuf:"bytes,4,rep,name=faults,proto3" json:"faults,omitempty"`
	// conditions on which to stop the experiment and mark it as failed
	// at least one must be specified
	FailureConditions []*FailureCondition `protobuf:"bytes,5,rep,name=failure_conditions,json=failureConditions,proto3" json:"failure_conditions,omitempty"`
	// the duration for which to run the experiment
	// if missing or set to 0 the experiment will run indefinitely
	// only Experiments with a timeout can succeed
	Duration             *time.Duration `protobuf:"bytes,6,opt,name=duration,proto3,stdduration" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExperimentSpec) Reset()         { *m = ExperimentSpec{} }
func (m *ExperimentSpec) String() string { return proto.CompactTextString(m) }
func (*ExperimentSpec) ProtoMessage()    {}
func (*ExperimentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{2}
}
func (m *ExperimentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExperimentSpec.Unmarshal(m, b)
}
func (m *ExperimentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExperimentSpec.Marshal(b, m, deterministic)
}
func (m *ExperimentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExperimentSpec.Merge(m, src)
}
func (m *ExperimentSpec) XXX_Size() int {
	return xxx_messageInfo_ExperimentSpec.Size(m)
}
func (m *ExperimentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExperimentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExperimentSpec proto.InternalMessageInfo

func (m *ExperimentSpec) GetFaults() []*ExperimentSpec_InjectedFault {
	if m != nil {
		return m.Faults
	}
	return nil
}

func (m *ExperimentSpec) GetFailureConditions() []*FailureCondition {
	if m != nil {
		return m.FailureConditions
	}
	return nil
}

func (m *ExperimentSpec) GetDuration() *time.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

// decribes a single fault to  inject
type ExperimentSpec_InjectedFault struct {
	// if specified, the fault will only apply to requests sent from these services
	OriginServices []*core.ResourceRef `protobuf:"bytes,1,rep,name=origin_services,json=originServices,proto3" json:"origin_services,omitempty"`
	// if specified, the fault will only apply to requests sent to these services
	DestinationServices []*core.ResourceRef `protobuf:"bytes,2,rep,name=destination_services,json=destinationServices,proto3" json:"destination_services,omitempty"`
	// the type of fault to inject
	Fault                *v1.FaultInjection `protobuf:"bytes,3,opt,name=fault,proto3" json:"fault,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExperimentSpec_InjectedFault) Reset()         { *m = ExperimentSpec_InjectedFault{} }
func (m *ExperimentSpec_InjectedFault) String() string { return proto.CompactTextString(m) }
func (*ExperimentSpec_InjectedFault) ProtoMessage()    {}
func (*ExperimentSpec_InjectedFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{2, 0}
}
func (m *ExperimentSpec_InjectedFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExperimentSpec_InjectedFault.Unmarshal(m, b)
}
func (m *ExperimentSpec_InjectedFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExperimentSpec_InjectedFault.Marshal(b, m, deterministic)
}
func (m *ExperimentSpec_InjectedFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExperimentSpec_InjectedFault.Merge(m, src)
}
func (m *ExperimentSpec_InjectedFault) XXX_Size() int {
	return xxx_messageInfo_ExperimentSpec_InjectedFault.Size(m)
}
func (m *ExperimentSpec_InjectedFault) XXX_DiscardUnknown() {
	xxx_messageInfo_ExperimentSpec_InjectedFault.DiscardUnknown(m)
}

var xxx_messageInfo_ExperimentSpec_InjectedFault proto.InternalMessageInfo

func (m *ExperimentSpec_InjectedFault) GetOriginServices() []*core.ResourceRef {
	if m != nil {
		return m.OriginServices
	}
	return nil
}

func (m *ExperimentSpec_InjectedFault) GetDestinationServices() []*core.ResourceRef {
	if m != nil {
		return m.DestinationServices
	}
	return nil
}

func (m *ExperimentSpec_InjectedFault) GetFault() *v1.FaultInjection {
	if m != nil {
		return m.Fault
	}
	return nil
}

// a condition based on an observed prometheus metric
type FailureCondition struct {
	// Types that are valid to be assigned to FailureTrigger:
	//	*FailureCondition_WebhookUrl
	//	*FailureCondition_PrometheusTrigger
	FailureTrigger       isFailureCondition_FailureTrigger `protobuf_oneof:"failure_trigger"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *FailureCondition) Reset()         { *m = FailureCondition{} }
func (m *FailureCondition) String() string { return proto.CompactTextString(m) }
func (*FailureCondition) ProtoMessage()    {}
func (*FailureCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{3}
}
func (m *FailureCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailureCondition.Unmarshal(m, b)
}
func (m *FailureCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailureCondition.Marshal(b, m, deterministic)
}
func (m *FailureCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailureCondition.Merge(m, src)
}
func (m *FailureCondition) XXX_Size() int {
	return xxx_messageInfo_FailureCondition.Size(m)
}
func (m *FailureCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_FailureCondition.DiscardUnknown(m)
}

var xxx_messageInfo_FailureCondition proto.InternalMessageInfo

type isFailureCondition_FailureTrigger interface {
	isFailureCondition_FailureTrigger()
	Equal(interface{}) bool
}

type FailureCondition_WebhookUrl struct {
	WebhookUrl string `protobuf:"bytes,1,opt,name=webhook_url,json=webhookUrl,proto3,oneof"`
}
type FailureCondition_PrometheusTrigger struct {
	PrometheusTrigger *PrometheusTrigger `protobuf:"bytes,2,opt,name=prometheus_trigger,json=prometheusTrigger,proto3,oneof"`
}

func (*FailureCondition_WebhookUrl) isFailureCondition_FailureTrigger()        {}
func (*FailureCondition_PrometheusTrigger) isFailureCondition_FailureTrigger() {}

func (m *FailureCondition) GetFailureTrigger() isFailureCondition_FailureTrigger {
	if m != nil {
		return m.FailureTrigger
	}
	return nil
}

func (m *FailureCondition) GetWebhookUrl() string {
	if x, ok := m.GetFailureTrigger().(*FailureCondition_WebhookUrl); ok {
		return x.WebhookUrl
	}
	return ""
}

func (m *FailureCondition) GetPrometheusTrigger() *PrometheusTrigger {
	if x, ok := m.GetFailureTrigger().(*FailureCondition_PrometheusTrigger); ok {
		return x.PrometheusTrigger
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FailureCondition) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FailureCondition_OneofMarshaler, _FailureCondition_OneofUnmarshaler, _FailureCondition_OneofSizer, []interface{}{
		(*FailureCondition_WebhookUrl)(nil),
		(*FailureCondition_PrometheusTrigger)(nil),
	}
}

func _FailureCondition_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FailureCondition)
	// failure_trigger
	switch x := m.FailureTrigger.(type) {
	case *FailureCondition_WebhookUrl:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.WebhookUrl)
	case *FailureCondition_PrometheusTrigger:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PrometheusTrigger); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FailureCondition.FailureTrigger has unexpected type %T", x)
	}
	return nil
}

func _FailureCondition_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FailureCondition)
	switch tag {
	case 1: // failure_trigger.webhook_url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.FailureTrigger = &FailureCondition_WebhookUrl{x}
		return true, err
	case 2: // failure_trigger.prometheus_trigger
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrometheusTrigger)
		err := b.DecodeMessage(msg)
		m.FailureTrigger = &FailureCondition_PrometheusTrigger{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FailureCondition_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FailureCondition)
	// failure_trigger
	switch x := m.FailureTrigger.(type) {
	case *FailureCondition_WebhookUrl:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.WebhookUrl)))
		n += len(x.WebhookUrl)
	case *FailureCondition_PrometheusTrigger:
		s := proto.Size(x.PrometheusTrigger)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PrometheusTrigger struct {
	// Types that are valid to be assigned to QueryType:
	//	*PrometheusTrigger_CustomQuery
	//	*PrometheusTrigger_MeshQuery_
	QueryType isPrometheusTrigger_QueryType `protobuf_oneof:"query_type"`
	// consider the failure condition met if the metric falls below this threshold
	ThresholdValue float64 `protobuf:"fixed64,3,opt,name=threshold_value,json=thresholdValue,proto3" json:"threshold_value,omitempty"`
	// the comparison operator to use when comparing the threshold and observed metric values
	// if the comparison evaluates to true, the failure condition will be considered met
	// possible values are '==', '>', '<', '>=', and '<='
	// defaults to '<'
	ComparisonOperator   string   `protobuf:"bytes,4,opt,name=comparison_operator,json=comparisonOperator,proto3" json:"comparison_operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrometheusTrigger) Reset()         { *m = PrometheusTrigger{} }
func (m *PrometheusTrigger) String() string { return proto.CompactTextString(m) }
func (*PrometheusTrigger) ProtoMessage()    {}
func (*PrometheusTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{4}
}
func (m *PrometheusTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrometheusTrigger.Unmarshal(m, b)
}
func (m *PrometheusTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrometheusTrigger.Marshal(b, m, deterministic)
}
func (m *PrometheusTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrometheusTrigger.Merge(m, src)
}
func (m *PrometheusTrigger) XXX_Size() int {
	return xxx_messageInfo_PrometheusTrigger.Size(m)
}
func (m *PrometheusTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_PrometheusTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_PrometheusTrigger proto.InternalMessageInfo

type isPrometheusTrigger_QueryType interface {
	isPrometheusTrigger_QueryType()
	Equal(interface{}) bool
}

type PrometheusTrigger_CustomQuery struct {
	CustomQuery string `protobuf:"bytes,1,opt,name=custom_query,json=customQuery,proto3,oneof"`
}
type PrometheusTrigger_MeshQuery_ struct {
	MeshQuery *PrometheusTrigger_MeshQuery `protobuf:"bytes,2,opt,name=mesh_query,json=meshQuery,proto3,oneof"`
}

func (*PrometheusTrigger_CustomQuery) isPrometheusTrigger_QueryType() {}
func (*PrometheusTrigger_MeshQuery_) isPrometheusTrigger_QueryType()  {}

func (m *PrometheusTrigger) GetQueryType() isPrometheusTrigger_QueryType {
	if m != nil {
		return m.QueryType
	}
	return nil
}

func (m *PrometheusTrigger) GetCustomQuery() string {
	if x, ok := m.GetQueryType().(*PrometheusTrigger_CustomQuery); ok {
		return x.CustomQuery
	}
	return ""
}

func (m *PrometheusTrigger) GetMeshQuery() *PrometheusTrigger_MeshQuery {
	if x, ok := m.GetQueryType().(*PrometheusTrigger_MeshQuery_); ok {
		return x.MeshQuery
	}
	return nil
}

func (m *PrometheusTrigger) GetThresholdValue() float64 {
	if m != nil {
		return m.ThresholdValue
	}
	return 0
}

func (m *PrometheusTrigger) GetComparisonOperator() string {
	if m != nil {
		return m.ComparisonOperator
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PrometheusTrigger) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PrometheusTrigger_OneofMarshaler, _PrometheusTrigger_OneofUnmarshaler, _PrometheusTrigger_OneofSizer, []interface{}{
		(*PrometheusTrigger_CustomQuery)(nil),
		(*PrometheusTrigger_MeshQuery_)(nil),
	}
}

func _PrometheusTrigger_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PrometheusTrigger)
	// query_type
	switch x := m.QueryType.(type) {
	case *PrometheusTrigger_CustomQuery:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.CustomQuery)
	case *PrometheusTrigger_MeshQuery_:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MeshQuery); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PrometheusTrigger.QueryType has unexpected type %T", x)
	}
	return nil
}

func _PrometheusTrigger_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PrometheusTrigger)
	switch tag {
	case 1: // query_type.custom_query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.QueryType = &PrometheusTrigger_CustomQuery{x}
		return true, err
	case 2: // query_type.mesh_query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrometheusTrigger_MeshQuery)
		err := b.DecodeMessage(msg)
		m.QueryType = &PrometheusTrigger_MeshQuery_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PrometheusTrigger_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PrometheusTrigger)
	// query_type
	switch x := m.QueryType.(type) {
	case *PrometheusTrigger_CustomQuery:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.CustomQuery)))
		n += len(x.CustomQuery)
	case *PrometheusTrigger_MeshQuery_:
		s := proto.Size(x.MeshQuery)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PrometheusTrigger_MeshQuery struct {
	// the name of the metric to monitor
	// possible values:
	// envoy_cluster_upstream_rq
	// istio_requests_total
	// istio_request_duration_seconds_bucket
	Metric string `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// the service for which the query should be made
	Service              *core.ResourceRef `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PrometheusTrigger_MeshQuery) Reset()         { *m = PrometheusTrigger_MeshQuery{} }
func (m *PrometheusTrigger_MeshQuery) String() string { return proto.CompactTextString(m) }
func (*PrometheusTrigger_MeshQuery) ProtoMessage()    {}
func (*PrometheusTrigger_MeshQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{4, 0}
}
func (m *PrometheusTrigger_MeshQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrometheusTrigger_MeshQuery.Unmarshal(m, b)
}
func (m *PrometheusTrigger_MeshQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrometheusTrigger_MeshQuery.Marshal(b, m, deterministic)
}
func (m *PrometheusTrigger_MeshQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrometheusTrigger_MeshQuery.Merge(m, src)
}
func (m *PrometheusTrigger_MeshQuery) XXX_Size() int {
	return xxx_messageInfo_PrometheusTrigger_MeshQuery.Size(m)
}
func (m *PrometheusTrigger_MeshQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PrometheusTrigger_MeshQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PrometheusTrigger_MeshQuery proto.InternalMessageInfo

func (m *PrometheusTrigger_MeshQuery) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *PrometheusTrigger_MeshQuery) GetService() *core.ResourceRef {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterEnum("glooshot.solo.io.ExperimentResult_State", ExperimentResult_State_name, ExperimentResult_State_value)
	proto.RegisterType((*Experiment)(nil), "glooshot.solo.io.Experiment")
	proto.RegisterType((*ExperimentResult)(nil), "glooshot.solo.io.ExperimentResult")
	proto.RegisterMapType((map[string]string)(nil), "glooshot.solo.io.ExperimentResult.FailureReportEntry")
	proto.RegisterType((*ExperimentSpec)(nil), "glooshot.solo.io.ExperimentSpec")
	proto.RegisterType((*ExperimentSpec_InjectedFault)(nil), "glooshot.solo.io.ExperimentSpec.InjectedFault")
	proto.RegisterType((*FailureCondition)(nil), "glooshot.solo.io.FailureCondition")
	proto.RegisterType((*PrometheusTrigger)(nil), "glooshot.solo.io.PrometheusTrigger")
	proto.RegisterType((*PrometheusTrigger_MeshQuery)(nil), "glooshot.solo.io.PrometheusTrigger.MeshQuery")
}

func init() {
	proto.RegisterFile("github.com/solo-io/glooshot/api/v1/glooshot.proto", fileDescriptor_b9da8418b9c75752)
}

var fileDescriptor_b9da8418b9c75752 = []byte{
	// 927 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xc7, 0xd7, 0xfb, 0xd5, 0xec, 0xd9, 0xec, 0x66, 0x33, 0x8d, 0xaa, 0xed, 0x5e, 0x34, 0xe9,
	0x56, 0x82, 0x5c, 0xb4, 0xb6, 0x92, 0x82, 0xa8, 0x0a, 0x42, 0xd5, 0x42, 0x43, 0x90, 0x28, 0xb4,
	0xb3, 0x01, 0x21, 0x84, 0x64, 0x39, 0xf6, 0x59, 0x7b, 0x88, 0xed, 0x31, 0x33, 0xe3, 0xd0, 0xdc,
	0xf6, 0x01, 0x10, 0x97, 0x5c, 0x71, 0xcd, 0x7b, 0x20, 0x24, 0xc4, 0x23, 0x70, 0x01, 0x12, 0x6f,
	0xd0, 0x37, 0x40, 0x1e, 0x8f, 0x9d, 0x64, 0x53, 0x25, 0xe1, 0x2a, 0x3e, 0x1f, 0xbf, 0x73, 0x8e,
	0xff, 0x73, 0xc6, 0x59, 0xd8, 0x09, 0x99, 0x8a, 0xf2, 0x43, 0xdb, 0xe7, 0x89, 0x23, 0x79, 0xcc,
	0x1f, 0x30, 0xee, 0x84, 0x31, 0xe7, 0x32, 0xe2, 0xca, 0xf1, 0x32, 0xe6, 0x1c, 0xef, 0xd4, 0xb6,
	0x9d, 0x09, 0xae, 0x38, 0x19, 0xd5, 0x76, 0x01, 0xd8, 0x8c, 0x4f, 0x36, 0x42, 0x1e, 0x72, 0x1d,
	0x74, 0x8a, 0xa7, 0x32, 0x6f, 0x72, 0x27, 0xe4, 0x3c, 0x8c, 0xd1, 0xd1, 0xd6, 0x61, 0xbe, 0x70,
	0x82, 0x5c, 0x78, 0x8a, 0xf1, 0xd4, 0xc4, 0x37, 0x97, 0xe3, 0x8a, 0x25, 0x28, 0x95, 0x97, 0x64,
	0x26, 0xc1, 0x79, 0xc3, 0x6c, 0xfa, 0xef, 0x11, 0xab, 0x67, 0x93, 0xca, 0x53, 0xb9, 0x34, 0xc0,
	0xce, 0x35, 0x80, 0x04, 0x95, 0x17, 0x78, 0xca, 0x33, 0xc8, 0xfd, 0x6b, 0x20, 0x02, 0x17, 0xff,
	0xa3, 0x41, 0x65, 0x5f, 0x86, 0xe4, 0x19, 0x8a, 0x42, 0xc5, 0xba, 0x03, 0xcf, 0x15, 0x4b, 0xc3,
	0x12, 0x99, 0xfe, 0xd8, 0x04, 0x78, 0xfa, 0x32, 0x43, 0xc1, 0x12, 0x4c, 0x15, 0x79, 0x04, 0x2b,
	0xd5, 0xd0, 0x63, 0x6b, 0xcb, 0xda, 0xee, 0xef, 0xde, 0xb2, 0x7d, 0x2e, 0xb0, 0x92, 0xdf, 0x7e,
	0x66, 0xa2, 0xb3, 0xf6, 0x1f, 0x7f, 0x6f, 0x36, 0x68, 0x9d, 0x4d, 0x76, 0xa1, 0x5b, 0xea, 0x33,
	0x6e, 0x69, 0x6e, 0xe3, 0x3c, 0x37, 0xd7, 0x31, 0x43, 0x99, 0x4c, 0xf2, 0x0e, 0xb4, 0x65, 0x86,
	0xfe, 0xb8, 0xa9, 0x89, 0x2d, 0x7b, 0xf9, 0xb0, 0xed, 0xd3, 0xc9, 0xe6, 0x19, 0xfa, 0x54, 0x67,
	0x93, 0x27, 0xd0, 0x15, 0x28, 0xf3, 0x58, 0x8d, 0xdb, 0x9a, 0x9b, 0x5e, 0xc6, 0x51, 0x9d, 0x59,
	0xf5, 0x2d, 0xb9, 0xc7, 0x93, 0x57, 0xaf, 0xdb, 0x1d, 0x68, 0xe1, 0xcb, 0xec, 0xd5, 0xeb, 0xf6,
	0x80, 0xf4, 0xb1, 0x4e, 0x97, 0xd3, 0xdf, 0x5a, 0x30, 0x5a, 0xc6, 0xc9, 0x87, 0xd0, 0x29, 0x46,
	0x46, 0xad, 0xc9, 0x70, 0x77, 0xfb, 0xea, 0x8e, 0xfa, 0x85, 0x91, 0x96, 0x18, 0xf9, 0x16, 0x86,
	0x0b, 0x8f, 0xc5, 0xb9, 0x40, 0x57, 0x60, 0xc6, 0x85, 0x1a, 0x37, 0xb7, 0x5a, 0xdb, 0xfd, 0xdd,
	0x77, 0xaf, 0x51, 0x68, 0xaf, 0x04, 0xa9, 0xe6, 0x9e, 0xa6, 0x4a, 0x9c, 0xd0, 0xc1, 0xe2, 0xac,
	0x8f, 0x7c, 0x02, 0xab, 0xc5, 0x3a, 0xbb, 0x52, 0x79, 0x42, 0x61, 0x60, 0x0e, 0x60, 0x62, 0x97,
	0x3b, 0x6f, 0x57, 0x3b, 0x6f, 0x1f, 0x54, 0x3b, 0x3f, 0x5b, 0x29, 0xe4, 0xf8, 0xe9, 0x9f, 0x4d,
	0x8b, 0xf6, 0x0b, 0x72, 0x5e, 0x82, 0x64, 0x1f, 0x06, 0xba, 0xd0, 0x82, 0xa5, 0x4c, 0x46, 0x18,
	0x18, 0x81, 0x6f, 0x5f, 0xa8, 0xf4, 0xb1, 0xb9, 0x5d, 0x67, 0x0a, 0xe9, 0x11, 0xf6, 0x0c, 0x38,
	0x79, 0x02, 0xe4, 0xe2, 0xdc, 0x64, 0x04, 0xad, 0x23, 0x3c, 0xd1, 0x22, 0xf6, 0x68, 0xf1, 0x48,
	0x36, 0xa0, 0x73, 0xec, 0xc5, 0x39, 0xea, 0x15, 0xe8, 0xd1, 0xd2, 0x78, 0xdc, 0x7c, 0x64, 0x4d,
	0x3f, 0x80, 0x8e, 0x96, 0x90, 0xf4, 0xe1, 0xc6, 0x73, 0x4c, 0x03, 0x96, 0x86, 0xa3, 0x46, 0x61,
	0x98, 0x61, 0x47, 0x16, 0x01, 0xe8, 0x16, 0x4d, 0x30, 0x18, 0x35, 0xc9, 0x00, 0x7a, 0xf3, 0xdc,
	0xf7, 0x11, 0x03, 0x0c, 0x46, 0xad, 0xe9, 0xef, 0x2d, 0x18, 0x9e, 0x5f, 0x1e, 0xb2, 0x07, 0xdd,
	0x85, 0x97, 0xc7, 0x4a, 0x8e, 0xdb, 0x5a, 0x7b, 0xfb, 0xaa, 0x75, 0xb3, 0x3f, 0x4d, 0xbf, 0x43,
	0x5f, 0x61, 0xb0, 0x57, 0x60, 0xd4, 0xd0, 0xe4, 0x05, 0x90, 0xea, 0x2c, 0x7d, 0x9e, 0x06, 0xac,
	0xd0, 0x41, 0x8e, 0x3b, 0xba, 0xe6, 0x1b, 0x56, 0xd1, 0xc8, 0xf0, 0x51, 0x95, 0x4a, 0xd7, 0x17,
	0x4b, 0x1e, 0x49, 0xde, 0x87, 0x95, 0xea, 0x7b, 0x35, 0xee, 0x5e, 0x25, 0x79, 0xfb, 0xe7, 0x42,
	0xee, 0x1a, 0x98, 0xfc, 0x65, 0xc1, 0xe0, 0xdc, 0xa4, 0x64, 0x06, 0x6b, 0x5c, 0xb0, 0x90, 0xa5,
	0xae, 0x44, 0x71, 0xcc, 0x7c, 0x94, 0x63, 0x4b, 0x8f, 0x77, 0xfb, 0xfc, 0x9d, 0xa4, 0x28, 0x79,
	0x2e, 0x7c, 0xa4, 0xb8, 0xa0, 0xc3, 0x92, 0x98, 0x1b, 0x80, 0x7c, 0x06, 0x1b, 0x01, 0x4a, 0xc5,
	0x52, 0xdd, 0xe4, 0xb4, 0x50, 0xf3, 0xaa, 0x42, 0x37, 0xcf, 0x60, 0x75, 0xb5, 0xf7, 0xa0, 0xa3,
	0xd5, 0x33, 0xab, 0x79, 0xd7, 0xae, 0xbf, 0x4a, 0x67, 0x74, 0xca, 0x63, 0x55, 0xbe, 0x47, 0xa1,
	0x52, 0x99, 0x3f, 0xfd, 0xc5, 0x82, 0xd1, 0xb2, 0x82, 0xe4, 0x2e, 0xf4, 0x7f, 0xc0, 0xc3, 0x88,
	0xf3, 0x23, 0x37, 0x17, 0x71, 0xb9, 0x4e, 0xfb, 0x0d, 0x0a, 0xc6, 0xf9, 0xa5, 0x88, 0xc9, 0x01,
	0x90, 0x4c, 0xf0, 0x04, 0x55, 0x84, 0xb9, 0x74, 0x95, 0x60, 0x61, 0x88, 0xc2, 0x7c, 0x67, 0xee,
	0x5d, 0x3c, 0xa4, 0xe7, 0x75, 0xee, 0x41, 0x99, 0xba, 0xdf, 0xa0, 0xeb, 0xd9, 0xb2, 0x73, 0xb6,
	0x0e, 0x6b, 0xd5, 0xd1, 0x9b, 0x92, 0xd3, 0x3f, 0x9b, 0xb0, 0x7e, 0x81, 0x26, 0xf7, 0x60, 0xd5,
	0xcf, 0xa5, 0xe2, 0x89, 0xfb, 0x7d, 0x8e, 0xe2, 0xa4, 0x1e, 0xb1, 0x5f, 0x7a, 0x5f, 0x14, 0x4e,
	0xf2, 0x39, 0x40, 0x82, 0x32, 0x32, 0x29, 0xe5, 0x6c, 0x0f, 0xae, 0x31, 0x9b, 0xfd, 0x0c, 0x65,
	0xa4, 0x4b, 0xec, 0x37, 0x68, 0x2f, 0xa9, 0x0c, 0xf2, 0x36, 0xac, 0xa9, 0x48, 0xa0, 0x8c, 0x78,
	0x1c, 0xb8, 0xe5, 0xad, 0x2a, 0xe4, 0xb6, 0xe8, 0xb0, 0x76, 0x7f, 0x55, 0x78, 0x89, 0x03, 0x37,
	0x7d, 0x9e, 0x64, 0x9e, 0x60, 0x92, 0xa7, 0x2e, 0xcf, 0x50, 0x78, 0x8a, 0x0b, 0x7d, 0xd9, 0x7b,
	0x94, 0x9c, 0x86, 0xbe, 0x30, 0x91, 0xc9, 0xd7, 0xd0, 0xab, 0x7b, 0x92, 0x5b, 0xd0, 0x4d, 0x50,
	0x09, 0xe6, 0x9b, 0x7b, 0x6c, 0x2c, 0xf2, 0x10, 0x6e, 0x98, 0x2d, 0x31, 0xef, 0x72, 0xc9, 0x92,
	0x54, 0x99, 0xb3, 0x55, 0x00, 0xfd, 0xfa, 0xae, 0x3a, 0xc9, 0x70, 0x76, 0xff, 0xd7, 0x7f, 0xef,
	0x58, 0xdf, 0xbc, 0x75, 0xd9, 0xcf, 0x84, 0xec, 0x28, 0x34, 0xff, 0xc8, 0x0e, 0xbb, 0xfa, 0x6e,
	0x3c, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x26, 0x72, 0x0a, 0x57, 0x08, 0x00, 0x00,
}

func (this *Experiment) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Experiment)
	if !ok {
		that2, ok := that.(Experiment)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !this.Result.Equal(&that1.Result) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ExperimentResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ExperimentResult)
	if !ok {
		that2, ok := that.(ExperimentResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if len(this.FailureReport) != len(that1.FailureReport) {
		return false
	}
	for i := range this.FailureReport {
		if this.FailureReport[i] != that1.FailureReport[i] {
			return false
		}
	}
	if !this.TimeStarted.Equal(that1.TimeStarted) {
		return false
	}
	if !this.TimeFinished.Equal(that1.TimeFinished) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ExperimentSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ExperimentSpec)
	if !ok {
		that2, ok := that.(ExperimentSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Faults) != len(that1.Faults) {
		return false
	}
	for i := range this.Faults {
		if !this.Faults[i].Equal(that1.Faults[i]) {
			return false
		}
	}
	if len(this.FailureConditions) != len(that1.FailureConditions) {
		return false
	}
	for i := range this.FailureConditions {
		if !this.FailureConditions[i].Equal(that1.FailureConditions[i]) {
			return false
		}
	}
	if this.Duration != nil && that1.Duration != nil {
		if *this.Duration != *that1.Duration {
			return false
		}
	} else if this.Duration != nil {
		return false
	} else if that1.Duration != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ExperimentSpec_InjectedFault) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ExperimentSpec_InjectedFault)
	if !ok {
		that2, ok := that.(ExperimentSpec_InjectedFault)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.OriginServices) != len(that1.OriginServices) {
		return false
	}
	for i := range this.OriginServices {
		if !this.OriginServices[i].Equal(that1.OriginServices[i]) {
			return false
		}
	}
	if len(this.DestinationServices) != len(that1.DestinationServices) {
		return false
	}
	for i := range this.DestinationServices {
		if !this.DestinationServices[i].Equal(that1.DestinationServices[i]) {
			return false
		}
	}
	if !this.Fault.Equal(that1.Fault) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailureCondition) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailureCondition)
	if !ok {
		that2, ok := that.(FailureCondition)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.FailureTrigger == nil {
		if this.FailureTrigger != nil {
			return false
		}
	} else if this.FailureTrigger == nil {
		return false
	} else if !this.FailureTrigger.Equal(that1.FailureTrigger) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailureCondition_WebhookUrl) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailureCondition_WebhookUrl)
	if !ok {
		that2, ok := that.(FailureCondition_WebhookUrl)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.WebhookUrl != that1.WebhookUrl {
		return false
	}
	return true
}
func (this *FailureCondition_PrometheusTrigger) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailureCondition_PrometheusTrigger)
	if !ok {
		that2, ok := that.(FailureCondition_PrometheusTrigger)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.PrometheusTrigger.Equal(that1.PrometheusTrigger) {
		return false
	}
	return true
}
func (this *PrometheusTrigger) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PrometheusTrigger)
	if !ok {
		that2, ok := that.(PrometheusTrigger)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.QueryType == nil {
		if this.QueryType != nil {
			return false
		}
	} else if this.QueryType == nil {
		return false
	} else if !this.QueryType.Equal(that1.QueryType) {
		return false
	}
	if this.ThresholdValue != that1.ThresholdValue {
		return false
	}
	if this.ComparisonOperator != that1.ComparisonOperator {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PrometheusTrigger_CustomQuery) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PrometheusTrigger_CustomQuery)
	if !ok {
		that2, ok := that.(PrometheusTrigger_CustomQuery)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CustomQuery != that1.CustomQuery {
		return false
	}
	return true
}
func (this *PrometheusTrigger_MeshQuery_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PrometheusTrigger_MeshQuery_)
	if !ok {
		that2, ok := that.(PrometheusTrigger_MeshQuery_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MeshQuery.Equal(that1.MeshQuery) {
		return false
	}
	return true
}
func (this *PrometheusTrigger_MeshQuery) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PrometheusTrigger_MeshQuery)
	if !ok {
		that2, ok := that.(PrometheusTrigger_MeshQuery)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Metric != that1.Metric {
		return false
	}
	if !this.Service.Equal(that1.Service) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
