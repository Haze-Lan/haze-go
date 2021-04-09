package option

//日志配置
type LoggingOptions struct {
	Level     map[string]string `properties:"logging.level,default=root=info"`
	FilePath  string            `properties:"logging.file.path,default=."`
	File      bool              `properties:"logging.file,default=false"`
	FileName  string            `properties:"logging.file.name,default=app.log"`
	FileLevel bool              `properties:"logging.file.level,default=false"`
	FileLimit string            `properties:"logging.file.limit,default=100MB"`
	FileRate  uint64            `properties:"logging.file.rate,default=30"`
}
