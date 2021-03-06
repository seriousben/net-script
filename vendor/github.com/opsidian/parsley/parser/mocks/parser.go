// Code generated by mockery v1.0.0
package mocks

import data "github.com/opsidian/parsley/data"
import mock "github.com/stretchr/testify/mock"
import parser "github.com/opsidian/parsley/parser"
import reader "github.com/opsidian/parsley/reader"

// Parser is an autogenerated mock type for the Parser type
type Parser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: h, leftRecCtx, r
func (_m *Parser) Parse(h *parser.History, leftRecCtx data.IntMap, r reader.Reader) (data.IntSet, parser.ResultSet, reader.Error) {
	ret := _m.Called(h, leftRecCtx, r)

	var r0 data.IntSet
	if rf, ok := ret.Get(0).(func(*parser.History, data.IntMap, reader.Reader) data.IntSet); ok {
		r0 = rf(h, leftRecCtx, r)
	} else {
		r0 = ret.Get(0).(data.IntSet)
	}

	var r1 parser.ResultSet
	if rf, ok := ret.Get(1).(func(*parser.History, data.IntMap, reader.Reader) parser.ResultSet); ok {
		r1 = rf(h, leftRecCtx, r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(parser.ResultSet)
		}
	}

	var r2 reader.Error
	if rf, ok := ret.Get(2).(func(*parser.History, data.IntMap, reader.Reader) reader.Error); ok {
		r2 = rf(h, leftRecCtx, r)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(reader.Error)
		}
	}

	return r0, r1, r2
}
