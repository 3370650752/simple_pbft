package consensus

//记录是谁请求的，请求的操作，请求的序列号
type RequestMsg struct {
	//时间戳
	Timestamp  int64  `json:"timestamp"`
	//客户端ID
	ClientID   string `json:"clientID"`
	//请求的操作
	Operation  string `json:"operation"`
	//给这个请求分配个序列号
	SequenceID int64  `json:"sequenceID"`
}

//记录是谁回复的，回复的结果，回复的序列号
type ReplyMsg struct {
	//视图编号
	ViewID    int64  `json:"viewID"`
	//时间戳
	Timestamp int64  `json:"timestamp"`
	//客户端ID
	ClientID  string `json:"clientID"`
	//节点ID，记录谁回复的
	NodeID    string `json:"nodeID"`
	//操作结果
	Result    string `json:"result"`
}

type PrePrepareMsg struct {
	//视图编号
	ViewID     int64       `json:"viewID"`
	//序列号,多个请求同时发起，通过序列号进行排序
	SequenceID int64       `json:"sequenceID"`
	//数字签名，记录合法性
	Digest     string      `json:"digest"`
	//包含原始请求。1.与Digest对应，验证消息完整性。2.简化消息流程，其他节点不需要再请求一轮requestmsg
	RequestMsg *RequestMsg `json:"requestMsg"`
}

type VoteMsg struct {
	//视图编号
	ViewID     int64  `json:"viewID"`
	//序列号
	SequenceID int64  `json:"sequenceID"`
	//数字签名，记录合法性
	Digest     string `json:"digest"`
	//节点ID，记录谁回复的
	NodeID     string `json:"nodeID"`
	//回复结果
	MsgType           `json:"msgType"`
}

//用MsgType替代int，是个别名，方便读代码
type MsgType int
//多个常量同时声明，如果忽略了值则表示和上面一行的值相同
const (
	PrepareMsg MsgType = iota
	CommitMsg
)
