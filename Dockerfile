FROM python:3.10

RUN pip install poetry

COPY . /app

WORKDIR /app

RUN poetry install --without dev

EXPOSE 8000

ENTRYPOINT [ "poetry", "run", "server" ]