// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/redhat-developer/service-binding-operator/pkg/reconcile/pipeline (interfaces: Context,Service,CRD,Application,ContextProvider,Handler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	binding "github.com/redhat-developer/service-binding-operator/pkg/binding"
	pipeline "github.com/redhat-developer/service-binding-operator/pkg/reconcile/pipeline"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// AddBindingItem mocks base method.
func (m *MockContext) AddBindingItem(arg0 *pipeline.BindingItem) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBindingItem", arg0)
}

// AddBindingItem indicates an expected call of AddBindingItem.
func (mr *MockContextMockRecorder) AddBindingItem(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBindingItem", reflect.TypeOf((*MockContext)(nil).AddBindingItem), arg0)
}

// AddBindings mocks base method.
func (m *MockContext) AddBindings(arg0 pipeline.Bindings) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBindings", arg0)
}

// AddBindings indicates an expected call of AddBindings.
func (mr *MockContextMockRecorder) AddBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBindings", reflect.TypeOf((*MockContext)(nil).AddBindings), arg0)
}

// Applications mocks base method.
func (m *MockContext) Applications() ([]pipeline.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Applications")
	ret0, _ := ret[0].([]pipeline.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Applications indicates an expected call of Applications.
func (mr *MockContextMockRecorder) Applications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Applications", reflect.TypeOf((*MockContext)(nil).Applications))
}

// BindAsFiles mocks base method.
func (m *MockContext) BindAsFiles() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindAsFiles")
	ret0, _ := ret[0].(bool)
	return ret0
}

// BindAsFiles indicates an expected call of BindAsFiles.
func (mr *MockContextMockRecorder) BindAsFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindAsFiles", reflect.TypeOf((*MockContext)(nil).BindAsFiles))
}

// BindingItems mocks base method.
func (m *MockContext) BindingItems() pipeline.BindingItems {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindingItems")
	ret0, _ := ret[0].(pipeline.BindingItems)
	return ret0
}

// BindingItems indicates an expected call of BindingItems.
func (mr *MockContextMockRecorder) BindingItems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindingItems", reflect.TypeOf((*MockContext)(nil).BindingItems))
}

// BindingName mocks base method.
func (m *MockContext) BindingName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindingName")
	ret0, _ := ret[0].(string)
	return ret0
}

// BindingName indicates an expected call of BindingName.
func (mr *MockContextMockRecorder) BindingName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindingName", reflect.TypeOf((*MockContext)(nil).BindingName))
}

// BindingSecretName mocks base method.
func (m *MockContext) BindingSecretName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindingSecretName")
	ret0, _ := ret[0].(string)
	return ret0
}

// BindingSecretName indicates an expected call of BindingSecretName.
func (mr *MockContextMockRecorder) BindingSecretName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindingSecretName", reflect.TypeOf((*MockContext)(nil).BindingSecretName))
}

// Close mocks base method.
func (m *MockContext) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockContextMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockContext)(nil).Close))
}

// EnvBindings mocks base method.
func (m *MockContext) EnvBindings() []*pipeline.EnvBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvBindings")
	ret0, _ := ret[0].([]*pipeline.EnvBinding)
	return ret0
}

// EnvBindings indicates an expected call of EnvBindings.
func (mr *MockContextMockRecorder) EnvBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvBindings", reflect.TypeOf((*MockContext)(nil).EnvBindings))
}

// Error mocks base method.
func (m *MockContext) Error(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", arg0)
}

// Error indicates an expected call of Error.
func (mr *MockContextMockRecorder) Error(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockContext)(nil).Error), arg0)
}

// FlowStatus mocks base method.
func (m *MockContext) FlowStatus() pipeline.FlowStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowStatus")
	ret0, _ := ret[0].(pipeline.FlowStatus)
	return ret0
}

// FlowStatus indicates an expected call of FlowStatus.
func (mr *MockContextMockRecorder) FlowStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowStatus", reflect.TypeOf((*MockContext)(nil).FlowStatus))
}

// Mappings mocks base method.
func (m *MockContext) Mappings() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mappings")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Mappings indicates an expected call of Mappings.
func (mr *MockContextMockRecorder) Mappings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mappings", reflect.TypeOf((*MockContext)(nil).Mappings))
}

// NamingTemplate mocks base method.
func (m *MockContext) NamingTemplate() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamingTemplate")
	ret0, _ := ret[0].(string)
	return ret0
}

// NamingTemplate indicates an expected call of NamingTemplate.
func (mr *MockContextMockRecorder) NamingTemplate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamingTemplate", reflect.TypeOf((*MockContext)(nil).NamingTemplate))
}

// ReadConfigMap mocks base method.
func (m *MockContext) ReadConfigMap(arg0, arg1 string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadConfigMap", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadConfigMap indicates an expected call of ReadConfigMap.
func (mr *MockContextMockRecorder) ReadConfigMap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConfigMap", reflect.TypeOf((*MockContext)(nil).ReadConfigMap), arg0, arg1)
}

// ReadSecret mocks base method.
func (m *MockContext) ReadSecret(arg0, arg1 string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSecret", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSecret indicates an expected call of ReadSecret.
func (mr *MockContextMockRecorder) ReadSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSecret", reflect.TypeOf((*MockContext)(nil).ReadSecret), arg0, arg1)
}

// RetryProcessing mocks base method.
func (m *MockContext) RetryProcessing(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RetryProcessing", arg0)
}

// RetryProcessing indicates an expected call of RetryProcessing.
func (mr *MockContextMockRecorder) RetryProcessing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryProcessing", reflect.TypeOf((*MockContext)(nil).RetryProcessing), arg0)
}

// Services mocks base method.
func (m *MockContext) Services() ([]pipeline.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Services")
	ret0, _ := ret[0].([]pipeline.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Services indicates an expected call of Services.
func (mr *MockContextMockRecorder) Services() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*MockContext)(nil).Services))
}

// SetCondition mocks base method.
func (m *MockContext) SetCondition(arg0 *v1.Condition) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCondition", arg0)
}

// SetCondition indicates an expected call of SetCondition.
func (mr *MockContextMockRecorder) SetCondition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCondition", reflect.TypeOf((*MockContext)(nil).SetCondition), arg0)
}

// StopProcessing mocks base method.
func (m *MockContext) StopProcessing() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopProcessing")
}

// StopProcessing indicates an expected call of StopProcessing.
func (mr *MockContextMockRecorder) StopProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopProcessing", reflect.TypeOf((*MockContext)(nil).StopProcessing))
}

// UnbindRequested mocks base method.
func (m *MockContext) UnbindRequested() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbindRequested")
	ret0, _ := ret[0].(bool)
	return ret0
}

// UnbindRequested indicates an expected call of UnbindRequested.
func (mr *MockContextMockRecorder) UnbindRequested() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindRequested", reflect.TypeOf((*MockContext)(nil).UnbindRequested))
}

// WorkloadResourceTemplate mocks base method.
func (m *MockContext) WorkloadResourceTemplate(arg0 *schema.GroupVersionResource, arg1 string) (*pipeline.WorkloadMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkloadResourceTemplate", arg0, arg1)
	ret0, _ := ret[0].(*pipeline.WorkloadMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkloadResourceTemplate indicates an expected call of WorkloadResourceTemplate.
func (mr *MockContextMockRecorder) WorkloadResourceTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkloadResourceTemplate", reflect.TypeOf((*MockContext)(nil).WorkloadResourceTemplate), arg0, arg1)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddBindingDef mocks base method.
func (m *MockService) AddBindingDef(arg0 binding.Definition) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBindingDef", arg0)
}

// AddBindingDef indicates an expected call of AddBindingDef.
func (mr *MockServiceMockRecorder) AddBindingDef(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBindingDef", reflect.TypeOf((*MockService)(nil).AddBindingDef), arg0)
}

// BindingDefs mocks base method.
func (m *MockService) BindingDefs() []binding.Definition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindingDefs")
	ret0, _ := ret[0].([]binding.Definition)
	return ret0
}

// BindingDefs indicates an expected call of BindingDefs.
func (mr *MockServiceMockRecorder) BindingDefs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindingDefs", reflect.TypeOf((*MockService)(nil).BindingDefs))
}

// CustomResourceDefinition mocks base method.
func (m *MockService) CustomResourceDefinition() (pipeline.CRD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CustomResourceDefinition")
	ret0, _ := ret[0].(pipeline.CRD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CustomResourceDefinition indicates an expected call of CustomResourceDefinition.
func (mr *MockServiceMockRecorder) CustomResourceDefinition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomResourceDefinition", reflect.TypeOf((*MockService)(nil).CustomResourceDefinition))
}

// Id mocks base method.
func (m *MockService) Id() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(*string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockServiceMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockService)(nil).Id))
}

// IsBindable mocks base method.
func (m *MockService) IsBindable() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBindable")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsBindable indicates an expected call of IsBindable.
func (mr *MockServiceMockRecorder) IsBindable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBindable", reflect.TypeOf((*MockService)(nil).IsBindable))
}

// OwnedResources mocks base method.
func (m *MockService) OwnedResources() ([]*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OwnedResources")
	ret0, _ := ret[0].([]*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OwnedResources indicates an expected call of OwnedResources.
func (mr *MockServiceMockRecorder) OwnedResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OwnedResources", reflect.TypeOf((*MockService)(nil).OwnedResources))
}

// Resource mocks base method.
func (m *MockService) Resource() *unstructured.Unstructured {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resource")
	ret0, _ := ret[0].(*unstructured.Unstructured)
	return ret0
}

// Resource indicates an expected call of Resource.
func (mr *MockServiceMockRecorder) Resource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resource", reflect.TypeOf((*MockService)(nil).Resource))
}

// MockCRD is a mock of CRD interface.
type MockCRD struct {
	ctrl     *gomock.Controller
	recorder *MockCRDMockRecorder
}

// MockCRDMockRecorder is the mock recorder for MockCRD.
type MockCRDMockRecorder struct {
	mock *MockCRD
}

// NewMockCRD creates a new mock instance.
func NewMockCRD(ctrl *gomock.Controller) *MockCRD {
	mock := &MockCRD{ctrl: ctrl}
	mock.recorder = &MockCRDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCRD) EXPECT() *MockCRDMockRecorder {
	return m.recorder
}

// Descriptor mocks base method.
func (m *MockCRD) Descriptor() (*pipeline.CRDDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].(*pipeline.CRDDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Descriptor indicates an expected call of Descriptor.
func (mr *MockCRDMockRecorder) Descriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*MockCRD)(nil).Descriptor))
}

// IsBindable mocks base method.
func (m *MockCRD) IsBindable() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBindable")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsBindable indicates an expected call of IsBindable.
func (mr *MockCRDMockRecorder) IsBindable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBindable", reflect.TypeOf((*MockCRD)(nil).IsBindable))
}

// Resource mocks base method.
func (m *MockCRD) Resource() *unstructured.Unstructured {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resource")
	ret0, _ := ret[0].(*unstructured.Unstructured)
	return ret0
}

// Resource indicates an expected call of Resource.
func (mr *MockCRDMockRecorder) Resource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resource", reflect.TypeOf((*MockCRD)(nil).Resource))
}

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// BindablePods mocks base method.
func (m *MockApplication) BindablePods() (*pipeline.MetaPodSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindablePods")
	ret0, _ := ret[0].(*pipeline.MetaPodSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindablePods indicates an expected call of BindablePods.
func (mr *MockApplicationMockRecorder) BindablePods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindablePods", reflect.TypeOf((*MockApplication)(nil).BindablePods))
}

// GroupVersionResource mocks base method.
func (m *MockApplication) GroupVersionResource() schema.GroupVersionResource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupVersionResource")
	ret0, _ := ret[0].(schema.GroupVersionResource)
	return ret0
}

// GroupVersionResource indicates an expected call of GroupVersionResource.
func (mr *MockApplicationMockRecorder) GroupVersionResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupVersionResource", reflect.TypeOf((*MockApplication)(nil).GroupVersionResource))
}

// IsUpdated mocks base method.
func (m *MockApplication) IsUpdated() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpdated")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUpdated indicates an expected call of IsUpdated.
func (mr *MockApplicationMockRecorder) IsUpdated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpdated", reflect.TypeOf((*MockApplication)(nil).IsUpdated))
}

// Resource mocks base method.
func (m *MockApplication) Resource() *unstructured.Unstructured {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resource")
	ret0, _ := ret[0].(*unstructured.Unstructured)
	return ret0
}

// Resource indicates an expected call of Resource.
func (mr *MockApplicationMockRecorder) Resource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resource", reflect.TypeOf((*MockApplication)(nil).Resource))
}

// SecretPath mocks base method.
func (m *MockApplication) SecretPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecretPath indicates an expected call of SecretPath.
func (mr *MockApplicationMockRecorder) SecretPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretPath", reflect.TypeOf((*MockApplication)(nil).SecretPath))
}

// SetMapping mocks base method.
func (m *MockApplication) SetMapping(arg0 pipeline.WorkloadMapping) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMapping", arg0)
}

// SetMapping indicates an expected call of SetMapping.
func (mr *MockApplicationMockRecorder) SetMapping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMapping", reflect.TypeOf((*MockApplication)(nil).SetMapping), arg0)
}

// MockContextProvider is a mock of ContextProvider interface.
type MockContextProvider struct {
	ctrl     *gomock.Controller
	recorder *MockContextProviderMockRecorder
}

// MockContextProviderMockRecorder is the mock recorder for MockContextProvider.
type MockContextProviderMockRecorder struct {
	mock *MockContextProvider
}

// NewMockContextProvider creates a new mock instance.
func NewMockContextProvider(ctrl *gomock.Controller) *MockContextProvider {
	mock := &MockContextProvider{ctrl: ctrl}
	mock.recorder = &MockContextProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContextProvider) EXPECT() *MockContextProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockContextProvider) Get(arg0 interface{}) (pipeline.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(pipeline.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockContextProviderMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContextProvider)(nil).Get), arg0)
}

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockHandler) Handle(arg0 pipeline.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", arg0)
}

// Handle indicates an expected call of Handle.
func (mr *MockHandlerMockRecorder) Handle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockHandler)(nil).Handle), arg0)
}
