package basicCtx

import (
	"context"
	"reflect"
	"testing"
)

func TestWithTraceID(t *testing.T) {
	type args struct {
		ctx     context.Context
		traceID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_trace_id",
			args: args{
				ctx:     context.Background(),
				traceID: "trace_id",
			},
			want: "trace_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := WithTraceID(tt.args.ctx, tt.args.traceID)
			if !reflect.DeepEqual(GetMeta(ctx).TraceID(), tt.want) {
				t.Errorf("WithTraceID() = %v, want %v", ctx, tt.want)
			}
		})
	}
}

func TestWithUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test_user_id",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := WithUserID(tt.args.ctx, tt.args.userID)
			if !reflect.DeepEqual(GetMeta(ctx).UserID(), tt.want) {
				t.Errorf("WithUserID() = %v, want %v", ctx, tt.want)
			}
		})
	}
}

func TestWithTraceIDAndUserID(t *testing.T) {
	type args struct {
		ctx     context.Context
		traceID string
		userID  int64
	}
	type want struct {
		traceID string
		userID  int64
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "test_trace_id_and_user_id",
			args: args{
				ctx:     context.Background(),
				traceID: "trace_id",
				userID:  1,
			},
			want: want{
				traceID: "trace_id",
				userID:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := WithTraceID(tt.args.ctx, tt.args.traceID)
			ctx = WithUserID(ctx, tt.args.userID)
			if !reflect.DeepEqual(GetMeta(ctx).TraceID(), tt.want.traceID) {
				t.Errorf("WithTraceID() = %v, want %v", ctx, tt.want)
			}
			if !reflect.DeepEqual(GetMeta(ctx).UserID(), tt.want.userID) {
				t.Errorf("WithUserID() = %v, want %v", ctx, tt.want)
			}
		})
	}
}
