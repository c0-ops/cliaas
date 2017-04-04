// This file was generated by counterfeiter
package azurefakes

import (
	"sync"

	"github.com/pivotal-cf/cliaas/iaas/azure"
)

type FakeBlobCopier struct {
	CopyBlobStub        func(container, name, sourceBlob string) error
	copyBlobMutex       sync.RWMutex
	copyBlobArgsForCall []struct {
		container  string
		name       string
		sourceBlob string
	}
	copyBlobReturns struct {
		result1 error
	}
	copyBlobReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlobCopier) CopyBlob(container string, name string, sourceBlob string) error {
	fake.copyBlobMutex.Lock()
	ret, specificReturn := fake.copyBlobReturnsOnCall[len(fake.copyBlobArgsForCall)]
	fake.copyBlobArgsForCall = append(fake.copyBlobArgsForCall, struct {
		container  string
		name       string
		sourceBlob string
	}{container, name, sourceBlob})
	fake.recordInvocation("CopyBlob", []interface{}{container, name, sourceBlob})
	fake.copyBlobMutex.Unlock()
	if fake.CopyBlobStub != nil {
		return fake.CopyBlobStub(container, name, sourceBlob)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyBlobReturns.result1
}

func (fake *FakeBlobCopier) CopyBlobCallCount() int {
	fake.copyBlobMutex.RLock()
	defer fake.copyBlobMutex.RUnlock()
	return len(fake.copyBlobArgsForCall)
}

func (fake *FakeBlobCopier) CopyBlobArgsForCall(i int) (string, string, string) {
	fake.copyBlobMutex.RLock()
	defer fake.copyBlobMutex.RUnlock()
	return fake.copyBlobArgsForCall[i].container, fake.copyBlobArgsForCall[i].name, fake.copyBlobArgsForCall[i].sourceBlob
}

func (fake *FakeBlobCopier) CopyBlobReturns(result1 error) {
	fake.CopyBlobStub = nil
	fake.copyBlobReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobCopier) CopyBlobReturnsOnCall(i int, result1 error) {
	fake.CopyBlobStub = nil
	if fake.copyBlobReturnsOnCall == nil {
		fake.copyBlobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyBlobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobCopier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.copyBlobMutex.RLock()
	defer fake.copyBlobMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBlobCopier) recordInvocation(key string, args []interface{}) {
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

var _ azure.BlobCopier = new(FakeBlobCopier)
