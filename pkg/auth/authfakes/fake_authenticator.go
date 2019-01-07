// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	sync "sync"

	auth "github.com/Peripli/service-manager-cli/pkg/auth"
)

type FakeAuthenticator struct {
	ClientCredentialsStub        func() (*auth.Token, error)
	clientCredentialsMutex       sync.RWMutex
	clientCredentialsArgsForCall []struct {
	}
	clientCredentialsReturns struct {
		result1 *auth.Token
		result2 error
	}
	clientCredentialsReturnsOnCall map[int]struct {
		result1 *auth.Token
		result2 error
	}
	PasswordCredentialsStub        func(string, string) (*auth.Token, error)
	passwordCredentialsMutex       sync.RWMutex
	passwordCredentialsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	passwordCredentialsReturns struct {
		result1 *auth.Token
		result2 error
	}
	passwordCredentialsReturnsOnCall map[int]struct {
		result1 *auth.Token
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthenticator) ClientCredentials() (*auth.Token, error) {
	fake.clientCredentialsMutex.Lock()
	ret, specificReturn := fake.clientCredentialsReturnsOnCall[len(fake.clientCredentialsArgsForCall)]
	fake.clientCredentialsArgsForCall = append(fake.clientCredentialsArgsForCall, struct {
	}{})
	fake.recordInvocation("ClientCredentials", []interface{}{})
	fake.clientCredentialsMutex.Unlock()
	if fake.ClientCredentialsStub != nil {
		return fake.ClientCredentialsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.clientCredentialsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthenticator) ClientCredentialsCallCount() int {
	fake.clientCredentialsMutex.RLock()
	defer fake.clientCredentialsMutex.RUnlock()
	return len(fake.clientCredentialsArgsForCall)
}

func (fake *FakeAuthenticator) ClientCredentialsCalls(stub func() (*auth.Token, error)) {
	fake.clientCredentialsMutex.Lock()
	defer fake.clientCredentialsMutex.Unlock()
	fake.ClientCredentialsStub = stub
}

func (fake *FakeAuthenticator) ClientCredentialsReturns(result1 *auth.Token, result2 error) {
	fake.clientCredentialsMutex.Lock()
	defer fake.clientCredentialsMutex.Unlock()
	fake.ClientCredentialsStub = nil
	fake.clientCredentialsReturns = struct {
		result1 *auth.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticator) ClientCredentialsReturnsOnCall(i int, result1 *auth.Token, result2 error) {
	fake.clientCredentialsMutex.Lock()
	defer fake.clientCredentialsMutex.Unlock()
	fake.ClientCredentialsStub = nil
	if fake.clientCredentialsReturnsOnCall == nil {
		fake.clientCredentialsReturnsOnCall = make(map[int]struct {
			result1 *auth.Token
			result2 error
		})
	}
	fake.clientCredentialsReturnsOnCall[i] = struct {
		result1 *auth.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticator) PasswordCredentials(arg1 string, arg2 string) (*auth.Token, error) {
	fake.passwordCredentialsMutex.Lock()
	ret, specificReturn := fake.passwordCredentialsReturnsOnCall[len(fake.passwordCredentialsArgsForCall)]
	fake.passwordCredentialsArgsForCall = append(fake.passwordCredentialsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("PasswordCredentials", []interface{}{arg1, arg2})
	fake.passwordCredentialsMutex.Unlock()
	if fake.PasswordCredentialsStub != nil {
		return fake.PasswordCredentialsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.passwordCredentialsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAuthenticator) PasswordCredentialsCallCount() int {
	fake.passwordCredentialsMutex.RLock()
	defer fake.passwordCredentialsMutex.RUnlock()
	return len(fake.passwordCredentialsArgsForCall)
}

func (fake *FakeAuthenticator) PasswordCredentialsCalls(stub func(string, string) (*auth.Token, error)) {
	fake.passwordCredentialsMutex.Lock()
	defer fake.passwordCredentialsMutex.Unlock()
	fake.PasswordCredentialsStub = stub
}

func (fake *FakeAuthenticator) PasswordCredentialsArgsForCall(i int) (string, string) {
	fake.passwordCredentialsMutex.RLock()
	defer fake.passwordCredentialsMutex.RUnlock()
	argsForCall := fake.passwordCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthenticator) PasswordCredentialsReturns(result1 *auth.Token, result2 error) {
	fake.passwordCredentialsMutex.Lock()
	defer fake.passwordCredentialsMutex.Unlock()
	fake.PasswordCredentialsStub = nil
	fake.passwordCredentialsReturns = struct {
		result1 *auth.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticator) PasswordCredentialsReturnsOnCall(i int, result1 *auth.Token, result2 error) {
	fake.passwordCredentialsMutex.Lock()
	defer fake.passwordCredentialsMutex.Unlock()
	fake.PasswordCredentialsStub = nil
	if fake.passwordCredentialsReturnsOnCall == nil {
		fake.passwordCredentialsReturnsOnCall = make(map[int]struct {
			result1 *auth.Token
			result2 error
		})
	}
	fake.passwordCredentialsReturnsOnCall[i] = struct {
		result1 *auth.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clientCredentialsMutex.RLock()
	defer fake.clientCredentialsMutex.RUnlock()
	fake.passwordCredentialsMutex.RLock()
	defer fake.passwordCredentialsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthenticator) recordInvocation(key string, args []interface{}) {
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

var _ auth.Authenticator = new(FakeAuthenticator)