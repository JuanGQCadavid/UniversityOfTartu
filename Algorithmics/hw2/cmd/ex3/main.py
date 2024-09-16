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

def quickSortVStack(arr):
    operations = []

    operations.append({
       "low": 0, 
       "high": len(numbers)-1 
    })


    while True:
        if len(operations) == 0:
            break

        op = operations.pop()
        # print(op)

        if (op["high"] <= op["low"]) : continue

        pivot = arr[op["high"]]     # pivot
        i = op["low"]
        j = op["high"]-1

        while( i <= j ) :
            if arr[i] < pivot:
                i = i+1
            else :
                arr[i], arr[j] = arr[j], arr[i]
                j = j-1
        pi = i
        arr[pi], arr[op["high"]] = arr[op["high"]], arr[pi]


        operations.append({
            "low": pi+1, 
            "high": op["high"]
        })

        operations.append({
            "low": op["low"], 
            "high": pi-1
        })


def quickSortV3(arr, low, high):
    if (high <= low) : return

    pivot = arr[high]     # pivot
    green = low - 1
    yellow = low - 1

    while( green < high and yellow < high) :
        green = green+1
        if arr[green] <= pivot:
            yellow = yellow + 1
            if green > yellow :
                arr[yellow], arr[green] = arr[green], arr[yellow]
        # print(arr, 'Green at: ', green, 'yellow at:', yellow)
    # return
    pi = yellow
    quickSortV3(arr, low, pi-1)
    quickSortV3(arr, pi+1, high)

# arr = [3,2,5,0,1,8,7,6,9,4]
# quickSortV3(arr=arr, low=0, high=len(arr)-1)
# print(arr)


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

print( "# Quicksort V3. Recursived")
for i in range(3):
  numbers = rnumbers.copy()
  start_time = time.time()
  quickSortV3( numbers, 0, len(numbers)-1 )
  end_time = time.time()
  print(f"Time to sort {len(numbers)} elements: {end_time - start_time:.2f} seconds")
  if numbers == snumbers:
    pass
    print("OK - Quicksort was correct")
  else:
    print("Error: Quicksort code was wrong")

print( "# Quicksort V. stacks")
for i in range(3):
  numbers = rnumbers.copy()
  start_time = time.time()
  quickSortVStack( numbers)
  end_time = time.time()
  print(f"Time to sort {len(numbers)} elements: {end_time - start_time:.2f} seconds")
  if numbers == snumbers:
    pass
    print("OK - Quicksort was correct")
  else:
    print("Error: Quicksort code was wrong")

print( "# Quicksort original")
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

# # print( "# Quicksort 2")
# # numbers = rnumbers.copy()
# # for i in range(3):
# #   start_time = time.time()
# #   quickSort( numbers, 0, len(numbers)-1 )
# #   end_time = time.time()
# #   print(f"Time to sort {len(numbers)} elements: {end_time - start_time:.2f} seconds")


