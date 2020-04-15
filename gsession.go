package gsession

func init() {
	COOKIEJ = make(map[string]string)
}

type session struct {
	Cookie Cookie
	gsessionAction
}

func Session() session {
	// 新增session要清空COOKIEJ
	COOKIEJ = make(map[string]string)

	sessionInit := func() gsessionAction {
		var ga gsessionAction
		ga = gsessionObject{}
		return ga
	}
	return session{Cookie{}, sessionInit()}
}

type gsessionAction interface {
	GET(o Options) (Response, error)
	POST(o Options) (Response, error)
	PUT(o Options) (Response, error)
	DELETE(o Options) (Response, error)
	OPTIONS(o Options) (Response, error)
	// GetAllCookies() map[string]string
}

type gsessionObject struct{}

//func sessionInit() gsessionAction {
//	var ga gsessionAction
//	ga = gsessionObject{}
//	return ga
//}
