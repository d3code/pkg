package api

import (
    "fmt"
    "net/url"
    "strconv"
)

func GetPageLimit(values url.Values, limitDefault int) (int, int, error) {
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
