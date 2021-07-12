module github.com/alpha-supsys/go-cm-ucloud-webhook

go 1.16

require (
	github.com/alpha-supsys/go-common v0.0.0-20210626160705-c4c4049064e6
	github.com/jetstack/cert-manager v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/ucloud/ucloud-sdk-go v0.21.9
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
)

replace golang.org/x/crypto => github.com/cert-manager/crypto v0.0.0-20210409161129-d4c19753215a

replace golang.org/x/net => golang.org/x/net v0.0.0-20210224082022-3d97a244fca7

// To be replaced once there is a release of kubernetes/apiserver that uses gnostic v0.5. See https://github.com/jetstack/cert-manager/pull/3926#issuecomment-828923436
replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1

// See https://github.com/kubernetes/kubernetes/issues/101567
replace k8s.io/code-generator => github.com/kmodules/code-generator v0.21.1-rc.0.0.20210428003838-7eafae069eb0

replace k8s.io/gengo => github.com/kmodules/gengo v0.0.0-20210428002657-a8850da697c2

// See https://github.com/kubernetes/kubernetes/pull/99817
replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
