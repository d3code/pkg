package eml

import (
    "testing"
)

type parseBodyTest struct {
    ct    string
    body  []byte
    parts []Part
}

var parseBodyTests = []parseBodyTest{
    {
        ct:   "text/plain",
        body: []byte(`This is some text.`),
        parts: []Part{
            {"text/plain", "UTF-8", []byte("This is some text."), nil},
        },
    },
    {
        ct: "multipart/alternative; boundary=90e6ba1efd30b0013a04b8d4970f",
        body: []byte(`--90e6ba1efd30b0013a04b8d4970f
Content-Type: text/plain; charset=ISO-8859-1

Some text.
--90e6ba1efd30b0013a04b8d4970f
Content-Type: text/html; charset=ISO-8859-1
Content-Transfer-Encoding: quoted-printable

Some other text.
--90e6ba1efd30b0013a04b8d4970f--
`),
        parts: []Part{
            {
                "text/plain; charset=ISO-8859-1",
                "ISO-8859-1",
                []byte("Some text."),
                map[string][]string{
                    "Content-Type": {
                        "text/plain; charset=ISO-8859-1",
                    },
                },
            },
            {
                "text/html; charset=ISO-8859-1",
                "ISO-8859-1",
                []byte("Some other text."),
                map[string][]string{
                    "Content-Type": {
                        "text/html; charset=ISO-8859-1",
                    },
                    "Content-Transfer-Encoding": {
                        "quoted-printable",
                    },
                },
            },
        },
    },
}

func TestParseBody(t *testing.T) {
    for _, testData := range parseBodyTests {
        _, err := parseBody(testData.ct, testData.body)

        if err != nil {
            t.Errorf("parseBody returned error for %#v: %#v", testData, err)
        }

        //else if !reflect.DeepEqual(parts, testData.parts) {
        //
        //    log.Printf("Parsed   :  %#v", parts)
        //    log.Printf("Original :  %#v", testData.parts)
        //
        //    t.Error("Parsed not match")
        //}
    }
}
