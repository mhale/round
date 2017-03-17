# Rounding Mode Utilities

These utilities display the same test values used by the round package, but they are rounded with the C standard library, Java and Go's [Float type](https://golang.org/pkg/math/big/#Float) from the math/big package as a comparison.

## Rounding Modes

This table describes equivalency between various rounding modes.

|<sub>Rounding Mode</sub>|<sub>IEEE 754</sub>|<sub>C</sub>|<sub>Java</sub>|<sub>Go math/big</sub>|
|-------------|-------------|-|----|-----------|
|<sub>Round to nearest neighbour, round ties to even</sub>|<sub>roundTiesToEven</sub>|<sub>FE_TONEAREST</sub>|<sub>RoundingMode.HALF_EVEN</sub>|<sub>big.ToNearestEven</sub>|
|<sub>Round to nearest neighbour, round ties toward zero</sub>|||<sub>RoundingMode.HALF_DOWN</sub>||
|<sub>Round to nearest neighbour, round ties away from zero</sub>|<sub>roundTiesToAway</sub>||<sub>RoundingMode.HALF_UP</sub>|<sub>big.ToNearestAway</sub>|
|<sub>Round toward positive infinity</sub>|<sub>roundTowardPositive</sub>|<sub>FE_UPWARD</sub>|<sub>RoundingMode.CEILING</sub>|<sub>big.ToPositiveInf</sub>|
|<sub>Round toward negative infinity</sub>|<sub>roundTowardNegative</sub>|<sub>FE_DOWNWARD</sub>|<sub>RoundingMode.FLOOR</sub>|<sub>big.ToNegativeInf</sub>|
|<sub>Round toward zero</sub>|<sub>roundTowardZero</sub>|<sub>FE_TOWARDZERO</sub>|<sub>RoundingMode.DOWN</sub>|<sub>big.ToZero</sub>|
|<sub>Round away from zero|||<sub>RoundingMode.UP</sub>|<sub>big.AwayFromZero</sub>|
|<sub>Do not round, assert exact result</sub>|||<sub>RoundingMode.UNNECESSARY</sub>||

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
