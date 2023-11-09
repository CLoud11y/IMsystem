package common

const (
	// client to server
	MsgIdPing uint32 = iota
	MsgIdWho
	MsgIdRename
	MsgIdPublic
	// server to client
	MsgIdShow
)

// 客户端输入指令与msgId的对应关系
var InstructionMap map[string]uint32

func init() {
	m := map[string]uint32{
		"ping":   MsgIdPing,
		"who":    MsgIdWho,
		"rename": MsgIdRename,
		"public": MsgIdPublic,
	}
	InstructionMap = m
}
