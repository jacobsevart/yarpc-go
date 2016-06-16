// Code generated by thriftrw

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

package thrifttestserver

import (
	"github.com/thriftrw/thriftrw-go/protocol"
	"github.com/thriftrw/thriftrw-go/wire"
	"github.com/yarpc/yarpc-go/crossdock/thrift/gauntlet"
	"github.com/yarpc/yarpc-go/crossdock/thrift/gauntlet/service/thrifttest"
	"github.com/yarpc/yarpc-go/encoding/thrift"
)

type Interface interface {
	TestBinary(reqMeta *thrift.ReqMeta, thing []byte) ([]byte, *thrift.ResMeta, error)
	TestByte(reqMeta *thrift.ReqMeta, thing *int8) (int8, *thrift.ResMeta, error)
	TestDouble(reqMeta *thrift.ReqMeta, thing *float64) (float64, *thrift.ResMeta, error)
	TestEnum(reqMeta *thrift.ReqMeta, thing *gauntlet.Numberz) (gauntlet.Numberz, *thrift.ResMeta, error)
	TestException(reqMeta *thrift.ReqMeta, arg *string) (*thrift.ResMeta, error)
	TestI32(reqMeta *thrift.ReqMeta, thing *int32) (int32, *thrift.ResMeta, error)
	TestI64(reqMeta *thrift.ReqMeta, thing *int64) (int64, *thrift.ResMeta, error)
	TestInsanity(reqMeta *thrift.ReqMeta, argument *gauntlet.Insanity) (map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, *thrift.ResMeta, error)
	TestList(reqMeta *thrift.ReqMeta, thing []int32) ([]int32, *thrift.ResMeta, error)
	TestMap(reqMeta *thrift.ReqMeta, thing map[int32]int32) (map[int32]int32, *thrift.ResMeta, error)
	TestMapMap(reqMeta *thrift.ReqMeta, hello *int32) (map[int32]map[int32]int32, *thrift.ResMeta, error)
	TestMulti(reqMeta *thrift.ReqMeta, arg0 *int8, arg1 *int32, arg2 *int64, arg3 map[int16]string, arg4 *gauntlet.Numberz, arg5 *gauntlet.UserId) (*gauntlet.Xtruct, *thrift.ResMeta, error)
	TestMultiException(reqMeta *thrift.ReqMeta, arg0 *string, arg1 *string) (*gauntlet.Xtruct, *thrift.ResMeta, error)
	TestNest(reqMeta *thrift.ReqMeta, thing *gauntlet.Xtruct2) (*gauntlet.Xtruct2, *thrift.ResMeta, error)
	TestSet(reqMeta *thrift.ReqMeta, thing map[int32]struct{}) (map[int32]struct{}, *thrift.ResMeta, error)
	TestString(reqMeta *thrift.ReqMeta, thing *string) (string, *thrift.ResMeta, error)
	TestStringMap(reqMeta *thrift.ReqMeta, thing map[string]string) (map[string]string, *thrift.ResMeta, error)
	TestStruct(reqMeta *thrift.ReqMeta, thing *gauntlet.Xtruct) (*gauntlet.Xtruct, *thrift.ResMeta, error)
	TestTypedef(reqMeta *thrift.ReqMeta, thing *gauntlet.UserId) (gauntlet.UserId, *thrift.ResMeta, error)
	TestVoid(reqMeta *thrift.ReqMeta) (*thrift.ResMeta, error)
}

func New(impl Interface) thrift.Service {
	return service{handler{impl}}
}

type service struct{ h handler }

func (service) Name() string {
	return "ThriftTest"
}

func (service) Protocol() protocol.Protocol {
	return protocol.Binary
}

func (s service) Handlers() map[string]thrift.Handler {
	return map[string]thrift.Handler{"testBinary": thrift.HandlerFunc(s.h.TestBinary), "testByte": thrift.HandlerFunc(s.h.TestByte), "testDouble": thrift.HandlerFunc(s.h.TestDouble), "testEnum": thrift.HandlerFunc(s.h.TestEnum), "testException": thrift.HandlerFunc(s.h.TestException), "testI32": thrift.HandlerFunc(s.h.TestI32), "testI64": thrift.HandlerFunc(s.h.TestI64), "testInsanity": thrift.HandlerFunc(s.h.TestInsanity), "testList": thrift.HandlerFunc(s.h.TestList), "testMap": thrift.HandlerFunc(s.h.TestMap), "testMapMap": thrift.HandlerFunc(s.h.TestMapMap), "testMulti": thrift.HandlerFunc(s.h.TestMulti), "testMultiException": thrift.HandlerFunc(s.h.TestMultiException), "testNest": thrift.HandlerFunc(s.h.TestNest), "testSet": thrift.HandlerFunc(s.h.TestSet), "testString": thrift.HandlerFunc(s.h.TestString), "testStringMap": thrift.HandlerFunc(s.h.TestStringMap), "testStruct": thrift.HandlerFunc(s.h.TestStruct), "testTypedef": thrift.HandlerFunc(s.h.TestTypedef), "testVoid": thrift.HandlerFunc(s.h.TestVoid)}
}

type handler struct{ impl Interface }

func (h handler) TestBinary(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestBinaryArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestBinary(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestBinaryHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestByte(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestByteArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestByte(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestByteHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestDouble(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestDoubleArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestDouble(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestDoubleHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestEnum(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestEnumArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestEnum(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestEnumHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestException(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestExceptionArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	resMeta, err := h.impl.TestException(reqMeta, args.Arg)
	hadError := err != nil
	result, err := thrifttest.TestExceptionHelper.WrapResponse(err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestI32(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestI32Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestI32(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestI32Helper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestI64(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestI64Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestI64(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestI64Helper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestInsanity(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestInsanityArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestInsanity(reqMeta, args.Argument)
	hadError := err != nil
	result, err := thrifttest.TestInsanityHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestList(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestListArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestList(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestListHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestMap(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestMapArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestMap(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestMapHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestMapMap(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestMapMapArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestMapMap(reqMeta, args.Hello)
	hadError := err != nil
	result, err := thrifttest.TestMapMapHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestMulti(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestMultiArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestMulti(reqMeta, args.Arg0, args.Arg1, args.Arg2, args.Arg3, args.Arg4, args.Arg5)
	hadError := err != nil
	result, err := thrifttest.TestMultiHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestMultiException(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestMultiExceptionArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestMultiException(reqMeta, args.Arg0, args.Arg1)
	hadError := err != nil
	result, err := thrifttest.TestMultiExceptionHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestNest(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestNestArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestNest(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestNestHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestSet(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestSetArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestSet(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestSetHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestString(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestStringArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestString(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestStringHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestStringMap(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestStringMapArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestStringMap(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestStringMapHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestStruct(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestStructArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestStruct(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestStructHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestTypedef(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestTypedefArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	success, resMeta, err := h.impl.TestTypedef(reqMeta, args.Thing)
	hadError := err != nil
	result, err := thrifttest.TestTypedefHelper.WrapResponse(success, err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}

func (h handler) TestVoid(reqMeta *thrift.ReqMeta, body wire.Value) (thrift.Response, error) {
	var args thrifttest.TestVoidArgs
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}
	resMeta, err := h.impl.TestVoid(reqMeta)
	hadError := err != nil
	result, err := thrifttest.TestVoidHelper.WrapResponse(err)
	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Meta = resMeta
		response.Body, err = result.ToWire()
	}
	return response, err
}