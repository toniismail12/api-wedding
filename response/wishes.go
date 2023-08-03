package response

type GetWishes struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Wishes string `json:"wishes"`
}

type FormWishes struct {
	App_id uint   `json:"app_id"`
	Name   string `json:"name"`
	Wishes string `json:"wishes"`
}
