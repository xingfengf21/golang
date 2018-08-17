# -*- coding: utf-8 -*-
import rsa
 
# 先生成一对密钥，然后保存.pem格式文件，当然也可以直接使用
# (pubkey, privkey) = rsa.newkeys(1024)
 
# pub = pubkey.save_pkcs1()
# print(type(pub))
# pubfile = open('public.pem','wb')

# pubfile.write(pub)
# pubfile.close()
 
# pri = privkey.save_pkcs1()

# prifile = open('private.pem','wb')
# prifile.write(pri)
# prifile.close()
 
# load公钥和密钥
message = b'lovesoo.org'
with open('public.pem','rb') as publickfile:
    
    p = publickfile.read()
    pubkey = rsa.PublicKey.load_pkcs1(p)
 
with open('private.pem','rb') as privatefile:
    p = privatefile.read()
    privkey = rsa.PrivateKey.load_pkcs1(p)
 
# 用公钥加密、再用私钥解密
#crypto = rsa.encrypt(message, pubkey)
crypto = b'\x9f!\x7f\xb29oH\x8f\x07#G\x1a\x89\xb7$\x82\x16\x14\xc0#r{\xd2\x96\xb6\xa8He\xd3i\xcc\xf4n\xb9$\xb4\\\xaa\xd8Mf\x02\xb0\x9d\x85&y\xbd\x0c\x00\xfcu\xfbI\x01x\x99\x10IE\xbc\xc0\x14\x9c\x1e\xb3D\xa6}\x15\xc6\xf0\xe6p\x1c\x11\xb9\xabI\x11Y6\xe1\x00Y\xa7d\x8a"1\xa0\xe5\xa5\xdf\xc3\x0c\x1cv?\xb6\xcb\xb5\x80\r\xab(fH\x81\xe2\xbeR\xbe\xb2\xae\xfe\t\x13A\x0c\x83\r\x0f\xb8\xc3\x1d8\x1e'
print("crypto", crypto)
message = rsa.decrypt(crypto, privkey)
print(message)
 
# sign 用私钥签名认证、再用公钥验证签名
signature = rsa.sign(message, privkey, 'SHA-1')
t = rsa.verify(b'lovesoo.org', signature, pubkey)
print(t)