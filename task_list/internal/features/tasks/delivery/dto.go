package delivery

type TaskResponse struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}
type TaskUpdate struct {
	Title     string `json:"title"`
	Desc      string `json:"description"`
	Completed bool   `json:"completed"`
}
