package decoder

import (
    "bytes"
    "github.com/d3code/pkg/zlog"
    "golang.org/x/text/encoding/charmap"
    "golang.org/x/text/transform"
    "io"
    "regexp"
    "strings"

    "encoding/base64"
    "mime/quotedprintable"
)

func UTF8(cs string, data []byte) ([]byte, error) {

    cs = strings.Replace(cs, "\"", "", -1)
    cs = strings.Replace(cs, "'", "", -1)

    if strings.ToUpper(cs) == "UTF-8" {
        return data, nil
    }

    if strings.ToUpper(cs) == "US-ASCII" {
        return data, nil
    }

    if strings.ToUpper(cs) == "WINDOWS-1251" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1251.NewDecoder()))
    }
    if strings.ToUpper(cs) == "KOI8-R" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.KOI8R.NewDecoder()))
    }
    if strings.ToUpper(cs) == "ISO-8859-1" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.ISO8859_1.NewDecoder()))
    }
    if strings.ToUpper(cs) == "ISO-8859-2" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.ISO8859_2.NewDecoder()))
    }
    if strings.ToUpper(cs) == "ISO-8859-5" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.ISO8859_5.NewDecoder()))
    }
    if strings.ToUpper(cs) == "ISO-8859-7" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.ISO8859_7.NewDecoder()))
    }
    if strings.ToUpper(cs) == "ISO-8859-9" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.ISO8859_9.NewDecoder()))
    }
    if strings.ToUpper(cs) == "ISO-8859-15" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.ISO8859_15.NewDecoder()))
    }
    if strings.ToUpper(cs) == "MACINTOSH" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Macintosh.NewDecoder()))
    }
    if strings.ToUpper(cs) == "IBM866" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.CodePage866.NewDecoder()))
    }
    if strings.ToUpper(cs) == "IBM855" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.CodePage855.NewDecoder()))
    }
    if strings.ToUpper(cs) == "IBM852" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.CodePage852.NewDecoder()))
    }
    if strings.ToUpper(cs) == "IBM437" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.CodePage437.NewDecoder()))
    }
    if strings.ToUpper(cs) == "IBM850" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.CodePage850.NewDecoder()))
    }
    if strings.ToUpper(cs) == "IBM858" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.CodePage858.NewDecoder()))
    }
    if strings.ToUpper(cs) == "IBM862" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.CodePage862.NewDecoder()))
    }
    if strings.ToUpper(cs) == "IBM864" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1251.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1252" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1252.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1253" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1253.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1254" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1254.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1255" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1255.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1256" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1256.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1257" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1257.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1258" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1258.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-874" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows874.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1250" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1250.NewDecoder()))
    }
    if strings.ToUpper(cs) == "WINDOWS-1251" {
        return io.ReadAll(transform.NewReader(bytes.NewReader(data), charmap.Windows1251.NewDecoder()))
    }

    zlog.Log.Errorf("Unknown charset: %s", strings.ToUpper(cs))
    return data, nil
}

func Parse(bstr []byte) ([]byte, error) {
    var err error
    strs := regexp.MustCompile("^=\\?(.*?)\\?(.*?)\\?(.*)\\?=$").FindAllStringSubmatch(string(bstr), -1)

    if len(strs) > 0 && len(strs[0]) == 4 {
        c := strs[0][1]
        e := strs[0][2]
        dstr := strs[0][3]

        bstr, err = Decode(e, []byte(dstr))
        if err != nil {
            return bstr, err
        }

        return UTF8(c, bstr)
    }
    return bstr, err

}

func Decode(e string, bstr []byte) ([]byte, error) {
    var err error
    switch strings.ToUpper(e) {
    case "Q":
        bstr, err = io.ReadAll(quotedprintable.NewReader(bytes.NewReader(bstr)))
    case "B":
        bstr, err = base64.StdEncoding.DecodeString(string(bstr))
    default:
        //not set encoding type

    }
    return bstr, err
}
