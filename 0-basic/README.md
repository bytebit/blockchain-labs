区块链基础环境搭建与测试网交互

| 项目  | 详细内容 |
| --- | --- |
| 参考课程 | 斯坦福CS251 2024 Lab1: Ethereum Basics & MetaMask Setup |
| 建议学时 | 1学时 |
| 前置知识 | 区块链基础定义、以太坊账户模型、哈希函数基础 |
| 核心目标 | 掌握区块链基础交互工具的使用，理解交易的完整生命周期，建立对链上数据的直观认知 |
| 实验环境 | Chrome/Firefox浏览器、MetaMask钱包、Sepolia以太坊测试网、Etherscan区块链浏览器 |
| 详细实验步骤 | 1. 安装MetaMask钱包，创建钱包账户，备份助记词与私钥，理解私钥-公钥-地址的映射关系<br>2. 配置Sepolia测试网，通过水龙头获取测试网ETH<br>3. 发起一笔测试网转账交易，设置Gas Limit与Gas Price，观察交易打包过程<br>4. 在Etherscan中查询该笔交易，解析交易的核心字段（Block Number、Nonce、From/To、Value、Gas Used、Input Data）<br>5. 对比交易打包前后的账户余额变化，理解Gas手续费的扣除逻辑 |
| 验收标准 | 1. 成功完成测试网转账，交易最终确认数≥6<br>2. 能完整解释交易中所有核心字段的含义，说明Gas消耗的原因<br>3. 能清晰说明私钥、助记词、地址的安全管理规范 |
| 拓展进阶任务 | 1. 用Remix部署一个极简的「存储合约」（包含set/get函数，可修改/读取一个数字），并通过MetaMask交互<br>2. 用Etherscan解析一笔以太坊主网的合约调用交易，说明Input Data的作用 |
