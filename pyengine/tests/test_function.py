import numpy as np
import requests
import pytest
import logging
import json

url = "http://127.0.0.1:5000"


# TODO: LOG and Assert
class TestEngineFunction():
    def test_1m_add(self):
        d = 4
        nb = 100
        nq = 1
        k = 10
        _, xb, xq = get_dataset(d, nb, 1, nq)

        groupid = "test_search_3"

        route_group = url + "/vector/group/" + groupid
        r = requests.post(route_group, json={"dimension": d})

        # import dataset
        vector_add_route = url + "/vector/add/" + groupid
        for i in xb:
            data = dict()
            data['vector'] = i.tolist()
            # print(data)
            r = requests.post(vector_add_route, json=data)
            print(r.json())

        # search dataset
        vector_search_route = url + "/vector/search/" + groupid
        data = dict()
        for i in xq:
            data['vector'] = i.tolist()
            data['limit'] = k
            # print(data)
            r = requests.get(vector_search_route, json=data)
            print(r.json())


def get_dataset(d, nb, nt, nq):
    d1 = 10  # intrinsic dimension (more or less)
    n = nb + nt + nq
    rs = np.random.RandomState(1338)
    x = rs.normal(size=(n, d1))
    x = np.dot(x, rs.rand(d1, d))
    x = x * (rs.rand(d) * 4 + 0.1)
    x = np.sin(x)
    x = x.astype('float32')
    return x[:nt], x[nt:-nq], x[-nq:]
