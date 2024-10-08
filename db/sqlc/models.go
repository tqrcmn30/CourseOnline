// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Cart struct {
	CartID       int32            `json:"cart_id"`
	CartUserID   *int32           `json:"cart_user_id"`
	CartCoursID  *int32           `json:"cart_cours_id"`
	CartQty      *int32           `json:"cart_qty"`
	CartPrice    pgtype.Numeric   `json:"cart_price"`
	CartModified pgtype.Timestamp `json:"cart_modified"`
	CartStatus   *string          `json:"cart_status"`
	CartCartID   *int32           `json:"cart_cart_id"`
}

type Category struct {
	CateID   int32   `json:"cate_id"`
	CateName *string `json:"cate_name"`
}

type Course struct {
	CoursID       int32            `json:"cours_id"`
	CoursName     *string          `json:"cours_name"`
	CoursDesc     *string          `json:"cours_desc"`
	CoursAuthor   *string          `json:"cours_author"`
	CoursPrice    pgtype.Numeric   `json:"cours_price"`
	CoursModified pgtype.Timestamp `json:"cours_modified"`
	CoursCateID   *int32           `json:"cours_cate_id"`
}

type CoursesImage struct {
	CoimID       int32   `json:"coim_id"`
	CoimFilename *string `json:"coim_filename"`
	CoimDefault  *string `json:"coim_default"`
	CoimRemeID   *int32  `json:"coim_reme_id"`
}

type OrderCourse struct {
	UscoID         int32            `json:"usco_id"`
	UscoPurchaseNo *string          `json:"usco_purchase_no"`
	UscoTax        pgtype.Numeric   `json:"usco_tax"`
	UscoSubtotal   *float32   `json:"usco_subtotal"`
	UscoPatrxNo    *string          `json:"usco_patrx_no"`
	UscoModified   pgtype.Timestamp `json:"usco_modified"`
	UscoUserID     *int32           `json:"usco_user_id"`
}

type OrderCoursesDetail struct {
	UcdeID         int32          `json:"ucde_id"`
	UcdeQty        *int32         `json:"ucde_qty"`
	UcdePrice      pgtype.Numeric `json:"ucde_price"`
	UcdeTotalPrice pgtype.Numeric `json:"ucde_total_price"`
	UcdeUscoID     *int32         `json:"ucde_usco_id"`
	UcdeCoursID    *int32         `json:"ucde_cours_id"`
}

type User struct {
	UserID       int32   `json:"user_id"`
	UserName     *string `json:"user_name"`
	UserPassword *string `json:"user_password"`
	UserEmail    *string `json:"user_email"`
	UserPhone    *string `json:"user_phone"`
	UserToken    *string `json:"user_token"`
}
