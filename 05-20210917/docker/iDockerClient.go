package docker

type iDockerClient interface {
	SetDocClient()
	SetIsAllContainer(bool)
	ListAll()
	StartContainer(string) error
	StopContainer(string) error
}

// Xây dựng đối tượng iDockerClient
func BuildIDockerClient(isAllContainer bool) iDockerClient {
	var docClient iDockerClient = GetInstanceSync()

	docClient.SetDocClient()
	docClient.SetIsAllContainer(isAllContainer)

	return docClient
}
