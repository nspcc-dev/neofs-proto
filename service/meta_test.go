package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockedRequest struct {
	msg     string
	name    string
	code    codes.Code
	handler TTLCondition
	RequestMetaHeader
}

func TestMetaRequest(t *testing.T) {
	tests := []mockedRequest{
		{
			name:              "direct to ir node",
			handler:           IRNonForwarding(InnerRingNode),
			RequestMetaHeader: RequestMetaHeader{TTL: NonForwardingTTL},
		},
		{
			code:              codes.InvalidArgument,
			msg:               ErrIncorrectTTL.Error(),
			name:              "direct to storage node",
			handler:           IRNonForwarding(StorageNode),
			RequestMetaHeader: RequestMetaHeader{TTL: NonForwardingTTL},
		},
		{
			msg:               ErrZeroTTL.Error(),
			code:              codes.InvalidArgument,
			name:              "zero ttl",
			handler:           IRNonForwarding(StorageNode),
			RequestMetaHeader: RequestMetaHeader{TTL: ZeroTTL},
		},
		{
			name:              "default to ir node",
			handler:           IRNonForwarding(InnerRingNode),
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
		},
		{
			name:              "default to storage node",
			handler:           IRNonForwarding(StorageNode),
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
		},
		{
			msg:               "not found",
			code:              codes.NotFound,
			name:              "custom status error",
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
			handler:           func(_ uint32) error { return status.Error(codes.NotFound, "not found") },
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			before := tt.GetTTL()
			err := ProcessRequestTTL(&tt, tt.handler)
			if tt.msg != "" {
				require.Errorf(t, err, tt.msg)

				state, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.code, state.Code())
				require.Equal(t, tt.msg, state.Message())
			} else {
				require.NoError(t, err)
				require.NotEqualf(t, before, tt.GetTTL(), "ttl should be changed: %d vs %d", before, tt.GetTTL())
			}
		})
	}
}