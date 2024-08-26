package services

import (
	"context"
	db "courseonline/db/sqlc"
	"log"
)

func (sm *StoreManager) CreateCartTx(ctx context.Context, args db.CreateOrderCourseParams) (*db.OrderCourse, error) {
	tx, err := sm.dbConn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	qtx := sm.Queries.WithTx(tx)

	//populate cart list
	carts, err := qtx.GetCartByID(ctx, *args.UscoUserID)
	if err != nil {
		return nil, err
	}

	log.Println(carts)

	//create order
	newOrder, err := qtx.CreateOrderCourse(ctx, args)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(context.Background()); err != nil {
		return nil, err
	}
	return newOrder, nil

}
