package gsession

import "net/http"

var COOKIEJ map[string]string

/*
通用setCookie
*/
func setCookie(c []*http.Cookie) {

	for _, v := range c {
		// v: *http.CookieJJ
		//fmt.Println("v.Value:", v.Value)
		//fmt.Println("v.Domain:", v.Domain) // string
		//fmt.Println("v.MaxAge:", v.MaxAge) // int
		COOKIEJ[v.Name] = v.Value
	}
}
