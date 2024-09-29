import matplotlib.pyplot as plt
import networkx as nx
import pandas as pd
import sys

# Constants for indices in the node
INFO = 0
LEFT = 1
RIGHT = 2
COLOR = 3
NULL = -1  # Using 0 as null

def load_tree_from_csv(file_path):
    df = pd.read_csv(file_path)
    tree = []  # Initialize tree with root placeholder

    for _, row in df.iterrows():
        info = row['info']
        left = row['left']
        if left == 0:
          left = NULL
        right = row['right']
        if right == 0:
          right = NULL
        color = row['color']
        tree.append([info, left, right, color])

    print("Tree Loaded:", tree)  # Debugging: print the loaded tree
    return tree

def visualize_tree(tree, filename):
    G = nx.DiGraph()  # Directed graph for tree
    root = 0
    pos = {}  # Dictionary to store positions of nodes

    def add_edges(node, x=0, y=0, x_offset=1):
        if node == NULL:
            return
        pos[tree[node][INFO]] = (x, y)  # Assign position based on x and y coordinates
        left_child = tree[node][LEFT]
        right_child = tree[node][RIGHT]

        if left_child != NULL:
            G.add_edge(tree[node][INFO], tree[left_child][INFO])
            add_edges(left_child, x - x_offset, y - 1, x_offset / 2)  # Move left child to the left
        if right_child != NULL:
            G.add_edge(tree[node][INFO], tree[right_child][INFO])
            add_edges(right_child, x + x_offset, y - 1, x_offset / 2)  # Move right child to the right

    # Start adding edges from the root
    add_edges(root)

    # Get colors for nodes based on their color attribute
    node_colors = ['gray' if tree[node][COLOR] == 'black' else 'lightcoral' if tree[node][COLOR] == 'red' else tree[node][COLOR] for node in range(0, len(tree)) if tree[node][INFO] != 0]
    print("Nodes in Graph:", G.nodes)
    print("Edges in Graph:", G.edges)
    print("Node Colors:", node_colors)
    # Draw the tree with specified positions
    nx.draw(G, pos, with_labels=True, node_size=500, node_color=node_colors, font_size=10, font_weight="bold", arrows=True)
    nx.draw_networkx_labels(G, pos, font_color='white', font_size=10, font_weight="bold")
    # plt.show()
    plt.savefig(filename)  # Save the figure to a file
    plt.close()  # Close the plot to free memory

if __name__ == "__main__":
    csv = sys.argv[1]
    print(csv)
    tree = load_tree_from_csv(csv)
    visualize_tree(tree, f"{csv.split('.csv')[0]}.png")