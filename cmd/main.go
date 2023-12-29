package main

import (
	"log"
	"os"

	"github.com/Ayomide-Daniel/go-hex/internal/adapters/app/api"
	"github.com/Ayomide-Daniel/go-hex/internal/adapters/core/arithmetic"
	"github.com/Ayomide-Daniel/go-hex/internal/adapters/framework/right/db"
	"github.com/Ayomide-Daniel/go-hex/internal/ports"

	gRPC "github.com/Ayomide-Daniel/go-hex/internal/adapters/framework/left/grpc"
)

func main() {
	// ports/ interfaces
	var core ports.ArithmeticPort
	var databaseAdapter ports.DBPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER")
	dbSourceName := os.Getenv("DS_NAME")

	databaseAdapter, err := db.NewAdapter(dbDriver, dbSourceName)

	if err != nil {
		log.Fatalf("Failed to initiate DB Connection: %v", err)
	}

	defer databaseAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(databaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
