from conf.config import CONF
from handler.api import query_stock,analyse_stock

from fastapi import FastAPI,status
from fastapi.responses import JSONResponse
import uvicorn
from pydantic import BaseModel

app = FastAPI()

class AnalyseItem(BaseModel):
    name: str
    symbol: str

@app.get("/query/{stock_id}")
def query(stock_id:str):
    result=query_stock(stock_id)
    return JSONResponse(status_code=status.HTTP_200_OK,content={
        "code":200,
        "data":result,
        "msg":"success"
    })

@app.post("/stock/analyse")
def analyse(stock:AnalyseItem):
    result=analyse_stock(stock.symbol,stock.name)
    return JSONResponse(status_code=status.HTTP_200_OK,content={
        "code":200,
        "data":result,
        "msg":"success"
    })

if __name__ == '__main__':
    port = CONF["port"]
    uvicorn.run("service:app", port=port, log_level="info")