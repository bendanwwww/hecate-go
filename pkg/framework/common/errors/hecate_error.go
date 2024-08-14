package errors

import (
	"fmt"
)

type ErrorCode ErrorCodeField

type ErrorCodeField struct {
	code    string
	message string
}

var (
	NodeNotSupport   ErrorCode = ErrorCode{code: "100001", message: "node type not support"}
	NodeUnreachable  ErrorCode = ErrorCode{code: "100002", message: "node is unreachable"}
	NodeSameName     ErrorCode = ErrorCode{code: "100003", message: "same node name"}
	CycleMap         ErrorCode = ErrorCode{code: "100004", message: "map existence cycle"}
	NullNodeName     ErrorCode = ErrorCode{code: "100005", message: "node name is null"}
	EndNameCanNotUse ErrorCode = ErrorCode{code: "100006", message: "node name \"End\" can not use"}
	NullNodeOperator ErrorCode = ErrorCode{code: "100007", message: "node operator is null"}
	ScenesNotFound   ErrorCode = ErrorCode{code: "200001", message: "map scenes not found"}
	LogRedisNotFound ErrorCode = ErrorCode{code: "999998", message: "log redis not found"}
	DefaultError     ErrorCode = ErrorCode{code: "999999", message: "system error"}
)

type HecateError struct {
	Code    string
	Message string
}

func NewHecateException(message string) *HecateError {
	return NewHecateExceptionWithCodeAndMsg(DefaultError, message)
}

func NewHecateExceptionWithCode(code ErrorCode) *HecateError {
	return NewHecateExceptionWithCodeAndMsg(code, code.message)
}

func NewHecateExceptionWithCodeAndMsg(code ErrorCode, message string) *HecateError {
	return &HecateError{
		Code:    code.code,
		Message: message,
	}
}

func (e *HecateError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

var _ error = (*HecateError)(nil)
