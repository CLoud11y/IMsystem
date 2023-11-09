package common

const (
	MsgIdPing uint32 = iota
	MsgIdPong
	MsgIdWho
	MsgIdBroadcast
)

// 客户端输入指令与msgId的对应关系
var InstructionMap map[string]uint32

func init() {
	m := map[string]uint32{
		"ping": MsgIdPing,
		"who":  MsgIdWho,
	}
	InstructionMap = m
}
