import argparse
import importlib.util
import os
import time


def run_functions(base_folder: str, days: list):
    for day_folder in os.listdir(base_folder):
        if days and day_folder not in days:
            continue

        folder_path = os.path.join(base_folder, day_folder)

        # Check if the folder contains a Python file matching the pattern
        if os.path.isdir(folder_path):
            file_path = os.path.join(folder_path, f"{day_folder}.py")
            data_path = os.path.join(folder_path, f"{day_folder}.txt")

            if os.path.exists(file_path):
                print(f"Running '{day_folder}'")
                print("========================")
                for part in ["part_1", "part_2"]:
                    # Import the module dynamically and run the function
                    module_name = f"{base_folder}.{day_folder}.{day_folder}"
                    module = importlib.import_module(module_name)

                    if hasattr(module, part):
                        start_time = time.time()
                        result = getattr(module, part)(data_path)
                        end_time = time.time()
                        elapsed_time = end_time - start_time
                        display = (
                            f"{part}: {result} (Time taken: {elapsed_time:.4f} seconds)"
                        )
                        print(display)
                    else:
                        print(f"Function '{part}' not found in {file_path}.")
        print()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Run specific days of Advent of Code solutions."
    )
    parser.add_argument("days", nargs="*", help="List of days to run (e.g., 1 2 3 4)")

    args = parser.parse_args()
    days = [f"day_{day}" for day in args.days]

    run_functions("days", days)
