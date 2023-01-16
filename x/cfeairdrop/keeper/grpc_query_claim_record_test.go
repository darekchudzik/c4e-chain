package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/chain4energy/c4e-chain/testutil/keeper"
	"github.com/chain4energy/c4e-chain/testutil/nullify"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestClaimRecordQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNClaimRecord(keeper, ctx, 2, 10, true, true)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryClaimRecordRequest
		response *types.QueryClaimRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryClaimRecordRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryClaimRecordResponse{UserAirdropEntries: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryClaimRecordRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryClaimRecordResponse{UserAirdropEntries: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryClaimRecordRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.UserAirdropEntries(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestClaimRecordQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CfeairdropKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNClaimRecord(keeper, ctx, 5, 0, false, false)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryClaimRecordsRequest {
		return &types.QueryClaimRecordsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ClaimRecords(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserAirdropEntries), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserAirdropEntries),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ClaimRecords(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserAirdropEntries), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserAirdropEntries),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ClaimRecords(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.UserAirdropEntries),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ClaimRecords(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
