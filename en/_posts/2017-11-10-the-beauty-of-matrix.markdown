---
title: "The Beauty of Matrix Rotations"
---

A matrix can be represented as an orientation of an object. If you came here because difficulties learning them as a game programmer, then read on. I had a similar pain that I want to share with you today.

## In Background

I'm developing [Engine4][engine4], a game engine in 4D. The biggest challenge when building the engine is how I should manipulate the rotation, which in this case I store it using matrix rotation. You may see this post is biased for 3D or 2D space, but it truly can be applied to all kind of space in general.

### Usability

Graphics libraries like [OpenGL][opengl] and [DirectX][directx] uses matrix to handle rotation. So it's best to follow what they use. Meanwhile, for 3D or higher space, I'd prefer matrix over Euler rotations overall to avoid pain with Gimbal Lock.

You may hear about [Quaternion][quaternion] as well. They're super efficient but hard to be understood and available only in 3D space. So I won't cover it here. Anyway, let's see which one is you prefer after reading this article.

### Aim

This article is written so you won't get confused with matrices, thereby can make use its analytical properties nicely. I assume you already knew what it is and the basic usage of matrix rotations and its connection to object transformation including terms like [*local* or *global*][localglobal] space.

Enough for the intro, let's dig in...

## Property of Matrix Rotation

What makes a matrix can be recognized as a matrix rotation?

### Same Order Number

Basically, the matrix must be square so that it won't change the number of dimensions of a vector.

```
2D matrix:      3D matrix:        4D matrix:
| 1 0 |         | 1 0 0 |         | 1 0 0 0 |
| 0 1 |         | 0 1 0 |         | 0 1 0 0 |
                | 0 0 1 |         | 0 0 1 0 |
                                  | 0 0 0 1 |
```

### Determinant = 1

A matrix rotation must always have a determinant of one, so it'll never change any vector magnitude. This kind of matrix is also called as [orthogonal matrix][orthomatrix], and it's mean good, as will be discussed later on.

When writing a program that involves matrix rotation, it's a good idea to pass some determinant test to ensure your matrix function works properly.

> **Important**: In [wikipedia][orthodeter] an orthogonal matrix can also have a determinant of -1, but I don't know if it also relevant to matrix rotation as well.

### Can't be added together

Matrix with the same order can be added or subtracted together, but this concept is no longer relevant for orthogonal matrix as it'll definitely change the matrix determinant even both is an orthogonal matrix as well.

Generally, you shouldn't write matrix addition or subtraction if it only means for rotation to avoid bad code ruins your software.

## Interesting Part

So what's interesting features of rotation matrices?

### Combining Rotation

This is basic stuff. If you already heard about quaternion then multiplication of matrix rotation have the exact same property with quaternion multiplication. You combine two rotation by multiplying it together.

```
C = A X B
```

You can also [inverse][minverse] one of the matrix. For example, to create a matrix that makes a point that relative to object B becomes relative to A.

```
C = inverse(A) * B
```

Remember that matrix multiplication is not commutative. We'll discuss below.

### Transpose = Inverse

This is my favorite. If you can guarantee that your code doesn't fail with the determinant test, then you can replace all inversion code with far cheaper [transpose][transpose] function as it is an identical operation if the matrix is orthogonal.

```
A CW matrix:                         A CCW matrix:
|  cos x   sin x | <- Transposed -> |  cos x  -sin x |
| -sin x   cos x | <- Transposed -> |  sin x   cos x |
```

If you haven't heard about this I suggest you change your code as this is a big opportunity for code optimization.

### Column as direction vector

It is very common to get orientation vector (i.e. normal) of an object. Most of the time, you do that by multiplying the desired vector by its rotation.

```
Result: Rotation:    Vector:
| 0 |   | 1  0 0 |   | 0 |
| 1 | = | 0  0 1 | X | 0 |
| 0 |   | 0 -1 0 |   | 1 |
  ^                    ^
  |                    +--- Local space
  +------------------------ World space
```

This operation, however, can be optimized. You only need to pick the **nth column** (not row) of the rotation matrix based on the vector. Note the third column of the example above.

```
Result: Rotation:    Vector:
| 0 |   | 1  0 0 |   | 0 |
| 1 | = | 0  0 1 | X | 0 |
| 0 |   | 0 -1 0 |   | 1 |
  ^            ^
  +------------+----------- World space (exact same)
```

Note if you pick from the nth row instead, then it's equivalent to multiplying the vector with the transposed version of the matrix.

### Sticking it with space vector

This is the cool part when using matrix. We know that an object must maintain not just its orientation, but its location as well. Game designers have a better way to do this by (make it look like we / actually) store the location data in the next column of the rotation:

```
|  1  0  0 10  |
|  0  0  1  1  |
|  0 -1  0  2  |
   <--+-->  ^
      |     +--- Position
      +--------- Rotation
```

And then we can multiply it using an extra row each. This works for both `matrix * vector` and `matrix * matrix`.

```
| 10 |   |  1  0  0 10 |   | 1 |
|  1 | = |  0  0  1  1 | X | 2 |
|  0 |   |  0 -1  0  2 |   | 0 |
|  1 |   |  0  0  0  1 |   | 1 |  <---  Extra row
```

## Challenging Part

There may some factors that prevent you to adapt with concept painlessly. You're not alone, I just covered some quirks for you ...

### Storage (Major order)

Yes, this is the #1 problem for many developers. This is happening because memory doesn't save it as a 2D array. They have to be mapped as a 1D array first.

This may sounds simple for you, but it's really not because there are two ways to do that:

```
Row major:  Column major:

| 0 1 2 |   | 0 3 6 |
| 3 4 5 |   | 1 4 7 |
| 6 7 8 |   | 2 5 8 |

*Numbers denote array index
```

And even worse, OpenGL maps the matrix as column major **while DirectX is the [opposite][matrixlayout]**.

Why this matter? Because picking row/column interchangeably could result in *bad code*, even worse, matrix multiplication rely heavily on that:

```c#
// simplest matrix * vector in pseudo code
vector operator *(matrix, vector){
    return vector {
        x = dot(matrix.row[0], vector)
        y = dot(matrix.row[1], vector)
        z = dot(matrix.row[2], vector)
    }
}
```

```c#
// simplest matrix * matrix in pseudo code
matrix operator *(matrix1, matrix2){
    return matrix {
        column[0] = matrix1 * matrix2.column[0]
        column[1] = matrix1 * matrix2.column[1]
        column[2] = matrix1 * matrix2.column[2]
    }
}
```

When reading someone code, make sure to know if their matrix is implemented in the row or column major. This is important because most people code their matrix members in vector variable like `matrix.ex` while you're not sure whether `ex` is an nth row or column of the `matrix`. This is also true when accessing matrix members via jagged array like `matrix[column][row]` or `matrix[row][column]`.

If you're using a third-party code and know there's matrix conflict, or translating OpenGL to DirectX and vice-versa, make sure to `transpose()` the matrix and everything will be okay then.

> When implementing your own matrix code, My suggestion is to use *row major* because it's our nature to read words row by row, and consistent with math notation when dealing with jagged arrays.

### Euler Conversion

This is usually don't have to be a big problem. Unless you're using a third-party code and you know there must be some Euler conversion involved.

In the nutshell, this is how 3D euler rotation converted to matrix (unoptimized):

```
| 1 1      0     |   | cos y 0 -sin y |   | cos z -sin z 0 |
| 0 cos x -sin x | X | 0     1  0     | X | sin z  cos z 0 |
| 0 sin x  cos x |   | sin y 0  cos y |   | 0      0     1 |
```

Notice that how easily the order can be rearranged. By convention, the arrangement above can be called as ZYX (Z then Y then X) Euler conversion. This is just one of [many conventions][eulconvent] available. Meanwhile, Euler to matrix conversion is just a start, converting back is really hard (it need kind of good guessing). FYI, did I also mention CW vs. CCW as well?

AFAIK, There is no widely used Euler conversion (except [in some professions][eulformal]). Everyone can pick which convention is best, as long as it's consistent with what they do.

> To avoid this problem with third parties, try to just use their built-in Euler conversion (game engines, for example). Or better, avoid passing Euler back and forth, just use their matrices instead (with watching the major order).

### Multiplication order

It's easy to make (but hard to spot) mistakes when multiplicating two matrices because of its non-commutativity.

For the starting point, everyone knows how to convert a vector direction from local to global space (and vice-versa). Because `matrix * vector` is allowed in the mathematical sense, but not `vector * matrix`.

```js
global = rotation * local;            // local to world
local = transpose(rotation) * global; // world to local
localInB = transpose(rotationB) * rotationA * localInA; // local of A becomes local of B
```

But the problem starts when *manipulating the matrix*. especially applying delta rotation or angular velocity.

In matrix world, you turn angular velocity to delta rotation with this:

```c#
delta = fromeuler(angularvelocity); // neat. just do euler conversion.
```

When integrating the delta to current rotation, most people who don't care about order will naively do this:

```c#
rotation *= delta; // a shorthand to "rotation = rotation * delta"
```

That is **not** entirely correct. While the example above perfectly fine for simple rotation (e.g. rotating turret, character sideways), that will make things broken for complex situations (e.g. physics simulation) as the delta orientation will be *changed* by the former (left-side) matrix.

Your option is reverse the order.

```js
rotation = delta * rotation; // This is the correct order. The readable saying is like "change rotation by delta" and not another way around.
```

That's fine. But in some (not often) situation, when you need the delta *is based on* an object orientation, all you have to do is wrap it around the object matrix. Some people refer this to a *sandwich operation*

```js
rotation = transpose(object) * delta * object * rotation; // A magic operation where "I want to adapt delta based to object but don't want it changes the rotation"
```

## All Covered

That's everything I know about matrices as an object orientation. Feel free to ask in the comment below.

[engine4]: http://wellosoft.net/engine4-doc/
[directx]: https://en.wikipedia.org/wiki/DirectX
[opengl]: https://en.wikipedia.org/wiki/OpenGL
[quaternion]: https://en.wikipedia.org/wiki/Quaternion
[localglobal]: https://teamtreehouse.com/community/global-space-vs-local-space
[orthomatrix]: https://en.wikipedia.org/wiki/Orthogonal_matrix
[orthodeter]: https://en.wikipedia.org/wiki/Orthogonal_matrix#Matrix_properties
[minverse]: https://www.mathsisfun.com/algebra/matrix-inverse.html
[transpose]: https://en.wikipedia.org/wiki/Transpose
[matrixlayout]: http://www.mindcontrol.org/~hplus/graphics/matrix-layout.html
[eulconvent]: https://ntrs.nasa.gov/search.jsp?R=19770019231
[eulformal]: https://en.wikipedia.org/wiki/Rotation_formalisms_in_three_dimensions#Euler_rotations