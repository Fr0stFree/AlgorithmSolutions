import contextlib
import random
import uuid
import psycopg2
import datetime as dt

from faker import Faker
from psycopg2.extras import execute_batch

PAGE_SIZE = 5000
PERSONS_COUNT = 100_000
ROLES = ("actor", "producer", "director")

dsn = {
    "host": "localhost",
    "port": 5432,
    "dbname": "movies_database",
    "user": "app",
    "password": "123qwe",
    "options": "-c search_path=content",
}
fake = Faker()


if __name__ == "__main__":
    with contextlib.closing(psycopg2.connect(**dsn)) as conn, conn.cursor() as cur:
        now = dt.datetime.now()

        # Заполнение таблицы Person
        persons_ids = [str(uuid.uuid4()) for _ in range(PERSONS_COUNT)]
        query = 'INSERT INTO person (id, full_name, created, modified) VALUES (%s, %s, %s, %s)'
        data = [(pk, fake.last_name(), now, now) for pk in persons_ids]
        execute_batch(cur, query, data, page_size=PAGE_SIZE)
        conn.commit()

        # Заполнение таблицы FilmWork
        film_work_data = []
        types = ['movie', 'tv_show', 'tv_series', 'video']
        for _ in range(100_00):
            title = fake.sentence(nb_words=3)
            description = fake.text()
            creation_date = fake.date_time_between(start_date='-30y', end_date='now')
            rating = random.randint(1, 10)
            film_work_data.append((
                str(uuid.uuid4()), title, description, creation_date, rating, random.choice(types), now, now
            ))

        query = 'INSERT INTO film_work (id, title, description, creation_date, rating, type, created, modified) VALUES (%s, %s, %s, %s, %s, %s, %s, %s)'
        execute_batch(cur, query, film_work_data, page_size=PAGE_SIZE)

        # Заполнение таблицы PersonFilmWork
        person_film_work_data = []
        roles = ['actor', 'producer', 'director']

        cur.execute('SELECT id FROM film_work')
        film_works_ids = [data[0] for data in cur.fetchall()]

        for film_work_id in film_works_ids:
            for person_id in random.sample(persons_ids, 5):
                role = random.choice(roles)
                person_film_work_data.append((
                    str(uuid.uuid4()), film_work_id, person_id, role, now
                ))

        query = 'INSERT INTO person_film_work (id, film_work_id, person_id, role, created) VALUES (%s, %s, %s, %s, %s)'
        execute_batch(cur, query, person_film_work_data, page_size=PAGE_SIZE)
        conn.commit()
