import random
import sys

def calculate_insertion_prob(queue_num, total_queues, time, max_time):
    # Insertion probability grows from 10% to 70% and peaks based on the queue/stack number
    peak_time = (queue_num / total_queues) * max_time
    if time <= peak_time:
        return 0.1 + 0.6 * (time / peak_time)  # Linear growth to peak
    else:
        return 0.7 - 0.7 * ((time - peak_time) / (max_time - peak_time))  # Linear decrease to 0%

def calculate_deletion_prob(queue_num, total_queues, time, max_time):
    # Deletion probability peaks 10% after insertion and decreases to 10% by the end
    peak_time = (queue_num / total_queues) * max_time * 1.1  # Peak happens 10% later
    if time <= peak_time:
        return 0.1 + 0.6 * (time / peak_time)  # Linear growth to peak
    else:
        return 0.7 - 0.6 * ((time - peak_time) / (max_time - peak_time)) + 0.1  # Decrease to 10%

def generate_operation_sequence(q, s, max_time=10000):
    operations = []
    unique_id = 1  # This will be the unique ID inserted into queues/stacks

    # Iterate through time from t00001 to t10000
    for time in range(1, max_time + 1):
        timestamp = f"t{time:05d}"

        # For each queue, decide insert or delete based on probabilities
        for queue_num in range(1, q + 1):
            insertion_prob = calculate_insertion_prob(queue_num, q, time, max_time)
            deletion_prob = calculate_deletion_prob(queue_num, q, time, max_time)

            if random.random() < insertion_prob:
                operations.append(f"{timestamp} enqueue Q{queue_num}, {unique_id}")
                unique_id += 1
            elif random.random() < deletion_prob:
                operations.append(f"{timestamp} dequeue Q{queue_num}")

        # For each stack, decide insert or delete based on probabilities
        for stack_num in range(1, s + 1):
            insertion_prob = calculate_insertion_prob(stack_num, s, time, max_time)
            deletion_prob = calculate_deletion_prob(stack_num, s, time, max_time)

            if random.random() < insertion_prob:
                operations.append(f"{timestamp} push S{stack_num}, {unique_id}")
                unique_id += 1
            elif random.random() < deletion_prob:
                operations.append(f"{timestamp} pop S{stack_num}")

    return operations

# Example usage
if __name__ == "__main__":
    
    q = int(sys.argv[1])
    s = int(sys.argv[2])
    operations = generate_operation_sequence(q, s)

    # Output the generated operations
    for op in operations:
        print(op)