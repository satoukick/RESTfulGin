package service

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/satoukick/webserver/config"
	"github.com/satoukick/webserver/model"
	"github.com/satoukick/webserver/proto"
)

// ToDoRequest handles dbRequest
type ToDoRequest struct {
	db *gorm.DB
}

func NewToDoRequestHandler() *ToDoRequest {
	pgconf := config.Conf.GetPGEnvString()
	db, _ := gorm.Open("postgres", pgconf)
	return &ToDoRequest{db: db}
}

// ToDoQuery implement server DBToDoQuery
func (r *ToDoRequest) ToDoQuery(ctx context.Context, req *proto.ToDoQueryRequest, rsp *proto.ToDoQueryResponse) error {
	id := req.GetId()
	todo := &model.TodoModel{}
	r.getTodomodel(int(id), todo)
	rsp.Completed = todo.Completed == 1
	rsp.Title = todo.Title
	return nil
}

func (r *ToDoRequest) getTodomodel(id int, model *model.TodoModel) {
	r.db.First(model, id)
}
