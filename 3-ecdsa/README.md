椭圆曲线数字签名与交易验证
【实验目的】
1. 理解公钥密码体系
2. 掌握数字签名流程
3. 模拟区块链交易签名验证

【实验环境准备】
pip install ecdsa

【实验步骤】
步骤1：生成公钥和私钥
步骤2：对交易信息进行签名
步骤3：使用公钥验证签名
步骤4：篡改交易验证签名失败

【参考代码】
from ecdsa import SigningKey, SECP256k1

sk = SigningKey.generate(curve=SECP256k1)
vk = sk.verifying_key
print("Private Key:",sk.to_string().hex())
print("Public Key:",vk.to_string().hex())

message = b"【你的学号】 pays Bob 5 BTC"
signature = sk.sign(message)

vk.verify(signature,message)
print("Signature verified")

【实验任务】
1. 完成交易签名程序
2. 修改verify交易信息（5 BTC改成6 BTC），验证签名失败
3. 分析签名机制在区块链中的作用

【思考题】
1. 为什么私钥必须保密？
2. 公钥为什么可以公开？
