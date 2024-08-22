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
