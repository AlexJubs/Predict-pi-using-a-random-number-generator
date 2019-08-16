import random
top =0
bot =0
for i in range(100000):
    x = random.randint(0,100000)/100000
    y = random.randint(0,100000)/100000
    if (x**2 + y**2 <= 1): top = top+1
    bot = bot+1
res = 4*(top/bot)
print("Pi is:" , res)
