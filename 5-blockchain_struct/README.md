构建简易区块链结构

实验目的
1. 理解区块链数据结构
2. 理解哈希指针链
3. 理解区块不可篡改性

区块结构
index
timestamp
data
previous_hash
hash

核心代码
import hashlib, time

class Block:
    def __init__(self,index,data,prev_hash):
        self.index=index
        self.timestamp=time.time()
        self.data=data
        self.prev_hash=prev_hash
        self.hash=self.calc()

    def calc(self):
        value=str(self.index)+str(self.timestamp)+self.data+self.prev_hash
        return hashlib.sha256(value.encode()).hexdigest()
