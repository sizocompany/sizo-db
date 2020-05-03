// Code generated by mockery v1.0.0. DO NOT EDIT.

package db

import (
	types "github.com/sizocompany/sizo-db/pkg/types"
	mock "github.com/stretchr/testify/mock"
	bbolt "go.etcd.io/bbolt"
)

// MockOperation is an autogenerated mock type for the Operation type
type MockOperation struct {
	mock.Mock
}

type OperationBatchUpdateArgs struct {
	Fn         func(*bbolt.Tx) error
	FnAnything bool
}

type OperationBatchUpdateReturns struct {
	Err error
}

type OperationBatchUpdateExpectation struct {
	Args    OperationBatchUpdateArgs
	Returns OperationBatchUpdateReturns
}

func (_m *MockOperation) ApplyBatchUpdateExpectation(e OperationBatchUpdateExpectation) {
	var args []interface{}
	if e.Args.FnAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Fn)
	}
	_m.On("BatchUpdate", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyBatchUpdateExpectations(expectations []OperationBatchUpdateExpectation) {
	for _, e := range expectations {
		_m.ApplyBatchUpdateExpectation(e)
	}
}

// BatchUpdate provides a mock function with given fields: fn
func (_m *MockOperation) BatchUpdate(fn func(*bbolt.Tx) error) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*bbolt.Tx) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationDeleteAdvisoryDetailBucketReturns struct {
	_a0 error
}

type OperationDeleteAdvisoryDetailBucketExpectation struct {
	Returns OperationDeleteAdvisoryDetailBucketReturns
}

func (_m *MockOperation) ApplyDeleteAdvisoryDetailBucketExpectation(e OperationDeleteAdvisoryDetailBucketExpectation) {
	var args []interface{}
	_m.On("DeleteAdvisoryDetailBucket", args...).Return(e.Returns._a0)
}

func (_m *MockOperation) ApplyDeleteAdvisoryDetailBucketExpectations(expectations []OperationDeleteAdvisoryDetailBucketExpectation) {
	for _, e := range expectations {
		_m.ApplyDeleteAdvisoryDetailBucketExpectation(e)
	}
}

// DeleteAdvisoryDetailBucket provides a mock function with given fields:
func (_m *MockOperation) DeleteAdvisoryDetailBucket() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationDeleteVulnerabilityDetailBucketReturns struct {
	Err error
}

type OperationDeleteVulnerabilityDetailBucketExpectation struct {
	Returns OperationDeleteVulnerabilityDetailBucketReturns
}

func (_m *MockOperation) ApplyDeleteVulnerabilityDetailBucketExpectation(e OperationDeleteVulnerabilityDetailBucketExpectation) {
	var args []interface{}
	_m.On("DeleteVulnerabilityDetailBucket", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyDeleteVulnerabilityDetailBucketExpectations(expectations []OperationDeleteVulnerabilityDetailBucketExpectation) {
	for _, e := range expectations {
		_m.ApplyDeleteVulnerabilityDetailBucketExpectation(e)
	}
}

// DeleteVulnerabilityDetailBucket provides a mock function with given fields:
func (_m *MockOperation) DeleteVulnerabilityDetailBucket() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationForEachAdvisoryArgs struct {
	Sources         []string
	SourcesAnything bool
	PkgName         string
	PkgNameAnything bool
}

type OperationForEachAdvisoryReturns struct {
	Value map[string]Value
	Err   error
}

type OperationForEachAdvisoryExpectation struct {
	Args    OperationForEachAdvisoryArgs
	Returns OperationForEachAdvisoryReturns
}

func (_m *MockOperation) ApplyForEachAdvisoryExpectation(e OperationForEachAdvisoryExpectation) {
	var args []interface{}
	if e.Args.SourcesAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Sources)
	}
	if e.Args.PkgNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.PkgName)
	}
	_m.On("ForEachAdvisory", args...).Return(e.Returns.Value, e.Returns.Err)
}

func (_m *MockOperation) ApplyForEachAdvisoryExpectations(expectations []OperationForEachAdvisoryExpectation) {
	for _, e := range expectations {
		_m.ApplyForEachAdvisoryExpectation(e)
	}
}

// ForEachAdvisory provides a mock function with given fields: sources, pkgName
func (_m *MockOperation) ForEachAdvisory(sources []string, pkgName string) (map[string]Value, error) {
	ret := _m.Called(sources, pkgName)

	var r0 map[string]Value
	if rf, ok := ret.Get(0).(func([]string, string) map[string]Value); ok {
		r0 = rf(sources, pkgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]Value)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, string) error); ok {
		r1 = rf(sources, pkgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OperationForEachVulnerabilityIDArgs struct {
	Fn         func(*bbolt.Tx, string) error
	FnAnything bool
}

type OperationForEachVulnerabilityIDReturns struct {
	Err error
}

type OperationForEachVulnerabilityIDExpectation struct {
	Args    OperationForEachVulnerabilityIDArgs
	Returns OperationForEachVulnerabilityIDReturns
}

func (_m *MockOperation) ApplyForEachVulnerabilityIDExpectation(e OperationForEachVulnerabilityIDExpectation) {
	var args []interface{}
	if e.Args.FnAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Fn)
	}
	_m.On("ForEachVulnerabilityID", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyForEachVulnerabilityIDExpectations(expectations []OperationForEachVulnerabilityIDExpectation) {
	for _, e := range expectations {
		_m.ApplyForEachVulnerabilityIDExpectation(e)
	}
}

// ForEachVulnerabilityID provides a mock function with given fields: fn
func (_m *MockOperation) ForEachVulnerabilityID(fn func(*bbolt.Tx, string) error) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*bbolt.Tx, string) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationGetAdvisoriesArgs struct {
	Source          string
	SourceAnything  bool
	PkgName         string
	PkgNameAnything bool
}

type OperationGetAdvisoriesReturns struct {
	Advisories []types.Advisory
	Err        error
}

type OperationGetAdvisoriesExpectation struct {
	Args    OperationGetAdvisoriesArgs
	Returns OperationGetAdvisoriesReturns
}

func (_m *MockOperation) ApplyGetAdvisoriesExpectation(e OperationGetAdvisoriesExpectation) {
	var args []interface{}
	if e.Args.SourceAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Source)
	}
	if e.Args.PkgNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.PkgName)
	}
	_m.On("GetAdvisories", args...).Return(e.Returns.Advisories, e.Returns.Err)
}

func (_m *MockOperation) ApplyGetAdvisoriesExpectations(expectations []OperationGetAdvisoriesExpectation) {
	for _, e := range expectations {
		_m.ApplyGetAdvisoriesExpectation(e)
	}
}

// GetAdvisories provides a mock function with given fields: source, pkgName
func (_m *MockOperation) GetAdvisories(source string, pkgName string) ([]types.Advisory, error) {
	ret := _m.Called(source, pkgName)

	var r0 []types.Advisory
	if rf, ok := ret.Get(0).(func(string, string) []types.Advisory); ok {
		r0 = rf(source, pkgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Advisory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(source, pkgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OperationGetVulnerabilityArgs struct {
	VulnerabilityID         string
	VulnerabilityIDAnything bool
}

type OperationGetVulnerabilityReturns struct {
	Vulnerability types.Vulnerability
	Err           error
}

type OperationGetVulnerabilityExpectation struct {
	Args    OperationGetVulnerabilityArgs
	Returns OperationGetVulnerabilityReturns
}

func (_m *MockOperation) ApplyGetVulnerabilityExpectation(e OperationGetVulnerabilityExpectation) {
	var args []interface{}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	_m.On("GetVulnerability", args...).Return(e.Returns.Vulnerability, e.Returns.Err)
}

func (_m *MockOperation) ApplyGetVulnerabilityExpectations(expectations []OperationGetVulnerabilityExpectation) {
	for _, e := range expectations {
		_m.ApplyGetVulnerabilityExpectation(e)
	}
}

// GetVulnerability provides a mock function with given fields: vulnerabilityID
func (_m *MockOperation) GetVulnerability(vulnerabilityID string) (types.Vulnerability, error) {
	ret := _m.Called(vulnerabilityID)

	var r0 types.Vulnerability
	if rf, ok := ret.Get(0).(func(string) types.Vulnerability); ok {
		r0 = rf(vulnerabilityID)
	} else {
		r0 = ret.Get(0).(types.Vulnerability)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(vulnerabilityID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OperationGetVulnerabilityDetailArgs struct {
	CveID         string
	CveIDAnything bool
}

type OperationGetVulnerabilityDetailReturns struct {
	Detail map[types.SourceID]types.VulnerabilityDetail
	Err    error
}

type OperationGetVulnerabilityDetailExpectation struct {
	Args    OperationGetVulnerabilityDetailArgs
	Returns OperationGetVulnerabilityDetailReturns
}

func (_m *MockOperation) ApplyGetVulnerabilityDetailExpectation(e OperationGetVulnerabilityDetailExpectation) {
	var args []interface{}
	if e.Args.CveIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CveID)
	}
	_m.On("GetVulnerabilityDetail", args...).Return(e.Returns.Detail, e.Returns.Err)
}

func (_m *MockOperation) ApplyGetVulnerabilityDetailExpectations(expectations []OperationGetVulnerabilityDetailExpectation) {
	for _, e := range expectations {
		_m.ApplyGetVulnerabilityDetailExpectation(e)
	}
}

// GetVulnerabilityDetail provides a mock function with given fields: cveID
func (_m *MockOperation) GetVulnerabilityDetail(cveID string) (map[types.SourceID]types.VulnerabilityDetail, error) {
	ret := _m.Called(cveID)

	var r0 map[types.SourceID]types.VulnerabilityDetail
	if rf, ok := ret.Get(0).(func(string) map[types.SourceID]types.VulnerabilityDetail); ok {
		r0 = rf(cveID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[types.SourceID]types.VulnerabilityDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cveID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OperationPutAdvisoryDetailArgs struct {
	Tx                      *bbolt.Tx
	TxAnything              bool
	VulnerabilityID         string
	VulnerabilityIDAnything bool
	PkgName                 string
	PkgNameAnything         bool
	NestedBktNames          []string
	NestedBktNamesAnything  bool
	Advisory                interface{}
	AdvisoryAnything        bool
}

type OperationPutAdvisoryDetailReturns struct {
	Err error
}

type OperationPutAdvisoryDetailExpectation struct {
	Args    OperationPutAdvisoryDetailArgs
	Returns OperationPutAdvisoryDetailReturns
}

func (_m *MockOperation) ApplyPutAdvisoryDetailExpectation(e OperationPutAdvisoryDetailExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	if e.Args.PkgNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.PkgName)
	}
	if e.Args.NestedBktNamesAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.NestedBktNames)
	}
	if e.Args.AdvisoryAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Advisory)
	}
	_m.On("PutAdvisoryDetail", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyPutAdvisoryDetailExpectations(expectations []OperationPutAdvisoryDetailExpectation) {
	for _, e := range expectations {
		_m.ApplyPutAdvisoryDetailExpectation(e)
	}
}

// PutAdvisoryDetail provides a mock function with given fields: tx, vulnerabilityID, pkgName, nestedBktNames, advisory
func (_m *MockOperation) PutAdvisoryDetail(tx *bbolt.Tx, vulnerabilityID string, pkgName string, nestedBktNames []string, advisory interface{}) error {
	ret := _m.Called(tx, vulnerabilityID, pkgName, nestedBktNames, advisory)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, string, []string, interface{}) error); ok {
		r0 = rf(tx, vulnerabilityID, pkgName, nestedBktNames, advisory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationPutDataSourceArgs struct {
	Tx              *bbolt.Tx
	TxAnything      bool
	BktName         string
	BktNameAnything bool
	Source          types.DataSource
	SourceAnything  bool
}

type OperationPutDataSourceReturns struct {
	Err error
}

type OperationPutDataSourceExpectation struct {
	Args    OperationPutDataSourceArgs
	Returns OperationPutDataSourceReturns
}

func (_m *MockOperation) ApplyPutDataSourceExpectation(e OperationPutDataSourceExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.BktNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.BktName)
	}
	if e.Args.SourceAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Source)
	}
	_m.On("PutDataSource", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyPutDataSourceExpectations(expectations []OperationPutDataSourceExpectation) {
	for _, e := range expectations {
		_m.ApplyPutDataSourceExpectation(e)
	}
}

// PutDataSource provides a mock function with given fields: tx, bktName, source
func (_m *MockOperation) PutDataSource(tx *bbolt.Tx, bktName string, source types.DataSource) error {
	ret := _m.Called(tx, bktName, source)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, types.DataSource) error); ok {
		r0 = rf(tx, bktName, source)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationPutRedHatCPEsArgs struct {
	Tx               *bbolt.Tx
	TxAnything       bool
	CpeIndex         int
	CpeIndexAnything bool
	Cpe              string
	CpeAnything      bool
}

type OperationPutRedHatCPEsReturns struct {
	Err error
}

type OperationPutRedHatCPEsExpectation struct {
	Args    OperationPutRedHatCPEsArgs
	Returns OperationPutRedHatCPEsReturns
}

func (_m *MockOperation) ApplyPutRedHatCPEsExpectation(e OperationPutRedHatCPEsExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.CpeIndexAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CpeIndex)
	}
	if e.Args.CpeAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Cpe)
	}
	_m.On("PutRedHatCPEs", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyPutRedHatCPEsExpectations(expectations []OperationPutRedHatCPEsExpectation) {
	for _, e := range expectations {
		_m.ApplyPutRedHatCPEsExpectation(e)
	}
}

// PutRedHatCPEs provides a mock function with given fields: tx, cpeIndex, cpe
func (_m *MockOperation) PutRedHatCPEs(tx *bbolt.Tx, cpeIndex int, cpe string) error {
	ret := _m.Called(tx, cpeIndex, cpe)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, int, string) error); ok {
		r0 = rf(tx, cpeIndex, cpe)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationPutRedHatNVRsArgs struct {
	Tx                 *bbolt.Tx
	TxAnything         bool
	Nvr                string
	NvrAnything        bool
	CpeIndices         []int
	CpeIndicesAnything bool
}

type OperationPutRedHatNVRsReturns struct {
	Err error
}

type OperationPutRedHatNVRsExpectation struct {
	Args    OperationPutRedHatNVRsArgs
	Returns OperationPutRedHatNVRsReturns
}

func (_m *MockOperation) ApplyPutRedHatNVRsExpectation(e OperationPutRedHatNVRsExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.NvrAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Nvr)
	}
	if e.Args.CpeIndicesAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CpeIndices)
	}
	_m.On("PutRedHatNVRs", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyPutRedHatNVRsExpectations(expectations []OperationPutRedHatNVRsExpectation) {
	for _, e := range expectations {
		_m.ApplyPutRedHatNVRsExpectation(e)
	}
}

// PutRedHatNVRs provides a mock function with given fields: tx, nvr, cpeIndices
func (_m *MockOperation) PutRedHatNVRs(tx *bbolt.Tx, nvr string, cpeIndices []int) error {
	ret := _m.Called(tx, nvr, cpeIndices)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, []int) error); ok {
		r0 = rf(tx, nvr, cpeIndices)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationPutRedHatRepositoriesArgs struct {
	Tx                 *bbolt.Tx
	TxAnything         bool
	Repository         string
	RepositoryAnything bool
	CpeIndices         []int
	CpeIndicesAnything bool
}

type OperationPutRedHatRepositoriesReturns struct {
	Err error
}

type OperationPutRedHatRepositoriesExpectation struct {
	Args    OperationPutRedHatRepositoriesArgs
	Returns OperationPutRedHatRepositoriesReturns
}

func (_m *MockOperation) ApplyPutRedHatRepositoriesExpectation(e OperationPutRedHatRepositoriesExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.RepositoryAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Repository)
	}
	if e.Args.CpeIndicesAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CpeIndices)
	}
	_m.On("PutRedHatRepositories", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyPutRedHatRepositoriesExpectations(expectations []OperationPutRedHatRepositoriesExpectation) {
	for _, e := range expectations {
		_m.ApplyPutRedHatRepositoriesExpectation(e)
	}
}

// PutRedHatRepositories provides a mock function with given fields: tx, repository, cpeIndices
func (_m *MockOperation) PutRedHatRepositories(tx *bbolt.Tx, repository string, cpeIndices []int) error {
	ret := _m.Called(tx, repository, cpeIndices)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, []int) error); ok {
		r0 = rf(tx, repository, cpeIndices)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationPutVulnerabilityArgs struct {
	Tx                      *bbolt.Tx
	TxAnything              bool
	VulnerabilityID         string
	VulnerabilityIDAnything bool
	Vulnerability           types.Vulnerability
	VulnerabilityAnything   bool
}

type OperationPutVulnerabilityReturns struct {
	Err error
}

type OperationPutVulnerabilityExpectation struct {
	Args    OperationPutVulnerabilityArgs
	Returns OperationPutVulnerabilityReturns
}

func (_m *MockOperation) ApplyPutVulnerabilityExpectation(e OperationPutVulnerabilityExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	if e.Args.VulnerabilityAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Vulnerability)
	}
	_m.On("PutVulnerability", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyPutVulnerabilityExpectations(expectations []OperationPutVulnerabilityExpectation) {
	for _, e := range expectations {
		_m.ApplyPutVulnerabilityExpectation(e)
	}
}

// PutVulnerability provides a mock function with given fields: tx, vulnerabilityID, vulnerability
func (_m *MockOperation) PutVulnerability(tx *bbolt.Tx, vulnerabilityID string, vulnerability types.Vulnerability) error {
	ret := _m.Called(tx, vulnerabilityID, vulnerability)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, types.Vulnerability) error); ok {
		r0 = rf(tx, vulnerabilityID, vulnerability)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationPutVulnerabilityDetailArgs struct {
	Tx                      *bbolt.Tx
	TxAnything              bool
	VulnerabilityID         string
	VulnerabilityIDAnything bool
	Source                  types.SourceID
	SourceAnything          bool
	Vulnerability           types.VulnerabilityDetail
	VulnerabilityAnything   bool
}

type OperationPutVulnerabilityDetailReturns struct {
	Err error
}

type OperationPutVulnerabilityDetailExpectation struct {
	Args    OperationPutVulnerabilityDetailArgs
	Returns OperationPutVulnerabilityDetailReturns
}

func (_m *MockOperation) ApplyPutVulnerabilityDetailExpectation(e OperationPutVulnerabilityDetailExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	if e.Args.SourceAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Source)
	}
	if e.Args.VulnerabilityAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Vulnerability)
	}
	_m.On("PutVulnerabilityDetail", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyPutVulnerabilityDetailExpectations(expectations []OperationPutVulnerabilityDetailExpectation) {
	for _, e := range expectations {
		_m.ApplyPutVulnerabilityDetailExpectation(e)
	}
}

// PutVulnerabilityDetail provides a mock function with given fields: tx, vulnerabilityID, source, vulnerability
func (_m *MockOperation) PutVulnerabilityDetail(tx *bbolt.Tx, vulnerabilityID string, source types.SourceID, vulnerability types.VulnerabilityDetail) error {
	ret := _m.Called(tx, vulnerabilityID, source, vulnerability)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, types.SourceID, types.VulnerabilityDetail) error); ok {
		r0 = rf(tx, vulnerabilityID, source, vulnerability)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationPutVulnerabilityIDArgs struct {
	Tx                      *bbolt.Tx
	TxAnything              bool
	VulnerabilityID         string
	VulnerabilityIDAnything bool
}

type OperationPutVulnerabilityIDReturns struct {
	Err error
}

type OperationPutVulnerabilityIDExpectation struct {
	Args    OperationPutVulnerabilityIDArgs
	Returns OperationPutVulnerabilityIDReturns
}

func (_m *MockOperation) ApplyPutVulnerabilityIDExpectation(e OperationPutVulnerabilityIDExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	_m.On("PutVulnerabilityID", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplyPutVulnerabilityIDExpectations(expectations []OperationPutVulnerabilityIDExpectation) {
	for _, e := range expectations {
		_m.ApplyPutVulnerabilityIDExpectation(e)
	}
}

// PutVulnerabilityID provides a mock function with given fields: tx, vulnerabilityID
func (_m *MockOperation) PutVulnerabilityID(tx *bbolt.Tx, vulnerabilityID string) error {
	ret := _m.Called(tx, vulnerabilityID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string) error); ok {
		r0 = rf(tx, vulnerabilityID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OperationRedHatNVRToCPEsArgs struct {
	Nvr         string
	NvrAnything bool
}

type OperationRedHatNVRToCPEsReturns struct {
	CpeIndices []int
	Err        error
}

type OperationRedHatNVRToCPEsExpectation struct {
	Args    OperationRedHatNVRToCPEsArgs
	Returns OperationRedHatNVRToCPEsReturns
}

func (_m *MockOperation) ApplyRedHatNVRToCPEsExpectation(e OperationRedHatNVRToCPEsExpectation) {
	var args []interface{}
	if e.Args.NvrAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Nvr)
	}
	_m.On("RedHatNVRToCPEs", args...).Return(e.Returns.CpeIndices, e.Returns.Err)
}

func (_m *MockOperation) ApplyRedHatNVRToCPEsExpectations(expectations []OperationRedHatNVRToCPEsExpectation) {
	for _, e := range expectations {
		_m.ApplyRedHatNVRToCPEsExpectation(e)
	}
}

// RedHatNVRToCPEs provides a mock function with given fields: nvr
func (_m *MockOperation) RedHatNVRToCPEs(nvr string) ([]int, error) {
	ret := _m.Called(nvr)

	var r0 []int
	if rf, ok := ret.Get(0).(func(string) []int); ok {
		r0 = rf(nvr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nvr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OperationRedHatRepoToCPEsArgs struct {
	Repository         string
	RepositoryAnything bool
}

type OperationRedHatRepoToCPEsReturns struct {
	CpeIndices []int
	Err        error
}

type OperationRedHatRepoToCPEsExpectation struct {
	Args    OperationRedHatRepoToCPEsArgs
	Returns OperationRedHatRepoToCPEsReturns
}

func (_m *MockOperation) ApplyRedHatRepoToCPEsExpectation(e OperationRedHatRepoToCPEsExpectation) {
	var args []interface{}
	if e.Args.RepositoryAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Repository)
	}
	_m.On("RedHatRepoToCPEs", args...).Return(e.Returns.CpeIndices, e.Returns.Err)
}

func (_m *MockOperation) ApplyRedHatRepoToCPEsExpectations(expectations []OperationRedHatRepoToCPEsExpectation) {
	for _, e := range expectations {
		_m.ApplyRedHatRepoToCPEsExpectation(e)
	}
}

// RedHatRepoToCPEs provides a mock function with given fields: repository
func (_m *MockOperation) RedHatRepoToCPEs(repository string) ([]int, error) {
	ret := _m.Called(repository)

	var r0 []int
	if rf, ok := ret.Get(0).(func(string) []int); ok {
		r0 = rf(repository)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repository)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OperationSaveAdvisoryDetailsArgs struct {
	Tx            *bbolt.Tx
	TxAnything    bool
	CveID         string
	CveIDAnything bool
}

type OperationSaveAdvisoryDetailsReturns struct {
	Err error
}

type OperationSaveAdvisoryDetailsExpectation struct {
	Args    OperationSaveAdvisoryDetailsArgs
	Returns OperationSaveAdvisoryDetailsReturns
}

func (_m *MockOperation) ApplySaveAdvisoryDetailsExpectation(e OperationSaveAdvisoryDetailsExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.CveIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CveID)
	}
	_m.On("SaveAdvisoryDetails", args...).Return(e.Returns.Err)
}

func (_m *MockOperation) ApplySaveAdvisoryDetailsExpectations(expectations []OperationSaveAdvisoryDetailsExpectation) {
	for _, e := range expectations {
		_m.ApplySaveAdvisoryDetailsExpectation(e)
	}
}

// SaveAdvisoryDetails provides a mock function with given fields: tx, cveID
func (_m *MockOperation) SaveAdvisoryDetails(tx *bbolt.Tx, cveID string) error {
	ret := _m.Called(tx, cveID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string) error); ok {
		r0 = rf(tx, cveID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}