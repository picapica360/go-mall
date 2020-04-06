package errcode

const (
	// AppInitError [1001000, 1001099]
	AppInitError                   = 1001000 // 系统初始化异常
	AppInitConfigFileNotFoundError = 1001001 // 初始化时加载配置文件出错
	AppInitDecodeConfigFileError   = 1001002 // 初始化时解析配置文件出错

	// SysError [1002000, 1002099]
	SysError                 = 1002000 // 系统级别异常
	SysPortUnavailableError  = 1002001 // 端口不可用
	SysDBConnectionError     = 1002002 // 数据库连接异常
	SysRedisUnavailableError = 1002003 // Redis 服务不可用
	SysESUnavailableError    = 1002004 // ElasticSearch 服务不可用
)
