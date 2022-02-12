package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
)

const (
	vsn                      = "2.0"
	serviceMethodSeparator   = "_"
	subscribeMethodSuffix    = "_subscribe"
	unsubscribeMethodSuffix  = "_unsubscribe"
	notificationMethodSuffix = "_subscription"

	defaultWriteTimeout = 10 * time.Second // used if context has no deadline
)

var null = json.RawMessage("null")

type subscriptionResult struct {
	Id     string          `json:"subscription"`
	Result json.RawMessage `json:"result,omitempty"`
}

// A value of this type can a JSON-RPC request, notification, successful response or
// error response. Which one it is depends on the fields.
type JsonMessage struct {
	JsonRpc string          `json:"jsonrpc,omitempty"`
	Id      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func (msg *JsonMessage) IsNotification() bool {
	return msg.Id == nil && msg.Method != ""
}

func (msg *JsonMessage) IsCall() bool {
	return msg.HasValidId() && msg.Method != ""
}

func (msg *JsonMessage) IsResponse() bool {
	return msg.HasValidId() && msg.Method == "" && msg.Params == nil && (msg.Result != nil || msg.Error != nil)
}

func (msg *JsonMessage) HasValidId() bool {
	return len(msg.Id) > 0 && msg.Id[0] != '{' && msg.Id[0] != '['
}

func (msg *JsonMessage) IsSubscribe() bool {
	return strings.HasSuffix(msg.Method, subscribeMethodSuffix)
}
func (msg *JsonMessage) IsError() bool {
	return msg.Error != nil
}
func (msg *JsonMessage) IsUnsubscribe() bool {
	return strings.HasSuffix(msg.Method, unsubscribeMethodSuffix)
}

func (msg *JsonMessage) namespace() string {
	elem := strings.SplitN(msg.Method, serviceMethodSeparator, 2)
	return elem[0]
}

func (msg *JsonMessage) String() string {
	b, _ := json.Marshal(msg)
	return string(b)
}

func (msg *JsonMessage) ErrorResponse(err error) *JsonMessage {
	resp := ErrorMessage(err)
	resp.Id = msg.Id
	return resp
}

func (msg *JsonMessage) Response(result interface{}) *JsonMessage {
	enc, err := json.Marshal(result)
	if err != nil {
		return msg.ErrorResponse(err)
	}

	resp := &JsonMessage{JsonRpc: vsn, Id: msg.Id, Result: enc}

	return resp
}
func (msg *JsonMessage) SubscriptionResult(result *JsonMessage) *JsonMessage {
	resp := &JsonMessage{JsonRpc: vsn, Id: msg.Id}
	err := json.Unmarshal(result.Result, &resp)
	if err != nil {
		return msg.ErrorResponse(err)
	}

	resp.JsonRpc = vsn
	resp.Id = msg.Id
	return resp
}

func ErrorMessage(err error) *JsonMessage {
	msg := &JsonMessage{JsonRpc: vsn, Id: null, Error: &jsonError{
		Code:    -1,
		Message: err.Error(),
	}}
	ec, ok := err.(Error)
	if ok {
		msg.Error.Code = ec.ErrorCode()
	}
	// de, ok := err.(DataError)
	// if ok {
	// 	msg.Error.Data = de.ErrorData()
	// }
	return msg
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err *jsonError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("json-rpc error %d", err.Code)
	}
	return err.Message
}

func (err *jsonError) ErrorCode() int {
	return err.Code
}

func (err *jsonError) ErrorData() interface{} {
	return err.Data
}

func parseMessage(raw json.RawMessage) ([]*JsonMessage, bool) {
	if !isBatch(raw) {
		msgs := []*JsonMessage{{}}
		json.Unmarshal(raw, &msgs[0])
		return msgs, false
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.Token() // skip '['
	var msgs []*JsonMessage
	for dec.More() {
		msgs = append(msgs, new(JsonMessage))
		dec.Decode(&msgs[len(msgs)-1])
	}
	return msgs, true
}

// isBatch returns true when the first non-whitespace characters is '['
func isBatch(raw json.RawMessage) bool {
	for _, c := range raw {
		// skip insignificant whitespace (http://www.ietf.org/rfc/rfc4627.txt)
		if c == 0x20 || c == 0x09 || c == 0x0a || c == 0x0d {
			continue
		}
		return c == '['
	}
	return false
}

// parsePositionalArguments tries to parse the given args to an array of values with the
// given types. It returns the parsed values or an error when the args could not be
// parsed. Missing optional arguments are returned as reflect.Zero values.
func parsePositionalArguments(rawArgs json.RawMessage, types []reflect.Type) ([]reflect.Value, error) {
	dec := json.NewDecoder(bytes.NewReader(rawArgs))
	var args []reflect.Value
	tok, err := dec.Token()
	switch {
	case err == io.EOF || tok == nil && err == nil:
		// "params" is optional and may be empty. Also allow "params":null even though it's
		// not in the spec because our own client used to send it.
	case err != nil:
		return nil, err
	case tok == json.Delim('['):
		// Read argument array.
		if args, err = parseArgumentArray(dec, types); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("non-array args")
	}
	// Set any missing args to nil.
	for i := len(args); i < len(types); i++ {
		if types[i].Kind() != reflect.Ptr {
			return nil, fmt.Errorf("missing value for required argument %d", i)
		}
		args = append(args, reflect.Zero(types[i]))
	}
	return args, nil
}

func parseArgumentArray(dec *json.Decoder, types []reflect.Type) ([]reflect.Value, error) {
	args := make([]reflect.Value, 0, len(types))
	for i := 0; dec.More(); i++ {
		if i >= len(types) {
			return args, fmt.Errorf("too many arguments, want at most %d", len(types))
		}
		argval := reflect.New(types[i])
		if err := dec.Decode(argval.Interface()); err != nil {
			return args, fmt.Errorf("invalid argument %d: %v", i, err)
		}
		if argval.IsNil() && types[i].Kind() != reflect.Ptr {
			return args, fmt.Errorf("missing value for required argument %d", i)
		}
		args = append(args, argval.Elem())
	}
	// Read end of args array.
	_, err := dec.Token()
	return args, err
}

// parseSubscriptionName extracts the subscription name from an encoded argument array.
func parseSubscriptionName(rawArgs json.RawMessage) (string, error) {
	dec := json.NewDecoder(bytes.NewReader(rawArgs))
	if tok, _ := dec.Token(); tok != json.Delim('[') {
		return "", errors.New("non-array args")
	}
	v, _ := dec.Token()
	method, ok := v.(string)
	if !ok {
		return "", errors.New("expected subscription name as first argument")
	}
	return method, nil
}

//todo: build msg with request
func ParseResult(data interface{}) *JsonMessage {
	body, err := json.Marshal(data)
	if err != nil {
		g.Log().Error(err)
		return ErrorMessage(err)
	}
	msg := &JsonMessage{}
	msg.Result = body
	return msg
}

//todo: build msg with request
func ParseMessageData(data interface{}) *JsonMessage {
	body, err := json.Marshal(data)
	if err != nil {
		g.Log().Error(err)
		return ErrorMessage(err)
	}
	return ParseMessage(body)
}

//todo: build msg with request
func ParseMessage(body []byte) *JsonMessage {
	req := &JsonMessage{}
	err := json.Unmarshal(body, req)

	if err != nil {
		g.Log().Error(err)
		return ErrorMessage(err)
	}
	g.Log().Debug("Api req:", req.Method)
	return req
}
