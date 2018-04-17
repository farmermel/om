// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ConfigureProductService struct {
	ListStagedProductsStub        func() (api.StagedProductsOutput, error)
	listStagedProductsMutex       sync.RWMutex
	listStagedProductsArgsForCall []struct{}
	listStagedProductsReturns     struct {
		result1 api.StagedProductsOutput
		result2 error
	}
	listStagedProductsReturnsOnCall map[int]struct {
		result1 api.StagedProductsOutput
		result2 error
	}
	ConfigureStub        func(api.ProductsConfigurationInput) error
	configureMutex       sync.RWMutex
	configureArgsForCall []struct {
		arg1 api.ProductsConfigurationInput
	}
	configureReturns struct {
		result1 error
	}
	configureReturnsOnCall map[int]struct {
		result1 error
	}
	ListStagedProductJobsStub        func(productGUID string) (map[string]string, error)
	listStagedProductJobsMutex       sync.RWMutex
	listStagedProductJobsArgsForCall []struct {
		productGUID string
	}
	listStagedProductJobsReturns struct {
		result1 map[string]string
		result2 error
	}
	listStagedProductJobsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetStagedProductJobResourceConfigStub        func(productGUID, jobGUID string) (api.JobProperties, error)
	getStagedProductJobResourceConfigMutex       sync.RWMutex
	getStagedProductJobResourceConfigArgsForCall []struct {
		productGUID string
		jobGUID     string
	}
	getStagedProductJobResourceConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	getStagedProductJobResourceConfigReturnsOnCall map[int]struct {
		result1 api.JobProperties
		result2 error
	}
	UpdateStagedProductJobResourceConfigStub        func(productGUID, jobGUID string, jobProperties api.JobProperties) error
	updateStagedProductJobResourceConfigMutex       sync.RWMutex
	updateStagedProductJobResourceConfigArgsForCall []struct {
		productGUID   string
		jobGUID       string
		jobProperties api.JobProperties
	}
	updateStagedProductJobResourceConfigReturns struct {
		result1 error
	}
	updateStagedProductJobResourceConfigReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigureProductService) ListStagedProducts() (api.StagedProductsOutput, error) {
	fake.listStagedProductsMutex.Lock()
	ret, specificReturn := fake.listStagedProductsReturnsOnCall[len(fake.listStagedProductsArgsForCall)]
	fake.listStagedProductsArgsForCall = append(fake.listStagedProductsArgsForCall, struct{}{})
	fake.recordInvocation("ListStagedProducts", []interface{}{})
	fake.listStagedProductsMutex.Unlock()
	if fake.ListStagedProductsStub != nil {
		return fake.ListStagedProductsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listStagedProductsReturns.result1, fake.listStagedProductsReturns.result2
}

func (fake *ConfigureProductService) ListStagedProductsCallCount() int {
	fake.listStagedProductsMutex.RLock()
	defer fake.listStagedProductsMutex.RUnlock()
	return len(fake.listStagedProductsArgsForCall)
}

func (fake *ConfigureProductService) ListStagedProductsReturns(result1 api.StagedProductsOutput, result2 error) {
	fake.ListStagedProductsStub = nil
	fake.listStagedProductsReturns = struct {
		result1 api.StagedProductsOutput
		result2 error
	}{result1, result2}
}

func (fake *ConfigureProductService) ListStagedProductsReturnsOnCall(i int, result1 api.StagedProductsOutput, result2 error) {
	fake.ListStagedProductsStub = nil
	if fake.listStagedProductsReturnsOnCall == nil {
		fake.listStagedProductsReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsOutput
			result2 error
		})
	}
	fake.listStagedProductsReturnsOnCall[i] = struct {
		result1 api.StagedProductsOutput
		result2 error
	}{result1, result2}
}

func (fake *ConfigureProductService) Configure(arg1 api.ProductsConfigurationInput) error {
	fake.configureMutex.Lock()
	ret, specificReturn := fake.configureReturnsOnCall[len(fake.configureArgsForCall)]
	fake.configureArgsForCall = append(fake.configureArgsForCall, struct {
		arg1 api.ProductsConfigurationInput
	}{arg1})
	fake.recordInvocation("Configure", []interface{}{arg1})
	fake.configureMutex.Unlock()
	if fake.ConfigureStub != nil {
		return fake.ConfigureStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.configureReturns.result1
}

func (fake *ConfigureProductService) ConfigureCallCount() int {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return len(fake.configureArgsForCall)
}

func (fake *ConfigureProductService) ConfigureArgsForCall(i int) api.ProductsConfigurationInput {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return fake.configureArgsForCall[i].arg1
}

func (fake *ConfigureProductService) ConfigureReturns(result1 error) {
	fake.ConfigureStub = nil
	fake.configureReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureProductService) ConfigureReturnsOnCall(i int, result1 error) {
	fake.ConfigureStub = nil
	if fake.configureReturnsOnCall == nil {
		fake.configureReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configureReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureProductService) ListStagedProductJobs(productGUID string) (map[string]string, error) {
	fake.listStagedProductJobsMutex.Lock()
	ret, specificReturn := fake.listStagedProductJobsReturnsOnCall[len(fake.listStagedProductJobsArgsForCall)]
	fake.listStagedProductJobsArgsForCall = append(fake.listStagedProductJobsArgsForCall, struct {
		productGUID string
	}{productGUID})
	fake.recordInvocation("ListStagedProductJobs", []interface{}{productGUID})
	fake.listStagedProductJobsMutex.Unlock()
	if fake.ListStagedProductJobsStub != nil {
		return fake.ListStagedProductJobsStub(productGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listStagedProductJobsReturns.result1, fake.listStagedProductJobsReturns.result2
}

func (fake *ConfigureProductService) ListStagedProductJobsCallCount() int {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	return len(fake.listStagedProductJobsArgsForCall)
}

func (fake *ConfigureProductService) ListStagedProductJobsArgsForCall(i int) string {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	return fake.listStagedProductJobsArgsForCall[i].productGUID
}

func (fake *ConfigureProductService) ListStagedProductJobsReturns(result1 map[string]string, result2 error) {
	fake.ListStagedProductJobsStub = nil
	fake.listStagedProductJobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *ConfigureProductService) ListStagedProductJobsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.ListStagedProductJobsStub = nil
	if fake.listStagedProductJobsReturnsOnCall == nil {
		fake.listStagedProductJobsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.listStagedProductJobsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *ConfigureProductService) GetStagedProductJobResourceConfig(productGUID string, jobGUID string) (api.JobProperties, error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	ret, specificReturn := fake.getStagedProductJobResourceConfigReturnsOnCall[len(fake.getStagedProductJobResourceConfigArgsForCall)]
	fake.getStagedProductJobResourceConfigArgsForCall = append(fake.getStagedProductJobResourceConfigArgsForCall, struct {
		productGUID string
		jobGUID     string
	}{productGUID, jobGUID})
	fake.recordInvocation("GetStagedProductJobResourceConfig", []interface{}{productGUID, jobGUID})
	fake.getStagedProductJobResourceConfigMutex.Unlock()
	if fake.GetStagedProductJobResourceConfigStub != nil {
		return fake.GetStagedProductJobResourceConfigStub(productGUID, jobGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductJobResourceConfigReturns.result1, fake.getStagedProductJobResourceConfigReturns.result2
}

func (fake *ConfigureProductService) GetStagedProductJobResourceConfigCallCount() int {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	return len(fake.getStagedProductJobResourceConfigArgsForCall)
}

func (fake *ConfigureProductService) GetStagedProductJobResourceConfigArgsForCall(i int) (string, string) {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	return fake.getStagedProductJobResourceConfigArgsForCall[i].productGUID, fake.getStagedProductJobResourceConfigArgsForCall[i].jobGUID
}

func (fake *ConfigureProductService) GetStagedProductJobResourceConfigReturns(result1 api.JobProperties, result2 error) {
	fake.GetStagedProductJobResourceConfigStub = nil
	fake.getStagedProductJobResourceConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *ConfigureProductService) GetStagedProductJobResourceConfigReturnsOnCall(i int, result1 api.JobProperties, result2 error) {
	fake.GetStagedProductJobResourceConfigStub = nil
	if fake.getStagedProductJobResourceConfigReturnsOnCall == nil {
		fake.getStagedProductJobResourceConfigReturnsOnCall = make(map[int]struct {
			result1 api.JobProperties
			result2 error
		})
	}
	fake.getStagedProductJobResourceConfigReturnsOnCall[i] = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *ConfigureProductService) UpdateStagedProductJobResourceConfig(productGUID string, jobGUID string, jobProperties api.JobProperties) error {
	fake.updateStagedProductJobResourceConfigMutex.Lock()
	ret, specificReturn := fake.updateStagedProductJobResourceConfigReturnsOnCall[len(fake.updateStagedProductJobResourceConfigArgsForCall)]
	fake.updateStagedProductJobResourceConfigArgsForCall = append(fake.updateStagedProductJobResourceConfigArgsForCall, struct {
		productGUID   string
		jobGUID       string
		jobProperties api.JobProperties
	}{productGUID, jobGUID, jobProperties})
	fake.recordInvocation("UpdateStagedProductJobResourceConfig", []interface{}{productGUID, jobGUID, jobProperties})
	fake.updateStagedProductJobResourceConfigMutex.Unlock()
	if fake.UpdateStagedProductJobResourceConfigStub != nil {
		return fake.UpdateStagedProductJobResourceConfigStub(productGUID, jobGUID, jobProperties)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedProductJobResourceConfigReturns.result1
}

func (fake *ConfigureProductService) UpdateStagedProductJobResourceConfigCallCount() int {
	fake.updateStagedProductJobResourceConfigMutex.RLock()
	defer fake.updateStagedProductJobResourceConfigMutex.RUnlock()
	return len(fake.updateStagedProductJobResourceConfigArgsForCall)
}

func (fake *ConfigureProductService) UpdateStagedProductJobResourceConfigArgsForCall(i int) (string, string, api.JobProperties) {
	fake.updateStagedProductJobResourceConfigMutex.RLock()
	defer fake.updateStagedProductJobResourceConfigMutex.RUnlock()
	return fake.updateStagedProductJobResourceConfigArgsForCall[i].productGUID, fake.updateStagedProductJobResourceConfigArgsForCall[i].jobGUID, fake.updateStagedProductJobResourceConfigArgsForCall[i].jobProperties
}

func (fake *ConfigureProductService) UpdateStagedProductJobResourceConfigReturns(result1 error) {
	fake.UpdateStagedProductJobResourceConfigStub = nil
	fake.updateStagedProductJobResourceConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureProductService) UpdateStagedProductJobResourceConfigReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedProductJobResourceConfigStub = nil
	if fake.updateStagedProductJobResourceConfigReturnsOnCall == nil {
		fake.updateStagedProductJobResourceConfigReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedProductJobResourceConfigReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureProductService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listStagedProductsMutex.RLock()
	defer fake.listStagedProductsMutex.RUnlock()
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	fake.updateStagedProductJobResourceConfigMutex.RLock()
	defer fake.updateStagedProductJobResourceConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigureProductService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}