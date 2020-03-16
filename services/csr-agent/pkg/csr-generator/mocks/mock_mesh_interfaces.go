// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_csr_generator is a generated GoMock package.
package mock_csr_generator

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/security.zephyr.solo.io/v1alpha1"
	types "github.com/solo-io/mesh-projects/pkg/api/security.zephyr.solo.io/v1alpha1/types"
	cert_secrets "github.com/solo-io/mesh-projects/pkg/security/secrets"
)

// MockCertClient is a mock of CertClient interface
type MockCertClient struct {
	ctrl     *gomock.Controller
	recorder *MockCertClientMockRecorder
}

// MockCertClientMockRecorder is the mock recorder for MockCertClient
type MockCertClientMockRecorder struct {
	mock *MockCertClient
}

// NewMockCertClient creates a new mock instance
func NewMockCertClient(ctrl *gomock.Controller) *MockCertClient {
	mock := &MockCertClient{ctrl: ctrl}
	mock.recorder = &MockCertClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertClient) EXPECT() *MockCertClientMockRecorder {
	return m.recorder
}

// EnsureSecretKey mocks base method
func (m *MockCertClient) EnsureSecretKey(ctx context.Context, obj *v1alpha1.MeshGroupCertificateSigningRequest) (*cert_secrets.RootCaData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureSecretKey", ctx, obj)
	ret0, _ := ret[0].(*cert_secrets.RootCaData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureSecretKey indicates an expected call of EnsureSecretKey
func (mr *MockCertClientMockRecorder) EnsureSecretKey(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSecretKey", reflect.TypeOf((*MockCertClient)(nil).EnsureSecretKey), ctx, obj)
}

// MockIstioCSRGenerator is a mock of IstioCSRGenerator interface
type MockIstioCSRGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIstioCSRGeneratorMockRecorder
}

// MockIstioCSRGeneratorMockRecorder is the mock recorder for MockIstioCSRGenerator
type MockIstioCSRGeneratorMockRecorder struct {
	mock *MockIstioCSRGenerator
}

// NewMockIstioCSRGenerator creates a new mock instance
func NewMockIstioCSRGenerator(ctrl *gomock.Controller) *MockIstioCSRGenerator {
	mock := &MockIstioCSRGenerator{ctrl: ctrl}
	mock.recorder = &MockIstioCSRGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIstioCSRGenerator) EXPECT() *MockIstioCSRGeneratorMockRecorder {
	return m.recorder
}

// GenerateIstioCSR mocks base method
func (m *MockIstioCSRGenerator) GenerateIstioCSR(ctx context.Context, obj *v1alpha1.MeshGroupCertificateSigningRequest) *types.MeshGroupCertificateSigningRequestStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIstioCSR", ctx, obj)
	ret0, _ := ret[0].(*types.MeshGroupCertificateSigningRequestStatus)
	return ret0
}

// GenerateIstioCSR indicates an expected call of GenerateIstioCSR
func (mr *MockIstioCSRGeneratorMockRecorder) GenerateIstioCSR(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIstioCSR", reflect.TypeOf((*MockIstioCSRGenerator)(nil).GenerateIstioCSR), ctx, obj)
}

// MockMeshGroupCSRProcessor is a mock of MeshGroupCSRProcessor interface
type MockMeshGroupCSRProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockMeshGroupCSRProcessorMockRecorder
}

// MockMeshGroupCSRProcessorMockRecorder is the mock recorder for MockMeshGroupCSRProcessor
type MockMeshGroupCSRProcessorMockRecorder struct {
	mock *MockMeshGroupCSRProcessor
}

// NewMockMeshGroupCSRProcessor creates a new mock instance
func NewMockMeshGroupCSRProcessor(ctrl *gomock.Controller) *MockMeshGroupCSRProcessor {
	mock := &MockMeshGroupCSRProcessor{ctrl: ctrl}
	mock.recorder = &MockMeshGroupCSRProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshGroupCSRProcessor) EXPECT() *MockMeshGroupCSRProcessorMockRecorder {
	return m.recorder
}

// ProcessUpsert mocks base method
func (m *MockMeshGroupCSRProcessor) ProcessUpsert(ctx context.Context, csr *v1alpha1.MeshGroupCertificateSigningRequest) *types.MeshGroupCertificateSigningRequestStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessUpsert", ctx, csr)
	ret0, _ := ret[0].(*types.MeshGroupCertificateSigningRequestStatus)
	return ret0
}

// ProcessUpsert indicates an expected call of ProcessUpsert
func (mr *MockMeshGroupCSRProcessorMockRecorder) ProcessUpsert(ctx, csr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessUpsert", reflect.TypeOf((*MockMeshGroupCSRProcessor)(nil).ProcessUpsert), ctx, csr)
}

// ProcessDelete mocks base method
func (m *MockMeshGroupCSRProcessor) ProcessDelete(ctx context.Context, csr *v1alpha1.MeshGroupCertificateSigningRequest) *types.MeshGroupCertificateSigningRequestStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessDelete", ctx, csr)
	ret0, _ := ret[0].(*types.MeshGroupCertificateSigningRequestStatus)
	return ret0
}

// ProcessDelete indicates an expected call of ProcessDelete
func (mr *MockMeshGroupCSRProcessorMockRecorder) ProcessDelete(ctx, csr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDelete", reflect.TypeOf((*MockMeshGroupCSRProcessor)(nil).ProcessDelete), ctx, csr)
}