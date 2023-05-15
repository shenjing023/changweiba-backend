from api import query_stock

def test_get():
    data=query_stock(688501)
    print(data)
    assert 0