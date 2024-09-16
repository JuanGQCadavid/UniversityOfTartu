import random
import time
import sys

sys.setrecursionlimit(1500)

# Generate a list of random numbers
rnumbers = random.sample(range(1, 10**8), 10**6)
# rnumbers = random.sample(range(1, 10**8), 10**3) # small data for initial test
snumbers = rnumbers.copy()
snumbers.sort()

def timeMesure(testName, funct):
   print(testName)
   for i in range(3):
    numbers = rnumbers.copy()
    start_time = time.time()
    funct( numbers, 0, len(numbers)-1 )
    end_time = time.time()
    print(f"Time to sort {len(numbers)} elements: {end_time - start_time:.2f} seconds")
    if numbers == snumbers:
        pass
        print("OK - was correct")
    else:
        print("Error: code was wrong")

def quickSortV3VStack(arr, low, high):
    operations = []

    operations.append({
       "low": 0, 
       "high": len(arr)-1 
    })
    ## TO HERE

    while True:
        if len(operations) == 0:
            break

        op = operations.pop()
        # print(op)

        if (op["high"] <= op["low"]) : continue

        pivot = arr[op["high"]]     # pivot
        green = op["low"] -1
        yellow = op["low"]-1

        while( green < op["high"] and yellow < op["high"]) :
            green = green+1
            if arr[green] <= pivot:
                yellow = yellow + 1
                if green > yellow :
                    arr[yellow], arr[green] = arr[green], arr[yellow]

        pi = yellow
        operations.append({
            "low": pi+1, 
            "high": op["high"]
        })

        operations.append({
            "low": op["low"], 
            "high": pi-1
        })

def quickSortVStack(arr, low, high):
    operations = []

    operations.append({
       "low": 0, 
       "high": len(arr)-1 
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
    pi = yellow
    quickSortV3(arr, low, pi-1)
    quickSortV3(arr, pi+1, high)

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

def median_of_three(arr, low, high):
    mid = low + (high - low) // 2

    if arr[low] > arr[mid]:
        arr[low], arr[mid] = arr[mid], arr[low]

    if arr[low] > arr[high]:
        arr[low], arr[high] = arr[high], arr[low]

    if arr[mid] > arr[high]:
        arr[mid], arr[high] = arr[high], arr[mid]

    arr[mid], arr[high-1] = arr[high-1], arr[mid]

    return arr[high-1]  # Return the pivot value

def quickSortMediamOfThree(arr, low, high):
    if high <= low:
        return
    pivot = median_of_three(arr, low, high)
    i = low
    j = high - 2  # Since pivot is at high-1
    while i <= j:
        if arr[i] < pivot:
            i += 1
        else:
            arr[i], arr[j] = arr[j], arr[i]
            j -= 1
    # Place pivot in its correct position
    pi = i
    arr[pi], arr[high-1] = arr[high-1], arr[pi]
    quickSortMediamOfThree(arr, low, pi - 1)
    quickSortMediamOfThree(arr, pi + 1, high)

timeMesure("# Quicksort original", quickSort)
timeMesure("# Quicksort Mediam of three", quickSortMediamOfThree)
timeMesure("# Quicksort V3. Recursived", quickSortV3)
timeMesure("# Quicksort V. stacks", quickSortVStack)
timeMesure("# Quicksort V. stacks and Updated", quickSortV3VStack)