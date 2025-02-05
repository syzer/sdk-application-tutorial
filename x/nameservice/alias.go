package nameservice

import (
	"github.com/cosmos/sdk-application-tutorial/x/nameservice/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgBuyName = types.NewMsgBuyLocation
	NewMsgSetName = types.NewMsgSetName
	NewWhois      = types.NewWhois
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	MsgSetName      = types.MsgSetDAG
	MsgBuyName      = types.MsgBuyLocation
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois
)
