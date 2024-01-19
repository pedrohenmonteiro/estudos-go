package main

import (
	"context"
	"database/sql"

	"sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/blog")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        sql.NullString{String: "Backend", Valid: true},
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name.String, category.Description.String)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "00000000-0000-0000-0000-000000000001",
	// 	Name:        sql.NullString{String: "Updated query", Valid: true},
	// 	Description: sql.NullString{String: "Query was updated", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name.String, category.Description.String)
	// }

	err = queries.DeleteCategory(ctx, "00000000-0000-0000-0000-000000000001")
	if err != nil {
		panic(err)
	}
}
