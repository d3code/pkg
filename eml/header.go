// Header parsing functionality.

package eml

import (
    "strings"
)

func split(ts []token, s token) [][]token {
    var r [][]token
    l := 0
    for i, t := range ts {
        if string(t) == string(s) {
            r = append(r, ts[l:i])
            l = i + 1
        }
    }
    if l != len(ts) {
        r = append(r, ts[l:])
    }
    return r
}

// BUG: We don't currently support domain literals with commas.
func parseAddressList(s []byte) ([]Address, error) {
    var al []Address
    ts, e := tokenize(s)
    if e != nil {
        return al, e
    }
    tts := split(ts, []byte{','})
    for _, ts := range tts {
        a, e := parseAddress(ts)
        if e != nil {
            return al, e
        }
        al = append(al, a)
    }
    return al, nil
}

func decodeRFC2047(word string) string {
    if strings.HasPrefix(word, "=?") && strings.HasSuffix(word, "?=") && strings.Count(word, "?") == 4 {
        return ""
    }
    return word
}
