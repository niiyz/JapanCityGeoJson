package geo

type TopoJson struct {
	Type     string    `json:"type"`
	Transform Transform `json:"transform"`
	Objects Objects `json:"objects"`
	Arcs  [][][][2]int32 `json:"arcs"`
}

type Transform struct{
	Scale [2]float64    `json:"scale"`
	Translate [2]float64    `json:"translate"`
}

type Objects map[string]struct{
	Type string   `json:"type"`
	Arcs [][][][1]int `json:"arcs"`
	Id int64 `json:"id"`
}