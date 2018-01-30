package responseentity

import "net/http"

//Entity 结构体
type Entity struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

//New 新建entity
func (r Entity) New(newCode int, newMsg string) Entity {
	r.Msg = newMsg
	r.Code = newCode
	return r
}

//WithMsg 设置msg
func (r *Entity) WithMsg(newMsg string) *Entity {
	r.Msg = newMsg
	return r
}

//WithAttachMsg 附加msg
func (r *Entity) WithAttachMsg(newMsg string) *Entity {
	r.Msg = r.Msg + newMsg
	return r
}

//WithCode 设置编码
func (r *Entity) WithCode(newCode int) *Entity {
	r.Code = newCode
	return r
}

//ResponseEntity 返回实体
type ResponseEntity struct {
	StatusCode int
	Data       interface{}
}

//NewBuild 新建
func (r *ResponseEntity) NewBuild(StatusCode int, Data interface{}) *ResponseEntity {
	r.StatusCode = StatusCode
	r.Data = Data
	return r
}

//Build 新建状态200
func (r *ResponseEntity) Build(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusOK
	r.Data = Data
	return r
}

//BuildError 新建状态400
func (r *ResponseEntity) BuildError(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusBadRequest
	r.Data = Data
	return r
}

//BuildFormatError  新建状态406
func (r *ResponseEntity) BuildFormatError(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusNotAcceptable
	r.Data = Data
	return r
}

//BuildPostAndPut 新建状态201
func (r *ResponseEntity) BuildPostAndPut(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusCreated
	r.Data = Data
	return r
}

//BuildDelete 新建状态204
func (r *ResponseEntity) BuildDelete(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusNoContent
	r.Data = Data
	return r
}

//BuildDeleteGone  新建状态410
func (r *ResponseEntity) BuildDeleteGone(Data interface{}) *ResponseEntity {
	r.StatusCode = http.StatusGone
	r.Data = Data
	return r
}

//BuildEntity 新建实体
func BuildEntity(newCode int, newMsg string) *Entity {
	return &Entity{newCode, newMsg}
}
