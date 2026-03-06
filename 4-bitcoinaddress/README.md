比特币地址生成

实验目的
1. 理解公钥到地址的转换
2. 理解Hash160
3. 理解Base58Check

地址生成流程
PublicKey -> SHA256 -> RIPEMD160 -> PubKeyHash -> Base58Check -> Address

核心代码
import hashlib, base58

sha = hashlib.sha256(pubkey).digest()
ripemd = hashlib.new('ripemd160')
ripemd.update(sha)
pubkey_hash = ripemd.digest()

最后用metamask的私钥和地址进行验证
