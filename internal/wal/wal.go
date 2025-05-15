package wal

import (
	pb "github.com/muhammadmontasir/simplewal/proto"
)

// WriteAheadLog struct
type WriteAheadLog struct {
	// In a real WAL, you'd have fields for file handles, current segment, etc.
	name string
}

// New creates a new WAL
func New(name string) *WriteAheadLog {
	return &WriteAheadLog{name: name}
}

func (w *WriteAheadLog) Append(entry *pb.LogEntry) (string, error) {
	formattedEntry := FormatEntry(entry)
	return "Appended (Simulated): " + formattedEntry, nil
}

// GetSimpleGreeting provides a basic greeting
func GetSimpleGreeting(target string) string {
	return "Hello, " + target + ", from the WAL!"
}
