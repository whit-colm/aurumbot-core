# Maths Equation Sheet

This paper contains complex markdown syntax and includes Forked GitHub-flavoured markdown, HTML, and LaTeX. It is recommended an application such as iA Writer is used to format this, as ViM will read HTML and LaTeX literally.

## Unit Circle

<img src="https://cdn.whitmans.io/unit_circle.png" alt="unit circle" width="500"/>

## Definitions:

### Unit 0:

**Standard Position** An angle is said to be in standard position **(sp)** if the vertex is at the origin of the coordinate system and the initial side is the \\(+x\\) axis.

**Pythagorean Theorem:** \\(a^2 + b^2 = c^2\\). Or in a more holistic sense: \\(c = \sqrt{a^2 + b^2}\\).

#### Angle Types:

Given angle \\(a\\):
- \\(a = 180˚\\) - Straight

- \\(a = 90˚\\) - Right

- \\(0˚ < a < 90˚\\) - Acute

- \\(90˚ < a < 180˚\\) - Obtuse

- \\(a_1 + a_2 = 90˚\\) - Complimentary

- \\(a_1 + a_2 = 120˚\\) - Supplimentary

#### Properties of Triangles

Given \\(\triangleABC\\):

<img src="https://cdn.whitmans.io/triangle.png" alt="unit circle" width="250"/>

1. \\(a + b > c\\), \\(a + c > b\\), \\(b + c > a\\)

2. Longer sides are always opposite to longer angles.

#### Special Triangles:

<img src="https://cdn.whitmans.io/45triangle.png" alt="The 45-45 Triangle" width="100"/> <img src="https://cdn.whitmans.io/3-4-5triangle.png" alt="The 3-4-5 Triangle" width="100"/>

### Complex Fractions
 
Complex fractions have a fraction in the numerator or the denominator: \\[\frac{\frac{-4}{3}+\frac{12}{5}}{1-\frac{-4}{3} * \frac{12}{5}}\\] To resolve this, multiply the overall numerator and overall denominator and multiply them each by the common denominator of all of the sub-fractions. So the previous fraction is resolved as: 
\\[\frac{\frac{-4}{3}+\frac{12}{5}}{1-\frac{48}{15}}* \frac{15}{15}\\]

\\[=\frac{-20+36}{15+48}\\]

\\[\frac{16}{63}\\]

### Degrees <-> Radians

\\(1˚=\frac{\pi}{180}r\\)

\\(1r=\frac{180˚}{\pi}\\)

If anything isn't suffixed with a `˚`, assume it's in radians.

## 6 Trigonometric Functions:

If \\(\theta\\) is an angle in standard position with \\(P(x, y)\\) being a point on the terminal side and \\(r = \sqrt{x^2 + y^2}\\) then:

\\(\sin\theta=\frac{y}r\\) | \\(\csc\theta=\frac{r}y\\)

\\(\cos\theta=\frac{x}r\\) | \\(\sec\theta=\frac{r}x\\)

\\(\tan\theta=\frac{y}x\\) | \\(\cot\theta=\frac{x}y\\)

If on the unit circle, it is always known that \\(r=1\\).

## Translating the Graphs of Sine and Cosine

To graph \\(y=a\cos(bx+c)+d\\) and \\(y=a\sin(bx+c)+d\\):

1. The **amplitude** \\(=|a|\\)

2. The **period** \\(=\frac{2\pi}{5}\\)

3. The **phase shift** \\(=bx+c = 0\\) as solved for \\(x\\).

4. The **vertical shift** is \\(d\\).

## Fundamental Identities

\\(2x + 8 = 12\\) → \\(2x = 4\\)

\\((x+1)^2 = x^2 + 2x + 1\\)

### Reciprocal Identities

\\(\frac{\sin\theta}{1}=\frac{1}{\csc\theta}\\)

\\(\frac{\cos\theta}{1}=\frac{1}{\sec\theta}\\)

\\(\frac{\tan\theta}{1}=\frac{1}{\cot\theta}\\)

### Quotant Identities

\\( \tan\theta=\frac{\sin\theta}{\cos\theta}\\)

\\( \cot\theta=\frac{\cos\theta}{\sin\theta}\\)

### Pythagorean Identities

\\(\cos^2\theta+\sin^2\theta=1\\)

\\(1+\tan^2\theta=\sec^2\theta\\)

\\(1+\cot^2\theta=\csc^2\theta\\)

### Symmetrical Identities

\\(\sin\theta=\frac{-y}{r}=-\frac{y}{r}=-\sin\theta\\)

\\(\cos-\theta=\frac{x}{r}=\sin\theta\\)

\\(\tan-\theta = -\frac{y}{x}=-\frac{b}{x}=\tan\theta\\)

### Sum/Difference identities

One cannot distribute \\(\sin\\), \\(\cos\\) and the like. Which is to say that \\(\sin(60˚ + 30˚) ≠ \sin60˚+\sin30˚\\).

\\(\sin(A±B)=\sinA\cosB±\cosA\sinB\\)

\\(\cos(A±B)=\cosA\cosB∓\sinA\sinB\\)

\\(\tan(A±B)=\frac{\tanA±\tanB}{1∓\tanA\tanB}\\)

## `/etc/`:

# EOF