package configx_test

import (
	"fmt"
	"os"
	redisx "sdkService/pkg/store/redis"
	"testing"

	"sdkService/core/jsonx"
	"sdkService/pkg/configx"
	"sdkService/pkg/logger"
	thirdXsolla "sdkService/pkg/xsollax"
	"sdkService/rpc"

	"github.com/stretchr/testify/require"
)

type Config struct {
	BaseURI string `json:"base,default=account/dev/api" config:"base,default=account/dev/api"`
	// 服务端口号
	ServerPort int `json:"port" config:"port,default=8080"`

	// mysql源
	MysqlDNS string `json:"mysqlDNS" config:"mysqlDNS,default=xjy:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"`

	// JwtKey 默认jwt key
	JwtKey string `json:"jwt_key" config:"jwt_key"`

	// JwtIss 默认iss
	JwtIss string `json:"jwt_iss" config:"jwt_iss"`

	// JwtExpireS jwtoken过期时间毫秒数默认7*24h
	JwtExpireS int64 `json:"jwt_expire_ms" config:"jwt_expire_s,default=604800"`

	// 实人验证码失效时间 默认1分钟
	CaptchaExpireS int64 `json:"captcha_expire_s" config:"captcha_expire_s,default=60"`
	// 验证码失效时间 默认3分钟
	VerifyCodeExpireS int64 `json:"vc_expire_s" config:"vc_expire_s,default=180"`

	// HTTPProxyURL http代理地址
	HTTPProxyURL string `json:"http_proxy_url" config:"HTTPProxyURL"`

	LogLevel int `json:"log_level" config:"log_level,default=1"`

	// 日志配置
	logger.Config
	XsollaConfig thirdXsolla.XsollaConfig `json:"XsollaConfig" config:"XsollaConfig"`

	UserMin3rd int `json:"user_min_3rd" config:"user_min_3rd"`

	Rpc rpc.Config `json:"rpc" config:"rpc"`

	// redis 配置
	Redis redisx.Config `json:"redis" config:"redis"`

	// event redis addr 用于异步消息通知redis队列
	RedisAddr string `json:"redisAddr" config:"redisAddr,default=redis:6379"`
	RedisDB   int    `json:"redisDb" config:"redisDb,default=0"`

	// 并发数 默认20
	TaskConcurrency int `json:"tc" config:"tc,default=20"`

	// 后台通知任务配置
	// 通知频率 默认10秒执行一次
	NotifyLoopInterval string `json:"nli" config:"nli,default=10s"`
	// 通知并发数 默认20
	NotifyConcurrency int `json:"nc" config:"nc,default=20"`
	// 钉钉消息通知accessToken
	DingTalkAccessToken string `json:"dtAccessToken" config:"dtAccessToken,default=6e8baea96cbb0b210608e73666387d4064570ed3d434d266ba2536c2ffd70711"`

	// 游戏侧后端接口域名
	GameServerApiDomain string `json:"game_server_api_domain" config:"GameServerApiDomain,default=http://localhost:8081"`
}

func TestConfigLoad(t *testing.T) {
	var c Config
	require.NoError(t, os.Setenv("IFUN_RPC_SERVER_NAME", "test"))
	configx.MustLoad("../../app/account/etc/config.yaml", &c)
	fmt.Println(jsonx.MarshalToStringWC(c))
}
