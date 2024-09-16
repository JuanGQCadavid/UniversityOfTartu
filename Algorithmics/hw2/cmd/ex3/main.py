import random
import time
import sys

sys.setrecursionlimit(1500)

# Generate a list of random numbers
rnumbers = random.sample(range(1, 10**8), 10**6)
# rnumbers = random.sample(range(1, 10**8), 10**3) # small data for initial test
snumbers = rnumbers.copy()
snumbers.sort()

# Sort the list using the quick sort algorithm

def forQuickSort(arr):
  high= len(arr) - 1
  low = 0

  pivot = arr[high]     # pivot
  i = low
  j = high-1

  while True:
    while True:
      if (high <= low) : break
      pivot = arr[high]     # pivot
      i = low
      j = high-1

      while( i <= j ) :
        if arr[i] < pivot:
            i = i+1
        else :
            arr[i], arr[j] = arr[j], arr[i]
            j = j-1
      pi = i
      arr[pi], arr[high] = arr[high], arr[pi]
      quickSort(arr, low, pi-1)
      quickSort(arr, pi+1, high)

def quickSort(arr, low, high):
    if (high <= low) : return

    pivot = arr[high]     # pivot
    i = low
    j = high-1

    while( i <= j ) :
        if arr[i] < pivot:
            i = i+1
        else :
            arr[i], arr[j] = arr[j], arr[i]
            j = j-1
    pi = i
    arr[pi], arr[high] = arr[high], arr[pi]

    quickSort(arr, low, pi-1)
    quickSort(arr, pi+1, high)

print( "# Quicksort 1")
for i in range(3):
  numbers = rnumbers.copy()
  start_time = time.time()
  quickSort( numbers, 0, len(numbers)-1 )
  end_time = time.time()
  print(f"Time to sort {len(numbers)} elements: {end_time - start_time:.2f} seconds")
  if numbers == snumbers:
    pass
    print("OK - Quicksort was correct")
  else:
    print("Error: Quicksort code was wrong")

# print( "# Quicksort 2")
# numbers = rnumbers.copy()
# for i in range(3):
#   start_time = time.time()
#   quickSort( numbers, 0, len(numbers)-1 )
#   end_time = time.time()
#   print(f"Time to sort {len(numbers)} elements: {end_time - start_time:.2f} seconds")


