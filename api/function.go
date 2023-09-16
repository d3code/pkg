package api

import (
    "net/url"
    "reflect"
    "strconv"
)

// ResponsePaginated returns a paginated response
// If the data is not an array, map or slice, the data is returned as is
func ResponsePaginated(data any, page int, limit int, total int, url *url.URL) Response {
    var lastPage int
    if limit == 0 {
        lastPage = 1
    } else {
        lastPage = (total / limit) + 1
    }

    if reflect.TypeOf(data).Kind() != reflect.Array &&
        reflect.TypeOf(data).Kind() != reflect.Map &&
        reflect.TypeOf(data).Kind() != reflect.Slice {

        return Response{
            Data: data,
        }
    }

    var first, prev, next, last *string
    if page > 1 {
        first = generateUrl(url, 1, limit)
        prev = generateUrl(url, page-1, limit)
    }
    if page > 1 && page > lastPage {
        prev = generateUrl(url, lastPage, limit)
    }
    if page < lastPage {
        next = generateUrl(url, page+1, limit)
        last = generateUrl(url, lastPage, limit)
    }
    if page < 1 {
        next = generateUrl(url, 1, limit)
    }

    response := Response{
        Data: data,
        Pagination: &pagination{
            Limit: limit,
            Page:  page,
            Pages: lastPage,
        },
        Count: &count{
            Returned: reflect.ValueOf(data).Len(),
            Total:    total,
        },
        Links: &links{
            Self:  url.String(),
            First: first,
            Last:  last,
            Next:  next,
            Prev:  prev,
        },
    }

    return response
}

func generateUrl(u *url.URL, page int, limit int) *string {
    pageUrl := url.URL{
        Path: u.Path,
    }

    // copy query values from u
    var v = url.Values{}
    for key, values := range u.Query() {
        for _, value := range values {
            v.Add(key, value)
        }
    }

    v.Set("page", strconv.Itoa(page))
    v.Set("limit", strconv.Itoa(limit))
    pageUrl.RawQuery = v.Encode()

    stringUtl := pageUrl.String()
    return &stringUtl
}
