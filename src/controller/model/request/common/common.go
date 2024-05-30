package requestCommon

type PaginationQuery struct {
	Page  string `form:"pagination_page" json:"pagination_page" query:"page" binding:"omitempty"`
	Size  string `form:"pagination_size" json:"pagination_size" query:"size" binding:"omitempty"`
	Order string `form:"pagination_order" json:"pagination_order" query:"order" binding:"omitempty"`
}
