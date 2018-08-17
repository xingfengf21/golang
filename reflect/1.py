#coding:utf-8
s='林' #当程序执行时，无需加u，'林'也会被以unicode形式保存新的内存空间中,

#s可以直接encode成任意编码格式
print(s)
print(s.encode('utf-8').decode("utf-8"))
print(s.encode('gbk').decode("gbk"))

print(type(s)) #<class 'str'>