package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"cosmossdk.io/core/address"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/client/cli"
)

// ExecTestCLICmd builds the client context, mocks the output and executes the command.
func ExecTestCLICmd(clientCtx client.Context, cmd *cobra.Command, extraArgs []string) (testutil.BufferWriter, error) {
	cmd.SetArgs(extraArgs)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	if err := cmd.ExecuteContext(ctx); err != nil {
		return out, err
	}

	return out, nil
}

func MsgSendExec(clientCtx client.Context, from, to sdk.AccAddress, amount fmt.Stringer, ac address.Codec, extraArgs ...string) (testutil.BufferWriter, error) {
	// args := []string{from.String(), to.String(), amount.String()}
	var args []string
	for _, addr := range []sdk.AccAddress{from, to} {
		addrStr, err := ac.BytesToString(addr.Bytes())
		if err != nil {
			return nil, fmt.Errorf("failed to convert address %s: %w", addr, err)
		}
		args = append(args, addrStr)
	}
	args = append(args, amount.String())
	args = append(args, extraArgs...)

	return ExecTestCLICmd(clientCtx, cli.NewSendTxCmd(ac), args)
}
