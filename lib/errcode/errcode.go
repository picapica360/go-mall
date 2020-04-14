package errcode

// ErrorCode error code
type ErrorCode int

const (
	// AppInitError [1001000, 1001099]
	AppInitError                   ErrorCode = 1001000 // 系统初始化异常
	AppInitConfigFileNotFoundError ErrorCode = 1001001 // 初始化时加载配置文件出错
	AppInitDecodeConfigFileError   ErrorCode = 1001002 // 初始化时解析配置文件出错

	// SysError [1002000, 1002099]
	SysError                 ErrorCode = 1002000 // 系统级别异常
	SysPortUnavailableError  ErrorCode = 1002001 // 端口不可用
	SysDBConnectionError     ErrorCode = 1002002 // 数据库连接异常
	SysRedisUnavailableError ErrorCode = 1002003 // Redis 服务不可用
	SysESUnavailableError    ErrorCode = 1002004 // ElasticSearch 服务不可用

	// HealthError [1003000, 1003099]
	HealthError ErrorCode = 1003000 // 监控检查异常
)
