package yanyue

type BrandResp struct {
	BrandList []*Brand `json:"brandlist"`
}

type ProductResp struct {
	Page        int32      `json:"page"`
	PageNum     int32      `json:"pagenum"`
	TotalNum    int32      `json:"totalnum"`
	ProductList []*Product `json:"productlist"`
}
