import sql from 'k6/x/sql';
import { check } from 'k6';

const db = sql.open('postgres', '');
const maxSize = 3891776803;
let size = 0;

function getdbsize() {
  const results = sql.query(db, "SELECT pg_database_size('testdb');");
  check(results, {'has result': (r) => r.length === 1})
  return results[0]['pg_database_size']
}
export function setup() {
  db.exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`);
  db.exec(`CREATE TABLE IF NOT EXISTS person (
           id uuid DEFAULT uuid_generate_v1(),
           email varchar NOT NULL,
           first_name varchar,
           last_name varchar,
           CONSTRAINT person_pkey PRIMARY KEY (id));`);
}

export function teardown() {
  // db.exec('DELETE FROM person;');
  // db.exec('DROP TABLE person;');
  db.close();
}

export default function () {
  size = getdbsize();
  while (size < maxSize) {
    db.exec("insert into person (email, first_name, last_name) select random()::text, random()::text, random()::text from generate_series(1, 80000);")
    size = db.getdbsize();
  }
}

