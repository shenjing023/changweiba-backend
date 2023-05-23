from conf.config import CONF
from handler.api import query_stock

from fastapi import FastAPI,status
from fastapi.responses import JSONResponse
import uvicorn

app = FastAPI()

@app.get("/query/{stock_id}")
def query(stock_id:str):
    result=query_stock(stock_id)
    return JSONResponse(status_code=status.HTTP_200_OK,content={
        "code":200,
        "data":result,
        "msg":"success"
    })


if __name__ == '__main__':
    port = CONF["port"]
    uvicorn.run("service:app", port=port, log_level="info")