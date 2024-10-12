package gormx

type Config struct {
	// mysql源
	MysqlDNS string `json:"mysqlDNS" config:"mysqlDNS,default=debian-sys-maint:l3BkmCtvLyjPCMlw@tcp(127.0.0.1:3306)/comment?charset=utf8mb4&parseTime=True&loc=Local"`
	// 是否打印mysql日志
	IsDebug bool `json:"isDebug" config:"isDebug,default=true"`
}
