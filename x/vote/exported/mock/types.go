// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/x/vote/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	"sync"
)

// Ensure, that PollMock does implement exported.Poll.
// If this is not the case, regenerate this file with moq.
var _ exported.Poll = &PollMock{}

// PollMock is a mock implementation of exported.Poll.
//
// 	func TestSomethingThatUsesPoll(t *testing.T) {
//
// 		// make and configure a mocked exported.Poll
// 		mockedPoll := &PollMock{
// 			AllowOverrideFunc: func()  {
// 				panic("mock out the AllowOverride method")
// 			},
// 			DeleteFunc: func() error {
// 				panic("mock out the Delete method")
// 			},
// 			GetKeyFunc: func() exported.PollKey {
// 				panic("mock out the GetKey method")
// 			},
// 			GetResultFunc: func() codec.ProtoMarshaler {
// 				panic("mock out the GetResult method")
// 			},
// 			GetTotalVotingPowerFunc: func() github_com_cosmos_cosmos_sdk_types.Int {
// 				panic("mock out the GetTotalVotingPower method")
// 			},
// 			GetVotersFunc: func() []exported.Voter {
// 				panic("mock out the GetVoters method")
// 			},
// 			HasVotedFunc: func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool {
// 				panic("mock out the HasVoted method")
// 			},
// 			HasVotedCorrectlyFunc: func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool {
// 				panic("mock out the HasVotedCorrectly method")
// 			},
// 			IsFunc: func(state exported.PollState) bool {
// 				panic("mock out the Is method")
// 			},
// 			VoteFunc: func(voter github_com_cosmos_cosmos_sdk_types.ValAddress, data codec.ProtoMarshaler) error {
// 				panic("mock out the Vote method")
// 			},
// 		}
//
// 		// use mockedPoll in code that requires exported.Poll
// 		// and then make assertions.
//
// 	}
type PollMock struct {
	// AllowOverrideFunc mocks the AllowOverride method.
	AllowOverrideFunc func()

	// DeleteFunc mocks the Delete method.
	DeleteFunc func() error

	// GetKeyFunc mocks the GetKey method.
	GetKeyFunc func() exported.PollKey

	// GetResultFunc mocks the GetResult method.
	GetResultFunc func() codec.ProtoMarshaler

	// GetTotalVotingPowerFunc mocks the GetTotalVotingPower method.
	GetTotalVotingPowerFunc func() github_com_cosmos_cosmos_sdk_types.Int

	// GetVotersFunc mocks the GetVoters method.
	GetVotersFunc func() []exported.Voter

	// HasVotedFunc mocks the HasVoted method.
	HasVotedFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool

	// HasVotedCorrectlyFunc mocks the HasVotedCorrectly method.
	HasVotedCorrectlyFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool

	// IsFunc mocks the Is method.
	IsFunc func(state exported.PollState) bool

	// VoteFunc mocks the Vote method.
	VoteFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress, data codec.ProtoMarshaler) error

	// calls tracks calls to the methods.
	calls struct {
		// AllowOverride holds details about calls to the AllowOverride method.
		AllowOverride []struct {
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
		}
		// GetKey holds details about calls to the GetKey method.
		GetKey []struct {
		}
		// GetResult holds details about calls to the GetResult method.
		GetResult []struct {
		}
		// GetTotalVotingPower holds details about calls to the GetTotalVotingPower method.
		GetTotalVotingPower []struct {
		}
		// GetVoters holds details about calls to the GetVoters method.
		GetVoters []struct {
		}
		// HasVoted holds details about calls to the HasVoted method.
		HasVoted []struct {
			// Voter is the voter argument value.
			Voter github_com_cosmos_cosmos_sdk_types.ValAddress
		}
		// HasVotedCorrectly holds details about calls to the HasVotedCorrectly method.
		HasVotedCorrectly []struct {
			// Voter is the voter argument value.
			Voter github_com_cosmos_cosmos_sdk_types.ValAddress
		}
		// Is holds details about calls to the Is method.
		Is []struct {
			// State is the state argument value.
			State exported.PollState
		}
		// Vote holds details about calls to the Vote method.
		Vote []struct {
			// Voter is the voter argument value.
			Voter github_com_cosmos_cosmos_sdk_types.ValAddress
			// Data is the data argument value.
			Data codec.ProtoMarshaler
		}
	}
	lockAllowOverride       sync.RWMutex
	lockDelete              sync.RWMutex
	lockGetKey              sync.RWMutex
	lockGetResult           sync.RWMutex
	lockGetTotalVotingPower sync.RWMutex
	lockGetVoters           sync.RWMutex
	lockHasVoted            sync.RWMutex
	lockHasVotedCorrectly   sync.RWMutex
	lockIs                  sync.RWMutex
	lockVote                sync.RWMutex
}

// AllowOverride calls AllowOverrideFunc.
func (mock *PollMock) AllowOverride() {
	if mock.AllowOverrideFunc == nil {
		panic("PollMock.AllowOverrideFunc: method is nil but Poll.AllowOverride was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAllowOverride.Lock()
	mock.calls.AllowOverride = append(mock.calls.AllowOverride, callInfo)
	mock.lockAllowOverride.Unlock()
	mock.AllowOverrideFunc()
}

// AllowOverrideCalls gets all the calls that were made to AllowOverride.
// Check the length with:
//     len(mockedPoll.AllowOverrideCalls())
func (mock *PollMock) AllowOverrideCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAllowOverride.RLock()
	calls = mock.calls.AllowOverride
	mock.lockAllowOverride.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *PollMock) Delete() error {
	if mock.DeleteFunc == nil {
		panic("PollMock.DeleteFunc: method is nil but Poll.Delete was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc()
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedPoll.DeleteCalls())
func (mock *PollMock) DeleteCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// GetKey calls GetKeyFunc.
func (mock *PollMock) GetKey() exported.PollKey {
	if mock.GetKeyFunc == nil {
		panic("PollMock.GetKeyFunc: method is nil but Poll.GetKey was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetKey.Lock()
	mock.calls.GetKey = append(mock.calls.GetKey, callInfo)
	mock.lockGetKey.Unlock()
	return mock.GetKeyFunc()
}

// GetKeyCalls gets all the calls that were made to GetKey.
// Check the length with:
//     len(mockedPoll.GetKeyCalls())
func (mock *PollMock) GetKeyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetKey.RLock()
	calls = mock.calls.GetKey
	mock.lockGetKey.RUnlock()
	return calls
}

// GetResult calls GetResultFunc.
func (mock *PollMock) GetResult() codec.ProtoMarshaler {
	if mock.GetResultFunc == nil {
		panic("PollMock.GetResultFunc: method is nil but Poll.GetResult was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetResult.Lock()
	mock.calls.GetResult = append(mock.calls.GetResult, callInfo)
	mock.lockGetResult.Unlock()
	return mock.GetResultFunc()
}

// GetResultCalls gets all the calls that were made to GetResult.
// Check the length with:
//     len(mockedPoll.GetResultCalls())
func (mock *PollMock) GetResultCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetResult.RLock()
	calls = mock.calls.GetResult
	mock.lockGetResult.RUnlock()
	return calls
}

// GetTotalVotingPower calls GetTotalVotingPowerFunc.
func (mock *PollMock) GetTotalVotingPower() github_com_cosmos_cosmos_sdk_types.Int {
	if mock.GetTotalVotingPowerFunc == nil {
		panic("PollMock.GetTotalVotingPowerFunc: method is nil but Poll.GetTotalVotingPower was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetTotalVotingPower.Lock()
	mock.calls.GetTotalVotingPower = append(mock.calls.GetTotalVotingPower, callInfo)
	mock.lockGetTotalVotingPower.Unlock()
	return mock.GetTotalVotingPowerFunc()
}

// GetTotalVotingPowerCalls gets all the calls that were made to GetTotalVotingPower.
// Check the length with:
//     len(mockedPoll.GetTotalVotingPowerCalls())
func (mock *PollMock) GetTotalVotingPowerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetTotalVotingPower.RLock()
	calls = mock.calls.GetTotalVotingPower
	mock.lockGetTotalVotingPower.RUnlock()
	return calls
}

// GetVoters calls GetVotersFunc.
func (mock *PollMock) GetVoters() []exported.Voter {
	if mock.GetVotersFunc == nil {
		panic("PollMock.GetVotersFunc: method is nil but Poll.GetVoters was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetVoters.Lock()
	mock.calls.GetVoters = append(mock.calls.GetVoters, callInfo)
	mock.lockGetVoters.Unlock()
	return mock.GetVotersFunc()
}

// GetVotersCalls gets all the calls that were made to GetVoters.
// Check the length with:
//     len(mockedPoll.GetVotersCalls())
func (mock *PollMock) GetVotersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetVoters.RLock()
	calls = mock.calls.GetVoters
	mock.lockGetVoters.RUnlock()
	return calls
}

// HasVoted calls HasVotedFunc.
func (mock *PollMock) HasVoted(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool {
	if mock.HasVotedFunc == nil {
		panic("PollMock.HasVotedFunc: method is nil but Poll.HasVoted was just called")
	}
	callInfo := struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}{
		Voter: voter,
	}
	mock.lockHasVoted.Lock()
	mock.calls.HasVoted = append(mock.calls.HasVoted, callInfo)
	mock.lockHasVoted.Unlock()
	return mock.HasVotedFunc(voter)
}

// HasVotedCalls gets all the calls that were made to HasVoted.
// Check the length with:
//     len(mockedPoll.HasVotedCalls())
func (mock *PollMock) HasVotedCalls() []struct {
	Voter github_com_cosmos_cosmos_sdk_types.ValAddress
} {
	var calls []struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}
	mock.lockHasVoted.RLock()
	calls = mock.calls.HasVoted
	mock.lockHasVoted.RUnlock()
	return calls
}

// HasVotedCorrectly calls HasVotedCorrectlyFunc.
func (mock *PollMock) HasVotedCorrectly(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool {
	if mock.HasVotedCorrectlyFunc == nil {
		panic("PollMock.HasVotedCorrectlyFunc: method is nil but Poll.HasVotedCorrectly was just called")
	}
	callInfo := struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}{
		Voter: voter,
	}
	mock.lockHasVotedCorrectly.Lock()
	mock.calls.HasVotedCorrectly = append(mock.calls.HasVotedCorrectly, callInfo)
	mock.lockHasVotedCorrectly.Unlock()
	return mock.HasVotedCorrectlyFunc(voter)
}

// HasVotedCorrectlyCalls gets all the calls that were made to HasVotedCorrectly.
// Check the length with:
//     len(mockedPoll.HasVotedCorrectlyCalls())
func (mock *PollMock) HasVotedCorrectlyCalls() []struct {
	Voter github_com_cosmos_cosmos_sdk_types.ValAddress
} {
	var calls []struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}
	mock.lockHasVotedCorrectly.RLock()
	calls = mock.calls.HasVotedCorrectly
	mock.lockHasVotedCorrectly.RUnlock()
	return calls
}

// Is calls IsFunc.
func (mock *PollMock) Is(state exported.PollState) bool {
	if mock.IsFunc == nil {
		panic("PollMock.IsFunc: method is nil but Poll.Is was just called")
	}
	callInfo := struct {
		State exported.PollState
	}{
		State: state,
	}
	mock.lockIs.Lock()
	mock.calls.Is = append(mock.calls.Is, callInfo)
	mock.lockIs.Unlock()
	return mock.IsFunc(state)
}

// IsCalls gets all the calls that were made to Is.
// Check the length with:
//     len(mockedPoll.IsCalls())
func (mock *PollMock) IsCalls() []struct {
	State exported.PollState
} {
	var calls []struct {
		State exported.PollState
	}
	mock.lockIs.RLock()
	calls = mock.calls.Is
	mock.lockIs.RUnlock()
	return calls
}

// Vote calls VoteFunc.
func (mock *PollMock) Vote(voter github_com_cosmos_cosmos_sdk_types.ValAddress, data codec.ProtoMarshaler) error {
	if mock.VoteFunc == nil {
		panic("PollMock.VoteFunc: method is nil but Poll.Vote was just called")
	}
	callInfo := struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
		Data  codec.ProtoMarshaler
	}{
		Voter: voter,
		Data:  data,
	}
	mock.lockVote.Lock()
	mock.calls.Vote = append(mock.calls.Vote, callInfo)
	mock.lockVote.Unlock()
	return mock.VoteFunc(voter, data)
}

// VoteCalls gets all the calls that were made to Vote.
// Check the length with:
//     len(mockedPoll.VoteCalls())
func (mock *PollMock) VoteCalls() []struct {
	Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	Data  codec.ProtoMarshaler
} {
	var calls []struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
		Data  codec.ProtoMarshaler
	}
	mock.lockVote.RLock()
	calls = mock.calls.Vote
	mock.lockVote.RUnlock()
	return calls
}
