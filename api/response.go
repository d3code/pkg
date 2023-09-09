package api

type Response struct {
    Data       any         `json:"data,omitempty"` // response data
    Count      *Count      `json:"count,omitempty"`
    Pagination *Pagination `json:"pagination,omitempty"`
    Links      Links       `json:"links"`
}

type Count struct {
    Returned int `json:"returned"` // items returned in response
    Total    int `json:"total"`    // total number of items
}

type Pagination struct {
    Limit int `json:"limit"` // number of items per page
    Page  int `json:"page"`  // current page
    Pages int `json:"pages"` // total number of pages
}

type Links struct {
    Self  string  `json:"_self"`           // current page
    First *string `json:"first,omitempty"` // first page
    Prev  *string `json:"prev,omitempty"`  // previous page
    Next  *string `json:"next,omitempty"`  // next page
    Last  *string `json:"last,omitempty"`  // last page
}

type ResponseError struct {
    Type    *string `json:"type,omitempty"`    // error type
    Message string  `json:"message"`           // error message
    Details any     `json:"details,omitempty"` // additional details about error
    Field   *string `json:"field,omitempty"`   // json field name of request in error if applicable
}
