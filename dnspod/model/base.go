package model

type DnspodCommonArg struct {
	LoginToken   string `query:"login_token"`
	Format       string `query:"format"`
	Lang         string `query:"lang"`
	ErrorOnEmpty string `query:"error_on_empty"`
	UserId       string `query:"user_id,omitempty"`
}
