// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package MyGame

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type InParentNamespaceT struct {
}

// InParentNamespaceT object pack function
func (t *InParentNamespaceT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}

	// pack process all field

	InParentNamespaceStart(builder)
	return InParentNamespaceEnd(builder)
}

// InParentNamespaceT object unpack function
func (rcv *InParentNamespace) UnPackTo(t *InParentNamespaceT) {
}

func (rcv *InParentNamespace) UnPack() *InParentNamespaceT {
	if rcv == nil {
		return nil
	}
	t := &InParentNamespaceT{}
	rcv.UnPackTo(t)
	return t
}

type InParentNamespace struct {
	_tab flatbuffers.Table
}

// GetRootAsInParentNamespace shortcut to access root table
func GetRootAsInParentNamespace(buf []byte, offset flatbuffers.UOffsetT) *InParentNamespace {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &InParentNamespace{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInParentNamespace(buf []byte, offset flatbuffers.UOffsetT) *InParentNamespace {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &InParentNamespace{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *InParentNamespace) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *InParentNamespace) Table() flatbuffers.Table {
	return rcv._tab
}

func InParentNamespaceStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}

func InParentNamespaceEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
