# import docker
from pymilvus import connections
from utils import *

connections.connect(host="127.0.0.1", port=19530, timeout=60)

get_collections()

load_and_search()

create_collections_and_insert_data()

create_index()

load_and_search()

