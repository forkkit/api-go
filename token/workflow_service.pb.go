// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/workflow_service.proto

package token

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	common "go.temporal.io/temporal-proto/common"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HistoryContinuationToken struct {
	RunId             string                        `protobuf:"bytes,1,opt,name=RunId,proto3" json:"RunId,omitempty"`
	FirstEventId      int64                         `protobuf:"varint,2,opt,name=FirstEventId,proto3" json:"FirstEventId,omitempty"`
	NextEventId       int64                         `protobuf:"varint,3,opt,name=NextEventId,proto3" json:"NextEventId,omitempty"`
	IsWorkflowRunning bool                          `protobuf:"varint,5,opt,name=IsWorkflowRunning,proto3" json:"IsWorkflowRunning,omitempty"`
	PersistenceToken  []byte                        `protobuf:"bytes,6,opt,name=PersistenceToken,proto3" json:"PersistenceToken,omitempty"`
	TransientDecision *common.TransientDecisionInfo `protobuf:"bytes,7,opt,name=TransientDecision,proto3" json:"TransientDecision,omitempty"`
	BranchToken       []byte                        `protobuf:"bytes,8,opt,name=BranchToken,proto3" json:"BranchToken,omitempty"`
}

func (m *HistoryContinuationToken) Reset()      { *m = HistoryContinuationToken{} }
func (*HistoryContinuationToken) ProtoMessage() {}
func (*HistoryContinuationToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_b44d6ad0aa805f5b, []int{0}
}
func (m *HistoryContinuationToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistoryContinuationToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HistoryContinuationToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HistoryContinuationToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryContinuationToken.Merge(m, src)
}
func (m *HistoryContinuationToken) XXX_Size() int {
	return m.Size()
}
func (m *HistoryContinuationToken) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryContinuationToken.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryContinuationToken proto.InternalMessageInfo

func (m *HistoryContinuationToken) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *HistoryContinuationToken) GetFirstEventId() int64 {
	if m != nil {
		return m.FirstEventId
	}
	return 0
}

func (m *HistoryContinuationToken) GetNextEventId() int64 {
	if m != nil {
		return m.NextEventId
	}
	return 0
}

func (m *HistoryContinuationToken) GetIsWorkflowRunning() bool {
	if m != nil {
		return m.IsWorkflowRunning
	}
	return false
}

func (m *HistoryContinuationToken) GetPersistenceToken() []byte {
	if m != nil {
		return m.PersistenceToken
	}
	return nil
}

func (m *HistoryContinuationToken) GetTransientDecision() *common.TransientDecisionInfo {
	if m != nil {
		return m.TransientDecision
	}
	return nil
}

func (m *HistoryContinuationToken) GetBranchToken() []byte {
	if m != nil {
		return m.BranchToken
	}
	return nil
}

type TaskToken struct {
	DomainId        []byte `protobuf:"bytes,1,opt,name=DomainId,proto3" json:"DomainId,omitempty"`
	WorkflowId      string `protobuf:"bytes,2,opt,name=WorkflowId,proto3" json:"WorkflowId,omitempty"`
	RunId           []byte `protobuf:"bytes,3,opt,name=RunId,proto3" json:"RunId,omitempty"`
	ScheduleId      int64  `protobuf:"varint,4,opt,name=ScheduleId,proto3" json:"ScheduleId,omitempty"`
	ScheduleAttempt int64  `protobuf:"varint,5,opt,name=ScheduleAttempt,proto3" json:"ScheduleAttempt,omitempty"`
	ActivityId      string `protobuf:"bytes,6,opt,name=ActivityId,proto3" json:"ActivityId,omitempty"`
}

func (m *TaskToken) Reset()      { *m = TaskToken{} }
func (*TaskToken) ProtoMessage() {}
func (*TaskToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_b44d6ad0aa805f5b, []int{1}
}
func (m *TaskToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskToken.Merge(m, src)
}
func (m *TaskToken) XXX_Size() int {
	return m.Size()
}
func (m *TaskToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskToken.DiscardUnknown(m)
}

var xxx_messageInfo_TaskToken proto.InternalMessageInfo

func (m *TaskToken) GetDomainId() []byte {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *TaskToken) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *TaskToken) GetRunId() []byte {
	if m != nil {
		return m.RunId
	}
	return nil
}

func (m *TaskToken) GetScheduleId() int64 {
	if m != nil {
		return m.ScheduleId
	}
	return 0
}

func (m *TaskToken) GetScheduleAttempt() int64 {
	if m != nil {
		return m.ScheduleAttempt
	}
	return 0
}

func (m *TaskToken) GetActivityId() string {
	if m != nil {
		return m.ActivityId
	}
	return ""
}

type QueryTaskToken struct {
	DomainID string `protobuf:"bytes,1,opt,name=DomainID,proto3" json:"DomainID,omitempty"`
	TaskList string `protobuf:"bytes,2,opt,name=TaskList,proto3" json:"TaskList,omitempty"`
	TaskID   string `protobuf:"bytes,3,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
}

func (m *QueryTaskToken) Reset()      { *m = QueryTaskToken{} }
func (*QueryTaskToken) ProtoMessage() {}
func (*QueryTaskToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_b44d6ad0aa805f5b, []int{2}
}
func (m *QueryTaskToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTaskToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTaskToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTaskToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTaskToken.Merge(m, src)
}
func (m *QueryTaskToken) XXX_Size() int {
	return m.Size()
}
func (m *QueryTaskToken) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTaskToken.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTaskToken proto.InternalMessageInfo

func (m *QueryTaskToken) GetDomainID() string {
	if m != nil {
		return m.DomainID
	}
	return ""
}

func (m *QueryTaskToken) GetTaskList() string {
	if m != nil {
		return m.TaskList
	}
	return ""
}

func (m *QueryTaskToken) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func init() {
	proto.RegisterType((*HistoryContinuationToken)(nil), "token.HistoryContinuationToken")
	proto.RegisterType((*TaskToken)(nil), "token.TaskToken")
	proto.RegisterType((*QueryTaskToken)(nil), "token.QueryTaskToken")
}

func init() { proto.RegisterFile("token/workflow_service.proto", fileDescriptor_b44d6ad0aa805f5b) }

var fileDescriptor_b44d6ad0aa805f5b = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xf6, 0x34, 0x24, 0xc4, 0xaf, 0x11, 0xd0, 0x11, 0x3f, 0x56, 0x05, 0x23, 0x2b, 0x6c, 0x2c,
	0x04, 0x8e, 0x04, 0x27, 0x68, 0x09, 0x08, 0x0b, 0x84, 0xca, 0x10, 0x09, 0x89, 0x0d, 0x18, 0x7b,
	0xda, 0x8e, 0x92, 0xcc, 0x54, 0x33, 0xe3, 0x94, 0xec, 0x38, 0x02, 0xc7, 0xe0, 0x16, 0x2c, 0x61,
	0x99, 0x65, 0x97, 0xc4, 0xd9, 0xb0, 0xcc, 0x11, 0x90, 0xc7, 0x4e, 0x30, 0xb8, 0x3b, 0x7f, 0x3f,
	0xcf, 0x6f, 0xbe, 0x4f, 0x0f, 0xee, 0x1a, 0x39, 0x66, 0x62, 0x70, 0x2e, 0xd5, 0xf8, 0x78, 0x22,
	0xcf, 0x3f, 0x68, 0xa6, 0x66, 0x3c, 0x61, 0xe1, 0x99, 0x92, 0x46, 0xe2, 0xb6, 0x55, 0xf7, 0x6f,
	0x25, 0x72, 0x3a, 0x95, 0x62, 0x90, 0xb2, 0x84, 0x6b, 0x2e, 0x45, 0xa9, 0xf6, 0xbf, 0xef, 0x80,
	0xf7, 0x82, 0x6b, 0x23, 0xd5, 0xfc, 0xa9, 0x14, 0x86, 0x8b, 0x2c, 0x36, 0x5c, 0x8a, 0x51, 0x31,
	0x83, 0x6f, 0x42, 0x9b, 0x66, 0x22, 0x4a, 0x3d, 0xe4, 0xa3, 0xc0, 0xa5, 0x25, 0xc0, 0x7d, 0xe8,
	0x3d, 0xe7, 0x4a, 0x9b, 0x67, 0x33, 0x26, 0x4c, 0x94, 0x7a, 0x3b, 0x3e, 0x0a, 0x5a, 0xf4, 0x1f,
	0x0e, 0xfb, 0xb0, 0xfb, 0x9a, 0x7d, 0xde, 0x5a, 0x5a, 0xd6, 0x52, 0xa7, 0xf0, 0x43, 0xd8, 0x8b,
	0xf4, 0xbb, 0xea, 0xc9, 0x34, 0x13, 0x82, 0x8b, 0x13, 0xaf, 0xed, 0xa3, 0xa0, 0x4b, 0x9b, 0x02,
	0x7e, 0x00, 0x37, 0x8e, 0x98, 0xd2, 0x5c, 0x1b, 0x26, 0x12, 0x66, 0x5f, 0xe7, 0x75, 0x7c, 0x14,
	0xf4, 0x68, 0x83, 0xc7, 0x2f, 0x61, 0x6f, 0xa4, 0x62, 0xa1, 0x39, 0x13, 0x66, 0x58, 0xa5, 0xf5,
	0xae, 0xfa, 0x28, 0xd8, 0x7d, 0x7c, 0x2f, 0x2c, 0x5b, 0x08, 0x1b, 0x86, 0x48, 0x1c, 0x4b, 0xda,
	0x9c, 0x2b, 0x82, 0x1c, 0xaa, 0x58, 0x24, 0xa7, 0xe5, 0xce, 0xae, 0xdd, 0x59, 0xa7, 0xfa, 0x3f,
	0x10, 0xb8, 0xa3, 0x58, 0x8f, 0xcb, 0xe5, 0xfb, 0xd0, 0x1d, 0xca, 0x69, 0xcc, 0x37, 0xad, 0xf5,
	0xe8, 0x16, 0x63, 0x02, 0xb0, 0xc9, 0x55, 0xd5, 0xe6, 0xd2, 0x1a, 0xf3, 0xb7, 0xee, 0x96, 0x1d,
	0xac, 0xea, 0x26, 0x00, 0x6f, 0x93, 0x53, 0x96, 0x66, 0x13, 0x16, 0xa5, 0xde, 0x15, 0xdb, 0x64,
	0x8d, 0xc1, 0x01, 0x5c, 0xdf, 0xa0, 0x03, 0x63, 0xd8, 0xf4, 0xcc, 0xd8, 0x1a, 0x5b, 0xf4, 0x7f,
	0xba, 0xf8, 0xd3, 0x41, 0x62, 0xf8, 0x8c, 0x9b, 0x79, 0x94, 0xda, 0xfa, 0x5c, 0x5a, 0x63, 0xfa,
	0x1f, 0xe1, 0xda, 0x9b, 0x8c, 0xa9, 0xf9, 0x65, 0x69, 0x86, 0xd5, 0x0d, 0x6c, 0x71, 0xa1, 0x15,
	0xc6, 0x57, 0x5c, 0x9b, 0x2a, 0xcb, 0x16, 0xe3, 0xdb, 0xd0, 0x29, 0xbe, 0xa3, 0xa1, 0x8d, 0xe2,
	0xd2, 0x0a, 0x1d, 0x9a, 0xc5, 0x92, 0x38, 0x17, 0x4b, 0xe2, 0xac, 0x97, 0x04, 0x7d, 0xc9, 0x09,
	0xfa, 0x96, 0x13, 0xf4, 0x33, 0x27, 0x68, 0x91, 0x13, 0xf4, 0x2b, 0x27, 0xe8, 0x77, 0x4e, 0x9c,
	0x75, 0x4e, 0xd0, 0xd7, 0x15, 0x71, 0x16, 0x2b, 0xe2, 0x5c, 0xac, 0x88, 0x03, 0x77, 0xb8, 0x0c,
	0x8b, 0x04, 0x52, 0xc5, 0x93, 0xf2, 0x6c, 0x43, 0x7b, 0xd3, 0x47, 0xe8, 0xfd, 0xfd, 0x93, 0x9a,
	0xc4, 0xe5, 0x60, 0xf3, 0xfd, 0xc8, 0xda, 0x06, 0xd6, 0xf6, 0xa9, 0x63, 0xc1, 0x93, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xae, 0xe0, 0x43, 0x28, 0x03, 0x00, 0x00,
}

func (this *HistoryContinuationToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HistoryContinuationToken)
	if !ok {
		that2, ok := that.(HistoryContinuationToken)
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
	if this.RunId != that1.RunId {
		return false
	}
	if this.FirstEventId != that1.FirstEventId {
		return false
	}
	if this.NextEventId != that1.NextEventId {
		return false
	}
	if this.IsWorkflowRunning != that1.IsWorkflowRunning {
		return false
	}
	if !bytes.Equal(this.PersistenceToken, that1.PersistenceToken) {
		return false
	}
	if !this.TransientDecision.Equal(that1.TransientDecision) {
		return false
	}
	if !bytes.Equal(this.BranchToken, that1.BranchToken) {
		return false
	}
	return true
}
func (this *TaskToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TaskToken)
	if !ok {
		that2, ok := that.(TaskToken)
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
	if !bytes.Equal(this.DomainId, that1.DomainId) {
		return false
	}
	if this.WorkflowId != that1.WorkflowId {
		return false
	}
	if !bytes.Equal(this.RunId, that1.RunId) {
		return false
	}
	if this.ScheduleId != that1.ScheduleId {
		return false
	}
	if this.ScheduleAttempt != that1.ScheduleAttempt {
		return false
	}
	if this.ActivityId != that1.ActivityId {
		return false
	}
	return true
}
func (this *QueryTaskToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryTaskToken)
	if !ok {
		that2, ok := that.(QueryTaskToken)
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
	if this.DomainID != that1.DomainID {
		return false
	}
	if this.TaskList != that1.TaskList {
		return false
	}
	if this.TaskID != that1.TaskID {
		return false
	}
	return true
}
func (this *HistoryContinuationToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&token.HistoryContinuationToken{")
	s = append(s, "RunId: "+fmt.Sprintf("%#v", this.RunId)+",\n")
	s = append(s, "FirstEventId: "+fmt.Sprintf("%#v", this.FirstEventId)+",\n")
	s = append(s, "NextEventId: "+fmt.Sprintf("%#v", this.NextEventId)+",\n")
	s = append(s, "IsWorkflowRunning: "+fmt.Sprintf("%#v", this.IsWorkflowRunning)+",\n")
	s = append(s, "PersistenceToken: "+fmt.Sprintf("%#v", this.PersistenceToken)+",\n")
	if this.TransientDecision != nil {
		s = append(s, "TransientDecision: "+fmt.Sprintf("%#v", this.TransientDecision)+",\n")
	}
	s = append(s, "BranchToken: "+fmt.Sprintf("%#v", this.BranchToken)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TaskToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&token.TaskToken{")
	s = append(s, "DomainId: "+fmt.Sprintf("%#v", this.DomainId)+",\n")
	s = append(s, "WorkflowId: "+fmt.Sprintf("%#v", this.WorkflowId)+",\n")
	s = append(s, "RunId: "+fmt.Sprintf("%#v", this.RunId)+",\n")
	s = append(s, "ScheduleId: "+fmt.Sprintf("%#v", this.ScheduleId)+",\n")
	s = append(s, "ScheduleAttempt: "+fmt.Sprintf("%#v", this.ScheduleAttempt)+",\n")
	s = append(s, "ActivityId: "+fmt.Sprintf("%#v", this.ActivityId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryTaskToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&token.QueryTaskToken{")
	s = append(s, "DomainID: "+fmt.Sprintf("%#v", this.DomainID)+",\n")
	s = append(s, "TaskList: "+fmt.Sprintf("%#v", this.TaskList)+",\n")
	s = append(s, "TaskID: "+fmt.Sprintf("%#v", this.TaskID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringWorkflowService(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HistoryContinuationToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistoryContinuationToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HistoryContinuationToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BranchToken) > 0 {
		i -= len(m.BranchToken)
		copy(dAtA[i:], m.BranchToken)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.BranchToken)))
		i--
		dAtA[i] = 0x42
	}
	if m.TransientDecision != nil {
		{
			size, err := m.TransientDecision.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkflowService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PersistenceToken) > 0 {
		i -= len(m.PersistenceToken)
		copy(dAtA[i:], m.PersistenceToken)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.PersistenceToken)))
		i--
		dAtA[i] = 0x32
	}
	if m.IsWorkflowRunning {
		i--
		if m.IsWorkflowRunning {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.NextEventId != 0 {
		i = encodeVarintWorkflowService(dAtA, i, uint64(m.NextEventId))
		i--
		dAtA[i] = 0x18
	}
	if m.FirstEventId != 0 {
		i = encodeVarintWorkflowService(dAtA, i, uint64(m.FirstEventId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RunId) > 0 {
		i -= len(m.RunId)
		copy(dAtA[i:], m.RunId)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.RunId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActivityId) > 0 {
		i -= len(m.ActivityId)
		copy(dAtA[i:], m.ActivityId)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.ActivityId)))
		i--
		dAtA[i] = 0x32
	}
	if m.ScheduleAttempt != 0 {
		i = encodeVarintWorkflowService(dAtA, i, uint64(m.ScheduleAttempt))
		i--
		dAtA[i] = 0x28
	}
	if m.ScheduleId != 0 {
		i = encodeVarintWorkflowService(dAtA, i, uint64(m.ScheduleId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RunId) > 0 {
		i -= len(m.RunId)
		copy(dAtA[i:], m.RunId)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.RunId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.WorkflowId) > 0 {
		i -= len(m.WorkflowId)
		copy(dAtA[i:], m.WorkflowId)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.WorkflowId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DomainId) > 0 {
		i -= len(m.DomainId)
		copy(dAtA[i:], m.DomainId)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.DomainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTaskToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTaskToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTaskToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TaskID) > 0 {
		i -= len(m.TaskID)
		copy(dAtA[i:], m.TaskID)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.TaskID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TaskList) > 0 {
		i -= len(m.TaskList)
		copy(dAtA[i:], m.TaskList)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.TaskList)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DomainID) > 0 {
		i -= len(m.DomainID)
		copy(dAtA[i:], m.DomainID)
		i = encodeVarintWorkflowService(dAtA, i, uint64(len(m.DomainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkflowService(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkflowService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HistoryContinuationToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RunId)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	if m.FirstEventId != 0 {
		n += 1 + sovWorkflowService(uint64(m.FirstEventId))
	}
	if m.NextEventId != 0 {
		n += 1 + sovWorkflowService(uint64(m.NextEventId))
	}
	if m.IsWorkflowRunning {
		n += 2
	}
	l = len(m.PersistenceToken)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	if m.TransientDecision != nil {
		l = m.TransientDecision.Size()
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	l = len(m.BranchToken)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	return n
}

func (m *TaskToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DomainId)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	l = len(m.WorkflowId)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	l = len(m.RunId)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	if m.ScheduleId != 0 {
		n += 1 + sovWorkflowService(uint64(m.ScheduleId))
	}
	if m.ScheduleAttempt != 0 {
		n += 1 + sovWorkflowService(uint64(m.ScheduleAttempt))
	}
	l = len(m.ActivityId)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	return n
}

func (m *QueryTaskToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DomainID)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	l = len(m.TaskList)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	l = len(m.TaskID)
	if l > 0 {
		n += 1 + l + sovWorkflowService(uint64(l))
	}
	return n
}

func sovWorkflowService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkflowService(x uint64) (n int) {
	return sovWorkflowService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HistoryContinuationToken) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HistoryContinuationToken{`,
		`RunId:` + fmt.Sprintf("%v", this.RunId) + `,`,
		`FirstEventId:` + fmt.Sprintf("%v", this.FirstEventId) + `,`,
		`NextEventId:` + fmt.Sprintf("%v", this.NextEventId) + `,`,
		`IsWorkflowRunning:` + fmt.Sprintf("%v", this.IsWorkflowRunning) + `,`,
		`PersistenceToken:` + fmt.Sprintf("%v", this.PersistenceToken) + `,`,
		`TransientDecision:` + strings.Replace(fmt.Sprintf("%v", this.TransientDecision), "TransientDecisionInfo", "common.TransientDecisionInfo", 1) + `,`,
		`BranchToken:` + fmt.Sprintf("%v", this.BranchToken) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TaskToken) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TaskToken{`,
		`DomainId:` + fmt.Sprintf("%v", this.DomainId) + `,`,
		`WorkflowId:` + fmt.Sprintf("%v", this.WorkflowId) + `,`,
		`RunId:` + fmt.Sprintf("%v", this.RunId) + `,`,
		`ScheduleId:` + fmt.Sprintf("%v", this.ScheduleId) + `,`,
		`ScheduleAttempt:` + fmt.Sprintf("%v", this.ScheduleAttempt) + `,`,
		`ActivityId:` + fmt.Sprintf("%v", this.ActivityId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryTaskToken) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryTaskToken{`,
		`DomainID:` + fmt.Sprintf("%v", this.DomainID) + `,`,
		`TaskList:` + fmt.Sprintf("%v", this.TaskList) + `,`,
		`TaskID:` + fmt.Sprintf("%v", this.TaskID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringWorkflowService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HistoryContinuationToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HistoryContinuationToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistoryContinuationToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstEventId", wireType)
			}
			m.FirstEventId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstEventId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextEventId", wireType)
			}
			m.NextEventId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextEventId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsWorkflowRunning", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsWorkflowRunning = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PersistenceToken", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PersistenceToken = append(m.PersistenceToken[:0], dAtA[iNdEx:postIndex]...)
			if m.PersistenceToken == nil {
				m.PersistenceToken = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransientDecision", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TransientDecision == nil {
				m.TransientDecision = &common.TransientDecisionInfo{}
			}
			if err := m.TransientDecision.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BranchToken", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BranchToken = append(m.BranchToken[:0], dAtA[iNdEx:postIndex]...)
			if m.BranchToken == nil {
				m.BranchToken = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaskToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DomainId = append(m.DomainId[:0], dAtA[iNdEx:postIndex]...)
			if m.DomainId == nil {
				m.DomainId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkflowId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkflowId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunId = append(m.RunId[:0], dAtA[iNdEx:postIndex]...)
			if m.RunId == nil {
				m.RunId = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduleId", wireType)
			}
			m.ScheduleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduleId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduleAttempt", wireType)
			}
			m.ScheduleAttempt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScheduleAttempt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivityId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActivityId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTaskToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTaskToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTaskToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DomainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskList = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkflowService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkflowService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWorkflowService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWorkflowService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorkflowService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorkflowService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorkflowService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorkflowService = fmt.Errorf("proto: unexpected end of group")
)