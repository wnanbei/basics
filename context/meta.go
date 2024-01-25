package basicCtx

// Meta 元数据信息
type Meta struct {
	traceID *string
	userID  *int64
}

// WithUserID 设置 user id
func (m *Meta) WithUserID(userID int64) {
	m.userID = &userID
}

// UserID 获取 user id
func (m *Meta) UserID() int64 {
	if m.userID == nil {
		return 0
	}
	return *m.userID
}

// UserIDExists
func (m *Meta) UserIDExists() (int64, bool) {
	if m.userID == nil {
		return 0, false
	}
	return *m.userID, true
}

// WithTraceID 设置 trace id
func (m *Meta) WithTraceID(traceID string) {
	m.traceID = &traceID
}

// TraceID 获取 trace id
func (m *Meta) TraceID() string {
	if m.traceID == nil {
		return ""
	}
	return *m.traceID
}

// TraceIDExists
func (m *Meta) TraceIDExists() (string, bool) {
	if m.traceID == nil {
		return "", false
	}
	return *m.traceID, true
}
