package immodels

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ ChatLogModel = (*customChatLogModel)(nil)

type (
	// ChatLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatLogModel.
	ChatLogModel interface {
		chatLogModel
	}

	customChatLogModel struct {
		*defaultChatLogModel
	}
)

// NewChatLogModel returns a model for the mongo.
//func NewChatLogModel(url, db, collection string) ChatLogModel {
//	conn := mon.MustNewModel(url, db, collection)
//	return &customChatLogModel{
//		defaultChatLogModel: newDefaultChatLogModel(conn),
//	}
//}

// 使用到集合是固定的，可以不从外面传递
func NewChatLogModel(url, db string) ChatLogModel {
	conn := mon.MustNewModel(url, db, "chat_log")
	return &customChatLogModel{
		defaultChatLogModel: newDefaultChatLogModel(conn),
	}
}
