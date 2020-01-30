module github.com/arutselvan15/estore-product-kube-client

require (
	github.com/imdario/mergo v0.3.8 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	k8s.io/api v0.17.2 // indirect
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.17.2
	k8s.io/utils v0.0.0-20200124190032-861946025e34 // indirect
)

// kubernetes libraries pinned to v1.16.0 tags
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190913080825-6f3bc4ba9215
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190912054826-cd179ad6a269
	k8s.io/kubernetes => k8s.io/kubernetes v1.16.0
)
