import random
import time
import matplotlib.pyplot as plt

# sys.setrecursionlimit(1500)

def timeMesure(testName, funct, nSize):
        rnumbers = random.sample(range(1, 10**8), nSize)
        snumbers = rnumbers.copy()
        snumbers.sort()
        print(testName)

        avg = 0
        for i in range(3):
            numbers = rnumbers.copy()
            start_time = time.time()
            funct( numbers, 0, len(numbers)-1 )
            end_time = time.time()

            print(f"Time to sort {len(numbers)} elements: {end_time - start_time:.2f} seconds")

            avg += end_time - start_time

            if numbers == snumbers:
                pass
                print("OK - was correct")
            else:
                print("Error: code was wrong")
        
        return avg/3

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




cases = [10**5, 10**6, 10**7]
# cases = [10**1, 10**2, 10**3]

methods = [
    (
       "Quicksort original", quickSort 
    ),
    (
       "Quicksort mediam of three", quickSortMediamOfThree
    ),
    (
       "Quicksort pivot starting from the left", quickSortV3
    ),
    (
       "Quicksort with stacks", quickSortVStack
    ),
    (
       "Quicksort with stacks and from the left", quickSortV3VStack
    ),
]

f = open("results_3.txt", "a")
to_plot_xs = {}
to_plot_ys = {}
to_plot_names = {}
for case in cases:
    for method in methods:
        timeAvg = timeMesure(method[0],method[1], case)
        print(f"Time to sort {case} elements in avg: {timeAvg:.2f} seconds\n")
        f.write(f"{method[0]}, {case}, {timeAvg}\n")

        if method[0] not in to_plot_xs:
            to_plot_xs[method[0]] = []  
        if method[0] not in to_plot_ys:
            to_plot_ys[method[0]] = [] 
        if method[0] not in to_plot_names:
            to_plot_names[method[0]] = method[0] 

        print(method[0])
        print(len(to_plot_xs))

        to_plot_xs[method[0]].append(case) 
        to_plot_ys[method[0]].append(timeAvg) 
        

f.close()

for key in to_plot_names:
    plt.plot(to_plot_xs[key], to_plot_ys[key], label = to_plot_names[key])
plt.title('N size vs Time for multiple quicksort implementations')
plt.xlabel('N size')
plt.ylabel('Time in seconds')
plt.legend()
plt.grid(True)
plt.show()

    

    # timeMesure("# Quicksort original", quickSort)
    # timeMesure("# Quicksort Mediam of three", quickSortMediamOfThree)
    # timeMesure("# Quicksort V3. Recursived", quickSortV3)
    # timeMesure("# Quicksort V. stacks", quickSortVStack)
    # timeMesure("# Quicksort V. stacks and Updated", quickSortV3VStack)