// 操作负载

package raft

import (
	math_bits "math/bits"
)

type EntryType int32

type Entry struct {
	Term  uint64
	Index uint64
	Type  EntryType
	Data  []byte
}

func sovRaft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func (m *Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.Type))
	n += 1 + sovRaft(uint64(m.Term))
	n += 1 + sovRaft(uint64(m.Index))
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovRaft(uint64(l))
	}
	return n
}
