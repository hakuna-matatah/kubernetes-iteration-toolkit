module github.com/prateekgogia/kit/cli

go 1.16

require (
	github.com/spf13/cobra v1.1.3
	go.uber.org/zap v1.16.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	helm.sh/helm/v3 v3.5.4
	honnef.co/go/tools v0.1.4 // indirect
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
	sigs.k8s.io/controller-runtime v0.7.0-alpha.3
	sigs.k8s.io/kind v0.10.0
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
