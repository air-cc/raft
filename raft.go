// raft 协议实体

package raft

// StateType represents the role of a node in a cluster.
type StateType uint64

type raft struct {
	// 节点 ID
	id uint64

	// 节点当前所处状态(角色)
	state StateType

	// 节点已知最新 term
	Term uint64

	// 节点投票候选者 ID
	Vote uint64

	// lead 节点ID
	lead uint64

	// 状态机操作日志
	logs []message
}
