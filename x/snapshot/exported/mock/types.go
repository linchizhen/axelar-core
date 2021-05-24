// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
)

// Ensure, that SDKValidatorMock does implement exported.SDKValidator.
// If this is not the case, regenerate this file with moq.
var _ exported.SDKValidator = &SDKValidatorMock{}

// SDKValidatorMock is a mock implementation of exported.SDKValidator.
//
// 	func TestSomethingThatUsesSDKValidator(t *testing.T) {
//
// 		// make and configure a mocked exported.SDKValidator
// 		mockedSDKValidator := &SDKValidatorMock{
// 			GetConsAddrFunc: func() (sdk.ConsAddress, error) {
// 				panic("mock out the GetConsAddr method")
// 			},
// 			GetConsensusPowerFunc: func() int64 {
// 				panic("mock out the GetConsensusPower method")
// 			},
// 			GetOperatorFunc: func() sdk.ValAddress {
// 				panic("mock out the GetOperator method")
// 			},
// 			IsJailedFunc: func() bool {
// 				panic("mock out the IsJailed method")
// 			},
// 			UnpackInterfacesFunc: func(unpacker codectypes.AnyUnpacker) error {
// 				panic("mock out the UnpackInterfaces method")
// 			},
// 		}
//
// 		// use mockedSDKValidator in code that requires exported.SDKValidator
// 		// and then make assertions.
//
// 	}
type SDKValidatorMock struct {
	// GetConsAddrFunc mocks the GetConsAddr method.
	GetConsAddrFunc func() (sdk.ConsAddress, error)

	// GetConsensusPowerFunc mocks the GetConsensusPower method.
	GetConsensusPowerFunc func() int64

	// GetOperatorFunc mocks the GetOperator method.
	GetOperatorFunc func() sdk.ValAddress

	// IsJailedFunc mocks the IsJailed method.
	IsJailedFunc func() bool

	// UnpackInterfacesFunc mocks the UnpackInterfaces method.
	UnpackInterfacesFunc func(unpacker codectypes.AnyUnpacker) error

	// calls tracks calls to the methods.
	calls struct {
		// GetConsAddr holds details about calls to the GetConsAddr method.
		GetConsAddr []struct {
		}
		// GetConsensusPower holds details about calls to the GetConsensusPower method.
		GetConsensusPower []struct {
		}
		// GetOperator holds details about calls to the GetOperator method.
		GetOperator []struct {
		}
		// IsJailed holds details about calls to the IsJailed method.
		IsJailed []struct {
		}
		// UnpackInterfaces holds details about calls to the UnpackInterfaces method.
		UnpackInterfaces []struct {
			// Unpacker is the unpacker argument value.
			Unpacker codectypes.AnyUnpacker
		}
	}
	lockGetConsAddr       sync.RWMutex
	lockGetConsensusPower sync.RWMutex
	lockGetOperator       sync.RWMutex
	lockIsJailed          sync.RWMutex
	lockUnpackInterfaces  sync.RWMutex
}

// GetConsAddr calls GetConsAddrFunc.
func (mock *SDKValidatorMock) GetConsAddr() (sdk.ConsAddress, error) {
	if mock.GetConsAddrFunc == nil {
		panic("SDKValidatorMock.GetConsAddrFunc: method is nil but SDKValidator.GetConsAddr was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetConsAddr.Lock()
	mock.calls.GetConsAddr = append(mock.calls.GetConsAddr, callInfo)
	mock.lockGetConsAddr.Unlock()
	return mock.GetConsAddrFunc()
}

// GetConsAddrCalls gets all the calls that were made to GetConsAddr.
// Check the length with:
//     len(mockedSDKValidator.GetConsAddrCalls())
func (mock *SDKValidatorMock) GetConsAddrCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetConsAddr.RLock()
	calls = mock.calls.GetConsAddr
	mock.lockGetConsAddr.RUnlock()
	return calls
}

// GetConsensusPower calls GetConsensusPowerFunc.
func (mock *SDKValidatorMock) GetConsensusPower() int64 {
	if mock.GetConsensusPowerFunc == nil {
		panic("SDKValidatorMock.GetConsensusPowerFunc: method is nil but SDKValidator.GetConsensusPower was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetConsensusPower.Lock()
	mock.calls.GetConsensusPower = append(mock.calls.GetConsensusPower, callInfo)
	mock.lockGetConsensusPower.Unlock()
	return mock.GetConsensusPowerFunc()
}

// GetConsensusPowerCalls gets all the calls that were made to GetConsensusPower.
// Check the length with:
//     len(mockedSDKValidator.GetConsensusPowerCalls())
func (mock *SDKValidatorMock) GetConsensusPowerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetConsensusPower.RLock()
	calls = mock.calls.GetConsensusPower
	mock.lockGetConsensusPower.RUnlock()
	return calls
}

// GetOperator calls GetOperatorFunc.
func (mock *SDKValidatorMock) GetOperator() sdk.ValAddress {
	if mock.GetOperatorFunc == nil {
		panic("SDKValidatorMock.GetOperatorFunc: method is nil but SDKValidator.GetOperator was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetOperator.Lock()
	mock.calls.GetOperator = append(mock.calls.GetOperator, callInfo)
	mock.lockGetOperator.Unlock()
	return mock.GetOperatorFunc()
}

// GetOperatorCalls gets all the calls that were made to GetOperator.
// Check the length with:
//     len(mockedSDKValidator.GetOperatorCalls())
func (mock *SDKValidatorMock) GetOperatorCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetOperator.RLock()
	calls = mock.calls.GetOperator
	mock.lockGetOperator.RUnlock()
	return calls
}

// IsJailed calls IsJailedFunc.
func (mock *SDKValidatorMock) IsJailed() bool {
	if mock.IsJailedFunc == nil {
		panic("SDKValidatorMock.IsJailedFunc: method is nil but SDKValidator.IsJailed was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsJailed.Lock()
	mock.calls.IsJailed = append(mock.calls.IsJailed, callInfo)
	mock.lockIsJailed.Unlock()
	return mock.IsJailedFunc()
}

// IsJailedCalls gets all the calls that were made to IsJailed.
// Check the length with:
//     len(mockedSDKValidator.IsJailedCalls())
func (mock *SDKValidatorMock) IsJailedCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsJailed.RLock()
	calls = mock.calls.IsJailed
	mock.lockIsJailed.RUnlock()
	return calls
}

// UnpackInterfaces calls UnpackInterfacesFunc.
func (mock *SDKValidatorMock) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if mock.UnpackInterfacesFunc == nil {
		panic("SDKValidatorMock.UnpackInterfacesFunc: method is nil but SDKValidator.UnpackInterfaces was just called")
	}
	callInfo := struct {
		Unpacker codectypes.AnyUnpacker
	}{
		Unpacker: unpacker,
	}
	mock.lockUnpackInterfaces.Lock()
	mock.calls.UnpackInterfaces = append(mock.calls.UnpackInterfaces, callInfo)
	mock.lockUnpackInterfaces.Unlock()
	return mock.UnpackInterfacesFunc(unpacker)
}

// UnpackInterfacesCalls gets all the calls that were made to UnpackInterfaces.
// Check the length with:
//     len(mockedSDKValidator.UnpackInterfacesCalls())
func (mock *SDKValidatorMock) UnpackInterfacesCalls() []struct {
	Unpacker codectypes.AnyUnpacker
} {
	var calls []struct {
		Unpacker codectypes.AnyUnpacker
	}
	mock.lockUnpackInterfaces.RLock()
	calls = mock.calls.UnpackInterfaces
	mock.lockUnpackInterfaces.RUnlock()
	return calls
}

// Ensure, that SnapshotterMock does implement exported.Snapshotter.
// If this is not the case, regenerate this file with moq.
var _ exported.Snapshotter = &SnapshotterMock{}

// SnapshotterMock is a mock implementation of exported.Snapshotter.
//
// 	func TestSomethingThatUsesSnapshotter(t *testing.T) {
//
// 		// make and configure a mocked exported.Snapshotter
// 		mockedSnapshotter := &SnapshotterMock{
// 			GetLatestCounterFunc: func(ctx sdk.Context) int64 {
// 				panic("mock out the GetLatestCounter method")
// 			},
// 			GetLatestSnapshotFunc: func(ctx sdk.Context) (exported.Snapshot, bool) {
// 				panic("mock out the GetLatestSnapshot method")
// 			},
// 			GetSnapshotFunc: func(ctx sdk.Context, counter int64) (exported.Snapshot, bool) {
// 				panic("mock out the GetSnapshot method")
// 			},
// 			TakeSnapshotFunc: func(ctx sdk.Context, subsetSize int64, keyShareDistributionPolicy tss.KeyShareDistributionPolicy) (sdk.Int, sdk.Int, error) {
// 				panic("mock out the TakeSnapshot method")
// 			},
// 		}
//
// 		// use mockedSnapshotter in code that requires exported.Snapshotter
// 		// and then make assertions.
//
// 	}
type SnapshotterMock struct {
	// GetLatestCounterFunc mocks the GetLatestCounter method.
	GetLatestCounterFunc func(ctx sdk.Context) int64

	// GetLatestSnapshotFunc mocks the GetLatestSnapshot method.
	GetLatestSnapshotFunc func(ctx sdk.Context) (exported.Snapshot, bool)

	// GetSnapshotFunc mocks the GetSnapshot method.
	GetSnapshotFunc func(ctx sdk.Context, counter int64) (exported.Snapshot, bool)

	// TakeSnapshotFunc mocks the TakeSnapshot method.
	TakeSnapshotFunc func(ctx sdk.Context, subsetSize int64, keyShareDistributionPolicy tss.KeyShareDistributionPolicy) (sdk.Int, sdk.Int, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetLatestCounter holds details about calls to the GetLatestCounter method.
		GetLatestCounter []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
		}
		// GetLatestSnapshot holds details about calls to the GetLatestSnapshot method.
		GetLatestSnapshot []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
		}
		// GetSnapshot holds details about calls to the GetSnapshot method.
		GetSnapshot []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Counter is the counter argument value.
			Counter int64
		}
		// TakeSnapshot holds details about calls to the TakeSnapshot method.
		TakeSnapshot []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// SubsetSize is the subsetSize argument value.
			SubsetSize int64
			// KeyShareDistributionPolicy is the keyShareDistributionPolicy argument value.
			KeyShareDistributionPolicy tss.KeyShareDistributionPolicy
		}
	}
	lockGetLatestCounter  sync.RWMutex
	lockGetLatestSnapshot sync.RWMutex
	lockGetSnapshot       sync.RWMutex
	lockTakeSnapshot      sync.RWMutex
}

// GetLatestCounter calls GetLatestCounterFunc.
func (mock *SnapshotterMock) GetLatestCounter(ctx sdk.Context) int64 {
	if mock.GetLatestCounterFunc == nil {
		panic("SnapshotterMock.GetLatestCounterFunc: method is nil but Snapshotter.GetLatestCounter was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetLatestCounter.Lock()
	mock.calls.GetLatestCounter = append(mock.calls.GetLatestCounter, callInfo)
	mock.lockGetLatestCounter.Unlock()
	return mock.GetLatestCounterFunc(ctx)
}

// GetLatestCounterCalls gets all the calls that were made to GetLatestCounter.
// Check the length with:
//     len(mockedSnapshotter.GetLatestCounterCalls())
func (mock *SnapshotterMock) GetLatestCounterCalls() []struct {
	Ctx sdk.Context
} {
	var calls []struct {
		Ctx sdk.Context
	}
	mock.lockGetLatestCounter.RLock()
	calls = mock.calls.GetLatestCounter
	mock.lockGetLatestCounter.RUnlock()
	return calls
}

// GetLatestSnapshot calls GetLatestSnapshotFunc.
func (mock *SnapshotterMock) GetLatestSnapshot(ctx sdk.Context) (exported.Snapshot, bool) {
	if mock.GetLatestSnapshotFunc == nil {
		panic("SnapshotterMock.GetLatestSnapshotFunc: method is nil but Snapshotter.GetLatestSnapshot was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetLatestSnapshot.Lock()
	mock.calls.GetLatestSnapshot = append(mock.calls.GetLatestSnapshot, callInfo)
	mock.lockGetLatestSnapshot.Unlock()
	return mock.GetLatestSnapshotFunc(ctx)
}

// GetLatestSnapshotCalls gets all the calls that were made to GetLatestSnapshot.
// Check the length with:
//     len(mockedSnapshotter.GetLatestSnapshotCalls())
func (mock *SnapshotterMock) GetLatestSnapshotCalls() []struct {
	Ctx sdk.Context
} {
	var calls []struct {
		Ctx sdk.Context
	}
	mock.lockGetLatestSnapshot.RLock()
	calls = mock.calls.GetLatestSnapshot
	mock.lockGetLatestSnapshot.RUnlock()
	return calls
}

// GetSnapshot calls GetSnapshotFunc.
func (mock *SnapshotterMock) GetSnapshot(ctx sdk.Context, counter int64) (exported.Snapshot, bool) {
	if mock.GetSnapshotFunc == nil {
		panic("SnapshotterMock.GetSnapshotFunc: method is nil but Snapshotter.GetSnapshot was just called")
	}
	callInfo := struct {
		Ctx     sdk.Context
		Counter int64
	}{
		Ctx:     ctx,
		Counter: counter,
	}
	mock.lockGetSnapshot.Lock()
	mock.calls.GetSnapshot = append(mock.calls.GetSnapshot, callInfo)
	mock.lockGetSnapshot.Unlock()
	return mock.GetSnapshotFunc(ctx, counter)
}

// GetSnapshotCalls gets all the calls that were made to GetSnapshot.
// Check the length with:
//     len(mockedSnapshotter.GetSnapshotCalls())
func (mock *SnapshotterMock) GetSnapshotCalls() []struct {
	Ctx     sdk.Context
	Counter int64
} {
	var calls []struct {
		Ctx     sdk.Context
		Counter int64
	}
	mock.lockGetSnapshot.RLock()
	calls = mock.calls.GetSnapshot
	mock.lockGetSnapshot.RUnlock()
	return calls
}

// TakeSnapshot calls TakeSnapshotFunc.
func (mock *SnapshotterMock) TakeSnapshot(ctx sdk.Context, subsetSize int64, keyShareDistributionPolicy tss.KeyShareDistributionPolicy) (sdk.Int, sdk.Int, error) {
	if mock.TakeSnapshotFunc == nil {
		panic("SnapshotterMock.TakeSnapshotFunc: method is nil but Snapshotter.TakeSnapshot was just called")
	}
	callInfo := struct {
		Ctx                        sdk.Context
		SubsetSize                 int64
		KeyShareDistributionPolicy tss.KeyShareDistributionPolicy
	}{
		Ctx:                        ctx,
		SubsetSize:                 subsetSize,
		KeyShareDistributionPolicy: keyShareDistributionPolicy,
	}
	mock.lockTakeSnapshot.Lock()
	mock.calls.TakeSnapshot = append(mock.calls.TakeSnapshot, callInfo)
	mock.lockTakeSnapshot.Unlock()
	return mock.TakeSnapshotFunc(ctx, subsetSize, keyShareDistributionPolicy)
}

// TakeSnapshotCalls gets all the calls that were made to TakeSnapshot.
// Check the length with:
//     len(mockedSnapshotter.TakeSnapshotCalls())
func (mock *SnapshotterMock) TakeSnapshotCalls() []struct {
	Ctx                        sdk.Context
	SubsetSize                 int64
	KeyShareDistributionPolicy tss.KeyShareDistributionPolicy
} {
	var calls []struct {
		Ctx                        sdk.Context
		SubsetSize                 int64
		KeyShareDistributionPolicy tss.KeyShareDistributionPolicy
	}
	mock.lockTakeSnapshot.RLock()
	calls = mock.calls.TakeSnapshot
	mock.lockTakeSnapshot.RUnlock()
	return calls
}

// Ensure, that SlasherMock does implement exported.Slasher.
// If this is not the case, regenerate this file with moq.
var _ exported.Slasher = &SlasherMock{}

// SlasherMock is a mock implementation of exported.Slasher.
//
// 	func TestSomethingThatUsesSlasher(t *testing.T) {
//
// 		// make and configure a mocked exported.Slasher
// 		mockedSlasher := &SlasherMock{
// 			GetValidatorSigningInfoFunc: func(ctx sdk.Context, address sdk.ConsAddress) (exported.ValidatorInfo, bool) {
// 				panic("mock out the GetValidatorSigningInfo method")
// 			},
// 		}
//
// 		// use mockedSlasher in code that requires exported.Slasher
// 		// and then make assertions.
//
// 	}
type SlasherMock struct {
	// GetValidatorSigningInfoFunc mocks the GetValidatorSigningInfo method.
	GetValidatorSigningInfoFunc func(ctx sdk.Context, address sdk.ConsAddress) (exported.ValidatorInfo, bool)

	// calls tracks calls to the methods.
	calls struct {
		// GetValidatorSigningInfo holds details about calls to the GetValidatorSigningInfo method.
		GetValidatorSigningInfo []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Address is the address argument value.
			Address sdk.ConsAddress
		}
	}
	lockGetValidatorSigningInfo sync.RWMutex
}

// GetValidatorSigningInfo calls GetValidatorSigningInfoFunc.
func (mock *SlasherMock) GetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress) (exported.ValidatorInfo, bool) {
	if mock.GetValidatorSigningInfoFunc == nil {
		panic("SlasherMock.GetValidatorSigningInfoFunc: method is nil but Slasher.GetValidatorSigningInfo was just called")
	}
	callInfo := struct {
		Ctx     sdk.Context
		Address sdk.ConsAddress
	}{
		Ctx:     ctx,
		Address: address,
	}
	mock.lockGetValidatorSigningInfo.Lock()
	mock.calls.GetValidatorSigningInfo = append(mock.calls.GetValidatorSigningInfo, callInfo)
	mock.lockGetValidatorSigningInfo.Unlock()
	return mock.GetValidatorSigningInfoFunc(ctx, address)
}

// GetValidatorSigningInfoCalls gets all the calls that were made to GetValidatorSigningInfo.
// Check the length with:
//     len(mockedSlasher.GetValidatorSigningInfoCalls())
func (mock *SlasherMock) GetValidatorSigningInfoCalls() []struct {
	Ctx     sdk.Context
	Address sdk.ConsAddress
} {
	var calls []struct {
		Ctx     sdk.Context
		Address sdk.ConsAddress
	}
	mock.lockGetValidatorSigningInfo.RLock()
	calls = mock.calls.GetValidatorSigningInfo
	mock.lockGetValidatorSigningInfo.RUnlock()
	return calls
}

// Ensure, that BroadcasterMock does implement exported.Broadcaster.
// If this is not the case, regenerate this file with moq.
var _ exported.Broadcaster = &BroadcasterMock{}

// BroadcasterMock is a mock implementation of exported.Broadcaster.
//
// 	func TestSomethingThatUsesBroadcaster(t *testing.T) {
//
// 		// make and configure a mocked exported.Broadcaster
// 		mockedBroadcaster := &BroadcasterMock{
// 			GetProxyFunc: func(ctx sdk.Context, principal sdk.ValAddress) sdk.AccAddress {
// 				panic("mock out the GetProxy method")
// 			},
// 		}
//
// 		// use mockedBroadcaster in code that requires exported.Broadcaster
// 		// and then make assertions.
//
// 	}
type BroadcasterMock struct {
	// GetProxyFunc mocks the GetProxy method.
	GetProxyFunc func(ctx sdk.Context, principal sdk.ValAddress) sdk.AccAddress

	// calls tracks calls to the methods.
	calls struct {
		// GetProxy holds details about calls to the GetProxy method.
		GetProxy []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Principal is the principal argument value.
			Principal sdk.ValAddress
		}
	}
	lockGetProxy sync.RWMutex
}

// GetProxy calls GetProxyFunc.
func (mock *BroadcasterMock) GetProxy(ctx sdk.Context, principal sdk.ValAddress) sdk.AccAddress {
	if mock.GetProxyFunc == nil {
		panic("BroadcasterMock.GetProxyFunc: method is nil but Broadcaster.GetProxy was just called")
	}
	callInfo := struct {
		Ctx       sdk.Context
		Principal sdk.ValAddress
	}{
		Ctx:       ctx,
		Principal: principal,
	}
	mock.lockGetProxy.Lock()
	mock.calls.GetProxy = append(mock.calls.GetProxy, callInfo)
	mock.lockGetProxy.Unlock()
	return mock.GetProxyFunc(ctx, principal)
}

// GetProxyCalls gets all the calls that were made to GetProxy.
// Check the length with:
//     len(mockedBroadcaster.GetProxyCalls())
func (mock *BroadcasterMock) GetProxyCalls() []struct {
	Ctx       sdk.Context
	Principal sdk.ValAddress
} {
	var calls []struct {
		Ctx       sdk.Context
		Principal sdk.ValAddress
	}
	mock.lockGetProxy.RLock()
	calls = mock.calls.GetProxy
	mock.lockGetProxy.RUnlock()
	return calls
}

// Ensure, that TssMock does implement exported.Tss.
// If this is not the case, regenerate this file with moq.
var _ exported.Tss = &TssMock{}

// TssMock is a mock implementation of exported.Tss.
//
// 	func TestSomethingThatUsesTss(t *testing.T) {
//
// 		// make and configure a mocked exported.Tss
// 		mockedTss := &TssMock{
// 			GetMinBondFractionPerShareFunc: func(ctx sdk.Context) utils.Threshold {
// 				panic("mock out the GetMinBondFractionPerShare method")
// 			},
// 			GetTssSuspendedUntilFunc: func(ctx sdk.Context, validator sdk.ValAddress) int64 {
// 				panic("mock out the GetTssSuspendedUntil method")
// 			},
// 			GetValidatorDeregisteredBlockHeightFunc: func(ctx sdk.Context, valAddr sdk.ValAddress) int64 {
// 				panic("mock out the GetValidatorDeregisteredBlockHeight method")
// 			},
// 		}
//
// 		// use mockedTss in code that requires exported.Tss
// 		// and then make assertions.
//
// 	}
type TssMock struct {
	// GetMinBondFractionPerShareFunc mocks the GetMinBondFractionPerShare method.
	GetMinBondFractionPerShareFunc func(ctx sdk.Context) utils.Threshold

	// GetTssSuspendedUntilFunc mocks the GetTssSuspendedUntil method.
	GetTssSuspendedUntilFunc func(ctx sdk.Context, validator sdk.ValAddress) int64

	// GetValidatorDeregisteredBlockHeightFunc mocks the GetValidatorDeregisteredBlockHeight method.
	GetValidatorDeregisteredBlockHeightFunc func(ctx sdk.Context, valAddr sdk.ValAddress) int64

	// calls tracks calls to the methods.
	calls struct {
		// GetMinBondFractionPerShare holds details about calls to the GetMinBondFractionPerShare method.
		GetMinBondFractionPerShare []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
		}
		// GetTssSuspendedUntil holds details about calls to the GetTssSuspendedUntil method.
		GetTssSuspendedUntil []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Validator is the validator argument value.
			Validator sdk.ValAddress
		}
		// GetValidatorDeregisteredBlockHeight holds details about calls to the GetValidatorDeregisteredBlockHeight method.
		GetValidatorDeregisteredBlockHeight []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ValAddr is the valAddr argument value.
			ValAddr sdk.ValAddress
		}
	}
	lockGetMinBondFractionPerShare          sync.RWMutex
	lockGetTssSuspendedUntil                sync.RWMutex
	lockGetValidatorDeregisteredBlockHeight sync.RWMutex
}

// GetMinBondFractionPerShare calls GetMinBondFractionPerShareFunc.
func (mock *TssMock) GetMinBondFractionPerShare(ctx sdk.Context) utils.Threshold {
	if mock.GetMinBondFractionPerShareFunc == nil {
		panic("TssMock.GetMinBondFractionPerShareFunc: method is nil but Tss.GetMinBondFractionPerShare was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetMinBondFractionPerShare.Lock()
	mock.calls.GetMinBondFractionPerShare = append(mock.calls.GetMinBondFractionPerShare, callInfo)
	mock.lockGetMinBondFractionPerShare.Unlock()
	return mock.GetMinBondFractionPerShareFunc(ctx)
}

// GetMinBondFractionPerShareCalls gets all the calls that were made to GetMinBondFractionPerShare.
// Check the length with:
//     len(mockedTss.GetMinBondFractionPerShareCalls())
func (mock *TssMock) GetMinBondFractionPerShareCalls() []struct {
	Ctx sdk.Context
} {
	var calls []struct {
		Ctx sdk.Context
	}
	mock.lockGetMinBondFractionPerShare.RLock()
	calls = mock.calls.GetMinBondFractionPerShare
	mock.lockGetMinBondFractionPerShare.RUnlock()
	return calls
}

// GetTssSuspendedUntil calls GetTssSuspendedUntilFunc.
func (mock *TssMock) GetTssSuspendedUntil(ctx sdk.Context, validator sdk.ValAddress) int64 {
	if mock.GetTssSuspendedUntilFunc == nil {
		panic("TssMock.GetTssSuspendedUntilFunc: method is nil but Tss.GetTssSuspendedUntil was just called")
	}
	callInfo := struct {
		Ctx       sdk.Context
		Validator sdk.ValAddress
	}{
		Ctx:       ctx,
		Validator: validator,
	}
	mock.lockGetTssSuspendedUntil.Lock()
	mock.calls.GetTssSuspendedUntil = append(mock.calls.GetTssSuspendedUntil, callInfo)
	mock.lockGetTssSuspendedUntil.Unlock()
	return mock.GetTssSuspendedUntilFunc(ctx, validator)
}

// GetTssSuspendedUntilCalls gets all the calls that were made to GetTssSuspendedUntil.
// Check the length with:
//     len(mockedTss.GetTssSuspendedUntilCalls())
func (mock *TssMock) GetTssSuspendedUntilCalls() []struct {
	Ctx       sdk.Context
	Validator sdk.ValAddress
} {
	var calls []struct {
		Ctx       sdk.Context
		Validator sdk.ValAddress
	}
	mock.lockGetTssSuspendedUntil.RLock()
	calls = mock.calls.GetTssSuspendedUntil
	mock.lockGetTssSuspendedUntil.RUnlock()
	return calls
}

// GetValidatorDeregisteredBlockHeight calls GetValidatorDeregisteredBlockHeightFunc.
func (mock *TssMock) GetValidatorDeregisteredBlockHeight(ctx sdk.Context, valAddr sdk.ValAddress) int64 {
	if mock.GetValidatorDeregisteredBlockHeightFunc == nil {
		panic("TssMock.GetValidatorDeregisteredBlockHeightFunc: method is nil but Tss.GetValidatorDeregisteredBlockHeight was just called")
	}
	callInfo := struct {
		Ctx     sdk.Context
		ValAddr sdk.ValAddress
	}{
		Ctx:     ctx,
		ValAddr: valAddr,
	}
	mock.lockGetValidatorDeregisteredBlockHeight.Lock()
	mock.calls.GetValidatorDeregisteredBlockHeight = append(mock.calls.GetValidatorDeregisteredBlockHeight, callInfo)
	mock.lockGetValidatorDeregisteredBlockHeight.Unlock()
	return mock.GetValidatorDeregisteredBlockHeightFunc(ctx, valAddr)
}

// GetValidatorDeregisteredBlockHeightCalls gets all the calls that were made to GetValidatorDeregisteredBlockHeight.
// Check the length with:
//     len(mockedTss.GetValidatorDeregisteredBlockHeightCalls())
func (mock *TssMock) GetValidatorDeregisteredBlockHeightCalls() []struct {
	Ctx     sdk.Context
	ValAddr sdk.ValAddress
} {
	var calls []struct {
		Ctx     sdk.Context
		ValAddr sdk.ValAddress
	}
	mock.lockGetValidatorDeregisteredBlockHeight.RLock()
	calls = mock.calls.GetValidatorDeregisteredBlockHeight
	mock.lockGetValidatorDeregisteredBlockHeight.RUnlock()
	return calls
}
