from fractions import Fraction
import timeit

def BenchmarkAddition(N=5712266):
	acc = Fraction(0, 1)
	step = Fraction(9223372036854775807, 1)
	for _ in range(N):
		acc += step
	return acc

def BenchmarkPush(N=5712266):
	acc = []
	for _ in range(N):
		acc.append(7)
	return acc


print(1000000000 * timeit.timeit(BenchmarkPush, number=5) / 5712266 / 5, "ns")