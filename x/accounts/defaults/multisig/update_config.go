package multisig

import (
	"context"

	v1 "cosmossdk.io/x/accounts/defaults/multisig/v1"
)

// UpdateConfig updates the configuration of the multisig account.
func (a Account) UpdateConfig(ctx context.Context, msg *v1.MsgUpdateConfigRequest) (*v1.MsgUpdateConfigResponse, error) {
	// set members
	for i := range msg.UpdateMembers {
		addrBz, err := a.addrCodec.StringToBytes(msg.UpdateMembers[i].Address)
		if err != nil {
			return nil, err
		}

		if msg.UpdateMembers[i].Weight == 0 {
			if err := a.Members.Remove(ctx, addrBz); err != nil {
				return nil, err
			}
			continue
		}
		if err := a.Members.Set(ctx, addrBz, msg.UpdateMembers[i].Weight); err != nil {
			return nil, err
		}
	}

	if msg.Config != nil {
		// set config
		if err := a.Config.Set(ctx, *msg.Config); err != nil {
			return nil, err
		}
	}

	// verify that the new set of members and config are valid
	// get the weight from the stored members
	totalWeight := uint64(0)
	err := a.Members.Walk(ctx, nil, func(_ []byte, value uint64) (stop bool, err error) {
		totalWeight += value
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	// get the config from state given that we might not have it in the message
	config, err := a.Config.Get(ctx)
	if err != nil {
		return nil, err
	}

	if err := validateConfig(config, totalWeight); err != nil {
		return nil, err
	}

	return &v1.MsgUpdateConfigResponse{}, nil
}
