require 'redis'

$redis = Redis.new(
  :host => '127.0.0.1',
  :port =>  6379
)

# String values
$redis.set('newkey', 'ofsomevalue')
value = $redis.get('newkey')
puts value

# Binary values
$redis.set('pages:about', 'about us')
value = $redis.get('pages:about')
puts value

# Counter values
$redis.set('counter', 100)
$redis.incr('counter')
value = $redis.get('counter')
puts value

$redis.incr('counter')
value = $redis.get('counter')
puts value

$redis.incrby('counter', 50)
value = $redis.get('counter')
puts value

# MSET/MGET 
$redis.mset('a', 10, 'b', 20, 'c', 30)
m_array = $redis.mget('a','b','c')
puts m_array

# ALTERING/QUERYING KEY
$redis.set('mykey', 'hello')
mykeyexists = $redis.exists('mykey')
puts mykeyexists

$redis.del('mykey')
mykeyexists = $redis.exists('mykey')
puts mykeyexists

keytype = $redis.type('mykey')
puts keytype

# SET TIMEOUT
$redis.set('mykey', 'somevalue')
$redis.expire('mykey', 5)
beforevalue = $redis.get('mykey')
puts 'before expiry: '  + beforevalue
sleep(6)
aftervalue = $redis.get('mykey')
puts 'after expiry: ' + aftervalue.to_s

# TTL
$redis.set('mykey', '100')
$redis.expire('mykey', 10)
ttl_value = $redis.ttl('mykey')
puts ttl_value

# LIST OPERATIONS

# remove existing values from list, if any
$redis.del('mylist') 

$redis.rpush('mylist', 'a')
$redis.rpush('mylist', 'b')
$redis.lpush('mylist', 'first')
puts $redis.lrange('mylist', 0, -1)

$redis.rpush('mylist', [ 1, 2, 3, 4, 5, "foo bar"])
puts $redis.lrange('mylist', 0, -1)

$redis.rpush('numberlist', [1, 2, 3])
puts 'Popped value: ' + $redis.rpop('numberlist')
puts 'Popped value: ' + $redis.rpop('numberlist')
puts 'Popped value: ' + $redis.rpop('numberlist')

# CAPPED LIST
$redis.rpush('mycappedlist', [1, 2, 3, 4, 5])
$redis.ltrim('mycappedlist', 0, 2)
puts 'Capped list:.... '
puts $redis.lrange('mycappedlist', 0, -1)

