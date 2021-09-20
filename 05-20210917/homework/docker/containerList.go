package docker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	docClient      client.Client
	isAllContainer bool
}

// Sử dụng sync để tạo mới đối tượng
var once sync.Once
var instance *DockerClient

// Hàm sử dụng sync để tạo mới đối tượng
func GetInstanceSync() *DockerClient {
	once.Do(func() {
		instance = &DockerClient{
			docClient:      client.Client{},
			isAllContainer: true,
		}
	})
	return instance
}

// Hàm lấy giá trị cho tham số docClient
func (dc *DockerClient) SetDocClient() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	dc.docClient = *cli
}

// Hàm lấy giá trị cho tham số isAllContainer
func (dc *DockerClient) SetIsAllContainer(isAllContainer bool) {
	dc.isAllContainer = isAllContainer
}

// In danh sách các container
func (dc *DockerClient) ListAll() {
	listOption := types.ContainerListOptions{}
	listOption.All = dc.isAllContainer // Lấy tất cả các container hoặc chỉ running container

	containers, err := dc.docClient.ContainerList(context.Background(), listOption)
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%12s : %-20s : %-20s : %-8s : ", container.ID[:10],
			container.Names[0],
			container.Image,
			container.State)
		if len(container.Ports) > 0 {
			fmt.Printf("%d:%d\n", container.Ports[0].PublicPort, container.Ports[0].PrivatePort)
		} else {
			fmt.Printf("\n")
		}
	}
}

// Hàm khởi động một Container theo ID
func (dc *DockerClient) StartContainer(conID string) error {
	return dc.docClient.ContainerStart(context.Background(), conID, types.ContainerStartOptions{})
}

// Hàm dừng một Container theo ID
func (dc *DockerClient) StopContainer(conID string) error {
	timeout, _ := time.ParseDuration("5s")
	return dc.docClient.ContainerStop(context.Background(), conID, &timeout)
}
