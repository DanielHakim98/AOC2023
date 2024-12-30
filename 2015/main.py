import importlib.util
import os


def run_functions(base_folder: str):
    for day_folder in os.listdir(base_folder):
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
                        display = f"{part}: {getattr(module, part)(data_path)}"
                        print(display)
                    else:
                        print(f"Function '{part}' not found in {file_path}.")
        print()


if __name__ == "__main__":
    run_functions("days")
