import requests as rq
import time

from handler.convert import convert
from handler.headers import headers
from conf.config import log
from handler.consts import BULL

def while_do(do, retry=5, sleep=3):
    count = 0
    while count < retry:
        time.sleep(sleep)
        try:
            return do()
        except:
            count += 1
    return None

def get_robot_data(stock_id: int):
    """
    get
    """
    data={
        'perpage': 50,
        'page': 1,
        'source': 'Ths_iwencai_Xuangu',
        'secondary_intent': "stock",
        'question': str(stock_id)
    }
    def do():
        res=rq.request(
            method='POST',
            url='http://www.iwencai.com/customized/chart/get-robot-data',
            json=data,
            headers=headers(),
            timeout=(5,10)
        )
        return convert(res)
        
    return while_do(do)

def query_stock(stock_id:int):
    result = get_robot_data(stock_id)
    if not result:
        log.info(f"query [stock:{stock_id}] failed")
        return {
            "bull": 0,
            "short": "---"
        }
    return {
        "bull":  BULL[result["bull"]],
        "short": result["short"]
    }