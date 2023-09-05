import contextlib
import random as rd
import uuid
import psycopg2
import datetime as dt

from faker import Faker
from psycopg2.extras import execute_batch

PAGE_SIZE = 5000
PERSONS_COUNT = 1000000
NOW = dt.datetime.utcnow()
ROLES = ("actor", "producer", "director")

dsn = {
    "host": "localhost",
    "port": 5432,
    "dbname": "movies_database",
    "user": "app",
    "password": "123qwe",
    "options": "-c search_path=content",
}
faker = Faker()

if __name__ == "__main__":
    with contextlib.closing(psycopg2.connect(**dsn)) as connection, connection.cursor() as cursor:

        statement = f"INSERT INTO person (id, full_name, created, modified) VALUES (%s, %s, %s, %s)"
        data = [(str(uuid.uuid4()), faker.full_name(), NOW, NOW) for _ in range(PERSONS_COUNT)]

        execute_batch(cursor, statement, data, PAGE_SIZE)
        connection.commit()
