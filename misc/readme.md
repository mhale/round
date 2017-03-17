# Rounding Mode Utilities

These utilities display the same test values used by the round package, but they are rounded with the C standard library, Java and Go's [Float type](https://golang.org/pkg/math/big/#Float) from the math/big package as a comparison.

## Rounding Modes

This table describes equivalency between various rounding modes.

|Rounding Mode|IEEE 754-2008|C|Java|Go math/big|
|-------------|-------------|-|----|-----------|
|Round to nearest neighbour, round ties to even|roundTiesToEven|FE_TONEAREST|RoundingMode.HALF_EVEN|big.ToNearestEven|
|Round to nearest neighbour, round ties toward zero|||RoundingMode.HALF_DOWN||
|Round to nearest neighbour, round ties away from zero|roundTiesToAway||RoundingMode.HALF_UP|big.ToNearestAway|
|Round toward positive infinity	|roundTowardPositive|FE_UPWARD|RoundingMode.CEILING|big.ToPositiveInf|
|Round toward negative infinity	|roundTowardNegative|FE_DOWNWARD|RoundingMode.FLOOR|big.ToNegativeInf|
|Round toward zero|roundTowardZero|FE_TOWARDZERO|RoundingMode.DOWN|big.ToZero|
|Round away from zero|||RoundingMode.UP|big.AwayFromZero|
|Do not round, assert exact result|||RoundingMode.UNNECESSARY||	
## C

Assumption: A C compiler, such as GCC or Clang, is installed.

To compile and run the C example on Linux or OS X, run:

```
cc -o rounding rounding.c -lm -frounding-math && ./rounding
```

Clang may emit a harmless warning or two during compilation.

There are other numeric packages for C, such as the [GNU MPFR library](http://www.mpfr.org/), which can perform more accurate rounding than the functions in the standard library.

## Java

Assumption: A recent Java Development Kit is installed.

To compile and run the Java example on Linux or OS X, run:

```
javac Rounding.java && java Rounding
```

## Go's math/big package

To run the Go example, run:

```
go run bigfloat.go 
```

Note that this Go example uses the only built-in way to perform floating-point rounding: the [SetPrec](https://golang.org/pkg/math/big/#Float) function in the math/big package.

The trouble with SetPrec is that it takes a number of bits of precision as a parameter, not a number of decimal places. This is not a user friendly method of rounding numbers.
