import multiprocessing
import numbers
import random
import time
import numpy
import pytest
import pandas as pd
from common import common_func as cf
from common import common_type as ct
from utils.util_log import test_log as log
from base.client_base import TestcaseBase
from common.common_type import CaseLabel, CheckTasks
from base.high_level_api_wrapper import HighLevelApiWrapper
client_w = HighLevelApiWrapper()


prefix = "client_rbac"
user_pre = "user"
role_pre = "role"
root_token = "root:Milvus"
default_nb = ct.default_nb
default_nq = ct.default_nq
default_dim = ct.default_dim
default_limit = ct.default_limit
default_search_exp = "id >= 0"
exp_res = "exp_res"
default_search_field = ct.default_float_vec_field_name
default_search_params = ct.default_search_params
default_primary_key_field_name = "id"
default_vector_field_name = "vector"
default_float_field_name = ct.default_float_field_name
default_bool_field_name = ct.default_bool_field_name
default_string_field_name = ct.default_string_field_name
default_int32_array_field_name = ct.default_int32_array_field_name
default_string_array_field_name = ct.default_string_array_field_name


@pytest.mark.tags(CaseLabel.RBAC)
class TestMilvusClientRbacBase(TestcaseBase):
    """ Test case of rbac interface """

    def teardown_method(self, method):
        """
        teardown method: drop role and user
        """
        log.info("[utility_teardown_method] Start teardown utility test cases ...")
        uri = f"http://{cf.param_info.param_host}:{cf.param_info.param_port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)

        # drop users
        users, _ = self.high_level_api_wrap.list_users()
        for user in users:
            if user != ct.default_user:
                self.high_level_api_wrap.drop_user(user)
        users, _ = self.high_level_api_wrap.list_users()
        assert len(users) == 1

        # drop roles
        roles, _ = self.high_level_api_wrap.list_roles()
        for role in roles:
            if role not in ['admin', 'public']:
                privileges, _ = self.high_level_api_wrap.describe_role(role)
                if privileges:
                    for privilege in privileges:
                        self.high_level_api_wrap.revoke_privilege(role, privilege["object_type"],
                                                                  privilege["privilege"], privilege["object_name"])
                self.high_level_api_wrap.drop_role(role)
        roles, _ = self.high_level_api_wrap.list_roles()
        assert len(roles) == 2

        super().teardown_method(method)

    def test_milvus_client_connect_using_token(self, host, port):
        """
        target: test init milvus client using token
        method: init milvus client with only token
        expected: init successfully
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        # check success link
        res = client_w.list_collections(client)[0]
        assert res == []

    def test_milvus_client_connect_using_user_password(self, host, port):
        """
        target: test init milvus client using user and password
        method: init milvus client with user and password
        expected: init successfully
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user,
                                                                password=ct.default_password)
        # check success link
        res = client_w.list_collections(client)[0]
        assert res == []

    def test_milvus_client_create_user(self, host, port):
        """
        target: test milvus client api create_user
        method: create user
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user, password=ct.default_password)
        user_name = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        # check
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, user=user_name, password=password)
        res = client_w.list_collections(client)[0]
        assert res == []

    def test_milvus_client_drop_user(self, host, port):
        """
        target: test milvus client api drop_user
        method: drop user
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user, password=ct.default_password)
        user_name = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        # drop user that exists
        self.high_level_api_wrap.drop_user(user_name=user_name)
        # drop user that not exists
        not_exist_user_name = cf.gen_unique_str(user_pre)
        self.high_level_api_wrap.drop_user(user_name=not_exist_user_name)

    def test_milvus_client_update_password(self, host, port):
        """
        target: test milvus client api update_password
        method: create a user and update password
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user, password=ct.default_password)
        user_name = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        new_password = cf.gen_str_by_length()
        self.high_level_api_wrap.update_password(user_name=user_name, old_password=password, new_password=new_password)
        # check
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, user=user_name, password=new_password)
        res = client_w.list_collections(client)[0]
        assert len(res) == 0
        self.high_level_api_wrap.init_milvus_client(uri=uri, user=user_name, password=password,
                                                    check_task=CheckTasks.check_permission_deny)

    def test_milvus_client_list_users(self, host, port):
        """
        target: test milvus client api list_users
        method: create a user and list users
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        user_name1 = cf.gen_unique_str(user_pre)
        user_name2 = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name1, password=password)
        self.high_level_api_wrap.create_user(user_name=user_name2, password=password)
        res = self.high_level_api_wrap.list_users()[0]
        assert {ct.default_user, user_name1, user_name2}.issubset(set(res)) is True

    def test_milvus_client_describe_user(self, host, port):
        """
        target: test milvus client api describe_user
        method: create a user and describe the user
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        user_name = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        # describe one self
        res, _ = self.high_level_api_wrap.describe_user(user_name=ct.default_user)
        assert res["user_name"] == ct.default_user
        # describe other users
        res, _ = self.high_level_api_wrap.describe_user(user_name=user_name)
        assert res["user_name"] == user_name
        # describe user that not exists
        user_not_exist = cf.gen_unique_str(user_pre)
        res, _ = self.high_level_api_wrap.describe_user(user_name=user_not_exist)
        assert res == {}

    def test_milvus_client_create_role(self, host, port):
        """
        target: test milvus client api create_role
        method: create a role
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        role_name = cf.gen_unique_str(role_pre)
        self.high_level_api_wrap.create_role(role_name=role_name)

    def test_milvus_client_drop_role(self, host, port):
        """
        target: test milvus client api drop_role
        method: create a role and drop
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        role_name = cf.gen_unique_str(role_pre)
        self.high_level_api_wrap.create_role(role_name=role_name)
        self.high_level_api_wrap.drop_role(role_name=role_name)

    def test_milvus_client_describe_role(self, host, port):
        """
        target: test milvus client api describe_role
        method: create a role and describe
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        role_name = cf.gen_unique_str(role_pre)
        self.high_level_api_wrap.create_role(role_name=role_name)
        # describe a role that exists
        self.high_level_api_wrap.describe_role(role_name=role_name)

    def test_milvus_client_list_roles(self, host, port):
        """
        target: test milvus client api list_roles
        method: create a role and list roles
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        role_name = cf.gen_unique_str(role_pre)
        self.high_level_api_wrap.create_role(role_name=role_name)
        res, _ = self.high_level_api_wrap.list_roles()
        assert role_name in res

    def test_milvus_client_grant_role(self, host, port):
        """
        target: test milvus client api grant_role
        method: create a role and a user, then grant role to the user
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        user_name = cf.gen_unique_str(user_pre)
        role_name = cf.gen_unique_str(role_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        self.high_level_api_wrap.create_role(role_name=role_name)
        self.high_level_api_wrap.grant_role(user_name=user_name, role_name=role_name)

    def test_milvus_client_revoke_role(self, host, port):
        """
        target: test milvus client api revoke_role
        method: create a role and a user, then grant role to the user, then revoke
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        user_name = cf.gen_unique_str(user_pre)
        role_name = cf.gen_unique_str(role_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        self.high_level_api_wrap.create_role(role_name=role_name)
        # revoke a user that does not exist
        self.high_level_api_wrap.revoke_role(user_name=user_name, role_name=role_name)
        # revoke a user that exists
        self.high_level_api_wrap.grant_role(user_name=user_name, role_name=role_name)
        self.high_level_api_wrap.revoke_role(user_name=user_name, role_name=role_name)

    def test_milvus_client_grant_privilege(self, host, port):
        """
        target: test milvus client api grant_privilege
        method: create a role and a user, then grant role to the user, grant a privilege to the role
        expected: succeed
        """
        # prepare a collection
        uri = f"http://{host}:{port}"
        client_root, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        coll_name = cf.gen_unique_str()
        client_w.create_collection(client_root, coll_name, default_dim, consistency_level="Strong")

        # create a new role and a new user ( no privilege)
        user_name = cf.gen_unique_str(user_pre)
        role_name = cf.gen_unique_str(role_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        self.high_level_api_wrap.create_role(role_name=role_name)
        self.high_level_api_wrap.grant_role(user_name=user_name, role_name=role_name)

        # check the role has no privilege of drop collection
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, user=user_name, password=password)
        self.high_level_api_wrap.drop_collection(client, coll_name, check_task=CheckTasks.check_permission_deny)

        # grant the role with the privilege of drop collection
        self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        self.high_level_api_wrap.grant_privilege(role_name, "Global", "*", "DropCollection")

        # check the role has privilege of drop collection
        self.high_level_api_wrap.drop_collection(client, coll_name)

    def test_milvus_client_revoke_privilege(self, host, port):
        """
        target: test milvus client api revoke_privilege
        method: create a role and a user, then grant role to the user, grant a privilege to the role, then revoke
        expected: succeed
        """
        # prepare a collection
        uri = f"http://{host}:{port}"
        client_root, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        coll_name = cf.gen_unique_str()

        # create a new role and a new user ( no privilege)
        user_name = cf.gen_unique_str(user_pre)
        role_name = cf.gen_unique_str(role_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        self.high_level_api_wrap.create_role(role_name=role_name)
        self.high_level_api_wrap.grant_role(user_name=user_name, role_name=role_name)
        self.high_level_api_wrap.grant_privilege(role_name, "Global", "*", "CreateCollection")
        time.sleep(60)

        # check the role has privilege of create collection
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, user=user_name, password=password)
        client_w.create_collection(client, coll_name, default_dim, consistency_level="Strong")

        # revoke the role with the privilege of create collection
        self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        self.high_level_api_wrap.revoke_privilege(role_name, "Global", "*", "CreateCollection")

        # check the role has no privilege of create collection
        self.high_level_api_wrap.create_collection(client, coll_name, default_dim, consistency_level="Strong",
                                                   check_task=CheckTasks.check_permission_deny)


@pytest.mark.tags(CaseLabel.RBAC)
class TestMilvusClientRbacInvalid(TestcaseBase):
    """ Test case of rbac interface """
    def test_milvus_client_init_token_invalid(self, host, port):
        """
        target: test milvus client api token invalid
        method: init milvus client using a wrong token
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        wrong_token = root_token + "kk"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=wrong_token,
                                                                check_task=CheckTasks.check_auth_failure)

    def test_milvus_client_init_username_invalid(self, host, port):
        """
        target: test milvus client api username invalid
        method: init milvus client using a wrong username
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        invalid_user_name = ct.default_user + "nn"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, user=invalid_user_name,
                                                                password=ct.default_password,
                                                                check_task=CheckTasks.check_auth_failure)

    def test_milvus_client_init_password_invalid(self, host, port):
        """
        target: test milvus client api password invalid
        method: init milvus client using a wrong password
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        wrong_password = ct.default_password + "kk"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user,
                                                                password=wrong_password,
                                                                check_task=CheckTasks.check_auth_failure)

    @pytest.mark.parametrize("invalid_name", ["", "0", "n@me", "h h"])
    def test_milvus_client_create_user_value_invalid(self, host, port, invalid_name):
        """
        target: test milvus client api create_user invalid
        method: create using a wrong username
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        self.high_level_api_wrap.create_user(invalid_name, ct.default_password,
                                             check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 1100,
                                                          ct.err_msg: "invalid user name"})

    @pytest.mark.parametrize("invalid_name", [1, [], None, {}])
    def test_milvus_client_create_user_type_invalid(self, host, port, invalid_name):
        """
        target: test milvus client api create_user invalid
        method: create using a wrong username
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        self.high_level_api_wrap.create_user(invalid_name, ct.default_password,
                                             check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 1,
                                                          ct.err_msg: "invalid user name"})

    def test_milvus_client_create_user_exist(self, host, port):
        """
        target: test milvus client api create_user invalid
        method: create using a wrong username
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        self.high_level_api_wrap.create_user("root", ct.default_password,
                                             check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 65535,
                                                          ct.err_msg: "user already exists: root"})

    @pytest.mark.parametrize("invalid_password", ["", "0", "p@ss", "h h", "1+1=2"])
    def test_milvus_client_create_user_password_invalid_value(self, host, port, invalid_password):
        """
        target: test milvus client api create_user invalid
        method: create using a wrong username
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        user_name = cf.gen_unique_str(user_pre)
        self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        self.high_level_api_wrap.create_user(user_name, invalid_password,
                                             check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 1100,
                                                          ct.err_msg: "invalid password"})

    @pytest.mark.parametrize("invalid_password", [1, [], None, {}])
    def test_milvus_client_create_user_password_invalid_type(self, host, port, invalid_password):
        """
        target: test milvus client api create_user invalid
        method: create using a wrong username
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        user_name = cf.gen_unique_str(user_pre)
        self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        self.high_level_api_wrap.create_user(user_name, invalid_password,
                                             check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 1,
                                                          ct.err_msg: "invalid password"})

    def test_milvus_client_update_password_user_not_exist(self, host, port):
        """
        target: test milvus client api update_password
        method: create a user and update password
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user, password=ct.default_password)
        user_name = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        new_password = cf.gen_str_by_length()
        self.high_level_api_wrap.update_password(user_name=user_name, old_password=password, new_password=new_password,
                                                 check_task=CheckTasks.err_res,
                                                 check_items={ct.err_code: 1400,
                                                              ct.err_msg: "old password not correct for %s: "
                                                                          "not authenticated" % user_name})

    def test_milvus_client_update_password_password_wrong(self, host, port):
        """
        target: test milvus client api update_password
        method: create a user and update password
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user, password=ct.default_password)
        user_name = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        new_password = cf.gen_str_by_length()
        wrong_password = password + 'kk'
        self.high_level_api_wrap.update_password(user_name=user_name, old_password=wrong_password,
                                                 new_password=new_password, check_task=CheckTasks.err_res,
                                                 check_items={ct.err_code: 1400,
                                                              ct.err_msg: "old password not correct for %s: "
                                                                          "not authenticated" % user_name})

    def test_milvus_client_update_password_new_password_same(self, host, port):
        """
        target: test milvus client api update_password
        method: create a user and update password
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user, password=ct.default_password)
        user_name = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        self.high_level_api_wrap.update_password(user_name=user_name, old_password=password, new_password=password)

    @pytest.mark.parametrize("invalid_password", ["", "0", "p@ss", "h h", "1+1=2"])
    def test_milvus_client_update_password_new_password_invalid(self, host, port, invalid_password):
        """
        target: test milvus client api update_password
        method: create a user and update password
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        self.high_level_api_wrap.init_milvus_client(uri=uri, user=ct.default_user, password=ct.default_password)
        user_name = cf.gen_unique_str(user_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        self.high_level_api_wrap.update_password(user_name=user_name, old_password=password,
                                                 new_password=invalid_password, check_task=CheckTasks.err_res,
                                                 check_items={ct.err_code: 1100,
                                                              ct.err_msg: "invalid password"})

    def test_milvus_client_create_role_invalid(self, host, port):
        """
        target: test milvus client api create_role
        method: create a role using invalid name
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        role_name = cf.gen_unique_str(role_pre)
        self.high_level_api_wrap.create_role(role_name=role_name)
        # create existed role
        error_msg = f"role [name:{role_pre}] already exists"
        self.high_level_api_wrap.create_role(role_name=role_name, check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 65535, ct.err_msg: error_msg})
        # create role public or admin
        self.high_level_api_wrap.create_role(role_name="public", check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 65535, ct.err_msg: error_msg})
        self.high_level_api_wrap.create_role(role_name="admin", check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 65535, ct.err_msg: error_msg})

    def test_milvus_client_drop_role_invalid(self, host, port):
        """
        target: test milvus client api drop_role
        method: create a role and drop
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        role_name = cf.gen_unique_str(role_pre)
        self.high_level_api_wrap.drop_role(role_name=role_name, check_task=CheckTasks.err_res,
                                           check_items={ct.err_code: 65535,
                                                        ct.err_msg: "not found the role, maybe the role isn't "
                                                                    "existed or internal system error"})

    def test_milvus_client_describe_role_invalid(self, host, port):
        """
        target: test milvus client api describe_role
        method: describe a role using invalid name
        expected: raise exception
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        # describe a role that does not exist
        role_not_exist = cf.gen_unique_str(role_pre)
        error_msg = "not found the role, maybe the role isn't existed or internal system error"
        self.high_level_api_wrap.describe_role(role_name=role_not_exist, check_task=CheckTasks.err_res,
                                             check_items={ct.err_code: 65535, ct.err_msg: error_msg})

    def test_milvus_client_grant_role_user_not_exist(self, host, port):
        """
        target: test milvus client api grant_role
        method: create a role and a user, then grant role to the user
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        user_name = cf.gen_unique_str(user_pre)
        role_name = cf.gen_unique_str(role_pre)
        self.high_level_api_wrap.create_role(role_name=role_name)
        self.high_level_api_wrap.grant_role(user_name=user_name, role_name=role_name,
                                            check_task=CheckTasks.err_res,
                                            check_items={ct.err_code: 65536,
                                                         ct.err_msg: "not found the user, maybe the user "
                                                                     "isn't existed or internal system error"})

    def test_milvus_client_grant_role_role_not_exist(self, host, port):
        """
        target: test milvus client api grant_role
        method: create a role and a user, then grant role to the user
        expected: succeed
        """
        uri = f"http://{host}:{port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)
        user_name = cf.gen_unique_str(user_pre)
        role_name = cf.gen_unique_str(role_pre)
        password = cf.gen_str_by_length()
        self.high_level_api_wrap.create_user(user_name=user_name, password=password)
        self.high_level_api_wrap.grant_role(user_name=user_name, role_name=role_name,
                                            check_task=CheckTasks.err_res,
                                            check_items={ct.err_code: 65536,
                                                         ct.err_msg: "not found the role, maybe the role "
                                                                     "isn't existed or internal system error"})


@pytest.mark.tags(CaseLabel.RBAC)
class TestMilvusClientRbacAdvance(TestcaseBase):
    """ Test case of rbac interface """

    def teardown_method(self, method):
        """
        teardown method: drop role and user
        """
        log.info("[utility_teardown_method] Start teardown utility test cases ...")
        uri = f"http://{cf.param_info.param_host}:{cf.param_info.param_port}"
        client, _ = self.high_level_api_wrap.init_milvus_client(uri=uri, token=root_token)

        # drop users
        users, _ = self.high_level_api_wrap.list_users()
        for user in users:
            if user != ct.default_user:
                self.high_level_api_wrap.drop_user(user)
        users, _ = self.high_level_api_wrap.list_users()
        assert len(users) == 1

        # drop roles
        roles, _ = self.high_level_api_wrap.list_roles()
        for role in roles:
            if role not in ['admin', 'public']:
                privileges, _ = self.high_level_api_wrap.describe_role(role)
                if privileges:
                    for privilege in privileges:
                        self.high_level_api_wrap.revoke_privilege(role, privilege["object_type"],
                                                                  privilege["privilege"], privilege["object_name"])
                self.high_level_api_wrap.drop_role(role)
        roles, _ = self.high_level_api_wrap.list_roles()
        assert len(roles) == 2

        super().teardown_method(method)
