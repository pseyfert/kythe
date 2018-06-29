// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/pipeline.proto

package pipeline_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import schema_go_proto "kythe.io/kythe/proto/schema_go_proto"
import serving_go_proto "kythe.io/kythe/proto/serving_go_proto"
import storage_go_proto "kythe.io/kythe/proto/storage_go_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Node struct {
	Source *storage_go_proto.VName `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Fact   []*Fact                 `protobuf:"bytes,2,rep,name=fact" json:"fact,omitempty"`
	Edge   []*Edge                 `protobuf:"bytes,3,rep,name=edge" json:"edge,omitempty"`
	// Types that are valid to be assigned to Kind:
	//	*Node_KytheKind
	//	*Node_GenericKind
	Kind isNode_Kind `protobuf_oneof:"kind"`
	// Types that are valid to be assigned to Subkind:
	//	*Node_KytheSubkind
	//	*Node_GenericSubkind
	Subkind              isNode_Subkind `protobuf_oneof:"subkind"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_31b70f6420ca9d7a, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

type isNode_Kind interface {
	isNode_Kind()
}
type isNode_Subkind interface {
	isNode_Subkind()
}

type Node_KytheKind struct {
	KytheKind schema_go_proto.NodeKind `protobuf:"varint,4,opt,name=kythe_kind,json=kytheKind,enum=kythe.proto.schema.NodeKind,oneof"`
}
type Node_GenericKind struct {
	GenericKind string `protobuf:"bytes,5,opt,name=generic_kind,json=genericKind,oneof"`
}
type Node_KytheSubkind struct {
	KytheSubkind schema_go_proto.Subkind `protobuf:"varint,6,opt,name=kythe_subkind,json=kytheSubkind,enum=kythe.proto.schema.Subkind,oneof"`
}
type Node_GenericSubkind struct {
	GenericSubkind string `protobuf:"bytes,7,opt,name=generic_subkind,json=genericSubkind,oneof"`
}

func (*Node_KytheKind) isNode_Kind()         {}
func (*Node_GenericKind) isNode_Kind()       {}
func (*Node_KytheSubkind) isNode_Subkind()   {}
func (*Node_GenericSubkind) isNode_Subkind() {}

func (m *Node) GetKind() isNode_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}
func (m *Node) GetSubkind() isNode_Subkind {
	if m != nil {
		return m.Subkind
	}
	return nil
}

func (m *Node) GetSource() *storage_go_proto.VName {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Node) GetFact() []*Fact {
	if m != nil {
		return m.Fact
	}
	return nil
}

func (m *Node) GetEdge() []*Edge {
	if m != nil {
		return m.Edge
	}
	return nil
}

func (m *Node) GetKytheKind() schema_go_proto.NodeKind {
	if x, ok := m.GetKind().(*Node_KytheKind); ok {
		return x.KytheKind
	}
	return schema_go_proto.NodeKind_UNKNOWN_NODE_KIND
}

func (m *Node) GetGenericKind() string {
	if x, ok := m.GetKind().(*Node_GenericKind); ok {
		return x.GenericKind
	}
	return ""
}

func (m *Node) GetKytheSubkind() schema_go_proto.Subkind {
	if x, ok := m.GetSubkind().(*Node_KytheSubkind); ok {
		return x.KytheSubkind
	}
	return schema_go_proto.Subkind_UNKNOWN_SUBKIND
}

func (m *Node) GetGenericSubkind() string {
	if x, ok := m.GetSubkind().(*Node_GenericSubkind); ok {
		return x.GenericSubkind
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Node) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Node_OneofMarshaler, _Node_OneofUnmarshaler, _Node_OneofSizer, []interface{}{
		(*Node_KytheKind)(nil),
		(*Node_GenericKind)(nil),
		(*Node_KytheSubkind)(nil),
		(*Node_GenericSubkind)(nil),
	}
}

func _Node_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Node)
	// kind
	switch x := m.Kind.(type) {
	case *Node_KytheKind:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.KytheKind))
	case *Node_GenericKind:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GenericKind)
	case nil:
	default:
		return fmt.Errorf("Node.Kind has unexpected type %T", x)
	}
	// subkind
	switch x := m.Subkind.(type) {
	case *Node_KytheSubkind:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.KytheSubkind))
	case *Node_GenericSubkind:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GenericSubkind)
	case nil:
	default:
		return fmt.Errorf("Node.Subkind has unexpected type %T", x)
	}
	return nil
}

func _Node_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Node)
	switch tag {
	case 4: // kind.kythe_kind
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Node_KytheKind{schema_go_proto.NodeKind(x)}
		return true, err
	case 5: // kind.generic_kind
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &Node_GenericKind{x}
		return true, err
	case 6: // subkind.kythe_subkind
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Subkind = &Node_KytheSubkind{schema_go_proto.Subkind(x)}
		return true, err
	case 7: // subkind.generic_subkind
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Subkind = &Node_GenericSubkind{x}
		return true, err
	default:
		return false, nil
	}
}

func _Node_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Node)
	// kind
	switch x := m.Kind.(type) {
	case *Node_KytheKind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.KytheKind))
	case *Node_GenericKind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.GenericKind)))
		n += len(x.GenericKind)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// subkind
	switch x := m.Subkind.(type) {
	case *Node_KytheSubkind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.KytheSubkind))
	case *Node_GenericSubkind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.GenericSubkind)))
		n += len(x.GenericSubkind)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Fact struct {
	Source *storage_go_proto.VName `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	// Types that are valid to be assigned to Name:
	//	*Fact_KytheName
	//	*Fact_GenericName
	Name                 isFact_Name `protobuf_oneof:"name"`
	Value                []byte      `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Fact) Reset()         { *m = Fact{} }
func (m *Fact) String() string { return proto.CompactTextString(m) }
func (*Fact) ProtoMessage()    {}
func (*Fact) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_31b70f6420ca9d7a, []int{1}
}
func (m *Fact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fact.Unmarshal(m, b)
}
func (m *Fact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fact.Marshal(b, m, deterministic)
}
func (dst *Fact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fact.Merge(dst, src)
}
func (m *Fact) XXX_Size() int {
	return xxx_messageInfo_Fact.Size(m)
}
func (m *Fact) XXX_DiscardUnknown() {
	xxx_messageInfo_Fact.DiscardUnknown(m)
}

var xxx_messageInfo_Fact proto.InternalMessageInfo

type isFact_Name interface {
	isFact_Name()
}

type Fact_KytheName struct {
	KytheName schema_go_proto.FactName `protobuf:"varint,2,opt,name=kythe_name,json=kytheName,enum=kythe.proto.schema.FactName,oneof"`
}
type Fact_GenericName struct {
	GenericName string `protobuf:"bytes,3,opt,name=generic_name,json=genericName,oneof"`
}

func (*Fact_KytheName) isFact_Name()   {}
func (*Fact_GenericName) isFact_Name() {}

func (m *Fact) GetName() isFact_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Fact) GetSource() *storage_go_proto.VName {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Fact) GetKytheName() schema_go_proto.FactName {
	if x, ok := m.GetName().(*Fact_KytheName); ok {
		return x.KytheName
	}
	return schema_go_proto.FactName_UNKNOWN_FACT_NAME
}

func (m *Fact) GetGenericName() string {
	if x, ok := m.GetName().(*Fact_GenericName); ok {
		return x.GenericName
	}
	return ""
}

func (m *Fact) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Fact) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Fact_OneofMarshaler, _Fact_OneofUnmarshaler, _Fact_OneofSizer, []interface{}{
		(*Fact_KytheName)(nil),
		(*Fact_GenericName)(nil),
	}
}

func _Fact_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Fact)
	// name
	switch x := m.Name.(type) {
	case *Fact_KytheName:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.KytheName))
	case *Fact_GenericName:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GenericName)
	case nil:
	default:
		return fmt.Errorf("Fact.Name has unexpected type %T", x)
	}
	return nil
}

func _Fact_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Fact)
	switch tag {
	case 2: // name.kythe_name
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Name = &Fact_KytheName{schema_go_proto.FactName(x)}
		return true, err
	case 3: // name.generic_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Name = &Fact_GenericName{x}
		return true, err
	default:
		return false, nil
	}
}

func _Fact_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Fact)
	// name
	switch x := m.Name.(type) {
	case *Fact_KytheName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.KytheName))
	case *Fact_GenericName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.GenericName)))
		n += len(x.GenericName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Edge struct {
	Source *storage_go_proto.VName `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Target *storage_go_proto.VName `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	// Types that are valid to be assigned to Kind:
	//	*Edge_KytheKind
	//	*Edge_GenericKind
	Kind                 isEdge_Kind `protobuf_oneof:"kind"`
	Ordinal              int32       `protobuf:"varint,5,opt,name=ordinal" json:"ordinal,omitempty"`
	SourceNode           *Node       `protobuf:"bytes,6,opt,name=source_node,json=sourceNode" json:"source_node,omitempty"`
	TargetNode           *Node       `protobuf:"bytes,7,opt,name=target_node,json=targetNode" json:"target_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Edge) Reset()         { *m = Edge{} }
func (m *Edge) String() string { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()    {}
func (*Edge) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_31b70f6420ca9d7a, []int{2}
}
func (m *Edge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Edge.Unmarshal(m, b)
}
func (m *Edge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Edge.Marshal(b, m, deterministic)
}
func (dst *Edge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Edge.Merge(dst, src)
}
func (m *Edge) XXX_Size() int {
	return xxx_messageInfo_Edge.Size(m)
}
func (m *Edge) XXX_DiscardUnknown() {
	xxx_messageInfo_Edge.DiscardUnknown(m)
}

var xxx_messageInfo_Edge proto.InternalMessageInfo

type isEdge_Kind interface {
	isEdge_Kind()
}

type Edge_KytheKind struct {
	KytheKind schema_go_proto.EdgeKind `protobuf:"varint,3,opt,name=kythe_kind,json=kytheKind,enum=kythe.proto.schema.EdgeKind,oneof"`
}
type Edge_GenericKind struct {
	GenericKind string `protobuf:"bytes,4,opt,name=generic_kind,json=genericKind,oneof"`
}

func (*Edge_KytheKind) isEdge_Kind()   {}
func (*Edge_GenericKind) isEdge_Kind() {}

func (m *Edge) GetKind() isEdge_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Edge) GetSource() *storage_go_proto.VName {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Edge) GetTarget() *storage_go_proto.VName {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Edge) GetKytheKind() schema_go_proto.EdgeKind {
	if x, ok := m.GetKind().(*Edge_KytheKind); ok {
		return x.KytheKind
	}
	return schema_go_proto.EdgeKind_UNKNOWN_EDGE_KIND
}

func (m *Edge) GetGenericKind() string {
	if x, ok := m.GetKind().(*Edge_GenericKind); ok {
		return x.GenericKind
	}
	return ""
}

func (m *Edge) GetOrdinal() int32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *Edge) GetSourceNode() *Node {
	if m != nil {
		return m.SourceNode
	}
	return nil
}

func (m *Edge) GetTargetNode() *Node {
	if m != nil {
		return m.TargetNode
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Edge) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Edge_OneofMarshaler, _Edge_OneofUnmarshaler, _Edge_OneofSizer, []interface{}{
		(*Edge_KytheKind)(nil),
		(*Edge_GenericKind)(nil),
	}
}

func _Edge_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Edge)
	// kind
	switch x := m.Kind.(type) {
	case *Edge_KytheKind:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.KytheKind))
	case *Edge_GenericKind:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GenericKind)
	case nil:
	default:
		return fmt.Errorf("Edge.Kind has unexpected type %T", x)
	}
	return nil
}

func _Edge_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Edge)
	switch tag {
	case 3: // kind.kythe_kind
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Edge_KytheKind{schema_go_proto.EdgeKind(x)}
		return true, err
	case 4: // kind.generic_kind
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &Edge_GenericKind{x}
		return true, err
	default:
		return false, nil
	}
}

func _Edge_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Edge)
	// kind
	switch x := m.Kind.(type) {
	case *Edge_KytheKind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.KytheKind))
	case *Edge_GenericKind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.GenericKind)))
		n += len(x.GenericKind)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Reference struct {
	Source *storage_go_proto.VName `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	// Types that are valid to be assigned to Kind:
	//	*Reference_KytheKind
	//	*Reference_GenericKind
	Kind                 isReference_Kind                 `protobuf_oneof:"kind"`
	Anchor               *serving_go_proto.ExpandedAnchor `protobuf:"bytes,4,opt,name=anchor" json:"anchor,omitempty"`
	Scope                *storage_go_proto.VName          `protobuf:"bytes,5,opt,name=scope" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *Reference) Reset()         { *m = Reference{} }
func (m *Reference) String() string { return proto.CompactTextString(m) }
func (*Reference) ProtoMessage()    {}
func (*Reference) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_31b70f6420ca9d7a, []int{3}
}
func (m *Reference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reference.Unmarshal(m, b)
}
func (m *Reference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reference.Marshal(b, m, deterministic)
}
func (dst *Reference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reference.Merge(dst, src)
}
func (m *Reference) XXX_Size() int {
	return xxx_messageInfo_Reference.Size(m)
}
func (m *Reference) XXX_DiscardUnknown() {
	xxx_messageInfo_Reference.DiscardUnknown(m)
}

var xxx_messageInfo_Reference proto.InternalMessageInfo

type isReference_Kind interface {
	isReference_Kind()
}

type Reference_KytheKind struct {
	KytheKind schema_go_proto.EdgeKind `protobuf:"varint,2,opt,name=kythe_kind,json=kytheKind,enum=kythe.proto.schema.EdgeKind,oneof"`
}
type Reference_GenericKind struct {
	GenericKind string `protobuf:"bytes,3,opt,name=generic_kind,json=genericKind,oneof"`
}

func (*Reference_KytheKind) isReference_Kind()   {}
func (*Reference_GenericKind) isReference_Kind() {}

func (m *Reference) GetKind() isReference_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Reference) GetSource() *storage_go_proto.VName {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Reference) GetKytheKind() schema_go_proto.EdgeKind {
	if x, ok := m.GetKind().(*Reference_KytheKind); ok {
		return x.KytheKind
	}
	return schema_go_proto.EdgeKind_UNKNOWN_EDGE_KIND
}

func (m *Reference) GetGenericKind() string {
	if x, ok := m.GetKind().(*Reference_GenericKind); ok {
		return x.GenericKind
	}
	return ""
}

func (m *Reference) GetAnchor() *serving_go_proto.ExpandedAnchor {
	if m != nil {
		return m.Anchor
	}
	return nil
}

func (m *Reference) GetScope() *storage_go_proto.VName {
	if m != nil {
		return m.Scope
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Reference) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Reference_OneofMarshaler, _Reference_OneofUnmarshaler, _Reference_OneofSizer, []interface{}{
		(*Reference_KytheKind)(nil),
		(*Reference_GenericKind)(nil),
	}
}

func _Reference_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Reference)
	// kind
	switch x := m.Kind.(type) {
	case *Reference_KytheKind:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.KytheKind))
	case *Reference_GenericKind:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GenericKind)
	case nil:
	default:
		return fmt.Errorf("Reference.Kind has unexpected type %T", x)
	}
	return nil
}

func _Reference_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Reference)
	switch tag {
	case 2: // kind.kythe_kind
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Reference_KytheKind{schema_go_proto.EdgeKind(x)}
		return true, err
	case 3: // kind.generic_kind
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &Reference_GenericKind{x}
		return true, err
	default:
		return false, nil
	}
}

func _Reference_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Reference)
	// kind
	switch x := m.Kind.(type) {
	case *Reference_KytheKind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.KytheKind))
	case *Reference_GenericKind:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.GenericKind)))
		n += len(x.GenericKind)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DecorationPiece struct {
	FileVName *storage_go_proto.VName `protobuf:"bytes,1,opt,name=file_v_name,json=fileVName" json:"file_v_name,omitempty"`
	// Types that are valid to be assigned to Piece:
	//	*DecorationPiece_File
	//	*DecorationPiece_Reference
	//	*DecorationPiece_Node
	//	*DecorationPiece_Definition_
	Piece                isDecorationPiece_Piece `protobuf_oneof:"piece"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DecorationPiece) Reset()         { *m = DecorationPiece{} }
func (m *DecorationPiece) String() string { return proto.CompactTextString(m) }
func (*DecorationPiece) ProtoMessage()    {}
func (*DecorationPiece) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_31b70f6420ca9d7a, []int{4}
}
func (m *DecorationPiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecorationPiece.Unmarshal(m, b)
}
func (m *DecorationPiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecorationPiece.Marshal(b, m, deterministic)
}
func (dst *DecorationPiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecorationPiece.Merge(dst, src)
}
func (m *DecorationPiece) XXX_Size() int {
	return xxx_messageInfo_DecorationPiece.Size(m)
}
func (m *DecorationPiece) XXX_DiscardUnknown() {
	xxx_messageInfo_DecorationPiece.DiscardUnknown(m)
}

var xxx_messageInfo_DecorationPiece proto.InternalMessageInfo

type isDecorationPiece_Piece interface {
	isDecorationPiece_Piece()
}

type DecorationPiece_File struct {
	File *serving_go_proto.File `protobuf:"bytes,2,opt,name=file,oneof"`
}
type DecorationPiece_Reference struct {
	Reference *Reference `protobuf:"bytes,3,opt,name=reference,oneof"`
}
type DecorationPiece_Node struct {
	Node *Node `protobuf:"bytes,4,opt,name=node,oneof"`
}
type DecorationPiece_Definition_ struct {
	Definition *DecorationPiece_Definition `protobuf:"bytes,5,opt,name=definition,oneof"`
}

func (*DecorationPiece_File) isDecorationPiece_Piece()        {}
func (*DecorationPiece_Reference) isDecorationPiece_Piece()   {}
func (*DecorationPiece_Node) isDecorationPiece_Piece()        {}
func (*DecorationPiece_Definition_) isDecorationPiece_Piece() {}

func (m *DecorationPiece) GetPiece() isDecorationPiece_Piece {
	if m != nil {
		return m.Piece
	}
	return nil
}

func (m *DecorationPiece) GetFileVName() *storage_go_proto.VName {
	if m != nil {
		return m.FileVName
	}
	return nil
}

func (m *DecorationPiece) GetFile() *serving_go_proto.File {
	if x, ok := m.GetPiece().(*DecorationPiece_File); ok {
		return x.File
	}
	return nil
}

func (m *DecorationPiece) GetReference() *Reference {
	if x, ok := m.GetPiece().(*DecorationPiece_Reference); ok {
		return x.Reference
	}
	return nil
}

func (m *DecorationPiece) GetNode() *Node {
	if x, ok := m.GetPiece().(*DecorationPiece_Node); ok {
		return x.Node
	}
	return nil
}

func (m *DecorationPiece) GetDefinition() *DecorationPiece_Definition {
	if x, ok := m.GetPiece().(*DecorationPiece_Definition_); ok {
		return x.Definition
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DecorationPiece) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DecorationPiece_OneofMarshaler, _DecorationPiece_OneofUnmarshaler, _DecorationPiece_OneofSizer, []interface{}{
		(*DecorationPiece_File)(nil),
		(*DecorationPiece_Reference)(nil),
		(*DecorationPiece_Node)(nil),
		(*DecorationPiece_Definition_)(nil),
	}
}

func _DecorationPiece_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DecorationPiece)
	// piece
	switch x := m.Piece.(type) {
	case *DecorationPiece_File:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.File); err != nil {
			return err
		}
	case *DecorationPiece_Reference:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Reference); err != nil {
			return err
		}
	case *DecorationPiece_Node:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Node); err != nil {
			return err
		}
	case *DecorationPiece_Definition_:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Definition); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DecorationPiece.Piece has unexpected type %T", x)
	}
	return nil
}

func _DecorationPiece_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DecorationPiece)
	switch tag {
	case 2: // piece.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(serving_go_proto.File)
		err := b.DecodeMessage(msg)
		m.Piece = &DecorationPiece_File{msg}
		return true, err
	case 3: // piece.reference
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Reference)
		err := b.DecodeMessage(msg)
		m.Piece = &DecorationPiece_Reference{msg}
		return true, err
	case 4: // piece.node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Node)
		err := b.DecodeMessage(msg)
		m.Piece = &DecorationPiece_Node{msg}
		return true, err
	case 5: // piece.definition
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DecorationPiece_Definition)
		err := b.DecodeMessage(msg)
		m.Piece = &DecorationPiece_Definition_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DecorationPiece_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DecorationPiece)
	// piece
	switch x := m.Piece.(type) {
	case *DecorationPiece_File:
		s := proto.Size(x.File)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DecorationPiece_Reference:
		s := proto.Size(x.Reference)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DecorationPiece_Node:
		s := proto.Size(x.Node)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DecorationPiece_Definition_:
		s := proto.Size(x.Definition)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DecorationPiece_Definition struct {
	Node                 *storage_go_proto.VName          `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Definition           *serving_go_proto.ExpandedAnchor `protobuf:"bytes,2,opt,name=definition" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *DecorationPiece_Definition) Reset()         { *m = DecorationPiece_Definition{} }
func (m *DecorationPiece_Definition) String() string { return proto.CompactTextString(m) }
func (*DecorationPiece_Definition) ProtoMessage()    {}
func (*DecorationPiece_Definition) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_31b70f6420ca9d7a, []int{4, 0}
}
func (m *DecorationPiece_Definition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecorationPiece_Definition.Unmarshal(m, b)
}
func (m *DecorationPiece_Definition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecorationPiece_Definition.Marshal(b, m, deterministic)
}
func (dst *DecorationPiece_Definition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecorationPiece_Definition.Merge(dst, src)
}
func (m *DecorationPiece_Definition) XXX_Size() int {
	return xxx_messageInfo_DecorationPiece_Definition.Size(m)
}
func (m *DecorationPiece_Definition) XXX_DiscardUnknown() {
	xxx_messageInfo_DecorationPiece_Definition.DiscardUnknown(m)
}

var xxx_messageInfo_DecorationPiece_Definition proto.InternalMessageInfo

func (m *DecorationPiece_Definition) GetNode() *storage_go_proto.VName {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DecorationPiece_Definition) GetDefinition() *serving_go_proto.ExpandedAnchor {
	if m != nil {
		return m.Definition
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "kythe.proto.pipeline.Node")
	proto.RegisterType((*Fact)(nil), "kythe.proto.pipeline.Fact")
	proto.RegisterType((*Edge)(nil), "kythe.proto.pipeline.Edge")
	proto.RegisterType((*Reference)(nil), "kythe.proto.pipeline.Reference")
	proto.RegisterType((*DecorationPiece)(nil), "kythe.proto.pipeline.DecorationPiece")
	proto.RegisterType((*DecorationPiece_Definition)(nil), "kythe.proto.pipeline.DecorationPiece.Definition")
}

func init() {
	proto.RegisterFile("kythe/proto/pipeline.proto", fileDescriptor_pipeline_31b70f6420ca9d7a)
}

var fileDescriptor_pipeline_31b70f6420ca9d7a = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6a, 0xdb, 0x4e,
	0x10, 0xc7, 0x23, 0x59, 0xb6, 0xf1, 0x28, 0xbf, 0x04, 0x96, 0x1c, 0x14, 0xff, 0x0a, 0x71, 0x1d,
	0x28, 0x6e, 0x0e, 0x4a, 0x70, 0x8f, 0xa1, 0x94, 0xba, 0x49, 0x30, 0x14, 0x42, 0xd9, 0x42, 0xaf,
	0x61, 0x23, 0x8d, 0x95, 0x25, 0xf2, 0xae, 0x91, 0x15, 0xd3, 0xbc, 0x42, 0x9f, 0xa7, 0x2f, 0xd0,
	0x17, 0xe9, 0x73, 0xf4, 0x58, 0x76, 0x76, 0xed, 0xda, 0x41, 0x6a, 0x9a, 0xf6, 0x24, 0x8d, 0xe7,
	0x33, 0xff, 0xbe, 0x33, 0x32, 0x74, 0x6f, 0xef, 0xcb, 0x1b, 0x3c, 0x9e, 0x15, 0xba, 0xd4, 0xc7,
	0x33, 0x39, 0xc3, 0x5c, 0x2a, 0x8c, 0xc9, 0x64, 0x7b, 0xe4, 0xb3, 0x46, 0xbc, 0xf4, 0x75, 0xa3,
	0xf5, 0x88, 0x79, 0x72, 0x83, 0x53, 0x61, 0x91, 0xee, 0xfe, 0x86, 0x07, 0x8b, 0x85, 0x54, 0x59,
	0xa5, 0xab, 0xd4, 0x85, 0xc8, 0x5c, 0xe2, 0xfe, 0x0f, 0x1f, 0x82, 0x4b, 0x9d, 0x22, 0x3b, 0x82,
	0xd6, 0x5c, 0xdf, 0x15, 0x09, 0x46, 0x5e, 0xcf, 0x1b, 0x84, 0x43, 0x16, 0xaf, 0xd7, 0xff, 0x74,
	0x29, 0xa6, 0xc8, 0x1d, 0xc1, 0x62, 0x08, 0x26, 0x22, 0x29, 0x23, 0xbf, 0xd7, 0x18, 0x84, 0xc3,
	0x6e, 0x5c, 0xd5, 0x69, 0x7c, 0x21, 0x92, 0x92, 0x13, 0x67, 0x78, 0x4c, 0x33, 0x8c, 0x1a, 0xbf,
	0xe3, 0xcf, 0xd3, 0x0c, 0x39, 0x71, 0xec, 0x35, 0x00, 0x21, 0x57, 0xb7, 0x52, 0xa5, 0x51, 0xd0,
	0xf3, 0x06, 0x3b, 0xc3, 0x67, 0x1b, 0x51, 0x6e, 0x72, 0xd3, 0xf9, 0x7b, 0xa9, 0xd2, 0xf1, 0x16,
	0xef, 0x90, 0xdb, 0x18, 0xec, 0x10, 0xb6, 0x33, 0x54, 0x58, 0xc8, 0xc4, 0x26, 0x68, 0xf6, 0xbc,
	0x41, 0x67, 0xbc, 0xc5, 0x43, 0xf7, 0x2b, 0x41, 0x23, 0xf8, 0xcf, 0xd6, 0x98, 0xdf, 0x5d, 0x13,
	0xd5, 0xa2, 0x32, 0xff, 0x57, 0x95, 0xf9, 0x68, 0x91, 0xb1, 0xc7, 0xb7, 0xc9, 0xeb, 0x6c, 0xf6,
	0x12, 0x76, 0x97, 0x85, 0x96, 0x59, 0xda, 0x54, 0xcb, 0xe3, 0x3b, 0xce, 0xe1, 0xd0, 0x51, 0x0b,
	0x02, 0x7a, 0x76, 0xa0, 0xed, 0xd0, 0xfe, 0x57, 0x0f, 0x02, 0x23, 0xd2, 0x93, 0xa4, 0x5f, 0x49,
	0xa3, 0xc4, 0x14, 0x23, 0xbf, 0x5e, 0x1a, 0x93, 0xd9, 0x44, 0xae, 0xa4, 0x31, 0xc6, 0xba, 0x34,
	0x94, 0xa0, 0xf1, 0x40, 0x1a, 0x82, 0xf6, 0xa0, 0xb9, 0x10, 0xf9, 0x1d, 0x92, 0xf2, 0xdb, 0xdc,
	0x1a, 0x66, 0x02, 0x13, 0xd2, 0xff, 0xee, 0x43, 0x60, 0x76, 0xf5, 0xa4, 0xb6, 0x8f, 0xa0, 0x55,
	0x8a, 0x22, 0xc3, 0x92, 0x5a, 0xae, 0x61, 0x2d, 0xf1, 0x60, 0xfb, 0x8d, 0xfa, 0x11, 0x4d, 0x17,
	0x8f, 0x6f, 0x3f, 0xa8, 0xda, 0x7e, 0x04, 0x6d, 0x5d, 0xa4, 0x52, 0x89, 0x9c, 0xae, 0xa3, 0xc9,
	0x97, 0x26, 0x3b, 0x85, 0xd0, 0xf6, 0x7c, 0xa5, 0x74, 0x8a, 0x74, 0x15, 0xb5, 0x27, 0x6b, 0xce,
	0x8f, 0x83, 0xc5, 0xe9, 0x23, 0x3a, 0x85, 0xd0, 0x0e, 0x61, 0x83, 0xdb, 0x8f, 0x07, 0x5b, 0xdc,
	0xbc, 0x2f, 0x4f, 0xa4, 0xff, 0xc5, 0x87, 0x0e, 0xc7, 0x09, 0x16, 0xa8, 0x12, 0xfc, 0xbb, 0xe3,
	0xa0, 0xc1, 0xfd, 0x7f, 0x55, 0xae, 0x51, 0xa5, 0xdc, 0x29, 0xb4, 0x84, 0x4a, 0x6e, 0x74, 0x41,
	0xc2, 0x86, 0xc3, 0xc3, 0xcd, 0xfc, 0xee, 0x7f, 0xe7, 0xfc, 0xf3, 0x4c, 0xa8, 0x14, 0xd3, 0xb7,
	0x84, 0x72, 0x17, 0xc2, 0x06, 0xd0, 0x9c, 0x27, 0x7a, 0x86, 0x24, 0x7a, 0xf5, 0x2c, 0x16, 0x58,
	0x89, 0xf1, 0xad, 0x01, 0xbb, 0x67, 0x98, 0xe8, 0x42, 0x94, 0x52, 0xab, 0x0f, 0x12, 0x13, 0x64,
	0x43, 0x08, 0x27, 0x32, 0xc7, 0xab, 0x85, 0xbd, 0xe1, 0x7a, 0x5d, 0x3a, 0x06, 0xa3, 0x57, 0x76,
	0x0c, 0x81, 0x31, 0xdc, 0xf9, 0xed, 0x57, 0x36, 0x7d, 0x21, 0x73, 0xf3, 0xb9, 0x10, 0xc8, 0xde,
	0x40, 0xa7, 0x58, 0x2e, 0x81, 0x94, 0x08, 0x87, 0x07, 0xd5, 0x8b, 0x5c, 0xed, 0xca, 0xa8, 0xb9,
	0x8a, 0x61, 0x27, 0x10, 0xd0, 0x11, 0x04, 0x8f, 0x1d, 0x81, 0x29, 0x69, 0x48, 0xc6, 0x01, 0x52,
	0x9c, 0x48, 0x25, 0xcd, 0xa8, 0x4e, 0xa2, 0x93, 0xea, 0xb8, 0x07, 0x92, 0xc4, 0x67, 0xab, 0xb8,
	0xf1, 0x16, 0x5f, 0xcb, 0xd2, 0xbd, 0x07, 0xf8, 0xe5, 0x63, 0x2f, 0x5c, 0x4f, 0xf5, 0x92, 0xd9,
	0x4e, 0xde, 0x6d, 0x74, 0xe2, 0xff, 0xf9, 0xa2, 0xd7, 0xc2, 0x46, 0x6d, 0x68, 0xce, 0x4c, 0x73,
	0xa3, 0xe7, 0x70, 0x90, 0xe8, 0x69, 0x9c, 0x69, 0x9d, 0xe5, 0x18, 0xa7, 0xb8, 0x28, 0xb5, 0xce,
	0xe7, 0xeb, 0xe9, 0xae, 0x5b, 0xf4, 0x78, 0xf5, 0x33, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xf7, 0xc4,
	0x76, 0x11, 0x07, 0x00, 0x00,
}