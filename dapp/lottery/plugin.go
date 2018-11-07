package lottery

import (
	"gitlab.33.cn/chain33/chain33/pluginmgr"
	"gitlab.33.cn/chain33/plugin/dapp/lottery/executor"
	"gitlab.33.cn/chain33/plugin/dapp/lottery/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.LotteryX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      nil,
		RPC:      nil,
	})
}
