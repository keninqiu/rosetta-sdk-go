// Code generated by mockery v1.0.0. DO NOT EDIT.

package modules

import (
	context "context"

	errgroup "github.com/neilotoole/errgroup"
	mock "github.com/stretchr/testify/mock"

	database "github.com/keninqiu/rosetta-sdk-go/storage/database"
	types "github.com/keninqiu/rosetta-sdk-go/types"
)

// BlockWorker is an autogenerated mock type for the BlockWorker type
type BlockWorker struct {
	mock.Mock
}

// AddingBlock provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *BlockWorker) AddingBlock(_a0 context.Context, _a1 *errgroup.Group, _a2 *types.Block, _a3 database.Transaction) (database.CommitWorker, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 database.CommitWorker
	if rf, ok := ret.Get(0).(func(context.Context, *errgroup.Group, *types.Block, database.Transaction) database.CommitWorker); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.CommitWorker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *errgroup.Group, *types.Block, database.Transaction) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemovingBlock provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *BlockWorker) RemovingBlock(_a0 context.Context, _a1 *errgroup.Group, _a2 *types.Block, _a3 database.Transaction) (database.CommitWorker, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 database.CommitWorker
	if rf, ok := ret.Get(0).(func(context.Context, *errgroup.Group, *types.Block, database.Transaction) database.CommitWorker); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.CommitWorker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *errgroup.Group, *types.Block, database.Transaction) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}