package main

import (
	"os"

	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// func validParams(ip string, port int) bool {
// 	ipRegex := `^(?:(?:[1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}(?:[1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
// 	regex, err := regexp.Compile(ipRegex)
// 	if err != nil {
// 		fmt.Println("Error compiling regular expression:", err)
// 		return false
// 	}

// 	if !regex.MatchString(ip) {
// 		fmt.Printf("The IP %s is invalid.\n", ip)
// 		return false
// 	}

// 	if port < 1024 || port > 65535 {
// 		fmt.Printf("The port %d is invalid.\n", port)
// 		return false
// 	}

// 	return true
// }

// func main() {
// 	ip := flag.String("ip", "127.0.0.1", "Set IP address of the peer")
// 	port := flag.Int("port", 8080, "Set port for TCP conection")
// 	turn_on := flag.Bool("turn-on", true, "Action to be performed, turn on the network node.")

// 	flag.Parse()

// 	if validParams(*ip, *port) {
// 		// Aquí puedes usar las variables parseadas para ejecutar tu funcionalidad
// 		if *turn_on {
// 			cmdLine := fmt.Sprintf("docker run --rm -it -v /home/leandro/study/test.py:/home/test.py --network spotify-net2 -p %s:%d:8080 mugenfier-api /bin/bash", *ip, *port)
// 			docker := exec.Command(cmdLine)

// 			err := docker.Run()
// 			if err != nil {
// 				fmt.Println("Error al ejecutar el comando:", err)
// 			} else {
// 				fmt.Println("Comando ejecutado exitosamente")
// 			}
// 		} else {
// 			fmt.Println("ip:", *ip)
// 			fmt.Println("port:", *port)
// 		}

//		}
//	}

func main() {
	ctx := context.Background()

	// Crea un cliente de Docker
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Crea el contenedor utilizando la imagen "docker-gs-ping"
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "mugenfier-api",
		Cmd:   []string{"ls"},
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	// Inicia el contenedor
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// Espera a que el contenedor se complete
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	// Obtiene los registros del contenedor
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
