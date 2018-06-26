
from datetime import datetime


# func TestNumber2() {
#     var number int64
#     number = 5000000000
#     var n int64
#     var i int64
#     fmt.Println(time.Now())
#     for i = 1; i <= number; i++ {
#         if i%2 == 0 {
#             n = n + 1
#         }
#     }
#     fmt.Println(time.Now().String())
#     fmt.Println(n)
# }

n = 0

number = 5000000000
print datetime.now()
for i in xrange(number):
    if i %2 == 0:
        n += 1


print datetime.now()
print n 


