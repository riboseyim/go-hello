
##

- 生成私钥
- 将私钥转换成PKCS8格式
- 生成公钥
```
openssl genrsa -out rsa_private_key.pem 2048

openssl pkcs8 -topk8 -inform PEM -in rsa_private_key.pem -outform PEM -nocrypt -out rsa_private_key_pkcs8.pem

openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
```

## 注意：Go crypto
2018/06/19 13:52:24 couldn't encrypt the secret message
 %s secrets must be 32 bytes long
2018/06/19 13:52:24 encrypt message:%s
panic: illegal base64 data at input byte 808 [recovered]
	panic: illegal base64 data at input byte 808
