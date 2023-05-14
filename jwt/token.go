package jwt

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/MicahParks/keyfunc"
    "github.com/d3code/zlog"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "strings"
    "time"
)

func GetValidTokenFromRequst(c *gin.Context, jwksUrl string, cacheExpiration time.Duration) (*jwt.Token, *jwt.MapClaims) {
    headerToken := GetTokenStringFromRequest(c)
    if headerToken == nil {
        return nil, nil
    }

    jwksString := GetJwks(jwksUrl, cacheExpiration)
    if jwksString == nil {
        zlog.Log.Error("Missing JWKS")
        return nil, nil
    }

    jwksJSON := json.RawMessage(*jwksString)
    jwks, jsonError := keyfunc.NewJSON(jwksJSON)
    if jsonError != nil {
        zlog.Log.Error(jsonError)
        return nil, nil
    }

    token, jwtParseError := jwt.Parse(*headerToken, jwks.Keyfunc)
    if jwtParseError != nil {
        zlog.Log.Error(jwtParseError)
        return nil, nil
    }

    claims, claimsError := GetClaims(*token)
    if claimsError != nil {
        zlog.Log.Error(claimsError)
        return token, nil
    }

    return token, &claims
}

func GetClaims(token jwt.Token) (jwt.MapClaims, error) {
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, errors.New("token not valid")
}

func GetTokenStringFromRequest(c *gin.Context) *string {

    authorizationHeader := c.GetHeader("Authorization")

    if authorizationHeader == "" {
        authorizationHeader = fmt.Sprintf("Bearer %s", c.Query("token"))
    }

    if authorizationHeader == "" {
        cookie, err := c.Cookie("token")
        if err == nil {
            authorizationHeader = fmt.Sprintf("Bearer %s", cookie)
        }
    }

    tokenString := strings.Split(authorizationHeader, " ")

    if len(tokenString) != 2 || tokenString[0] != "Bearer" {
        zlog.Log.Warnw("Invalid Authorization header", "value", authorizationHeader)
        return nil
    }

    return &tokenString[1]
}
