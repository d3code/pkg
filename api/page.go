package api

import (
    "fmt"
    "net/http"
    "net/url"
    "strconv"
)

// GetPageLimit returns the page and limit values from the query params
// If the page or limit is not set, the default values are used
func GetPageLimit(req *http.Request, limitDefault int) (int, int, error) {
    values, err := url.ParseQuery(req.URL.RawQuery)

    paramPage := values.Get("page")
    paramLimit := values.Get("limit")

    if paramPage == "" {
        paramPage = "1"
    }

    if paramLimit == "" {
        paramLimit = strconv.Itoa(limitDefault)
    }

    limit, err := strconv.Atoi(paramLimit)
    if err != nil || limit < 0 {
        return 0, 0, fmt.Errorf("invalid query param 'limit' [ %s ]", paramLimit)
    }

    page, err := strconv.Atoi(paramPage)
    if err != nil || page < 1 {
        return 0, 0, fmt.Errorf("invalid query param 'page' [ %s ]", paramPage)
    }

    return page, limit, nil
}
