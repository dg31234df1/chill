from utils.util_log import test_log as log
from common import common_type as ct
from common.common_type import CheckTasks, Connect_Object_Name
# from common.code_mapping import ErrorCode, ErrorMessage
from pymilvus_orm import Collection, Partition
from utils.api_request import Error
import check.param_check as pc


class ResponseChecker:
    def __init__(self, response, func_name, check_task, check_items, is_succ=True, **kwargs):
        self.response = response  # response of api request
        self.func_name = func_name  # api function name
        self.check_task = check_task  # task to check response of the api request
        self.check_items = check_items  # check items and expectations that to be checked in check task
        self.succ = is_succ  # api responses successful or not

        self.kwargs_dict = {}  # not used for now, just for extension
        for key, value in kwargs.items():
            self.kwargs_dict[key] = value
        self.keys = self.kwargs_dict.keys()

    def run(self):
        """
        Method: start response checking for milvus API call
        """
        result = True
        if self.check_task is None:
            result = self.assert_succ(self.succ, True)

        elif self.check_task == CheckTasks.err_res:
            result = self.assert_exception(self.response, self.succ, self.check_items)

        elif self.check_task == CheckTasks.check_connection_result:
            result = self.check_value_equal(self.response, self.func_name, self.check_items)

        elif self.check_task == CheckTasks.check_collection_property:
            result = self.check_collection_property(self.response, self.func_name, self.check_items)

        elif self.check_task == CheckTasks.check_partition_property:
            result = self.check_partition_property(self.response, self.func_name, self.check_items)

        elif self.check_task == CheckTasks.check_search_results:
            result = self.check_search_results(self.response, self.check_items)

        # Add check_items here if something new need verify

        return result

    @staticmethod
    def assert_succ(actual, expect):
        assert actual is expect
        return True

    @staticmethod
    def assert_exception(res, actual=True, error_dict=None):
        assert actual is False
        assert len(error_dict) > 0
        if isinstance(res, Error):
            error_code = error_dict[ct.err_code]
            assert res.code == error_code or error_dict[ct.err_msg] in res.message
        else:
            log.error("[CheckFunc] Response of API is not an error: %s" % str(res))
            assert False
        return True

    @staticmethod
    def check_value_equal(res, func_name, params):
        """ check response of connection interface that result is normal """

        if func_name == "list_connections":
            if not isinstance(res, list):
                log.error("[CheckFunc] Response of list_connections is not a list: %s" % str(res))
                assert False

            list_content = params.get("list_content", None)
            if not isinstance(list_content, list):
                log.error("[CheckFunc] Check param of list_content is not a list: %s" % str(list_content))
                assert False

            new_res = pc.get_connect_object_name(res)
            assert pc.list_equal_check(new_res, list_content)

        if func_name == "get_connection_addr":
            dict_content = params.get("dict_content", None)
            assert pc.dict_equal_check(res, dict_content)

        if func_name == "connect":
            class_obj = Connect_Object_Name
            res_obj = type(res).__name__
            assert res_obj == class_obj

        if func_name == "get_connection":
            value_content = params.get("value_content", None)
            res_obj = type(res).__name__ if res is not None else None
            assert res_obj == value_content

        return True

    @staticmethod
    def check_collection_property(collection, func_name, check_items):
        exp_func_name = "init_collection"
        if func_name != exp_func_name:
            log.warning("The function name is {} rather than {}".format(func_name, exp_func_name))
        if not isinstance(collection, Collection):
            raise Exception("The result to check isn't collection type object")
        if len(check_items) == 0:
            raise Exception("No expect values found in the check task")
        name = check_items.get("name", None)
        schema = check_items.get("schema", None)
        num_entities = check_items.get("num_entities", 0)
        primary = check_items.get("primary", None)
        if name:
            assert collection.name == name
        if schema:
            assert collection.schema == schema
        if num_entities == 0:
            assert collection.is_empty
        assert collection.num_entities == num_entities
        assert collection.primary_field == primary
        return True

    @staticmethod
    def check_partition_property(partition, func_name, check_items):
        exp_func_name = "_init_partition"
        if func_name != exp_func_name:
            log.warning("The function name is {} rather than {}".format(func_name, exp_func_name))
        if not isinstance(partition, Partition):
            raise Exception("The result to check isn't partition type object")
        if len(check_items) == 0:
            raise Exception("No expect values found in the check task")
        if check_items.get("name", None):
            assert partition.name == check_items["name"]
        if check_items.get("description", None):
            assert partition.description == check_items["description"]
        if check_items.get("is_empty", None):
            assert partition.is_empty == check_items["is_empty"]
        if check_items.get("num_entities", None):
            assert partition.num_entities == check_items["num_entities"]
        return True

    @staticmethod
    def check_search_results(search_res, check_items):
        """
        target: check the search results
        method: 1. check the query number
                2. check the limit(topK)
                3. check the distance
        expected: check the search is ok
        """
        log.info("search_results_check: checking the searching results")
        if len(search_res) != check_items["nq"]:
            log.error("search_results_check: Numbers of query searched (%d) "
                      "is not equal with expected (%d)"
                      % (len(search_res), check_items["nq"]))
            assert len(search_res) == check_items["nq"]
        else:
            log.info("search_results_check: Numbers of query searched is correct")
        for hits in search_res:
            if len(hits) != check_items["limit"]:
                log.error("search_results_check: limit(topK) searched (%d) "
                          "is not equal with expected (%d)"
                          % (len(hits), check_items["limit"]))
                assert len(hits) == check_items["limit"]
                assert len(hits.ids) == check_items["limit"]
            else:
                log.info("search_results_check: limit (topK) "
                         "searched for each query is correct")
        log.info("search_results_check: search_results_check: checked the searching results")
