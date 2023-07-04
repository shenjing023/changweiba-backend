import sys
import os

sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
from api import query_stock,analyse_stock
from hot import get_trade_data

# def test_get():
#     data=query_stock("688501")
#     print(data)
#     assert 0
    
def test_analyse():
    result=analyse_stock("sz300657","弘信电子")
    print(result)
    assert 0
    
# def test_trade_data():
#     data=get_trade_data("sz002527",10)
#     print(data)
#     # for d in data:
#     #     print(d)
#     assert 0