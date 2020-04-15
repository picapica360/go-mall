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

	// InputParamsError [2001000, 2001999]
	InputParamsError        ErrorCode = 2001000 // 输入参数错误
	InputParamsMissingError ErrorCode = 2001001 // 输入参数缺失错误

	DataTypeError                ErrorCode = 2002020 // 数据类型验证错误
	DataTypeIntegerError         ErrorCode = 2002021 // 数据类型必须为整型
	DataTypeNumberError          ErrorCode = 2002022 // 数据类型必须为数字类型
	DataTypeDateError            ErrorCode = 2002023 // 数据类型必须为date类型
	DataTypeDateTimeError        ErrorCode = 2002024 // 数据类型必须为datetime类型
	DataNumberError              ErrorCode = 2002030 // 数据值错误
	DataNumberPositiveError      ErrorCode = 2002031 // 数据值必须为正数
	DataNumberNonNegativeError   ErrorCode = 2002032 // 数据值必须为非负数
	DataRegexError               ErrorCode = 2002040 // 数据规则匹配验证错误
	DataRegexMailError           ErrorCode = 2002041 // 数据必须为邮件格式
	DataRegexMobileError         ErrorCode = 2002042 // 数据必须为移动电话格式
	DataRegexIdentityError       ErrorCode = 2002043 // 数据必须为身份证格式
	DataRegexEnNumUnderlineError ErrorCode = 2002044 // 数据必须为英文、数字或下划线的一种或几种组合
)
