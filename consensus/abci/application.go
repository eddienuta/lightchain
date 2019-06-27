package abci

import "github.com/tendermint/tendermint/abci/types"

var _ types.Application = (*Application)(nil)

type Application struct {
	types.BaseApplication
}

func (app *Application) BeginBlock(req types.Request_BeginBlock) (types.ResponseBeginBlock, error) {
	return types.ResponseBeginBlock{}, nil
} 


