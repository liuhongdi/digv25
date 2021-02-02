package sign

import (
	"errors"
	"github.com/liuhongdi/digv25/pkg/md5"
	"strconv"
	"time"
)

// 验证签名
func VerifySign(appId,timestamp,sign,nonce,serverAppId,serverAppSecret,serverApiVersion string,signExpire int64) error {

    //验证appid
	if (serverAppId != appId) {
		return errors.New("appId Error")
	}

	// 验证过期时间
	timestampNow := time.Now().Unix()
	//fmt.Println("now:")
	//fmt.Println(timestampNow)
	tsInt, _  := strconv.ParseInt(timestamp, 10, 64)
	//fmt.Println("timeint:")
	//fmt.Println(tsInt)
	if tsInt > timestampNow || timestampNow - tsInt >= signExpire {
		return errors.New("timestamp Error")
	}

	//得到正确的sign供检验用
	origin := appId + serverAppSecret + timestamp + nonce + serverApiVersion;
	signEcrypt := md5.MD5(origin);
	if (signEcrypt != sign) {
		return errors.New("sign Error")
	}
    //未出错,返回
	return nil
}
