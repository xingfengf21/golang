import jwt
import random 
import time 

# random.randint(0,2**32-1)

def newAuthorKey(mySigningKey , kid , user ):


    d =  {"iss":      "bfupdate",
        "username": user,
        "iat":      int(round(time.time() * 100)),
        "kid":      kid,
        "rnd":      1298498081,
    }

    print(d)

    encoded = jwt.encode(d, 'secret', algorithm='HS256')
    return encoded



if __name__ == '__main__':
    token = newAuthorKey("mySigningKey", "kid", "user")
    print(token)
    print(len(token))
    token = str(token, encoding = "utf8")  
    print(len(token),type(token))
