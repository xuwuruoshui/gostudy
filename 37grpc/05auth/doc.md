# Openssl
## CA证书生成
go 1.15后需要支持 SANs
SANs: SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。
```shell
vim ca.conf

[ req ]
default_bits       = 4096
distinguished_name = req_distinguished_name

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = Sichuan
localityName                = Locality Name (eg, city)
localityName_default        = Chengdu
organizationName            = Organization Name (eg, company)
organizationName_default    = Test
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = MyServer

# 生成key
openssl genrsa -out ca.key 4096

# 生成csr
openssl req -new -sha256 -out ca.csr -key ca.key -config ca.conf

# 生成crt
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt
```

## 终端证书
```shell
vim server.conf
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name
req_extensions     = req_ext

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = Sichuan
localityName                = Locality Name (eg, city)
localityName_default        = Chengdu
organizationName            = Organization Name (eg, company)
organizationName_default    = Test
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = xwrs

[ req_ext ]
subjectAltName = @alt_names

[alt_names]
DNS.1   = www.haha.com
IP      = 192.168.0.105


# 生成server.key
openssl genrsa -out server.key 2048
# 生成server.csr
openssl req -new -sha256 -out server.csr -key server.key -config server.conf
# 生成server.crt
openssl x509 -req -days 3650 -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.pem -extensions req_ext -extfile server.conf
```