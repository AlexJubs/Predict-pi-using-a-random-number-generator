# Predict-pi-using-a-random-number-generator
A small python Project to estimate the value of pi using a random number generator

The results:

<img src="https://user-images.githubusercontent.com/41933907/97347148-bd7ed680-1862-11eb-992a-1d79398528c1.png">


How it works:

Let's say we generate 2 float values at random between 0 and 1, the 2 values represent a coordinate on a cartesian plane.
Suppose these 2 values are x and y. Thus, 0 <= x <= 1 and 0 <= y <= 1.

We know that the distance of the coordinate which these values represent will be a maximum of (x^2 + y^2)^0.5 = (1+1)^0.5 = sqrt(2)

Now suppose we have a sector of angle pi/4 which is enclosed by a square of dimentions 1 by 1:


<p align="center">
  <img src="https://aquantumofknowledge.files.wordpress.com/2011/03/circle1.png">
</p>

As we generate the coordinates, we see that they can be mapped within the image above. Every single coordinate will be within the square as 0 <= x <= 1 and 0 <= y <= 1. Since x and y are random numbers, the coordinates will be mapped randomly in the square.be

We can say that the probability of x and y being within the quarter circle is equal to the number of coordinates which are within the quarter circle divided by the total number of coordinates. This ratio will always be less than or equal to 1.

We can also say that the ratio between the area of the quarter circle and the area of the square must be equal to the probability of a random coordinate landing within the quarter circle.

Using this information, we can say:
(area of quarter circle)/(area of square) = (number of coordinates in the quarter circle)/(total number of coordinates)

We know the area of the quarter circle is (pi * r^2)/4 and we know that the area of the square is l^2. But r = l =1 so we have the lefthand side of the equation as pi/4 / 1 = pi/4 and the righht and side is just the ratio of coordinates within the quarter circle.

We can characterize coordinates which land inside the quarter circle by the distance from origin formula. We know that if (x^2 + y^2)^0.5 is less than 1, then the distance from the origin of this coordinate will be less than the radius of the quarter circle. This means that the coordinate will be within the quarter circle. However, (x^2 + y^2)^0.5 is being less than 1 is the same as x^2 + y^2 being less than 1. (1^2 = 1)

Therefore, we have the right hand side as (number of coordinates in which x^2 + y^2 < 1)/(total number of coordinates)

We equate the 2 sides of the equations and we get pi/4 = (number of coordinates in which x^2 + y^2 < 1)/(total number of coordinates)

Thus, pi = 4*(number of coordinates in which x^2 + y^2 < 1)/(total number of coordinates)

As we generate more and more coordinates, the accuracy of the ratio improves and hence the error on pi decreases. For this particular example, I generated 100,000 coordinates.
