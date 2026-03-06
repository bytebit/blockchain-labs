 P2PKH交易脚本验证

实验目的
1. 理解ScriptSig与ScriptPubKey
2. 理解堆栈脚本执行机制
3. 理解P2PKH交易验证流程

ScriptSig
<Sig> <PubKey>

ScriptPubKey
OP_DUP OP_HASH160 <PubKeyHash> OP_EQUALVERIFY OP_CHECKSIG

学生任务
1. 构造解锁脚本
2. 执行脚本堆栈
3. 验证签名
