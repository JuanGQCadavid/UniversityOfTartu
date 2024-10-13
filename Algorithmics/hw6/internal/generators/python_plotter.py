import pandas as pd
import matplotlib.pyplot as plt
import sys

def plot_times(filename):
    # Load CSV data
    data = pd.read_csv(filename)
    
    # Extract data
    nsizes = data['NSizes']
    heapify_times = data['HeapifyTimes']
    bubble_sort_times = data['BubbleSortTimes']
    
    # Plotting
    plt.figure(figsize=(10, 6))
    plt.plot(nsizes, heapify_times, marker='o', label='Heapify')
    plt.plot(nsizes, bubble_sort_times, marker='x', label='Bubble')
    plt.xlabel('N size')
    plt.ylabel('Time in seconds')
    plt.title('Heapify vs bubble sort')
    plt.legend()
    plt.grid(True)
    
    # Save plot
    output_file = filename.replace('.csv', '.png')
    plt.savefig(output_file)
    # plt.show()
    print(f"Plot saved as {output_file}")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python plot_times.py <csv_file>")
        sys.exit(1)
    
    filename = sys.argv[1]
    plot_times(filename)
