FROM nikolaik/python-nodejs:latest

WORKDIR /app

COPY ./requirements.txt /app/requirements.txt

RUN pip install --no-cache-dir --upgrade -r /app/requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple/

COPY . /app/

EXPOSE 8022

ENTRYPOINT ["python", "service.py"]