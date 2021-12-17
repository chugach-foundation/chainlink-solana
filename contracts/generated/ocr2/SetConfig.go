// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ocr_2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetConfig is the `setConfig` instruction.
type SetConfig struct {
	NewOracles *[]NewOracle
	F          *uint8

	// [0] = [WRITE] state
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetConfigInstructionBuilder creates a new `SetConfig` instruction builder.
func NewSetConfigInstructionBuilder() *SetConfig {
	nd := &SetConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewOracles sets the "newOracles" parameter.
func (inst *SetConfig) SetNewOracles(newOracles []NewOracle) *SetConfig {
	inst.NewOracles = &newOracles
	return inst
}

// SetF sets the "f" parameter.
func (inst *SetConfig) SetF(f uint8) *SetConfig {
	inst.F = &f
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *SetConfig) SetStateAccount(state ag_solanago.PublicKey) *SetConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *SetConfig) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst SetConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewOracles == nil {
			return errors.New("NewOracles parameter is not set")
		}
		if inst.F == nil {
			return errors.New("F parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *SetConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewOracles", *inst.NewOracles))
						paramsBranch.Child(ag_format.Param("         F", *inst.F))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj SetConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewOracles` param:
	err = encoder.Encode(obj.NewOracles)
	if err != nil {
		return err
	}
	// Serialize `F` param:
	err = encoder.Encode(obj.F)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewOracles`:
	err = decoder.Decode(&obj.NewOracles)
	if err != nil {
		return err
	}
	// Deserialize `F`:
	err = decoder.Decode(&obj.F)
	if err != nil {
		return err
	}
	return nil
}

// NewSetConfigInstruction declares a new SetConfig instruction with the provided parameters and accounts.
func NewSetConfigInstruction(
	// Parameters:
	newOracles []NewOracle,
	f uint8,
	// Accounts:
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetConfig {
	return NewSetConfigInstructionBuilder().
		SetNewOracles(newOracles).
		SetF(f).
		SetStateAccount(state).
		SetAuthorityAccount(authority)
}
