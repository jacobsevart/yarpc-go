// Code generated by thriftrw
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package thrifttest

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type TestListArgs struct {
	Thing []int32 `json:"thing"`
}

type _List_I32_ValueList []int32

func (v _List_I32_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := wire.NewValueI32(x), error(nil)
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_I32_ValueList) Size() int {
	return len(v)
}

func (_List_I32_ValueList) ValueType() wire.Type {
	return wire.TI32
}

func (_List_I32_ValueList) Close() {
}

func (v *TestListArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Thing != nil {
		w, err = wire.NewValueList(_List_I32_ValueList(v.Thing)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _List_I32_Read(l wire.ValueList) ([]int32, error) {
	if l.ValueType() != wire.TI32 {
		return nil, nil
	}
	o := make([]int32, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := x.GetI32(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func (v *TestListArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.Thing, err = _List_I32_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *TestListArgs) String() string {
	var fields [1]string
	i := 0
	if v.Thing != nil {
		fields[i] = fmt.Sprintf("Thing: %v", v.Thing)
		i++
	}
	return fmt.Sprintf("TestListArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestListArgs) MethodName() string {
	return "testList"
}

func (v *TestListArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type TestListResult struct {
	Success []int32 `json:"success"`
}

func (v *TestListResult) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueList(_List_I32_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("TestListResult should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *TestListResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TList {
				v.Success, err = _List_I32_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("TestListResult should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *TestListResult) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("TestListResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestListResult) MethodName() string {
	return "testList"
}

func (v *TestListResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var TestListHelper = struct {
	IsException    func(error) bool
	Args           func(thing []int32) *TestListArgs
	WrapResponse   func([]int32, error) (*TestListResult, error)
	UnwrapResponse func(*TestListResult) ([]int32, error)
}{}

func init() {
	TestListHelper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	TestListHelper.Args = func(thing []int32) *TestListArgs {
		return &TestListArgs{Thing: thing}
	}
	TestListHelper.WrapResponse = func(success []int32, err error) (*TestListResult, error) {
		if err == nil {
			return &TestListResult{Success: success}, nil
		}
		return nil, err
	}
	TestListHelper.UnwrapResponse = func(result *TestListResult) (success []int32, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}
