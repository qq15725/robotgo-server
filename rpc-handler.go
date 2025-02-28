package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sourcegraph/jsonrpc2"
	"log"
	"reflect"
)

type RPCHandler struct{}

func (s *RPCHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	methodName := req.Method
	var params []interface{}
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		log.Println("Error unmarshalling params:", err)
		conn.ReplyWithError(ctx, req.ID, &jsonrpc2.Error{
			Code:    jsonrpc2.CodeInvalidParams,
			Message: "Invalid parameters",
		})
		return
	}

	method := reflect.ValueOf(methodMap[methodName])
	if !method.IsValid() {
		conn.ReplyWithError(ctx, req.ID, &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: fmt.Sprintf("Method %s not found in robotgo\n", methodName),
		})
		return
	}

	methodParams := make([]reflect.Value, len(params))
	for i, param := range params {
		methodParam := reflect.ValueOf(param)
		switch methodParam.Kind() {
		case reflect.Float32, reflect.Float64:
			methodParams[i] = reflect.ValueOf(int(methodParam.Float()))
		default:
			methodParams[i] = methodParam
		}
	}

	methodResult := method.Call(methodParams)

	result := make([]interface{}, len(methodResult))
	for _, value := range methodResult {
		result = append(result, reflect.ValueOf(value))
	}

	conn.Reply(ctx, req.ID, result)
}
