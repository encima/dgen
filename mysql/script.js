import sql from 'k6/x/sql';
import { check } from 'k6';

const db = sql.open('mysql', '');
const maxSize = 36384;
let size = 0;

function getdbsize() {
  const results = sql.query(db, "SELECT SUM(data_length + index_length) 'size' FROM information_schema.tables where table_schema='defaultdb';");
  check(results, {'has result': (r) => r.length === 1})
  return Number(String.fromCharCode(...results[0]['size']))
}

export function setup() {
//  db.exec(`CREATE TABLE IF NOT EXISTS person (id int not null primary key auto_increment, email varchar(200), first_name varchar(200), last_name varchar(200));`)
/*  db.exec(`DELIMITER $$
  CREATE PROCEDURE IF NOT EXISTS InsertRand(IN NumRows INT)
      BEGIN
          DECLARE i INT;
          SET i = 1;
          START TRANSACTION;
          WHILE i <= NumRows DO
              INSERT INTO person (email, first_name, last_name) select MD5(RAND()), MD5(RAND()), MD5(RAND());
              SET i = i + 1;
          END WHILE;
          COMMIT;
      END$$
  DELIMITER ;
  `);*/
}

export function teardown() {
  // db.exec('DELETE FROM person;');
  // db.exec('DROP TABLE person;');
  db.close();
}

export default function () {
  size = getdbsize();
  console.log(size)
  while (size < maxSize) {
    db.exec("CALL InsertRand(90000);")
    setTimeout(function() {
      size = getdbsize();
    }, 903000);
  }
}

