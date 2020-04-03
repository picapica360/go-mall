package errcode

const (
	// AppInit [1001000, 1000000]
	AppInit                      = 1001000
	AppInitConfigFileNotFound    = 1001001 // 初始化时加载配置文件出错
	AppInitDecodeConfigFileError = 1001002 // 初始化时解析配置文件出错
)
