//这个文件用于设置系统服务
package server

func Run(server *Server) error {
	return server.Start()
}
