package eml

import (
    "bytes"
    "crypto/sha1"
    "encoding/base64"
    "errors"
    "fmt"
    "github.com/d3code/pkg/eml/decoder"
    "io"
    "mime/quotedprintable"
    "regexp"
    "strings"
    "time"
)

var urlEncoding = base64.URLEncoding

type HeaderInfo struct {
    FullHeaders []Header // all headers
    OptHeaders  []Header // unprocessed headers

    MessageId   string
    Id          string
    Date        time.Time
    From        []Address
    Sender      Address
    ReplyTo     []Address
    To          []Address
    Cc          []Address
    Bcc         []Address
    Subject     string
    Comments    []string
    Keywords    []string
    ContentType string

    InReply    []string
    References []string
}

type Message struct {
    HeaderInfo
    Body        []byte
    Text        string
    Html        string
    Attachments []Attachment
    Parts       []Part
}

type Attachment struct {
    Filename string
    Data     []byte
}

type Header struct {
    Key, Value string
}

func Parse(s []byte) (*Message, error) {
    rawMessage, err := ParseRaw(s)
    if err != nil {
        return nil, err
    }

    return Process(rawMessage)
}

func Process(rawMessage RawMessage) (message *Message, err error) {

    message.FullHeaders = []Header{}
    message.OptHeaders = []Header{}

    for _, rawHeader := range rawMessage.RawHeaders {

        header := Header{string(rawHeader.Key), string(rawHeader.Value)}
        message.FullHeaders = append(message.FullHeaders, header)

        switch string(rawHeader.Key) {

        case `Content-Type`:
            message.ContentType = string(rawHeader.Value)

        case `Message-ID`:
            v := bytes.Trim(rawHeader.Value, `<>`)
            message.MessageId = string(v)
            message.Id = makeId(v)

        case `In-Reply-To`:
            ids := strings.Fields(string(rawHeader.Value))
            for _, id := range ids {
                message.InReply = append(message.InReply, strings.Trim(id, `<> `))
            }

        case `References`:
            ids := strings.Fields(string(rawHeader.Value))
            for _, id := range ids {
                message.References = append(message.References, strings.Trim(id, `<> `))
            }

        case `Date`:
            message.Date = ParseDate(string(rawHeader.Value))

        case `From`:
            message.From, err = parseAddressList(rawHeader.Value)

        case `Sender`:
            message.Sender, err = ParseAddress(rawHeader.Value)

        case `Reply-To`:
            message.ReplyTo, err = parseAddressList(rawHeader.Value)

        case `To`:
            message.To, err = parseAddressList(rawHeader.Value)

        case `Cc`:
            message.Cc, err = parseAddressList(rawHeader.Value)

        case `Bcc`:
            message.Bcc, err = parseAddressList(rawHeader.Value)

        case `Subject`:
            subject, err := decoder.Parse(rawHeader.Value)
            if err != nil {
                fmt.Println("Failed decode subject", err)
            }

            message.Subject = string(subject)

        case `Comments`:
            message.Comments = append(message.Comments, string(rawHeader.Value))

        case `Keywords`:
            ks := strings.Split(string(rawHeader.Value), ",")
            for _, k := range ks {
                message.Keywords = append(message.Keywords, strings.TrimSpace(k))
            }

        default:
            message.OptHeaders = append(message.OptHeaders, header)
        }
        if err != nil {
            return
        }
    }

    if message.Sender == nil && len(message.From) > 0 {
        message.Sender = message.From[0]
    }

    if message.ContentType != "" {

        parts, er := parseBody(message.ContentType, rawMessage.Body)
        if er != nil {
            err = er
            return
        }

        for _, part := range parts {
            switch {
            case strings.Contains(part.Type, "text/plain"):

                data, err := decoder.UTF8(part.Charset, part.Data)
                if err != nil {
                    message.Text = string(part.Data)
                } else {
                    message.Text = string(data)
                }

            case strings.Contains(part.Type, "text/html"):

                data, err := decoder.UTF8(part.Charset, part.Data)
                if err != nil {
                    message.Html = string(part.Data)
                } else {
                    message.Html = string(data)
                }

            default:

                if cd, ok := part.Headers["Content-Disposition"]; ok {
                    if strings.Contains(cd[0], "attachment") {
                        filename := regexp.MustCompile("(?msi)name=\"(.*?)\"").FindStringSubmatch(cd[0]) //.FindString(cd[0])
                        if len(filename) < 2 {
                            fmt.Println("failed get filename from header content-disposition")
                            break
                        }

                        dfilename, err := decoder.Parse([]byte(filename[1]))
                        if err != nil {
                            fmt.Println("Failed decode filename of attachment", err)
                        } else {
                            filename[1] = string(dfilename)
                        }

                        if encoding, ok := part.Headers["Content-Transfer-Encoding"]; ok {
                            switch strings.ToLower(encoding[0]) {
                            case "base64":
                                part.Data, er = base64.StdEncoding.DecodeString(string(part.Data))
                                if er != nil {
                                    fmt.Println(er, "failed decode base64")
                                }
                            case "quoted-printable":
                                part.Data, _ = io.ReadAll(quotedprintable.NewReader(bytes.NewReader(part.Data)))
                            }
                        }
                        message.Attachments = append(message.Attachments, Attachment{filename[1], part.Data})

                    }
                }
            }
        }

        message.Parts = parts
        message.ContentType = parts[0].Type
        message.Text = string(parts[0].Data)

    } else {
        message.Text = string(rawMessage.Body)
    }

    return
}

type RawHeader struct {
    Key, Value []byte
}

type RawMessage struct {
    RawHeaders []RawHeader
    Body       []byte
}

func isWSP(b byte) bool {
    return b == ' ' || b == '\t'
}

func ParseRaw(s []byte) (m RawMessage, e error) {
    // parser states
    const (
        READY = iota
        HKEY
        HVWS
        HVAL
    )

    const (
        CR = '\r'
        LF = '\n'
    )

    CRLF := []byte{CR, LF}

    state := READY
    kstart, kend, vstart := 0, 0, 0
    done := false

    m.RawHeaders = []RawHeader{}

    for i := 0; i < len(s); i++ {
        b := s[i]
        switch state {
        case READY:
            if b == CR && i < len(s)-1 && s[i+1] == LF {
                // we are at the beginning of an empty header
                m.Body = s[i+2:]
                done = true
                goto Done
            }
            if b == LF {
                m.Body = s[i+1:]
                done = true
                goto Done
            }
            // otherwise this character is the first in a header
            // key
            kstart = i
            state = HKEY
        case HKEY:
            if b == ':' {
                kend = i
                state = HVWS
            }
        case HVWS:
            if !isWSP(b) {
                vstart = i
                state = HVAL
            }
        case HVAL:
            if b == CR && i < len(s)-2 && s[i+1] == LF && !isWSP(s[i+2]) {
                v := bytes.Replace(s[vstart:i], CRLF, nil, -1)
                hdr := RawHeader{s[kstart:kend], v}
                m.RawHeaders = append(m.RawHeaders, hdr)
                state = READY
                i++
            } else if b == LF && i < len(s)-1 && !isWSP(s[i+1]) {
                v := bytes.Replace(s[vstart:i], CRLF, nil, -1)
                hdr := RawHeader{s[kstart:kend], v}
                m.RawHeaders = append(m.RawHeaders, hdr)
                state = READY
            }
        }
    }
Done:
    if !done {
        e = errors.New("unexpected EOF")
    }
    return
}

func makeId(headerBytes []byte) string {
    h := sha1.New()
    h.Write(headerBytes)

    hash := h.Sum(nil)
    encodeToString := urlEncoding.EncodeToString(hash)

    return encodeToString[0:20]
}
