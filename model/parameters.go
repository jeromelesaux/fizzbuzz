package model

import (
	"errors"
	"net/http"
	"strconv"
)

type Parameters struct {
	Int1  int64  `json:"int1"`
	Int2  int64  `json:"int2"`
	Limit int64  `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Hits  int64  `json:"hits"`
}

func parseInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return i, err
	}
	return i, nil
}

func ParseParameters(r *http.Request) (p *Parameters, err error) {
	p = &Parameters{}
	i1, err := parseInt64(r.URL.Query().Get("int1"))
	if err != nil {
		return
	}

	if i1 == 0 {
		return p, errors.New("int1 must differ from 0")
	}
	p.Int1 = i1

	i2, err := parseInt64(r.URL.Query().Get("int2"))
	if err != nil {
		return
	}

	if i2 == 0 {
		return p, errors.New("int2 must differ from 0")
	}
	p.Int2 = i2

	limit, err := parseInt64(r.URL.Query().Get("limit"))
	if err != nil {
		return
	}
	p.Limit = limit

	p.Str1 = r.URL.Query().Get("str1")
	p.Str2 = r.URL.Query().Get("str2")
	return p, nil
}
