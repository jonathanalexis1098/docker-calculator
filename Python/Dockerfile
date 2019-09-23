FROM alpine:3.10

RUN apk add --no-cache python3-dev \
    && pip3 install --upgrade pip

WORKDIR /app

COPY . /app

RUN pip3 --no-cache-dir install -r requirements.txt

CMD [ "python3", "src/app.py" ]



#python3-dev instala python 3 y pip



# alpine es una imagen minimalista de docker basada en una 
# distribucion de linux llamada alpine, es una distro pequeña.

# Es decir que con esto vamos a tener por debajo un sistema linux
# y despues del sistema linux vamos a instalar python


