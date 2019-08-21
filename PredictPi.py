import random
#A counter for the total number of coordinates and a counter for the number of coordinates which are within the quarter circle
top =0
bot =0
for i in range(1000000):
    x = random.randint(0,100000)/100000
    y = random.randint(0,100000)/100000
    #This will only run if the coordinate is within the quarter circle
    if (x**2 + y**2 <= 1): top = top+1
    bot = bot+1
res = 4*(top/bot)
print("Pi is:" , res)
