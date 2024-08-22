package models

type CreateUserReq struct {
	UserName     *string `json:"user_name" binding:"required"`
	UserPassword *string `json:"user_password" binding:"required"`
}

type UserResponse struct {
	UserID       int32   `json:"user_id"`
	UserName     *string `json:"user_name"`
	UserPassword *string `json:"user_password"`
	UserPhone    *string `json:"user_phone"`
	UserToken    *string `json:"user_token"`
}

type CategoryUpdateReq struct {
	CateID   int32   `json:"cate_id"`
	CateName *string `json:"cate_name"`
}

type CategoryPostReq struct {
	CateName *string `json:"cate_name"`
}

type CoursePostReq struct {
	CoursName   *string  `json:"cours_name"`
	CoursDesc   *string  `json:"cours_desc"`
	CoursAuthor *string  `json:"cours_author"`
	CoursPrice  *float32 `json:"cours_price"`
	CoursCateID *int32   `json:"cours_cate_id"`
}

type CourseUpdateReq struct {
	CoursID     int32    `json:"cours_id"`
	CoursName   *string  `json:"cours_name"`
	CoursDesc   *string  `json:"cours_desc"`
	CoursAuthor *string  `json:"cours_author"`
	CoursPrice  *float32 `json:"cours_price"`
	CoursCateID *int32   `json:"cours_cate_id"`
}

type CourseImagesPostReq struct {
	CoimFilename *string `json:"coim_filename"`
	CoimDefault  *string `json:"coim_default"`
	CoimRemeID   *int32  `json:"coim_reme_id"`
}

type CourseImagesUpdateReq struct {
	CoimFilename *string `json:"coim_filename"`
	CoimDefault  *string `json:"coim_default"`
	CoimRemeID   *int32  `json:"coim_reme_id"`
}

type OrderCoursesDetailPostReq struct {
	UcdeQty        *int32   `json:"ucde_qty"`
	UcdePrice      *float32 `json:"ucde_price"`
	UcdeTotalPrice *float32 `json:"ucde_total_price"`
	UcdeUscoID     *int32   `json:"ucde_usco_id"`
	UcdeCoursID    *int32   `json:"ucde_cours_id"`
}

type OrderCoursesDetailUpdateReq struct {
	UcdeID         int32    `json:"ucde_id"`
	UcdeQty        *int32   `json:"ucde_qty"`
	UcdePrice      *float32 `json:"ucde_price"`
	UcdeTotalPrice *float32 `json:"ucde_total_price"`
	UcdeUscoID     *int32   `json:"ucde_usco_id"`
	UcdeCoursID    *int32   `json:"ucde_cours_id"`
}

type OrderCoursePostReq struct {
	UscoPurchaseNo *string  `json:"usco_purchase_no"`
	UscoTax        *float32 `json:"usco_tax"`
	UscoSubtotal   *float32 `json:"usco_subtotal"`
	UscoPatrxNo    *string  `json:"usco_patrx_no"`
	UscoUserID     *int32   `json:"usco_user_id"`
}

type OrderCourseUpdateReq struct {
	UscoID         int32    `json:"usco_id"`
	UscoPurchaseNo *string  `json:"usco_purchase_no"`
	UscoTax        *float32 `json:"usco_tax"`
	UscoSubtotal   *float32 `json:"usco_subtotal"`
	UscoPatrxNo    *string  `json:"usco_patrx_no"`
	UscoUserID     *int32   `json:"usco_user_id"`
}

type CourseImagePostReq struct {
	CoimFilename *string `form:"coim_filename"`
	CoimDefault  *string `form:"coim_default"`
	CoimRemeID   *int32  `form:"coim_reme_id"`
}

type CourseImageUpdateReq struct {
	CoimID       int32   `form:"coim_id"`
	CoimFilename *string `form:"coim_filename"`
	CoimDefault  *string `form:"coim_default"`
	CoimRemeID   *int32  `form:"coim_reme_id"`
}
