import pytest

from base.client_base import TestcaseBase
from common import common_func as cf
from common import common_type as ct
from utils.util_log import test_log as log
from common.common_type import CaseLabel, CheckTasks

prefix = "delete"
half_nb = ct.default_nb // 2
tmp_nb = 100
tmp_expr = f'{ct.default_int64_field_name} in {[0]}'


@pytest.mark.skip(reason="Delete function is not implemented")
class TestDeleteParams(TestcaseBase):
    """
    Test case of delete interface
    def delete(expr, partition_name=None, timeout=None, **kwargs)
    return MutationResult
    Only the `in` operator is supported in the expr
    """

    @pytest.mark.skip(reason="Issues #10273")
    @pytest.mark.tags(CaseLabel.L0)
    @pytest.mark.parametrize('is_binary', [False, True])
    def test_delete_entities(self, is_binary):
        """
        target: test delete data from collection
        method: 1.create and insert nb
                2. delete half of nb
        expected: assert num entities
        """
        # init collection with default_nb default data
        collection_w, _, _, ids = self.init_collection_general(prefix, insert_data=True, is_binary=is_binary)[0:4]
        expr = f'{ct.default_int64_field_name} in {ids[:half_nb]}'

        # delete half of data
        del_res = collection_w.delete(expr)[0]
        assert del_res.delete_count == half_nb
        collection_w.num_entities

        # query with deleted ids
        query_res = collection_w.query(expr)[0]
        assert len(query_res) == 0

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_without_connection(self):
        """
        target: test delete without connect
        method: delete after remove connection
        expected: raise exception
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]

        # remove connection and delete
        self.connection_wrap.remove_connection(ct.default_alias)
        res_list, _ = self.connection_wrap.list_connections()
        assert ct.default_alias not in res_list
        error = {ct.err_code: 0, ct.err_msg: "should create connect first"}
        collection_w.delete(expr=tmp_expr, check_task=CheckTasks.err_res, check_items=error)

    @pytest.mark.skip(reason="Issue #10271")
    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_expr_none(self):
        """
        target: test delete with None expr
        method: delete with None expr
        expected: raise exception
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        error = {ct.err_code: 0, ct.err_msg: "todo"}
        collection_w.delete(expr=None, check_task=CheckTasks.err_res, check_items=error)
        assert collection_w.num_entities == tmp_nb

    @pytest.mark.tags(CaseLabel.L2)
    @pytest.mark.parametrize("expr", [1, "12-s", "中文", [], ()])
    @pytest.mark.skip(reason="Issues #10271")
    def test_delete_expr_non_string(self, expr):
        """
        target: test delete with non-string expression
        method: delete with non-string expr
        expected: raise exception
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(expr, check_task=CheckTasks.err_res, check_items=error)

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_expr_empty_value(self):
        """
        target: test delete with empty array expr
        method: delete with expr: "id in []"
        expected: assert num entities
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        expr = f'{ct.default_int64_field_name} in {[]}'

        # delete empty entities
        collection_w.delete(expr)
        assert collection_w.num_entities == tmp_nb

    @pytest.mark.skip(reason="Issues #10273")
    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_expr_single(self):
        """
        target: test delete with one value
        method: delete with expr: "id in [0]"
        expected: Describe num entities by one
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        expr = f'{ct.default_int64_field_name} in {[0]}'
        del_res, _ = collection_w.delete(expr)
        assert del_res.delete_count == 1
        assert collection_w.num_entities == tmp_nb - 1

    @pytest.mark.skip(reason="Issues #10273")
    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_expr_all_values(self):
        """
        target: test delete with all values
        method: delete with expr: "id in [all]"
        expected: num entities becomes zero
        """
        # init collection with default_nb default data
        collection_w, _, _, ids = self.init_collection_general(prefix, insert_data=True)[0:4]
        expr = f'{ct.default_int64_field_name} in {ids}'
        del_res, _ = collection_w.delete(expr)
        assert del_res.delete_count == ct.default_nb

        assert collection_w.num_entities == 0
        assert collection_w.is_empty

    @pytest.mark.skip(reason="Issues #10277")
    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_not_existed_values(self):
        """
        target: test delete not existed values
        method: delete data not in the collection
        expected: raise exception
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        expr = f'{ct.default_int64_field_name} in {[tmp_nb]}'

        # raise exception
        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(expr=expr, check_task=CheckTasks.err_res, check_items=error)

    @pytest.mark.tags(CaseLabel.L2)
    @pytest.mark.skip(reason="Issues #10277")
    def test_delete_part_existed_values(self):
        """
        target: test delete with part not existed values
        method: delete data part not in the collection
        expected: delete any entities
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        expr = f'{ct.default_int64_field_name} in {[0, tmp_nb]}'
        res, _ = collection_w.delete(expr)
        assert res.delete_count == 0
        assert collection_w.num_entities == tmp_nb

    @pytest.mark.tags(CaseLabel.L1)
    @pytest.mark.skip(reason="Issues #10271")
    def test_delete_expr_inconsistent_values(self):
        """
        target: test delete with inconsistent type values
        method: delete with non-int64 type values
        expected: raise exception
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        expr = f'{ct.default_int64_field_name} in {[0.0, 1.0]}'

        # raise exception
        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(expr=expr, check_task=CheckTasks.err_res, check_items=error)

    @pytest.mark.tags(CaseLabel.L2)
    @pytest.mark.skip(reason="Issues #10271")
    def test_delete_expr_mix_values(self):
        """
        target: test delete with mix type values
        method: delete with int64 and float values
        expected: raise exception
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        expr = f'{ct.default_int64_field_name} in {[0, 1.0]}'

        # raise exception
        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(expr=expr, check_task=CheckTasks.err_res, check_items=error)

    @pytest.mark.tags(CaseLabel.L0)
    def test_delete_partition(self):
        """
        target: test delete from partition
        method: delete with partition names
        expected: verify partition entities is deleted
        """
        # init collection and partition
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        partition_w = self.init_partition_wrap(collection_wrap=collection_w)

        # insert data to partition
        df = cf.gen_default_dataframe_data(tmp_nb)
        partition_w.insert(df)
        assert partition_w.num_entities == tmp_nb
        collection_w.load()
        del_res, _ = collection_w.delete(tmp_expr, partition_name=[partition_w.name])

        # verify partition num entities
        assert del_res.delete_cnt == 1
        assert partition_w.num_entities == tmp_nb - 1
        assert collection_w.num_entities == tmp_nb - 1

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_default_partition(self):
        """
        target: test delete from default partition
        method: delete with partition names is _default
        expected: assert delete successfully
        """
        # init collection with tmp_nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        del_res, _ = collection_w.delete(tmp_expr, partition_name=[ct.default_partition_name])
        assert del_res.delete_cnt == 1
        assert collection_w.num_entities == tmp_nb - 1

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_empty_partition_names(self):
        """
        target: test delete none partition
        method: delete with None partition names
        expected: delete from all partitions
        """
        # init collection and partition
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        partition_w = self.init_partition_wrap(collection_wrap=collection_w)

        # insert data to partition
        df = cf.gen_default_dataframe_data(tmp_nb)
        partition_w.insert(df)
        collection_w.insert(df)
        assert collection_w.num_entities == tmp_nb * 2
        collection_w.load()
        del_res, _ = collection_w.delete(tmp_expr, partition_name=[])

        # verify partition num entities
        assert del_res.delete_cnt == 0
        assert collection_w.num_entities == tmp_nb


@pytest.mark.skip(reason="Waiting for development")
class TestDeleteOperation(TestcaseBase):
    """
    ******************************************************************
      The following cases are used to test delete interface operations
    ******************************************************************
    """

    @pytest.mark.skip(reason="Issues #10277")
    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_empty_collection(self):
        """
        target: test delete entities from an empty collection
        method: create an collection and delete entities
        expected: raise exception
        """
        c_name = cf.gen_unique_str(prefix)
        collection_w = self.init_collection_wrap(name=c_name)
        # raise exception
        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(expr=tmp_expr, check_task=CheckTasks.err_res, check_items=error)
        collection_w.delete(tmp_expr)

    @pytest.mark.skip(reason="Delete function is not implemented")
    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_entities_repeatedly(self):
        """
        target: test delete entities twice
        method: delete with same expr twice
        expected: raise exception
        """
        # init collection with nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        # assert delete successfully
        collection_w.delete(expr=tmp_expr)
        # raise exception
        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(expr=tmp_expr, check_task=CheckTasks.err_res, check_items=error)

    @pytest.mark.skip(reason="Delete function is not implemented")
    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_after_index(self):
        """
        target: test delete after creating index
        method: 1.create index 2.delete entities
        expected: assert index and num entities
        """
        # init collection with nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]

        # create index
        index_params = {"index_type": "IVF_SQ8", "metric_type": "L2", "params": {"nlist": 64}}
        collection_w.create_index(ct.default_float_vec_field_name, index_params)
        assert collection_w.has_index()[0]
        # delete entity
        collection_w.delete(tmp_expr)
        assert collection_w.num_entities == tmp_nb - 1
        assert collection_w.has_index()

    @pytest.mark.skip(reason="Delete function is not implemented")
    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_query(self):
        """
        target: test delete and query
        method: query entity after it was deleted
        expected: query result is empty
        """
        # init collection with nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        # assert delete successfully
        collection_w.delete(expr=tmp_expr)
        res = collection_w.query(expr=tmp_expr)[0]
        assert len(res) == 0

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_search(self):
        """
        target: test delete and search
        method: search entities after it was deleted
        expected: deleted entity is not in the search result
        """
        # init collection with nb default data
        collection_w = self.init_collection_general(prefix, insert_data=True)[0]
        entity, _ = collection_w.query(tmp_expr, output_fields=["%"])
        default_search_exp = "int64 >= 0"
        search_res, _ = collection_w.search(entity[ct.default_float_vec_field_name], ct.default_float_vec_field_name,
                                            ct.default_search_params, limit=1)
        # assert search results contains entity
        # todo
        del_res, _ = collection_w.delete(tmp_expr)
        assert del_res.delete_cnt == 1
        search_res_2, _ = collection_w.search(entity[ct.default_float_vec_field_name], ct.default_float_vec_field_name,
                                              ct.default_search_params, limit=1)
        # assert search result is not equal to entity
        # todo

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_expr_repeated_values(self):
        """
        target: test delete with repeated values
        method: 1.insert data with unique primary keys
                2.delete with repeated values: 'id in [0, 0]'
        expected: delete one entity
        """
        # init collection with nb default data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]
        expr = f'{ct.default_int64_field_name} in {[0, 0, 0]}'
        del_res, _ = collection_w.delete(expr)
        assert del_res.delete_count == 1
        assert collection_w.num_entities == tmp_nb - 1

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_duplicate_primary_keys(self):
        """
        target: test delete from duplicate primary keys
        method: 1.insert data with dup ids
                2.delete with repeated or not values
        expected: delete all entities
        """
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        df = cf.gen_default_dataframe_data(nb=tmp_nb)
        df[ct.default_int64_field_name] = 0
        collection_w.insert(df)
        assert collection_w.num_entities == tmp_nb
        collection_w.load()
        del_res, _ = collection_w.delete(tmp_expr)
        assert del_res.delete_cnt == tmp_nb
        assert collection_w.is_empty

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_empty_partition(self):
        """
        target: test delete empty partition
        method: delete from an empty partition
        expected: raise exception
        """
        # init collection and partition
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        partition_w = self.init_partition_wrap(collection_wrap=collection_w)

        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(tmp_expr, partition_name=[partition_w.name], check_task=CheckTasks.err_res,
                            check_items=error)

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_not_existed_partition(self):
        """
        target: test delete from an not existed partition
        method: delete from an fake partition
        expected: raise exception
        """
        # init collection with tmp_nb data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]

        # raise exception
        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(tmp_expr, partition_name=[ct.default_tag], check_task=CheckTasks.err_res, check_items=error)

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_part_not_existed_partition(self):
        """
        target: test delete from part not existed partitions
        method: delete with part not existed partition name
        expected: raise exception
        """
        # init collection with tmp_nb data
        collection_w = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True)[0]

        # raise exception
        error = {ct.err_code: 0, ct.err_msg: "..."}
        collection_w.delete(tmp_expr, partition_name=[ct.default_partition_name, ct.default_tag],
                            check_task=CheckTasks.err_res, check_items=error)
        assert collection_w.num_entities == tmp_nb

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_multi_partitions(self):
        """
        target: delete entities from multi partitions
        method: 1.insert pk (primary keys) [0,half) to tag partition
                2.insert pk [half,nb) to default partition
                3.delete pk values 0 and half from two partition
        expected: two entities are deleted
        """
        half = ct.default_nb // 2
        collection_w, partition_w, _, _ = self.insert_entities_into_two_partitions_in_half(half)
        expr = f'{ct.default_int64_field_name} in {[0, half]}'
        collection_w.delete(expr, partition_name=[ct.default_partition_name, partition_w.name])
        assert collection_w.num_entities == ct.default_nb - 2

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_from_partition_with_another_ids(self):
        """
        target: delete another partition entities from partition
        method: 1.insert nb entities into two partitions in half
                2.delete entities from partition_1 with partition_2 values
        expected: No entities will be deleted
        """
        half = ct.default_nb // 2
        collection_w, partition_w, _, _ = self.insert_entities_into_two_partitions_in_half(half)

        # delete entities from another partition
        expr = f'{ct.default_int64_field_name} in {[0]}'
        collection_w.delete(expr, partition_name=[ct.default_partition_name])
        assert collection_w.num_entities == ct.default_nb

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_from_partition_with_own_ids(self):
        """
        target: test delete own pk from partition
        method: 1.insert nb entities into two partitions in half
                2.delete entities from partition_1 with partition_1 values
        expected: verify entities is deleted
        """
        half = ct.default_nb // 2
        collection_w, partition_w, _, _ = self.insert_entities_into_two_partitions_in_half(half)

        # delete entities from another partition
        expr = f'{ct.default_int64_field_name} in {[0]}'
        collection_w.delete(expr, partition_name=[partition_w.name])
        assert collection_w.num_entities == ct.default_nb - 1

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_from_partitions_with_same_ids(self):
        """
        target: test delete same ids from two partitions with same data
        method: 1.insert same nb data into two partitions
                2.delete same ids from two partitions
        expected: The data in both partitions will be deleted
        """
        # init collection and partition
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        partition_w = self.init_partition_wrap(collection_wrap=collection_w)

        # insert same data into partition_w and default partition
        df = cf.gen_default_dataframe_data(tmp_nb)
        collection_w.insert(df)
        partition_w.insert(df)
        assert collection_w.num_entities == tmp_nb * 2

        # delete same ids from two partition
        collection_w.delete(tmp_expr, partition_name=[ct.default_partition_name, partition_w.name])
        assert collection_w.num_entities == tmp_nb - 2

    @pytest.mark.tags(CaseLabel.L0)
    def test_delete_auto_id_collection(self):
        """
        target: test delete from auto_id collection
        method: delete entities from auto_id=true collection
        expected: versify delete successfully
        """
        # init an auto_id collection and insert tmp_nb data
        collection_w, _, _, ids = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True, auto_id=True)

        # delete with insert ids
        expr = f'{ct.default_int64_field_name} in {[ids[0]]}'
        res, _ = collection_w.delete(expr)

        # verify delete result
        collection_w.load()
        assert res.delete_count == 1
        query_res, _ = collection_w.query(expr)
        assert len(query_res) == 0

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_query_without_loading(self):
        """
        target: test delete and query without loading
        method: 1.insert and flush data
                2.delete ids
                3.query without loading
        expected: Raise exception
        """
        # create collection, insert data without flush
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        df = cf.gen_default_dataframe_data(tmp_nb)
        collection_w.insert(df)
        assert collection_w.num_entities == tmp_nb

        # query without loading and raise exception
        error = {ct.err_code: 1, ct.err_msg: f"collection {collection_w.name} was not loaded into memory"}
        collection_w.delete(expr=tmp_expr, check_task=CheckTasks.err_res, check_items=error)

    @pytest.mark.tags(CaseLabel.L1)
    def test_delete_without_flush(self):
        """
        target: test delete without flush
        method: 1.insert data and no flush
                2.delete ids from collection
                3.load and query with id
        expected: No query result
        """
        # create collection, insert data without flush
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        df = cf.gen_default_dataframe_data(tmp_nb)
        collection_w.insert(df)

        # delete
        del_res, _ = collection_w.delete(tmp_expr)
        assert del_res.delete_count == 1

        # load and query with id
        collection_w.load()
        query_res = collection_w.query(tmp_expr)[0]
        assert len(query_res) == 0

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_data_from_growing_segment(self):
        """
        target: test delete entities from growing segment
        method: 1.create collection
                2.load collection
                3.insert data and delete ids
                4.query deleted ids
        expected: No query result
        """
        # create collection
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        # load collection and the queryNode watch the insertChannel
        collection_w.load()
        # insert data
        df = cf.gen_default_dataframe_data(tmp_nb)
        collection_w.insert(df)
        # delete id 0
        del_res = collection_w.delete(tmp_expr)[0]
        assert del_res.delete_count == 1
        # query id 0
        query_res = collection_w.query(tmp_expr)[0]
        assert len(query_res) == 0

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_after_load_sealed_segment(self):
        """
        target: test delete sealed data and get deleteMsg from insertChannel
        method: 1.create, insert and flush data
                2.load collection
                3.delete id without flush
                4.query deleted ids (queryNode get deleted ids from channel not persistence)
        expected: Delete successfully and no query result
        """
        # create collection and insert flush data
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        df = cf.gen_default_dataframe_data(tmp_nb)
        collection_w.insert(df)
        assert collection_w.num_entities == tmp_nb

        # load collection and queryNode subscribe channel
        collection_w.load()

        # delete ids and query
        collection_w.delete(tmp_expr)
        query_res = collection_w.query(tmp_expr)[0]
        assert len(query_res) == 0

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_sealed_segment_with_flush(self):
        """
        target: test delete data from sealed segment and flush delta log
        method: 1.create and insert and flush data
                2.delete entities and flush (delta log persistence)
                3.load collection (load data and delta log)
                4.query deleted ids
        expected: No query result
        """
        # create collection
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        # insert and flush data
        df = cf.gen_default_dataframe_data(tmp_nb)
        collection_w.insert(df)
        assert collection_w.num_entities == tmp_nb
        # delete id 0 and flush
        del_res = collection_w.delete(tmp_expr)[0]
        assert del_res.delete_count == 1
        collection_w.num_entities
        # load and query id 0
        collection_w.load()
        query_res = collection_w.query(tmp_expr)[0]
        assert len(query_res) == 0

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_insert_same_entity(self):
        """
        target: test delete and insert same entity
        method: 1.delete entity one
                2.insert entity one
                3.query  entity one
        expected: verify query result
        """
        # init collection and insert data without flush
        collection_w = self.init_collection_wrap(name=cf.gen_unique_str(prefix))
        df = cf.gen_default_dataframe_data(tmp_nb)
        collection_w.insert(df)

        # delete
        del_res, _ = collection_w.delete(tmp_expr)
        assert del_res.delete_count == 1
        assert collection_w.num_entities == tmp_nb

        # insert entity with primary key 0
        collection_w.insert(df[:1])

        # query entity one
        collection_w.load()
        res = df.iloc[0:1, :1].to_dict('records')
        collection_w.query(tmp_expr, check_task=CheckTasks.check_query_results, check_items={'exp_res': res})

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_entity_loop(self):
        """
        target: test delete all entities one by one in a loop
        method: delete data one by one for a loop
        expected: No exception
        """
        # init an auto_id collection and insert tmp_nb data
        collection_w, _, _, ids = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True, auto_id=True)

        for del_id in ids:
            expr = f'{ct.default_int64_field_name} in {[del_id]}'
            res, _ = collection_w.delete(expr)
            assert res.delete_count == 1

        # query with first and last id
        collection_w.load()
        first_expr = f'{ct.default_int64_field_name} in {ids[0]}'
        first_res, _ = collection_w.query(first_expr)
        assert len(first_res) == 0
        last_expr = f'{ct.default_int64_field_name} in {ids[-1]}'
        last_res, _ = collection_w.query(last_expr)
        assert len(last_res) == 0

    @pytest.mark.tags(CaseLabel.L2)
    def test_delete_flush_loop(self):
        """
        target: test delete and flush in a loop
        method: in a loop, delete batch and flush, until delete all entities
        expected: No exception
        """
        # init an auto_id collection and insert tmp_nb data
        collection_w, _, _, ids = self.init_collection_general(prefix, nb=tmp_nb, insert_data=True, auto_id=True)

        batch = 10
        for i in range(tmp_nb // batch):
            expr = f'{ct.default_int64_field_name} in {ids[i * batch: (i + 1) * batch]}'
            res, _ = collection_w.delete(expr)
            assert res.delete_count == batch
            assert collection_w.num_entities == tmp_nb

        # query with first and last id
        collection_w.load()
        first_expr = f'{ct.default_int64_field_name} in {ids[0]}'
        first_res, _ = collection_w.query(first_expr)
        assert len(first_res) == 0
        last_expr = f'{ct.default_int64_field_name} in {ids[-1]}'
        last_res, _ = collection_w.query(last_expr)
        assert len(last_res) == 0

    @pytest.mark.tags(CaseLabel.L2)
    @pytest.mark.skip(reason="TODO")
    def test_delete_multi_threading(self):
        """
        target: test delete multi threading
        method: delete multi threading
        expected: delete successfully
        """
        pass
