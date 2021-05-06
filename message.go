// 操作日志

package raft

// MessageType 操作日志类型
type MessageType int32

const (
	MessageHup            MessageType = 0
	MessageBeat           MessageType = 1
	MessageProp           MessageType = 2
	MessageApp            MessageType = 3
	MessageAppResp        MessageType = 4
	MessageVote           MessageType = 5
	MessageVoteResp       MessageType = 6
	MessageSnap           MessageType = 7
	MessageHeartbeat      MessageType = 8
	MessageHeartbeatResp  MessageType = 9
	MessageUnreachable    MessageType = 10
	MessageSnapStatus     MessageType = 11
	MessageCheckQuorum    MessageType = 12
	MessageTransferLeader MessageType = 13
	MessageTimeoutNow     MessageType = 14
	MessageReadIndex      MessageType = 15
	MessageReadIndexResp  MessageType = 16
	MessagePreVote        MessageType = 17
	MessagePreVoteResp    MessageType = 18
)

// message raft 状态机命令
type message struct {
	Index uint64

	Type MessageType

	From uint64

	To uint64
}
