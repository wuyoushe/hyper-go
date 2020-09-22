package mdw

import "github.com/wuyoushe/hyper-go/service/jwt-iris"

/**
 * jwt中间件
 */
func JwtHandler() *jwt.Middleware {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret123123"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
