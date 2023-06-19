package id

import (
	"github.com/rxlwin/g-pkg/jm/sm/md5"
)

func GetAppidAndSecret(id int64) (string, string) {
	if id == 0 {
		id = GetId()
	}
	appid0 := Int64ToStr(id)
	appid1 := md5.ToStr32(appid0)
	appid := appid1[5:14]
	secret := md5.ToStr32(appid0 + appid)
	return appid, secret
}
