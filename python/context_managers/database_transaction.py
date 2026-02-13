# class DatabaseTransaction:
#     def __enter__(self):
#         print("Starting Transaction...")
#         return self

#     def __exit__(self, exc_type, exc_val, exc_tb):
#         if exc_val is None:
#             print("Commit")
#         else:
#             print("Rollback")
#             return False


# try:
#     with DatabaseTransaction():
#         print("Writing data...")
#         raise ValueError("Oops! Something went wrong.")
#         print("This will never print.")
# except Exception as e:
#     print(f"Caught outside: {e}")

from contextlib import contextmanager


@contextmanager
def database_transaction():
    print("Starting Transaction...")
    try:
        yield
        print("Commit")
    except Exception as e:
        print("Rollback")
        raise e


try:
    with database_transaction():
        print("Writing data...")
        raise ValueError("Oops! Something went wrong.")
        print("This will never print.")
except Exception as e:
    print(f"Caught outside: {e}")