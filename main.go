package main

import (
	"context"
	"joshuaswickirl/sql-poc/internal/db"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func run() error {
	ctx := context.Background()

	connstring := "postgresql://app:password@localhost:5432/db?sslmode=disable"
	conn, err := pgx.Connect(ctx, connstring)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))

	// if err := queries.DeleteAuthor(ctx, insertedAuthor.ID); err != nil {
	// 	return err
	// }

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
