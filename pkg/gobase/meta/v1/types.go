package v1

type ListOptions struct {
	LabelSelector string `json:"labelSelector,omitempty" form:"labelSelector"`
	FieldSelector string `json:"fieldSelector,omitempty" form:"fieldSelector"`
	Offset        *int64 `json:"offset,omitempty" form:"offset"`
	Limit         *int64 `json:"limit,omitempty" form:"limit"`
}
