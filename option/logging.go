package option

//日志配置
type LoggingOptions struct {
	Level     string `properties:"logging.level,default=info"`
	FilePath  string `properties:"logging.file.path,default=."`
	File      bool   `properties:"logging.file,default=false"`
	FileName  string `properties:"logging.file.name,default=app.log"`
	FileLevel bool   `properties:"logging.file.level,default=false"`
	FileLimit string `properties:"logging.file.limit,default=100MB"`
	FileRate  uint64 `properties:"logging.file.rate,default=30"`
}

type LoggingOptionsFun interface {
	Apply(*LoggingOptions)
}
type funcLoggingOptionsFun struct {
	f func(*LoggingOptions)
}

func (fdo *funcLoggingOptionsFun) Apply(do *LoggingOptions) {
	fdo.f(do)
}

func newFuncLoggingOptions(f func(*LoggingOptions)) *funcLoggingOptionsFun {
	return &funcLoggingOptionsFun{
		f: f,
	}
}
func WithFilePath(s string) LoggingOptionsFun {
	return newFuncLoggingOptions(func(o *LoggingOptions) {
		o.FilePath = s + "\\" + o.FilePath
	})
}
