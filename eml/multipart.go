package eml

import (
    "bytes"
    "errors"
    "io"
    "mime"
    "mime/multipart"
    "regexp"
    "strings"
)

type Part struct {
    Type    string
    Charset string
    Data    []byte
    Headers map[string][]string
}

// Parse the body of a message, using the given content-type. If the content
// type is multipart, the parts slice will contain an entry for each part
// present; otherwise, it will contain a single entry, with the entire (raw)
// message contents.
func parseBody(contentType string, body []byte) ([]Part, error) {
    mediaType, params, err := mime.ParseMediaType(contentType)

    var parts []Part

    if err != nil {
        return nil, err
    }

    if !strings.HasPrefix(mediaType, "multipart/") {
        part := Part{mediaType, params["charset"], body, nil}
        parts = append(parts, part)

        return parts, nil
    }

    boundary, ok := params["boundary"]
    if !ok {
        return nil, errors.New("multipart specified without boundary")
    }

    r := multipart.NewReader(bytes.NewReader(body), boundary)
    p, err := r.NextPart()

    for err == nil {
        data, _ := io.ReadAll(p) // ignore error
        var subparts []Part
        subparts, err = parseBody(p.Header["Content-Type"][0], data)
        //if err == nil then body have sub multipart, and append him
        if err == nil {
            parts = append(parts, subparts...)
        } else {
            contenttype := regexp.MustCompile("(?is)charset=(.*)").FindStringSubmatch(p.Header["Content-Type"][0])
            charset := "UTF-8"
            if len(contenttype) > 1 {
                charset = contenttype[1]
            }
            part := Part{p.Header["Content-Type"][0], charset, data, p.Header}
            parts = append(parts, part)
        }
        p, err = r.NextPart()
    }
    if err == io.EOF {
        err = nil
    }

    return parts, err
}
