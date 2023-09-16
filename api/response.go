package api

type Response struct {
    Data       any         `json:"data,omitempty"` // response data
    Count      *count      `json:"count,omitempty"`
    Pagination *pagination `json:"pagination,omitempty"`
    Links      *links      `json:"links,omitempty"`
}

type count struct {
    Returned int `json:"returned"` // items returned in response
    Total    int `json:"total"`    // total number of items
}

type pagination struct {
    Limit int `json:"limit"` // number of items per page
    Page  int `json:"page"`  // current page
    Pages int `json:"pages"` // total number of pages
}

type links struct {
    Self  string  `json:"_self"`           // current page
    First *string `json:"first,omitempty"` // first page
    Prev  *string `json:"prev,omitempty"`  // previous page
    Next  *string `json:"next,omitempty"`  // next page
    Last  *string `json:"last,omitempty"`  // last page
}
