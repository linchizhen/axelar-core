package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/axelarnetwork/axelar-core/x/tss/exported"
	"github.com/axelarnetwork/axelar-core/x/tss/types"
)

// InitGenesis initializes the tss module's state from a given genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, snapshotter types.Snapshotter, genState *types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	k.SetGovernanceKey(ctx, *genState.GovernanceKey)

	rotationCountMap := make(map[string]map[exported.KeyRole]int64)
	for _, key := range genState.Keys {
		if _, ok := k.GetKey(ctx, key.ID); ok {
			panic(fmt.Errorf("key %s already set", key.ID))
		}

		k.setKey(ctx, key)

		if key.RotationCount > 0 {
			if _, ok := k.getKeyID(ctx, key.Chain, key.RotationCount, key.Role); ok {
				panic(fmt.Errorf("key ID for chain %s, rotation count %d and role %s already set", key.Chain, key.RotationCount, key.Role.SimpleString()))
			}

			k.setKeyID(ctx, key.Chain, key.RotationCount, key.Role, key.ID)
		}

		if key.Role != exported.ExternalKey {
			if _, ok := snapshotter.GetSnapshot(ctx, key.SnapshotCounter); !ok {
				panic(fmt.Errorf("snapshot %d for key %s not found", key.SnapshotCounter, key.ID))
			}

			k.setSnapshotCounterForKeyID(ctx, key.ID, key.SnapshotCounter)
		}

		rotationCount, ok := rotationCountMap[key.Chain][key.Role]
		if !ok || key.RotationCount > rotationCount {
			rotationCountMap[key.Chain][key.Role] = key.RotationCount
		}
	}

	for chain, keyRoleToRotationCount := range rotationCountMap {
		for keyRole, rotationCount := range keyRoleToRotationCount {
			k.setRotationCount(ctx, chain, keyRole, rotationCount)
		}
	}

	for _, keyRecoveryInfo := range genState.KeyRecoveryInfos {
		if key, ok := k.GetKey(ctx, keyRecoveryInfo.KeyID); ok && key.Type != exported.Threshold {
			panic(fmt.Errorf("no threshold key %s found", keyRecoveryInfo.KeyID))
		}

		if _, ok := k.getKeyRecoveryInfo(ctx, keyRecoveryInfo.KeyID); ok {
			panic(fmt.Errorf("key recovery info %s already set", keyRecoveryInfo.KeyID))
		}

		k.setKeyRecoveryInfo(ctx, keyRecoveryInfo)
	}

	for _, multisigInfo := range genState.MultisigInfos {
		if _, ok := k.GetMultisigKeygenInfo(ctx, exported.KeyID(multisigInfo.ID)); ok {
			panic(fmt.Errorf("multisig keygen info %s already set", multisigInfo.ID))
		}

		k.SetMultisigKeygenInfo(ctx, multisigInfo)
	}

	for _, externalKeys := range genState.ExternalKeys {
		for _, externalKeyID := range externalKeys.KeyIDs {
			if _, ok := k.GetKey(ctx, externalKeyID); !ok {
				panic(fmt.Errorf("no key %s found", externalKeyID))
			}
		}

		if _, ok := k.getExternalKeyIDs(ctx, externalKeys.Chain); ok {
			panic(fmt.Errorf("external key IDs for chain %s already set", externalKeys.Chain))
		}

		k.setExternalKeys(ctx, externalKeys)
	}
}

// ExportGenesis returns the tss module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	keys := k.getKeys(ctx)

	governanceKey, ok := k.GetGovernanceKey(ctx)
	if !ok {
		return types.NewGenesisState(
			k.GetParams(ctx),
			nil,
			k.getKeyRecoveryInfos(ctx),
			keys,
			k.getCompletedMultisigKeygenInfos(ctx),
			k.getAllExternalKeys(ctx),
		)
	}

	return types.NewGenesisState(
		k.GetParams(ctx),
		&governanceKey,
		k.getKeyRecoveryInfos(ctx),
		keys,
		k.getCompletedMultisigKeygenInfos(ctx),
		k.getAllExternalKeys(ctx),
	)
}
