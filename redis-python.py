import redis

r = redis.StrictRedis(host='localhost', port=6379, db=0, decode_responses=True)

# String values
r.set('newkey', 'ofsomevalue')
value = r.get('newkey') 
print(value)

# Binary values
r.set('pages:about', 'about us')
value = r.get('pages:about') 
print(value)

# Counter values
r.set('counter', 100)
r.incr('counter')
value = r.get('counter') 
print(value)

r.incr('counter')
value = r.get('counter')
print(value)

r.incrby('counter', 50)
value = r.get('counter')
print(value)

# MSET/MGET 
r.mset(a=10, b=20, c=30)
m_array = r.mget('a','b','c')
print(m_array)

# ALTERING/QUERYING KEY
r.set('mykey', 'hello')
mykeyexists = r.exists('mykey')
print(mykeyexists)

r.delete('mykey')
mykeyexists = r.exists('mykey')
print(mykeyexists)

keytype = r.type('mykey')
print(keytype)

# SET TIMEOUT
r.set('mykey', 'somevalue')
r.expire('mykey', 5)
beforevalue = r.get('mykey')
print('before expiry: '  + beforevalue)

import time 
time.sleep(6)

aftervalue = r.get('mykey') or ''
print('after expiry: ' + aftervalue)

# TTL
r.set('mykey', '100')
r.expire('mykey', 10)
ttl_value = r.ttl('mykey')
print(ttl_value)

# LIST OPERATIONS

# remove existing values from list, if any
r.delete('mylist') 

r.rpush('mylist', 'a')
r.rpush('mylist', 'b')
r.lpush('mylist', 'first')
print(r.lrange('mylist', 0, -1))

r.rpush('mylist', *[ 1, 2, 3, 4, 5, "foo bar"])
print(r.lrange('mylist', 0, -1))

r.rpush('numberlist', *[1, 2, 3])
print('Popped value: ' + r.rpop('numberlist'))
print('Popped value: ' + r.rpop('numberlist'))
print('Popped value: ' + r.rpop('numberlist'))

# CAPPED LIST
r.rpush('mycappedlist', *[1, 2, 3, 4, 5])
r.ltrim('mycappedlist', 0, 2)
print('Capped list:.... ')
print(r.lrange('mycappedlist', 0, -1))
