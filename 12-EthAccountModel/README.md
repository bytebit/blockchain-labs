| 项目  | 详细内容 |
| --- | --- |
| 参考课程 | 斯坦福CS251 Lab2: Ethereum Account Model & Transaction Structure |
| 建议学时 | 2学时 |
| 前置知识 | 以太坊EOA/合约账户模型、交易结构、Gas机制 |
| 核心目标 | 掌握以太坊与比特币的核心差异，理解以太坊交易的执行逻辑与Gas消耗规则 |
| 实验环境 | MetaMask、Sepolia测试网、Etherscan、Remix IDE |
| 详细实验步骤 | 1. 区分EOA账户与合约账户：在Etherscan中查询典型EOA地址、ERC20合约地址、NFT合约地址，对比账户结构差异<br>2. 解析以太坊交易的完整结构：<br> - 核心字段：chainId、nonce、maxPriorityFeePerGas、maxFeePerGas、gasLimit、to、value、data、v/r/s签名<br> - 理解nonce的防重放作用，模拟nonce错误导致的交易失败场景<br>3. Gas机制深度验证：<br> - 发起不同类型的交易（ETH转账、ERC20转账、合约调用），对比Gas消耗差异<br> - 分析Gas Limit、Gas Price对交易打包优先级的影响，模拟Gas不足导致的交易失败<br>4. 合约部署交易解析：在Remix部署一个简单合约，解析合约部署交易的to字段为空、data字段为合约字节码的逻辑 |
| 验收标准 | 1. 能清晰区分EOA与合约账户的核心差异，说明两类账户的状态字段<br>2. 能完整解析以太坊交易的所有字段，解释不同交易类型的Gas消耗差异<br>3. 能说明nonce的作用，以及交易签名的验证逻辑 |
| 拓展进阶任务 | 1. 用Etherscan的解码工具解析一笔Uniswap V3的兑换交易，说明Input Data的函数调用逻辑<br>2. 模拟以太坊交易的签名与验签流程，理解EIP-155防重放攻击的实现 |
