package jwt

import (
    "github.com/d3code/zlog"
    "github.com/patrickmn/go-cache"
    "io"
    "net/http"
    "sync"
    "time"
)

var (
    jwksCache *cache.Cache
    onceJwks  sync.Once
)

func GetJwks(jwksUrl string, cacheExpiration time.Duration) *string {
    onceJwks.Do(func() {
        jwksCache = cache.New(cacheExpiration, cacheExpiration)
    })

    if cachedJwks, found := jwksCache.Get(jwksUrl); found {
        cachedJwksString := cachedJwks.(string)
        return &cachedJwksString
    }

    // Call JWKS endpoint
    jwksString := getRequest(jwksUrl)
    if jwksString != nil {
        jwksCache.Set(jwksUrl, *jwksString, cache.DefaultExpiration)
    }

    return jwksString
}

func getRequest(jwksUrl string) *string {
    response, err := http.Get(jwksUrl)
    if err != nil {
        zlog.Log.Error(err)
        return nil
    }

    body, err := io.ReadAll(response.Body)
    if err != nil {
        zlog.Log.Error(err)
        return nil
    }

    stringResponse := string(body)
    return &stringResponse
}
