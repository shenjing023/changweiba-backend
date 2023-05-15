import os
import time
import execjs
from fake_useragent import UserAgent

ua = UserAgent()

def get_token():
    '''获取token'''
    with open(os.path.join(os.path.dirname(__file__), 'hexin-v.js'), 'r') as f:
        jscontent = f.read()
    context = execjs.compile(jscontent)
    return context.call("v")

def headers(cookie=None):
    # t1 = time.perf_counter()
    hexin_v = get_token()
    # t2 =time.perf_counter()
    # print('生成token:%s毫秒' % ((t2 - t1)*1000))
    return {
        'hexin-v': hexin_v,
        'User-Agent': ua.random,
        'cookie': cookie
    }
