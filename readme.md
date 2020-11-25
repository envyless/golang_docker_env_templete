# Golang Dev Container with Docker

생성일: 2020년 11월 17일 오후 1:14
태그: #Golang, #devcontainer.json, #개발환경

# golang_docker_env_templete

golang, vscode, devcontainer, work with docker containers

![Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled.png](Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled.png)

Vscode에서 해당 Remote Development를 설치한다.

devcontainer.json 에서는 포트포워딩, 앱 포트 등등 퍼블리시 포트를 설정 할 수 있다.

```cpp
{
	"name": "Existing Dockerfile",

	// Sets the run context to one level up instead of the .devcontainer folder.
	"context": "..",

	// Update the 'dockerFile' property if you aren't using the standard 'Dockerfile' filename.
	"dockerFile": "../Dockerfile",

	// Set *default* container specific settings.json values on container create.
	"settings": { 
		"terminal.integrated.shell.linux": null
	},

	// Add the IDs of extensions you want installed when the container is created.
	"extensions": [],

	// Use 'forwardPorts' to make a list of ports inside the container available locally.
	//"forwardPorts": [3000],

	"appPort": ["8080:3000" ]
	// Uncomment the next line to run commands after the container is created - for example installing curl.
	// "postCreateCommand": "apt-get update && apt-get install -y curl",

	// Uncomment when using a ptrace-based debugger like C++, Go, and Rust
	// "runArgs": [ "--cap-add=SYS_PTRACE", "--security-opt", "seccomp=unconfined" ],

	// Uncomment to use the Docker CLI from inside the container. See https://aka.ms/vscode-remote/samples/docker-from-docker.
	// "mounts": [ "source=/var/run/docker.sock,target=/var/run/docker.sock,type=bind" ],

	// Uncomment to connect as a non-root user if you've added one. See https://aka.ms/vscode-remote/containers/non-root.
	// "remoteUser": "vscode"
}
```

devcontainer.json

![Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled%201.png](Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled%201.png)

>< 이 마크를 누루면

![Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled%202.png](Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled%202.png)

Reopen Locally 를 통해 Json값이 변경되면 다시 열어 주어야 한다.

![Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled%203.png](Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled%203.png)

open folder in container, 등등 in container를 통해서 해당 remote 환경이 설정된다.

![Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled%204.png](Golang%20Dev%20Container%20with%20Docker%200166dc25682141669ef3f867e44eefed/Untitled%204.png)

이런식으로 소스코드를 수정할 수 있다.