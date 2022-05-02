## Running

1. Deploy an Aiven for Redis service and set the Service URI
2. Run `generate -l X` where X is the number of lines to generate
3. Run `cat sample.txt | redis-cli -u $SVC --pipe` where $SVC is the Service URI
