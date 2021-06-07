from pymilvus_orm import utility
import sys

sys.path.append("..")
from check.param_check import *
from check.func_check import *
from utils.api_request import api_request


class ApiUtilityWrapper:
    """ Method of encapsulating utility files """

    ut = utility

    def loading_progress(self, collection_name, partition_names=[], using="default", check_res=None, check_params=None):
        func_name = sys._getframe().f_code.co_name
        res, check = api_request([self.ut.loading_progress, collection_name, partition_names, using])
        check_result = ResponseChecker(res, func_name, check_res, check_params, check, collection_name=collection_name,
                                       partition_names=partition_names, using=using).run()
        return res, check_result

    def wait_for_loading_complete(self, collection_name, partition_names=[], timeout=None, using="default", check_res=None, check_params=None):
        func_name = sys._getframe().f_code.co_name
        res, check = api_request([self.ut.wait_for_loading_complete, collection_name, partition_names, timeout, using])
        check_result = ResponseChecker(res, func_name, check_res, check_params, check, collection_name=collection_name,
                                       partition_names=partition_names, timeout=timeout, using=using).run()
        return res, check_result

    def index_building_progress(self, collection_name, index_name="", using="default", check_res=None, check_params=None):
        func_name = sys._getframe().f_code.co_name
        res, check = api_request([self.ut.index_building_progress, collection_name, index_name, using])
        check_result = ResponseChecker(res, func_name, check_res, check_params, check, collection_name=collection_name, index_name=index_name,
                                       using=using).run()
        return res, check_result

    def wait_for_index_building_complete(self, collection_name, index_name="", timeout=None, using="default", check_res=None, check_params=None):
        func_name = sys._getframe().f_code.co_name
        res, check = api_request([self.ut.wait_for_loading_complete, collection_name, index_name, timeout, using])
        check_result = ResponseChecker(res, func_name, check_res, check_params, check, collection_name=collection_name, index_name=index_name,
                                       timeout=timeout, using=using).run()
        return res, check_result

    def has_collection(self, collection_name, using="default", check_res=None, check_params=None):
        func_name = sys._getframe().f_code.co_name
        res, check = api_request([self.ut.has_collection, collection_name, using])
        check_result = ResponseChecker(res, func_name, check_res, check_params, check, collection_name=collection_name, using=using).run()
        return res, check_result

    def has_partition(self, collection_name, partition_name, using="default", check_res=None, check_params=None):
        func_name = sys._getframe().f_code.co_name
        res, check = api_request([self.ut.has_partition, collection_name, partition_name, using])
        check_result = ResponseChecker(res, func_name, check_res, check_params, check, collection_name=collection_name,
                                       partition_name=partition_name, using=using).run()
        return res, check_result

    def list_collections(self, timeout=None, using="default", check_res=None, check_params=None):
        func_name = sys._getframe().f_code.co_name
        res, check = api_request([self.ut.list_collections, timeout, using])
        check_result = ResponseChecker(res, func_name, check_res, check_params, check, timeout=timeout, using=using).run()
        return res, check_result
