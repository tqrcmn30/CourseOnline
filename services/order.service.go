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
	carts, err := qtx.GetCartByUserID(ctx, args.UscoUserID)
	if err != nil {
		return nil, err
	}

	log.Println(carts)

	//create order
	newOrder, err := qtx.CreateOrderCourse(ctx, args)
	if err != nil {
		return nil, err
	}

	var subtotal float32
	var totalPrice float32
	var tax float32
	tax = 0.1
	for _, v := range carts {
		totalPrice = float32(*v.CartPrice) * float32(*v.CartQty)
		subtotal = totalPrice + (totalPrice * tax)
		argsCoursedetail := db.CreateOrderCoursesDetailParams{
			UcdeQty:        v.CartQty,
			UcdePrice:      v.CartPrice,
			UcdeTotalPrice: &totalPrice,
			UcdeUscoID:     &newOrder.UscoID,
			UcdeCoursID:    v.CartCoursID,
		}
		_, err := qtx.CreateOrderCoursesDetail(ctx, argsCoursedetail)
		if err != nil {
			return nil, err
		}
	}

	argsUpdate := db.UpdateOrderCourseParams{
		UscoID:         newOrder.UscoID,
		UscoPurchaseNo: newOrder.UscoPurchaseNo,
		UscoTax:        &tax,
		UscoSubtotal:   &subtotal,
		UscoUserID:     newOrder.UscoUserID,
	}

	_, errUpdate := qtx.UpdateOrderCourse(ctx, argsUpdate)
	if errUpdate != nil {
		return nil, errUpdate
	}
	newOrder.UscoSubtotal = &subtotal
	if err = tx.Commit(context.Background()); err != nil {
		return nil, err
	}
	return newOrder, nil

}
