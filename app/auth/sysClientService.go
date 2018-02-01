package auth

//GetClientService 获取客户端信息
func GetClientService(clientID string) (SysClient, error) {
	return getClient(clientID)
}
