# go-cm-ucloud-webhook

cert-manager 关于 ucloud的 webhook

# 使用

## 修改并安装
修改deploy/all.yaml
1、所有yourcompany为自己域名
2、yourPrivatekeyBase64、yourPublickeyBase64为ucloud api服务提供，自行调整

kubectl apply

## 配置
```
apiVersion: cert-manager.io/v1alpha2
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    email: youremail@xxx.com
    privateKeySecretRef:
      name: letsencrypt-prod
    server: 'https://acme-v02.api.letsencrypt.org/directory'
    solvers:
      - dns01:
          webhook:
            config:
              dn: yourcompany.com
              region: cn-gd
            groupName: acme.yourcompany.com
            solverName: udns
```

```
apiVersion: cert-manager.io/v1alpha2
kind: Certificate
metadata:
  name: example-tls
  namespace: xxxx
spec:
  secretName: example-com-tls
  commonName: example.com
  dnsNames:
  - example.com
  - "*.example.com"
  issuerRef:
    name: letsencrypt-prod
    kind: ClusterIssuer
```