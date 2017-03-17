# round

A floating-point number rounding package for Go.

The Go standard library does not provide a rounding function in the math package, which requires Gophers to write their own. Here are some rounding functions you can use instead of doing that.

The default rounding mode is "banker's rounding". That is, values are rounded to the nearest integer value, with ties rounded to even integers.

For a C version of this functionality, see the [cround](http://github.com/mhale/cround) package.

## Usage

Import round as you would any other package, and then call one of the exposed functions. See the source code for a listing.

If you are not sure which function to use, you should use Round. A more explicit function is ToNearestEven.

### Round

Round is a convenience wrapper for ToNearestEven.

```
value := round.Round(2.5)
// value is 2.0
```

### RoundTo

RoundTo is a convenience wrapper for ToNearestEven, which rounds to a specified number of decimal places.

```
value := round.Round(1234.5678, 2)
// value is 1234.57
```

### ToNearestEven

ToNearestEven rounds to the nearest integer value, with ties rounded to even integers.

```
value := round.ToNearestEven(2.5)
// value is 2.0
```

## Miscellaneous

These links may be useful to learn about floating-point rounding modes.

* [IEEE Standard for Floating-Point Arithmetic: Rounding rules (IEEE 754)](https://en.wikipedia.org/wiki/IEEE_floating_point#Rounding_rules)
* [The GNU C Library: Rounding](https://www.gnu.org/software/libc/manual/html_node/Rounding.html)
* [RoundingMode (Java Platform SE 8)](https://docs.oracle.com/javase/8/docs/api/java/math/RoundingMode.html)

There are some additional utilities in the [misc](/misc) subdirectory. These utilities display the same test values used by this package, but rounded with C, Java and Go's [Float type](https://golang.org/pkg/math/big/#Float) from the math/big package as a comparison.

## Testing

The tests cover a range of test values, including the zero and infinity special cases. NaN is not included because a comparison between two NaN values always returns false.

There is an additional set of more rigorous tests in the [cround](/cround) subdirectory. These tests compare the round package results to the results returned from the mature rounding functions implemented in the C standard library.

Please submit any additional test cases that demonstrate errors or inaccuracies.


## Licensing

This code is in the public domain. Anyone is free to copy, modify, publish, use, compile, sell, or distribute the original code, either in source code form or as a compiled binary, for any purpose, commercial or non-commercial, and by any means.