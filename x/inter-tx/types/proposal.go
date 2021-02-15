package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ProposalTypeRegisterInterchainAccount defines the type for a
	ProposalTypeRegisterInterchainAccount = "RegisterInterchainAccount"
)

var _ govtypes.Content = &MsgRegisterProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterInterchainAccount)
	govtypes.RegisterProposalTypeCodec(&MsgRegisterProposal{}, "cosmos-sdk/MsgRegisterProposal")
}

func NewMsgRegisterProposal(title, description, sourcePort, sourceChannel string, sender sdk.AccAddress) *MsgRegisterProposal {
	return &MsgRegisterProposal{title, description, sourcePort, sourceChannel, sender}
}

func (csp *MsgRegisterProposal) ProposalRoute() string { return RouterKey }

func (csp *MsgRegisterProposal) ProposalType() string { return ProposalTypeRegisterInterchainAccount }

func (csp *MsgRegisterProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(csp)
	if err != nil {
		return err
	}
	//if !csp.Amount.IsValid() {
	//	return ErrInvalidProposalAmount
	//}
	//if csp.Recipient == "" {
	//	return ErrEmptyProposalRecipient
	//}

	return nil
}
