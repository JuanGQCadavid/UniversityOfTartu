import numpy as np
from PIL import Image, ImageDraw, ImageFont
import datetime
import math
import sys


def draw_tsp_path(original_file, method):
    path = []
    with open(original_file, 'r') as f:
        lines = f.readlines()[1:]
        for line in lines:
            _, x, y = line.strip().split()
            path.append((int(x), int(y)))

    print(len(path))
    image_size = 1100
    path_width = 3
    border = 25

    # Calculate path length
    path_length = sum(math.dist(path[i], path[i + 1]) for i in range(len(path) - 1))
    path_length += math.dist(path[-1], path[0])  # Closing the loop

    # Create image for the path
    image = Image.new("RGB", (image_size, image_size), "white")
    draw = ImageDraw.Draw(image)

    # Draw bounding box
    bounding_box = [border, border, image_size - border, image_size - border]
    draw.rectangle(bounding_box, outline="lightgray", width=1)

    # Draw points
    for x, y in path:
        upper_left = (x - 7 // 2, y - 7 // 2)
        lower_right = (x + 7 // 2, y + 7 // 2)
        draw.ellipse([upper_left, lower_right], fill="black")

    # Draw path
    for i in range(len(path) - 1):
        draw.line([path[i], path[i + 1]], fill="blue", width=path_width)
    draw.line([path[-1], path[0]], fill="blue", width=path_width)  # Closing the loop

    # Add path length and method to the image
    try:
        font = ImageFont.truetype("arial.ttf", 12)
        large_font = ImageFont.truetype("arial.ttf", 24)  # Larger font for filename
    except IOError:
        font = ImageFont.load_default()
        large_font = ImageFont.load_default()

    # Save the image
    num_points = len(path)
    num_points_str = f"{num_points:04d}"
    image_path = f"output/tsp_points_image_{num_points_str}_{method}.png"


    # Print filename and path length/method at the top of the image
    draw.text((50, 10), f"tsp_points_{method}_order_{num_points_str}", fill="black", font=large_font)
    draw.text((50, 40), f"{method.upper()} - Path Length: {path_length:.2f}", fill="black", font=font)

    image.save(image_path)
    # Output details
    print(f"{method.upper()} image generated:", image_path)
    print(f"File: tsp_points_{method}_order_{num_points_str}, Method: {method.upper()}, Distance: {path_length:.2f}")

    # Save path order to file
    path_file_path = f"output/tsp_points_{method}_order_{num_points_str}.txt"
    with open(path_file_path, 'w') as f:
        f.write(f"# {method.upper()} Method - Path Length: {path_length:.2f}\n")
        for i, (x, y) in enumerate(path):
            f.write(f"p{i+1} {x} {y}\n")
    print(f"{method.upper()} path order saved:", path_file_path)


# Run the main program
if __name__ == "__main__":
    fileName = sys.argv[1]
    version = sys.argv[2]
    print(fileName, version)
    draw_tsp_path(fileName, version)