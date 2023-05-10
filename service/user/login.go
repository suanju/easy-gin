package user

import (
	"crypto/md5"
	"easy-gin/global"
	receive "easy-gin/interaction/receive/user"
	response "easy-gin/interaction/response/user"
	"easy-gin/models/common"
	userModel "easy-video-net/models/users"
	"easy-video-net/utils/jwt"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

func Register(data *receive.RegisterReceive) (results interface{}, err error) {
	//判断邮箱是否唯一
	users := new(userModel.User)
	if users.IsExistByField("email", data.Email) {
		return nil, fmt.Errorf("邮箱已被注册")
	}
	//判断验证码是否正确
	verCode, err := global.RedisDb.Get(fmt.Sprintf("%s%s", consts.RegEmailVerCode, data.Email)).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("验证码过期！")
	}

	if verCode != data.VerificationCode {
		return nil, fmt.Errorf("验证码错误")
	}
	//生成密码盐 8 位
	salt := make([]byte, 6)
	for i := range salt {
		salt[i] = jwt.SaltStr[rand.Int63()%int64(len(jwt.SaltStr))]
	}
	password := []byte(fmt.Sprintf("%s%s%s", salt, data.Password, salt))
	passwordMd5 := fmt.Sprintf("%x", md5.Sum(password))
	photo, _ := json.Marshal(common.Img{
		Src: "",
		Tp:  "local",
	})
	registerData := &userModel.User{
		Email:     data.Email,
		Username:  data.UserName,
		Salt:      string(salt),
		Password:  passwordMd5,
		Photo:     photo,
		BirthDate: time.Now(),
	}
	registerRes := registerData.Create()
	if !registerRes {
		return nil, fmt.Errorf("注册失败")
	}
	//注册token
	tokenString := jwt.NextToken(registerData.ID)
	results = response.UserInfoResponse(registerData, tokenString)

	return results, nil

}

func Login(data *receive.LoginReceive) (results interface{}, err error) {
	users := new(userModel.User)
	if !users.IsExistByField("username", data.Username) {
		return nil, fmt.Errorf("账号不存在")
	}
	if !users.IfPasswordCorrect(data.Password) {
		return nil, fmt.Errorf("密码错误")
	}
	//注册token
	tokenString := jwt.NextToken(users.ID)
	userInfo := response.UserInfoResponse(users, tokenString)
	return userInfo, nil
}
