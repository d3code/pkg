package jwt

import (
    "github.com/patrickmn/go-cache"
    "io"
    "net/http"
    "time"
)

var jwksCache *cache.Cache

func init() {
    jwksCache = cache.New(1*time.Millisecond, 1*time.Millisecond)
}

func GetJwks(jwksUrl string) *string {
    if cachedJwks, found := jwksCache.Get(jwksUrl); found {
        cachedJwksString := cachedJwks.(string)
        return &cachedJwksString
    }

    jwksString := getRequest(jwksUrl)
    if jwksString != nil {
        jwksCache.Set(jwksUrl, *jwksString, cache.DefaultExpiration)
    }

    return jwksString
}

func getRequest(jwksUrl string) *string {
    response, err := http.Get(jwksUrl)
    if err != nil {
        return nil
    }

    body, err := io.ReadAll(response.Body)
    if err != nil {
        return nil
    }

    stringResponse := string(body)
    return &stringResponse
}
