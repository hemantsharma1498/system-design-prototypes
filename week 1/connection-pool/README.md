# Db Connection Pool with a bounded blocking queue


Prototyped a db connection pool:
1. Used sqlite3 as the db.
2. Created a bounded blocking queue using a go channel.
3. Same channel acts as a semaphore
