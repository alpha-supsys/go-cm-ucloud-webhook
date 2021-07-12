package ucloud

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alpha-supsys/go-common/config"
	"github.com/jetstack/cert-manager/pkg/acme/webhook"
	whapi "github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var pkNotExists = errors.New("pk not exists")
var pukNotExists = errors.New("puk not exist")

type Config struct {
	Region string `json:"region"`
	Dn     string `json:"dn"`
	Secret string `json:"secret"`
}

type Solver struct {
	K8sClient  *kubernetes.Clientset
	UClient    *UClient
	Namespace  string
	SecretName string
}

func NewSolver(cfg config.Config) webhook.Solver {
	namespace := cfg.GetString("Namespace", "")
	secretName := cfg.GetString("SecretName", "")

	return &Solver{
		Namespace:  namespace,
		SecretName: secretName,
	}
}

func (s *Solver) Name() string {

	return "udns"
}

func (s *Solver) Present(ch *whapi.ChallengeRequest) error {
	cfg := &Config{}
	if err := json.Unmarshal(ch.Config.Raw, cfg); err != nil {
		return err
	}

	fmt.Println(cfg.Region, cfg.Dn, ch.Key)
	err := s.UClient.UdnrDomainDNSAdd(cfg.Region, cfg.Dn, &DnsRecord{
		DnsType:    "TXT",
		RecordName: "_acme-challenge." + cfg.Dn,
		Content:    ch.Key,
		Prio:       "-",
		TTL:        "600",
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Solver) CleanUp(ch *whapi.ChallengeRequest) error {
	cfg := &Config{}
	if err := json.Unmarshal(ch.Config.Raw, cfg); err != nil {
		return err
	}
	err := s.UClient.UdnrDeleteDnsRecord(cfg.Region, cfg.Dn, &DnsRecord{
		DnsType:    "TXT",
		RecordName: "_acme-challenge." + cfg.Dn,
		Content:    ch.Key,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Solver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	cl, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return err
	}
	s.K8sClient = cl

	secret, err := s.K8sClient.CoreV1().Secrets(s.Namespace).Get(context.TODO(), s.SecretName, metav1.GetOptions{})
	if err != nil {
		return errors.Wrapf(err, "failed to load secret %q", s.Namespace+"/"+s.SecretName)
	}

	var pk string
	var puk string
	if pkbs, ok := secret.Data["pk"]; ok {
		pk = string(pkbs)
	} else {
		return pkNotExists
	}
	if pukbs, ok := secret.Data["puk"]; ok {
		puk = string(pukbs)
	} else {
		return pukNotExists
	}

	uclient := NewClientFromKeys(puk, pk)
	s.UClient = uclient

	return nil
}
