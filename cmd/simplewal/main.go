package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/muhammadmontasir/simplewal/internal/wal"
	pb "github.com/muhammadmontasir/simplewal/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("SimpleWAL Client Tool: Greetings!")

	// Use the internal WAL package for a greeting
	greeting := wal.GetSimpleGreeting("Bazel User")
	fmt.Println(greeting)

	// Simulate using the WAL
	myWal := wal.New("mydata.wal")
	sampleData := []byte("This is a test log entry from the client.")
	entry := wal.NewLogEntry(1, sampleData)

	simulatedAppendMsg, err := myWal.Append(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error simulating append: %v\n", err)
	} else {
		fmt.Println(simulatedAppendMsg)
	}

	// --- Client Tool Specific Logic with Protobuf ---
	if len(os.Args) > 1 {
		command := os.Args[1]
		clientRequest := &pb.ClientRequest{Command: command}
		var clientResponse *pb.ClientResponse

		fmt.Printf("\nClient sending command: %s\n", command)

		if strings.ToLower(command) == "add" {
			var entryData string
			if len(os.Args) > 2 {
				entryData = strings.Join(os.Args[2:], " ")
			} else {
				entryData = "Default client data"
			}
			// Parse the ID argument as uint64, fallback to 1 if not provided or invalid
			var entryID uint64 = 1
			if len(os.Args) > 2 {
				if parsedID, err := strconv.ParseUint(os.Args[2], 10, 64); err == nil {
					entryID = parsedID
					entryData = strings.Join(os.Args[3:], " ")
				}
			}
			clientRequest.Entry = wal.NewLogEntry(entryID, []byte(entryData)) // Simple ID generation
			// In a real app, this request would be sent to a server.
			// Here, we simulate processing it.
			fmt.Printf("Client simulating 'add' with data: %s\n", entryData)
			clientResponse = &pb.ClientResponse{
				Success: true,
				Message: "Entry added (simulated).",
				Entry:   clientRequest.Entry, // Echo back the entry
			}
		} else {
			clientResponse = &pb.ClientResponse{
				Success: false,
				Message: fmt.Sprintf("Unknown command: %s. Try 'add'.", command),
			}
		}

		fmt.Printf("Client received response: Success=%t, Message='%s'\n", clientResponse.GetSuccess(), clientResponse.GetMessage())
		if clientResponse.GetEntry() != nil {
			fmt.Printf("Response Entry: %s\n", wal.FormatEntry(clientResponse.GetEntry()))
		}

		// Demonstrate marshalling the response (optional)
		marshaledResponse, protoErr := proto.Marshal(clientResponse)
		if protoErr != nil {
			fmt.Fprintf(os.Stderr, "Failed to marshal client response: %v\n", protoErr)
		} else {
			fmt.Printf("Marshaled client response size: %d bytes\n", len(marshaledResponse))
		}

	} else {
		fmt.Println("\nClient Tool Usage: ./simplewal <command> [data...]")
		fmt.Println("Example: ./simplewal add My log message here")
	}
}
