package wal

import (
	"fmt"
	"time"

	pb "github.com/muhammadmontasir/simplewal/proto"
)

// NewLogEntry creates a new log entry protobuf message
func NewLogEntry(id uint64, data []byte) *pb.LogEntry {
	return &pb.LogEntry{
		Id:          id,
		Data:        data,
		TimestampMs: time.Now().UnixMilli(),
	}
}

// FormatEntry creates a string representation of a LogEntry.
func FormatEntry(entry *pb.LogEntry) string {
	if entry == nil {
		return "nil LogEntry"
	}

	return fmt.Sprint("ID: %d, Timestamp: %d, Data: %s", entry.GetId(), entry.GetTimestampMs(), string(entry.GetData()))
}
