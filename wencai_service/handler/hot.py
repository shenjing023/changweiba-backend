import uuid
import requests
import akshare as ak
import pandas as pd
import json
import time

from conf.config import log,CONF

class KLineData:
    def __init__(self,open: float,close: float,high: float,low: float,volume: float,date: str) -> None:
        self.open = open
        self.close = close
        self.high = high
        self.low = low
        self.volume = volume
        self.date=date
        
    def __str__(self) -> str:
        return f"{self.open},{self.close},{self.high},{self.low},{self.volume},{self.date}"
    
def get_hot_new(symbol:str,name:str):
    print("symbol: ",symbol)
    news=ak.stock_news_em(symbol)
    news.drop_duplicates(subset='新闻标题',inplace=True)
    news['发布时间']=pd.to_datetime(news['发布时间'])
    news['新闻标题']=news['发布时间'].dt.strftime('%Y-%m-%d ')+news['新闻标题'].str.replace('%s：'%name,'')
    news = news[~news['新闻标题'].str.contains('股|主力|机构|资金流')]
    # news['新闻内容']=news['新闻标题'].str.cat(news['新闻内容'].str.split('。').str[0], sep=' ')
    news = news[news['新闻内容'].str.contains(name)]
    news.drop_duplicates(subset='新闻标题',inplace=True)
    news.sort_values(by=['发布时间'],ascending=False,inplace=True)
    # print("news: ",news.to_string())
    return news[:10]
    
def get_trade_data(symbol:str,day:int):
    url=f"http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayqfq&param={symbol},day,,,{day},qfq&r=0.{int(round(time.time()*1000))}"
    resp=requests.get(url)
    data=resp.text.split("=")[1]
    _json=json.loads(data)
    if _json["code"]!=0:
        return None
    data=[]
    for d in _json["data"][symbol]["qfqday"]:
        data.append(KLineData(d[1],d[2],d[3],d[4],d[5],d[0]))
    return json.dumps(data,default=lambda obj: obj.__dict__,indent=4)
    # return data
    
def chatgpt(prompt:str):
    url = f"{CONF['chatgpt']['host']}:{CONF['chatgpt']['port']}/api/conversation/talk"
    payload = {
        "prompt": prompt, 
        "model": "text-davinci-002-render-sha",
        "message_id": str(uuid.uuid4()),
        "parent_message_id": str(uuid.uuid4()),
        "stream":False,
        "conversation_id":""
    }
    headers = {"content-type":"application/json"}

    response = requests.post(url, json=payload,headers=headers)
    # response.content.decode('unicode-escape')

    if response.status_code == 200:
        data=json.loads(response.text)
        result=data["message"]["content"]["parts"][0]
        delete_chatgpt_conversation(data['conversation_id'])
        if data['error'] is not None:
            log.info('POST request failed:', data)
            return None
        return result
    else:
        log.info('POST request failed:', json.loads(response.text))
        return None
    
def delete_chatgpt_conversation(conversation_id:str):
    url = f"{CONF['chatgpt']['host']}:{CONF['chatgpt']['port']}/api/conversation/{conversation_id}"

    response = requests.delete(url)
    if response.status_code != 200:
        log.info('POST request failed:', json.loads(response.text))
