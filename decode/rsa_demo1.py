
import rsa
#rsa加密

publicKey =''' -----BEGIN RSA PRIVATE KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END RSA PRIVATE KEY-----'''


# 
        

def rsaEncrypt(str):
    #生成公钥、私钥
    (pubkey, privkey) = rsa.newkeys(1024)
    print("pubkey",pubkey)
    print("privkey",privkey)

    # privkey = rsa.PrivateKey.load_pkcs1(publicKey.encode())

    #明文编码格式
    content = str.encode('utf-8')
    #公钥加密
    crypto = rsa.encrypt(content,pubkey)
    return (crypto,privkey)




def rsaDecrypt(str,pk):
    #私钥解密
    content = rsa.decrypt(str,pk)
    con=content.decode('utf-8')
    return con

(a,b) = rsaEncrypt("hello")
print('加密后密文：')
print(a)
content = rsaDecrypt(a,b)
print('解密后明文：')
print(content)

# publicKey.n = 5
# publicKey.k = 3 
rsa.encrypt(b'hello', publicKey)



