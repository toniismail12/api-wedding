package table

type Wishes struct {
	Id     uint   `json:"id"`
	App_id uint   `json:"app_id"`
	Name   string `json:"name"`
	Wishes string `json:"wishes"`
}
