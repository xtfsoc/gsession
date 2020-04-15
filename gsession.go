package gsession

func init() {
	COOKIEJ = make(map[string]string)
}

type session struct {
	gsessionObject
	Cookie //组合People
}

func Session() session {
	return session{sessionInit(), Cookie{}}
}

type gsessionObject interface {
	GET(o Options) (Response, error)
	POST(o Options) (Response, error)
	PUT(o Options) (Response, error)
	DELETE(o Options) (Response, error)
	OPTIONS(o Options) (Response, error)
	//GetAllCookies() map[string]string
}

type gsob struct {
}

func sessionInit() gsessionObject {
	var obj gsessionObject
	obj = gsob{}
	return obj
}
