# NOTE
This method is not the most effective and can cause timeouts. Ideally need to 
refresh stats more often or find another way to get DB size

## Running

1. Create an Aiven for MySQL service and add the Service URI to the `script.js`
2. Set the `Advanced Config` -> `mysql.information_schema_stats_expiry` to the lowest setting of 900s (15 mins)
2. `../k6 run script.js`
