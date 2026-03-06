工作量证明 PoW 挖矿模拟

实验目的
1. 理解工作量证明机制
2. 理解区块难度
3. 模拟挖矿过程

核心思想
寻找nonce使得

SHA256(block_header + nonce)

满足指定前导零。

核心代码
import hashlib

nonce=0
while True:
    text="block"+str(nonce)
    hash=hashlib.sha256(text.encode()).hexdigest()
    if hash.startswith("0000"):
        print("found",nonce,hash)
        break
    nonce+=1
