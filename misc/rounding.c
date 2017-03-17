#include <fenv.h>
#include <math.h>
#include <stdio.h>

#pragma STDC FENV_ACCESS ON

int main()
{
	double values[] = {5.5, 2.5, 1.6, 1.1, 1.0, -1.0, -1.1, -1.6, -2.5, -5.5};

	printf("Default rounding mode: ");
	switch (fegetround()) {
		case FE_TONEAREST:  printf("FE_TONEAREST\n");  break;
		case FE_DOWNWARD:   printf("FE_DOWNWARD\n");   break;
		case FE_UPWARD:     printf("FE_UPWARD\n");     break;
		case FE_TOWARDZERO: printf("FE_TOWARDZERO\n"); break;
		default:            printf("unknown\n");
	};

	printf("   x  FE_TONEAREST  FE_DOWNWARD  FE_UPWARD  FE_TOWARDZERO\n");

	for (int i = 0; i < (sizeof(values) / sizeof(double)); i++) {
		printf("%4.1f", values[i]);

		fesetround(FE_TONEAREST);
		printf("%14.1f", rint(values[i]));

		fesetround(FE_DOWNWARD);
		printf("%13.1f", rint(values[i]));

		fesetround(FE_UPWARD);
		printf("%11.1f", rint(values[i]));

		fesetround(FE_TOWARDZERO);
		printf("%15.1f\n", rint(values[i]));
	}

	return 0;
}
