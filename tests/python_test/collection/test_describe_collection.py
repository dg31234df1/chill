import pytest
import logging
import time
from utils import *
from constants import *

uid = "describe_collection"


class TestDescribeCollection:

    @pytest.fixture(
        scope="function",
        params=gen_single_filter_fields()
    )
    def get_filter_field(self, request):
        yield request.param

    @pytest.fixture(
        scope="function",
        params=gen_single_vector_fields()
    )
    def get_vector_field(self, request):
        yield request.param

    @pytest.fixture(
        scope="function",
        params=gen_simple_index()
    )
    def get_simple_index(self, request, connect):
        logging.getLogger().info(request.param)
        if str(connect._cmd("mode")) == "CPU":
            if request.param["index_type"] in index_cpu_not_support():
                pytest.skip("sq8h not support in CPU mode")
        return request.param

    """
    ******************************************************************
      The following cases are used to test `describe_collection` function, no data in collection
    ******************************************************************
    """

    def test_collection_fields(self, connect, get_filter_field, get_vector_field):
        '''
        target: test create normal collection with different fields, check info returned
        method: create collection with diff fields: metric/field_type/..., calling `describe_collection`
        expected: no exception raised, and value returned correct
        '''
        filter_field = get_filter_field
        vector_field = get_vector_field
        collection_name = gen_unique_str(uid)
        fields = {
            "fields": [filter_field, vector_field],
            # "segment_row_limit": default_segment_row_limit
        }
        connect.create_collection(collection_name, fields)
        res = connect.describe_collection(collection_name)
        assert res['auto_id'] == True
        # assert res['segment_row_limit'] == default_segment_row_limit
        assert len(res["fields"]) == 2
        for field in res["fields"]:
            if field["type"] == filter_field:
                assert field["name"] == filter_field["name"]
            elif field["type"] == vector_field:
                assert field["name"] == vector_field["name"]
                assert field["params"] == vector_field["params"]

    def test_describe_collection_after_index_created(self, connect, collection, get_simple_index):
        connect.create_index(collection, default_float_vec_field_name, get_simple_index)
        res = connect.describe_collection(collection)
        for field in res["fields"]:
            if field["name"] == default_float_vec_field_name:
                index = field["indexes"][0]
                assert index["index_type"] == get_simple_index["index_type"]
                assert index["metric_type"] == get_simple_index["metric_type"]

    @pytest.mark.level(2)
    def test_describe_collection_without_connection(self, collection, dis_connect):
        '''
        target: test get collection info, without connection
        method: calling get collection info with correct params, with a disconnected instance
        expected: get collection info raise exception
        '''
        with pytest.raises(Exception) as e:
            dis_connect.describe_collection(collection)

    def test_describe_collection_not_existed(self, connect):
        '''
        target: test if collection not created
        method: random a collection name, create this collection then drop it,
            assert the value returned by describe_collection method
        expected: False
        '''
        collection_name = gen_unique_str(uid)
        connect.create_collection(collection_name, default_fields)
        connect.describe_collection(collection_name)
        connect.drop_collection(collection_name)
        with pytest.raises(Exception) as e:
            connect.describe_collection(collection_name)

    @pytest.mark.level(2)
    def test_describe_collection_multithread(self, connect):
        '''
        target: test create collection with multithread
        method: create collection using multithread,
        expected: collections are created
        '''
        threads_num = 4
        threads = []
        collection_name = gen_unique_str(uid)
        connect.create_collection(collection_name, default_fields)

        def get_info():
            connect.describe_collection(collection_name)

        for i in range(threads_num):
            t = TestThread(target=get_info)
            threads.append(t)
            t.start()
            time.sleep(0.2)
        for t in threads:
            t.join()

    """
    ******************************************************************
      The following cases are used to test `describe_collection` function, and insert data in collection
    ******************************************************************
    """

    def test_describe_collection_fields_after_insert(self, connect, get_filter_field, get_vector_field):
        '''
        target: test create normal collection with different fields, check info returned
        method: create collection with diff fields: metric/field_type/..., calling `describe_collection`
        expected: no exception raised, and value returned correct
        '''
        filter_field = get_filter_field
        vector_field = get_vector_field
        collection_name = gen_unique_str(uid)
        fields = {
            "fields": [filter_field, vector_field],
            # "segment_row_limit": default_segment_row_limit
        }
        connect.create_collection(collection_name, fields)
        entities = gen_entities_by_fields(fields["fields"], default_nb, vector_field["params"]["dim"])
        res_ids = connect.insert(collection_name, entities)
        connect.flush([collection_name])
        res = connect.describe_collection(collection_name)
        assert res['auto_id'] == True
        # assert res['segment_row_limit'] == default_segment_row_limit
        assert len(res["fields"]) == 2
        for field in res["fields"]:
            if field["type"] == filter_field:
                assert field["name"] == filter_field["name"]
            elif field["type"] == vector_field:
                assert field["name"] == vector_field["name"]
                assert field["params"] == vector_field["params"]



class TestDescribeCollectionInvalid(object):
    """
    Test describe collection with invalid params
    """
    @pytest.fixture(
        scope="function",
        params=gen_invalid_strs()
    )
    def get_collection_name(self, request):
        yield request.param

    @pytest.mark.level(2)
    def test_describe_collection_with_invalid_collection_name(self, connect, get_collection_name):
        collection_name = get_collection_name
        with pytest.raises(Exception) as e:
            connect.describe_collection(collection_name)

    @pytest.mark.level(2)
    def test_describe_collection_with_empty_collection_name(self, connect):
        collection_name = ''
        with pytest.raises(Exception) as e:
            connect.describe_collection(collection_name)

    @pytest.mark.level(2)
    def test_describe_collection_with_none_collection_name(self, connect):
        collection_name = None
        with pytest.raises(Exception) as e:
            connect.describe_collection(collection_name)
