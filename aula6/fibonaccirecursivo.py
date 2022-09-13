import time

def fibonacci_recursivo(n):
    if (n ==0) or (n == 1):
        return n
    else:
        return fibonacci_recursivo(n-1) + fibonacci_recursivo(n-2)


n = 38
start_time = time.time()
f = fibonacci_recursivo(n)
print("--- %s segundos -- " % (time.time() - start_time))
print("Fibonacci de %d eh: %d" % (n, f))
