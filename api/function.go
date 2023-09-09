package api

import (
    "net/url"
    "strconv"
)

func ResponseArray(items any, page int, limit int, total int, url *url.URL) Response {

    lastPage := (total / limit) + 1

    var length Count
    if itemsArray, isArray := items.([]any); isArray {
        length = Count{
            Returned: len(itemsArray),
            Total:    total,
        }
    } else if itemsMap, isMap := items.(map[any]any); isMap {
        length = Count{
            Returned: len(itemsMap),
            Total:    total,
        }
    }
    response := Response{
        Data: items,
        Pagination: &Pagination{
            Limit: limit,
            Page:  page,
            Pages: lastPage,
        },
        Count: &length,
        Links: Links{
            Self:  *generateUrl(url, page, limit),
            First: generateUrl(url, 1, limit),
            Last:  generateUrl(url, lastPage, limit),
            Next:  generateUrl(url, page+1, limit),
            Prev:  generateUrl(url, page-1, limit),
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
