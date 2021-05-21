# Intersection of two circles

The following note describes how to find the intersection point(s) between two circles on a plane, the following notation is used. The aim is to find the two points
![](<https://render.githubusercontent.com/render/math?math=P_{3}=(x_{3},y_{3})>) if they exist.

![](http://paulbourke.net/geometry/circlesphere/2circle1.gif)

First calculate the distance d between the center of the circles. ![](https://render.githubusercontent.com/render/math?math=d=||P1-P0||)

If ![](https://render.githubusercontent.com/render/math?math=d>r_{0}+r_{1}) then there are no solutions, the circles are separate.
If ![](https://render.githubusercontent.com/render/math?math=d<|r0-r1|) then there are no solutions because one circle is contained within the other.
If ![](https://render.githubusercontent.com/render/math?math=d=0) and ![](https://render.githubusercontent.com/render/math?math=r0=r1) then the circles are coincident and there are an infinite number of solutions.
Considering the two triangles ![](https://render.githubusercontent.com/render/math?math=P_{0}P_{2}P_{3}) and ![](https://render.githubusercontent.com/render/math?math=P_{1}P_{2}P_{3}) we can write
![](https://render.githubusercontent.com/render/math?math=a{2}+h^{2}=r^{0}^{2}) and ![](https://render.githubusercontent.com/render/math?math=b_{2}+h_{2}=r_{12})
Using ![](https://render.githubusercontent.com/render/math?math=d=a+b) we can solve for ![](https://render.githubusercontent.com/render/math?math=a),

![](<https://render.githubusercontent.com/render/math?math=a=(r_{0}^{2}-r_{1}^{2}+d^{2})/(2*d)>),
It can be readily shown that this reduces to r0 when the two circles touch at one point, ie: ![](https://render.githubusercontent.com/render/math?math=d=r_{0}±r_{1})

Solve for h by substituting a into the first equation, ![](https://render.githubusercontent.com/render/math?math=d=h2=r_{0}^{2}-a^{2})

So
![](<https://render.githubusercontent.com/render/math?math=P_{2}=P_{0}+a(P_{1}-P_{0})/d>)
And finally, P3=(x3,y3) in terms of P0=(x0,y0), P1=(x1,y1) and P2=(x2,y2), is

![](<https://render.githubusercontent.com/render/math?math=x_{3}=x_{2}+-h(y_{1}-y_{0})/d>)
![](<https://render.githubusercontent.com/render/math?math=y_{3}=y_{2}-+h(x_{1}-x_{0})/d>)
