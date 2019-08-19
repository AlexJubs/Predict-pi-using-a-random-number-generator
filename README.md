# Predict-pi-using-a-random-number-generator
A small python Project to estimate the value of pi using a random number generator

How it works:

Let's say we generate 2 float values at random between 0 and 1, the 2 values represent a coordinate on a cartesian plane.
Suppose these 2 values are x and y. Thus, 0 <= x <= 1 and 0 <= y <= 1.

We know that the distance of the coordinate which these values represent will be a maximum of (x^2 + y^2)^0.5 = (1+1)^0.5 = sqrt(2)

Now suppose we have a sector of angle pi/4 which is enclosed by a square of dimentions 1 by 1:

![alt text](https://aquantumofknowledge.files.wordpress.com/2011/03/circle1.png)

As we generate the coordinates, we see that they can be mapped within the image above. Every single coordinate will be within the square.
