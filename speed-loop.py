import time

def main():
  cycles = 1000000000
  start = time.perf_counter()

  for i in range(0, cycles):
    pass # This is an empty loop

  end = time.perf_counter()

  duration = end - start

  print("Python 3 looped %s time in %s seconds" % (cycles, duration))

main()