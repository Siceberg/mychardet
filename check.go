package chardet

import "sort"

type detect interface {
	String() string
	Feed(byte) bool
	Priority() float64
}

func check(data []byte, lst []detect) []detect {
	for _, c := range data {
		for i, l := 0, len(lst); i < l; {
			if !lst[i].Feed(c) {
				copy(lst[i:], lst[i+1:])
				l--
				lst = lst[:l]
			} else {
				i++
			}
		}
	}
	if len(lst) == 0 {
		return nil
	}
	return lst
}

// Mostlike - 本函数返回文本最可能的编码格式
func Mostlike(data []byte) string {
	if s := checkbom(data); s != "" {
		return s
	}
	lb := check(data, []detect{&utf8{}, &utf16BE{}, &utf16LE{}, &utf32BE{}, &utf32LE{}, &hzgb2312{}})
	if len(lb) > 0 {
		x, y := -1, -100.0
		for i, l := range lb {
			if r := l.Priority(); y < r {
				x, y = i, r
			}
		}
		return lb[x].String()
	}
	lp := check(data, []detect{&gbk{}, &big5{}, &eucJP{}, &shiftJIS{}, &iso2022JP{}, &eucKR{}, &gb18030{}})
	if len(lp) > 0 {
		x, y := -1, -100.0
		for i, l := range lp {
			if r := l.Priority(); y < r {
				x, y = i, r
			}
		}
		return lp[x].String()
	}
	return ""
}

// Mostlikein - 本函数返回指定编码格式中文本最可能的编码格式
func Mostlikein(data []byte, detectNames []string) string {
	var lbDetects, lpDetects []detect

	for _, detectName := range detectNames {
		switch detectName {
		default:
			panic("unexpected type %T")
		case "utf8":
			lbDetects = append(lbDetects, &utf8{})
		case "utf16BE":
			lbDetects = append(lbDetects, &utf16BE{})
		case "utf16LE":
			lbDetects = append(lbDetects, &utf16LE{})
		case "utf32BE":
			lbDetects = append(lbDetects, &utf32BE{})
		case "utf32LE":
			lbDetects = append(lbDetects, &utf32LE{})
		case "hzgb2312":
			lbDetects = append(lbDetects, &hzgb2312{})
		case "gbk":
			lpDetects = append(lpDetects, &gbk{})
		case "big5":
			lpDetects = append(lpDetects, &big5{})
		case "eucJP":
			lpDetects = append(lpDetects, &eucJP{})
		case "shiftJIS":
			lpDetects = append(lpDetects, &shiftJIS{})
		case "iso2022JP":
			lpDetects = append(lpDetects, &iso2022JP{})
		case "eucKR":
			lpDetects = append(lpDetects, &eucKR{})
		case "gb18030":
			lpDetects = append(lpDetects, &gb18030{})
		}
	}

	if s := checkbom(data); s != "" {
		return s
	}
	lb := check(data, lbDetects)
	if len(lb) > 0 {
		x, y := -1, -100.0
		for i, l := range lb {
			if r := l.Priority(); y < r {
				x, y = i, r
			}
		}
		return lb[x].String()
	}
	lp := check(data, lpDetects)
	if len(lp) > 0 {
		x, y := -1, -100.0
		for i, l := range lp {
			if r := l.Priority(); y < r {
				x, y = i, r
			}
		}
		return lp[x].String()
	}
	return ""
}

// Possible - 本函数返回文本所有可能的编码格式，可能性越高越靠前
func Possible(data []byte) []string {
	if s := checkbom(data); s != "" {
		return []string{s}
	}
	lb := check(data, []detect{
		&utf8{}, &utf16BE{}, &utf16LE{}, &utf32BE{}, &utf32LE{},
		&hzgb2312{}, &gbk{}, &big5{}, &gb18030{},
		&eucKR{}, &eucJP{}, &shiftJIS{}, &iso2022JP{}})
	if l := len(lb); l > 0 {
		x := make(tks, l)
		for i, e := range lb {
			x[i] = tk{e.Priority(), e.String()}
		}
		sort.Stable(sort.Reverse(x))
		s := make([]string, l)
		for i, y := range x {
			s[i] = y.s
		}
		return s
	}
	return nil
}

type tk struct {
	f float64
	s string
}

type tks []tk

func (t tks) Len() int           { return len(t) }
func (t tks) Less(i, j int) bool { return t[i].f < t[j].f }
func (t tks) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
