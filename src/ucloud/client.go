package ucloud

import (
	"encoding/json"
	"fmt"

	"github.com/alpha-supsys/go-common/config"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type UClient struct {
	UClient *uhost.UHostClient
}

func NewClient(cfg config.Config) *UClient {
	cred := &auth.Credential{
		PublicKey:  cfg.GetString("PublicKey", ""),
		PrivateKey: cfg.GetString("PrivateKey", ""),
	}

	ucfg := ucloud.NewConfig()
	// ucfg.Region = cfg.GetString("Region", "")

	uhostClient := uhost.NewClient(&ucfg, cred)
	// d := "Action=DescribeUHostInstance&Limit=10&Region=cn-gd"
	// fmt.Println(cred.CreateSign(d))

	return &UClient{
		UClient: uhostClient,
	}
}

func NewClientFromKeys(puk, pk string) *UClient {
	ucfg := ucloud.NewConfig()
	return &UClient{
		UClient: uhost.NewClient(&ucfg, &auth.Credential{
			PublicKey:  puk,
			PrivateKey: pk,
		}),
	}
}

func (s *UClient) UdnrDomainDNSQuery(Region string, Dn string) ([]*DnsRecord, error) {
	req := s.UClient.NewGenericRequest()
	req.SetPayload(map[string]interface{}{
		"Action": "UdnrDomainDNSQuery",
		"Dn":     Dn,
		"Region": Region,
	})
	res, err := s.UClient.GenericInvoke(req)
	res.GetPayload()
	if err != nil {
		return nil, err
	}
	payload := &Payload{}
	err = res.Unmarshal(payload)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (s *UClient) UdnrDomainDNSAdd(Region string, Dn string, record *DnsRecord) error {
	req := s.UClient.NewGenericRequest()
	req.SetPayload(map[string]interface{}{
		"Action":     "UdnrDomainDNSAdd",
		"Dn":         Dn,
		"Region":     Region,
		"RecordName": record.RecordName,
		"DnsType":    record.DnsType,
		"Content":    record.Content,
		"Prio":       record.Prio,
		"TTL":        record.TTL,
	})
	res, err := s.UClient.GenericInvoke(req)
	if err != nil {
		return err
	}
	fmt.Println(res.GetPayload())
	return nil
}

func (s *UClient) UdnrDeleteDnsRecord(Region string, Dn string, record *DnsRecord) error {
	req := s.UClient.NewGenericRequest()
	req.SetPayload(map[string]interface{}{
		"Action":     "UdnrDeleteDnsRecord",
		"Dn":         Dn,
		"Region":     Region,
		"RecordName": record.RecordName,
		"DnsType":    record.DnsType,
		"Content":    record.Content,
	})
	res, err := s.UClient.GenericInvoke(req)
	if err != nil {
		return err
	}
	fmt.Println(res.GetPayload())
	return nil
}

type DnsRecord struct {
	DnsType    string `json:"DnsType"`
	RecordName string `json:"RecordName"`
	Content    string `json:"Content"`
	Prio       string `json:"Prio"`
	TTL        string `json:"TTL"`
}

func (s *DnsRecord) String() string {
	bs, _ := json.Marshal(s)
	return string(bs)
}

type Payload struct {
	RetCode int          `json:"RetCode"`
	Action  string       `json:"Action"`
	Message string       `json:"Message"`
	Data    []*DnsRecord `json:"Data"`
}
