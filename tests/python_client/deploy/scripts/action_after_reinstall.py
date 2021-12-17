from pymilvus import connections
from utils import *


def task_1():
    """
    task_1:
        before reinstall: create collection and insert data
        after reinstall: get collection, load, search, create index, load, and search
    """
    prefix = "task_1_"
    connections.connect(host="127.0.0.1", port=19530, timeout=60)
    get_collections(prefix)
    load_and_search(prefix)
    create_collections_and_insert_data(prefix)


def task_2():
    """
    task_2:
        before reinstall: create collection, insert data and create index
        after reinstall: get collection, load, search, insert data, create index, load, and search
    """
    prefix = "task_2_"
    connections.connect(host="127.0.0.1", port=19530, timeout=60)
    get_collections(prefix)
    load_and_search(prefix)
    create_collections_and_insert_data(prefix)
    create_index(prefix)
    load_and_search(prefix)


if __name__ == '__main__':
    task_1()
    task_2()