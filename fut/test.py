from EAHashingAlgorithm import *

hash = EAHashingAlgorithm()

print hash.EAHash("565656")
print hash.zero_fill_right_shit(123456789,2)
print hash.num2hex(123456789)
print hash.chunkMessage("123456789")
print hash.add(12345,12345)
print hash.bitwiseRotate(12345,10)
print hash.cmn(1,2,3,4,5,6)
print hash.md5_f(1,2,3,4,5,6,7)
print hash.md5_g(1,2,3,4,5,6,7)
print hash.md5_h(1,2,3,4,5,6,7)
print hash.md5_i(1,2,3,4,5,6,7)
