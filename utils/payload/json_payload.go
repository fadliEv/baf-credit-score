package payload

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type PagedResponse struct {
	Status Status        `json:"status"`
	Data   []interface{} `json:"data,omitempty"`
	Paging Paging        `json:"paging,omitempty"`
}

type Paging struct {
	Page        int   `json:"paging"`
	RowsPerPage int   `json:"rowsPerPage"`
	TotalRows   int64 `json:"totalRows"`
	TotalPages  int   `json:"totalPages"`
}