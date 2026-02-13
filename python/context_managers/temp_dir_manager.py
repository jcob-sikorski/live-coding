# The Task: Create a manager that creates a temporary folder on entry,
# provides the path to the user, and deletes the folder (and all its contents)
# on exit.

from contextlib import contextmanager
import shutil
from pathlib import Path


@contextmanager
def temporary_dir_manager(dir_path: str):
    p = Path(dir_path)

    p.mkdir(parents=True, exist_ok=True)

    try:
        yield p
    finally:
        if p.exists():
            shutil.rmtree(p)
            print(f"Successfully cleaned up: {p}")


folder_name = "my_temp_workspace"
with temporary_dir_manager(folder_name) as tmp_path:
    # Now tmp_path is the Path object we yielded
    print(f"Working inside: {tmp_path.absolute()}")

    # Create a dummy file to test recursive deletion
    (tmp_path / "test_file.txt").write_text("Hello World")
    print("Created a file inside the temp directory.")

print("Context exited. The folder is gone.")
