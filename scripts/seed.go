package main

import (
	"context"

	"github.com/haleyrc/phir/internal/fake"
	"github.com/haleyrc/phir/lib/sqlite"
)

func main() {
	ctx := context.Background()

	conn := sqlite.MustOpen("dev.db")
	defer conn.Close()

	faker := fake.New(conn)

	faker.CreatePatient(ctx, fake.Patient{
		FirstName: "Jane",
		LastName:  "Doe",
	})
	faker.CreatePatient(ctx, fake.Patient{
		FirstName: "John",
		LastName:  "Smith",
	})
}
