package piscine

func Swap(a *int, b *int) {
	/*
	   a, b = swap(a, b)

	   	Before swapping:

	   a = 10
	   b = 20
	   After swapping:
	   a = 20
	   b = 10
	*/
	*a, *b = *b, *a
}
