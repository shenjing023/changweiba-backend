import requests as rq
import time

from handler.convert import convert
from handler.headers import headers
from conf.config import log
from handler.consts import BULL
from handler.hot import get_hot_new,chatgpt,get_trade_data

def while_do(do, retry=5, sleep=3):
    count = 0
    while count < retry:
        time.sleep(sleep)
        try:
            return do()
        except:
            count += 1
    return None

def get_robot_data(stock_symbol: str):
    """
    get
    """
    data={
        'perpage': 50,
        'page': 1,
        'source': 'Ths_iwencai_Xuangu',
        'secondary_intent': "stock",
        'question': stock_symbol
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

def query_stock(stock_symbol:str):
    result = get_robot_data(stock_symbol)
    if not result:
        log.info(f"query [stock:{stock_symbol}] failed")
        return {
            "bull": 0,
            "short": "---"
        }
    return {
        "bull":  BULL[result["bull"]],
        "short": result["short"]
    }
    
def analyse_stock(symbol:str,name:str):
    return {
        "result":"",
    }
    # symbol = symbol.lower()
    # news=get_hot_new(symbol[2:],name)
    # new_content='\n'.join(news['新闻内容'])
    # trade_data=get_trade_data(symbol,10)
    # prompt1="""我想让你扮演一个具有丰富经验的金融市场专业知识的股票投资者，我这有一只股票近10天的交易数据(json格式)，其中字段date表示交易日期，字段open表示开盘价，字段close表示收盘价，字段high表示最高价，字段low表示最低价，字段volume表示成交量。交易数据如下：'''%s''',
    # """%trade_data
    # prompt2="""还有这只股票近期的相关资讯:{'%s相关资讯':'''%s''',\n}\n
    # """%(name,new_content)
    # prompt3="""请分析总结机会点和风险点和持仓建议，其中持仓建议请从这几项选项中选择：'''[卖出，减持，观望，增持，买入]'''，并给出该股票所属的题材标签。输出格式为{'机会':'''1..\n2..\n...''',\n'风险':'''1..\n2..\n...''',\n'持仓建议':[持仓建议],\n'题材标签':[标签]}
    # """
    # prompt=prompt1+prompt2+prompt3
    # result=chatgpt(prompt)
    # if result is None:
    #     return {
    #         "result":"",
    #     }
    # return {
    #     "result":result
    # }
    