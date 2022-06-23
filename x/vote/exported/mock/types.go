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
// 			DeleteFunc: func()  {
// 				panic("mock out the Delete method")
// 			},
// 			GetIDFunc: func() exported.PollID {
// 				panic("mock out the GetID method")
// 			},
// 			GetModuleMetadataFunc: func() exported.PollModuleMetadata {
// 				panic("mock out the GetModuleMetadata method")
// 			},
// 			GetResultFunc: func() codec.ProtoMarshaler {
// 				panic("mock out the GetResult method")
// 			},
// 			GetRewardPoolNameFunc: func() (string, bool) {
// 				panic("mock out the GetRewardPoolName method")
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
// 			HasVotedLateFunc: func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool {
// 				panic("mock out the HasVotedLate method")
// 			},
// 			IsFunc: func(state exported.PollState) bool {
// 				panic("mock out the Is method")
// 			},
// 			VoteFunc: func(voter github_com_cosmos_cosmos_sdk_types.ValAddress, blockHeight int64, data codec.ProtoMarshaler) (codec.ProtoMarshaler, bool, error) {
// 				panic("mock out the Vote method")
// 			},
// 		}
//
// 		// use mockedPoll in code that requires exported.Poll
// 		// and then make assertions.
//
// 	}
type PollMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func()

	// GetIDFunc mocks the GetID method.
	GetIDFunc func() exported.PollID

	// GetModuleMetadataFunc mocks the GetModuleMetadata method.
	GetModuleMetadataFunc func() exported.PollModuleMetadata

	// GetResultFunc mocks the GetResult method.
	GetResultFunc func() codec.ProtoMarshaler

	// GetRewardPoolNameFunc mocks the GetRewardPoolName method.
	GetRewardPoolNameFunc func() (string, bool)

	// GetTotalVotingPowerFunc mocks the GetTotalVotingPower method.
	GetTotalVotingPowerFunc func() github_com_cosmos_cosmos_sdk_types.Int

	// GetVotersFunc mocks the GetVoters method.
	GetVotersFunc func() []exported.Voter

	// HasVotedFunc mocks the HasVoted method.
	HasVotedFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool

	// HasVotedCorrectlyFunc mocks the HasVotedCorrectly method.
	HasVotedCorrectlyFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool

	// HasVotedLateFunc mocks the HasVotedLate method.
	HasVotedLateFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool

	// IsFunc mocks the Is method.
	IsFunc func(state exported.PollState) bool

	// VoteFunc mocks the Vote method.
	VoteFunc func(voter github_com_cosmos_cosmos_sdk_types.ValAddress, blockHeight int64, data codec.ProtoMarshaler) (codec.ProtoMarshaler, bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
		}
		// GetID holds details about calls to the GetID method.
		GetID []struct {
		}
		// GetModuleMetadata holds details about calls to the GetModuleMetadata method.
		GetModuleMetadata []struct {
		}
		// GetResult holds details about calls to the GetResult method.
		GetResult []struct {
		}
		// GetRewardPoolName holds details about calls to the GetRewardPoolName method.
		GetRewardPoolName []struct {
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
		// HasVotedLate holds details about calls to the HasVotedLate method.
		HasVotedLate []struct {
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
			// BlockHeight is the blockHeight argument value.
			BlockHeight int64
			// Data is the data argument value.
			Data codec.ProtoMarshaler
		}
	}
	lockDelete              sync.RWMutex
	lockGetID               sync.RWMutex
	lockGetModuleMetadata   sync.RWMutex
	lockGetResult           sync.RWMutex
	lockGetRewardPoolName   sync.RWMutex
	lockGetTotalVotingPower sync.RWMutex
	lockGetVoters           sync.RWMutex
	lockHasVoted            sync.RWMutex
	lockHasVotedCorrectly   sync.RWMutex
	lockHasVotedLate        sync.RWMutex
	lockIs                  sync.RWMutex
	lockVote                sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *PollMock) Delete() {
	if mock.DeleteFunc == nil {
		panic("PollMock.DeleteFunc: method is nil but Poll.Delete was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	mock.DeleteFunc()
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

// GetID calls GetIDFunc.
func (mock *PollMock) GetID() exported.PollID {
	if mock.GetIDFunc == nil {
		panic("PollMock.GetIDFunc: method is nil but Poll.GetID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetID.Lock()
	mock.calls.GetID = append(mock.calls.GetID, callInfo)
	mock.lockGetID.Unlock()
	return mock.GetIDFunc()
}

// GetIDCalls gets all the calls that were made to GetID.
// Check the length with:
//     len(mockedPoll.GetIDCalls())
func (mock *PollMock) GetIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetID.RLock()
	calls = mock.calls.GetID
	mock.lockGetID.RUnlock()
	return calls
}

// GetModuleMetadata calls GetModuleMetadataFunc.
func (mock *PollMock) GetModuleMetadata() exported.PollModuleMetadata {
	if mock.GetModuleMetadataFunc == nil {
		panic("PollMock.GetModuleMetadataFunc: method is nil but Poll.GetModuleMetadata was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetModuleMetadata.Lock()
	mock.calls.GetModuleMetadata = append(mock.calls.GetModuleMetadata, callInfo)
	mock.lockGetModuleMetadata.Unlock()
	return mock.GetModuleMetadataFunc()
}

// GetModuleMetadataCalls gets all the calls that were made to GetModuleMetadata.
// Check the length with:
//     len(mockedPoll.GetModuleMetadataCalls())
func (mock *PollMock) GetModuleMetadataCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetModuleMetadata.RLock()
	calls = mock.calls.GetModuleMetadata
	mock.lockGetModuleMetadata.RUnlock()
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

// GetRewardPoolName calls GetRewardPoolNameFunc.
func (mock *PollMock) GetRewardPoolName() (string, bool) {
	if mock.GetRewardPoolNameFunc == nil {
		panic("PollMock.GetRewardPoolNameFunc: method is nil but Poll.GetRewardPoolName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetRewardPoolName.Lock()
	mock.calls.GetRewardPoolName = append(mock.calls.GetRewardPoolName, callInfo)
	mock.lockGetRewardPoolName.Unlock()
	return mock.GetRewardPoolNameFunc()
}

// GetRewardPoolNameCalls gets all the calls that were made to GetRewardPoolName.
// Check the length with:
//     len(mockedPoll.GetRewardPoolNameCalls())
func (mock *PollMock) GetRewardPoolNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetRewardPoolName.RLock()
	calls = mock.calls.GetRewardPoolName
	mock.lockGetRewardPoolName.RUnlock()
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

// HasVotedLate calls HasVotedLateFunc.
func (mock *PollMock) HasVotedLate(voter github_com_cosmos_cosmos_sdk_types.ValAddress) bool {
	if mock.HasVotedLateFunc == nil {
		panic("PollMock.HasVotedLateFunc: method is nil but Poll.HasVotedLate was just called")
	}
	callInfo := struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}{
		Voter: voter,
	}
	mock.lockHasVotedLate.Lock()
	mock.calls.HasVotedLate = append(mock.calls.HasVotedLate, callInfo)
	mock.lockHasVotedLate.Unlock()
	return mock.HasVotedLateFunc(voter)
}

// HasVotedLateCalls gets all the calls that were made to HasVotedLate.
// Check the length with:
//     len(mockedPoll.HasVotedLateCalls())
func (mock *PollMock) HasVotedLateCalls() []struct {
	Voter github_com_cosmos_cosmos_sdk_types.ValAddress
} {
	var calls []struct {
		Voter github_com_cosmos_cosmos_sdk_types.ValAddress
	}
	mock.lockHasVotedLate.RLock()
	calls = mock.calls.HasVotedLate
	mock.lockHasVotedLate.RUnlock()
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
func (mock *PollMock) Vote(voter github_com_cosmos_cosmos_sdk_types.ValAddress, blockHeight int64, data codec.ProtoMarshaler) (codec.ProtoMarshaler, bool, error) {
	if mock.VoteFunc == nil {
		panic("PollMock.VoteFunc: method is nil but Poll.Vote was just called")
	}
	callInfo := struct {
		Voter       github_com_cosmos_cosmos_sdk_types.ValAddress
		BlockHeight int64
		Data        codec.ProtoMarshaler
	}{
		Voter:       voter,
		BlockHeight: blockHeight,
		Data:        data,
	}
	mock.lockVote.Lock()
	mock.calls.Vote = append(mock.calls.Vote, callInfo)
	mock.lockVote.Unlock()
	return mock.VoteFunc(voter, blockHeight, data)
}

// VoteCalls gets all the calls that were made to Vote.
// Check the length with:
//     len(mockedPoll.VoteCalls())
func (mock *PollMock) VoteCalls() []struct {
	Voter       github_com_cosmos_cosmos_sdk_types.ValAddress
	BlockHeight int64
	Data        codec.ProtoMarshaler
} {
	var calls []struct {
		Voter       github_com_cosmos_cosmos_sdk_types.ValAddress
		BlockHeight int64
		Data        codec.ProtoMarshaler
	}
	mock.lockVote.RLock()
	calls = mock.calls.Vote
	mock.lockVote.RUnlock()
	return calls
}

// Ensure, that VoteHandlerMock does implement exported.VoteHandler.
// If this is not the case, regenerate this file with moq.
var _ exported.VoteHandler = &VoteHandlerMock{}

// VoteHandlerMock is a mock implementation of exported.VoteHandler.
//
// 	func TestSomethingThatUsesVoteHandler(t *testing.T) {
//
// 		// make and configure a mocked exported.VoteHandler
// 		mockedVoteHandler := &VoteHandlerMock{
// 			HandleCompletedPollFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported.Poll) error {
// 				panic("mock out the HandleCompletedPoll method")
// 			},
// 			HandleExpiredPollFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported.Poll) error {
// 				panic("mock out the HandleExpiredPoll method")
// 			},
// 			HandleResultFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, result codec.ProtoMarshaler) error {
// 				panic("mock out the HandleResult method")
// 			},
// 			IsFalsyResultFunc: func(result codec.ProtoMarshaler) bool {
// 				panic("mock out the IsFalsyResult method")
// 			},
// 		}
//
// 		// use mockedVoteHandler in code that requires exported.VoteHandler
// 		// and then make assertions.
//
// 	}
type VoteHandlerMock struct {
	// HandleCompletedPollFunc mocks the HandleCompletedPoll method.
	HandleCompletedPollFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported.Poll) error

	// HandleExpiredPollFunc mocks the HandleExpiredPoll method.
	HandleExpiredPollFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported.Poll) error

	// HandleResultFunc mocks the HandleResult method.
	HandleResultFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, result codec.ProtoMarshaler) error

	// IsFalsyResultFunc mocks the IsFalsyResult method.
	IsFalsyResultFunc func(result codec.ProtoMarshaler) bool

	// calls tracks calls to the methods.
	calls struct {
		// HandleCompletedPoll holds details about calls to the HandleCompletedPoll method.
		HandleCompletedPoll []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Poll is the poll argument value.
			Poll exported.Poll
		}
		// HandleExpiredPoll holds details about calls to the HandleExpiredPoll method.
		HandleExpiredPoll []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Poll is the poll argument value.
			Poll exported.Poll
		}
		// HandleResult holds details about calls to the HandleResult method.
		HandleResult []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Result is the result argument value.
			Result codec.ProtoMarshaler
		}
		// IsFalsyResult holds details about calls to the IsFalsyResult method.
		IsFalsyResult []struct {
			// Result is the result argument value.
			Result codec.ProtoMarshaler
		}
	}
	lockHandleCompletedPoll sync.RWMutex
	lockHandleExpiredPoll   sync.RWMutex
	lockHandleResult        sync.RWMutex
	lockIsFalsyResult       sync.RWMutex
}

// HandleCompletedPoll calls HandleCompletedPollFunc.
func (mock *VoteHandlerMock) HandleCompletedPoll(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported.Poll) error {
	if mock.HandleCompletedPollFunc == nil {
		panic("VoteHandlerMock.HandleCompletedPollFunc: method is nil but VoteHandler.HandleCompletedPoll was just called")
	}
	callInfo := struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Poll exported.Poll
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockHandleCompletedPoll.Lock()
	mock.calls.HandleCompletedPoll = append(mock.calls.HandleCompletedPoll, callInfo)
	mock.lockHandleCompletedPoll.Unlock()
	return mock.HandleCompletedPollFunc(ctx, poll)
}

// HandleCompletedPollCalls gets all the calls that were made to HandleCompletedPoll.
// Check the length with:
//     len(mockedVoteHandler.HandleCompletedPollCalls())
func (mock *VoteHandlerMock) HandleCompletedPollCalls() []struct {
	Ctx  github_com_cosmos_cosmos_sdk_types.Context
	Poll exported.Poll
} {
	var calls []struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Poll exported.Poll
	}
	mock.lockHandleCompletedPoll.RLock()
	calls = mock.calls.HandleCompletedPoll
	mock.lockHandleCompletedPoll.RUnlock()
	return calls
}

// HandleExpiredPoll calls HandleExpiredPollFunc.
func (mock *VoteHandlerMock) HandleExpiredPoll(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported.Poll) error {
	if mock.HandleExpiredPollFunc == nil {
		panic("VoteHandlerMock.HandleExpiredPollFunc: method is nil but VoteHandler.HandleExpiredPoll was just called")
	}
	callInfo := struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Poll exported.Poll
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockHandleExpiredPoll.Lock()
	mock.calls.HandleExpiredPoll = append(mock.calls.HandleExpiredPoll, callInfo)
	mock.lockHandleExpiredPoll.Unlock()
	return mock.HandleExpiredPollFunc(ctx, poll)
}

// HandleExpiredPollCalls gets all the calls that were made to HandleExpiredPoll.
// Check the length with:
//     len(mockedVoteHandler.HandleExpiredPollCalls())
func (mock *VoteHandlerMock) HandleExpiredPollCalls() []struct {
	Ctx  github_com_cosmos_cosmos_sdk_types.Context
	Poll exported.Poll
} {
	var calls []struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Poll exported.Poll
	}
	mock.lockHandleExpiredPoll.RLock()
	calls = mock.calls.HandleExpiredPoll
	mock.lockHandleExpiredPoll.RUnlock()
	return calls
}

// HandleResult calls HandleResultFunc.
func (mock *VoteHandlerMock) HandleResult(ctx github_com_cosmos_cosmos_sdk_types.Context, result codec.ProtoMarshaler) error {
	if mock.HandleResultFunc == nil {
		panic("VoteHandlerMock.HandleResultFunc: method is nil but VoteHandler.HandleResult was just called")
	}
	callInfo := struct {
		Ctx    github_com_cosmos_cosmos_sdk_types.Context
		Result codec.ProtoMarshaler
	}{
		Ctx:    ctx,
		Result: result,
	}
	mock.lockHandleResult.Lock()
	mock.calls.HandleResult = append(mock.calls.HandleResult, callInfo)
	mock.lockHandleResult.Unlock()
	return mock.HandleResultFunc(ctx, result)
}

// HandleResultCalls gets all the calls that were made to HandleResult.
// Check the length with:
//     len(mockedVoteHandler.HandleResultCalls())
func (mock *VoteHandlerMock) HandleResultCalls() []struct {
	Ctx    github_com_cosmos_cosmos_sdk_types.Context
	Result codec.ProtoMarshaler
} {
	var calls []struct {
		Ctx    github_com_cosmos_cosmos_sdk_types.Context
		Result codec.ProtoMarshaler
	}
	mock.lockHandleResult.RLock()
	calls = mock.calls.HandleResult
	mock.lockHandleResult.RUnlock()
	return calls
}

// IsFalsyResult calls IsFalsyResultFunc.
func (mock *VoteHandlerMock) IsFalsyResult(result codec.ProtoMarshaler) bool {
	if mock.IsFalsyResultFunc == nil {
		panic("VoteHandlerMock.IsFalsyResultFunc: method is nil but VoteHandler.IsFalsyResult was just called")
	}
	callInfo := struct {
		Result codec.ProtoMarshaler
	}{
		Result: result,
	}
	mock.lockIsFalsyResult.Lock()
	mock.calls.IsFalsyResult = append(mock.calls.IsFalsyResult, callInfo)
	mock.lockIsFalsyResult.Unlock()
	return mock.IsFalsyResultFunc(result)
}

// IsFalsyResultCalls gets all the calls that were made to IsFalsyResult.
// Check the length with:
//     len(mockedVoteHandler.IsFalsyResultCalls())
func (mock *VoteHandlerMock) IsFalsyResultCalls() []struct {
	Result codec.ProtoMarshaler
} {
	var calls []struct {
		Result codec.ProtoMarshaler
	}
	mock.lockIsFalsyResult.RLock()
	calls = mock.calls.IsFalsyResult
	mock.lockIsFalsyResult.RUnlock()
	return calls
}
