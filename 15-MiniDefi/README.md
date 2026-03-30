# 🎓 期末项目实验

# **Mini-DeFi：构建一个简化去中心化金融系统**

---

# 一、项目目标

学生完成一个 **可运行的 DeFi 原型系统**：

✅ 编写 Solidity 合约
✅ 使用 Foundry / Hardhat 开发
✅ 实现 Token + 存款 + 利息
✅ 完成前端调用
✅ 部署测试网

最终成果 ≈ 一个“迷你 Aave / Compound”。

---

# 二、学生将学会什么（课程闭环）

| 知识             | 在项目中的体现 |
| -------------- | ------- |
| 哈希与签名          | 钱包交易    |
| P2PKH思想        | 地址控制资产  |
| 智能合约           | DeFi协议  |
| UTXO → Account | 账户余额    |
| Gas            | 交易成本    |
| Web3           | 前端调用    |

---

# 三、系统架构（核心教学图）

```
用户钱包 (MetaMask)
        ↓
React 前端 (ethers.js)
        ↓
MiniDeFi 合约
        ↓
ERC20 Token 合约
        ↓
Ethereum Testnet
```

---

# 四、项目功能要求（必须完成）

## ⭐ 模块1：ERC20 Token（基础资产）

学生实现：

```solidity
MiniToken.sol
```

功能：

* mint()
* transfer()
* approve()
* transferFrom()

学习目标：

👉 Token 即链上资产。

---

## ⭐ 模块2：DeFi 存款协议

```solidity
MiniDeFi.sol
```

核心逻辑：

### 存款

```solidity
function deposit(uint amount)
```

流程：

1. 用户 approve
2. 合约 transferFrom
3. 记录存款余额

---

### 提款

```solidity
function withdraw(uint amount)
```

检查：

* 用户余额
* 合约流动性

---

### 利息机制（重点）

```solidity
interest = deposit * rate * time
```

示例：

```solidity
balance += balance * 5 / 100;
```

教学点：

✅ DeFi = 状态机 + 数学规则

---

## ⭐ 模块3：收益计算

记录：

```solidity
mapping(address => uint) depositTime;
```

计算：

```solidity
block.timestamp - depositTime[user]
```

理解：

👉 区块链时间。

---

## ⭐ 模块4：事件日志

```solidity
event Deposit(address user, uint amount);
event Withdraw(address user, uint amount);
```

学生学会：

* 链上日志
* 前端监听

---

# 五、Hardhat / Foundry 开发任务

---

## Task 1：Hardhat 部署

```bash
npx hardhat run scripts/deploy.js
```

要求：

* 部署 Token
* 部署 MiniDeFi
* 连接地址

---

## Task 2：Foundry 测试

必须编写：

```solidity
testDeposit()
testWithdraw()
testInterest()
```

加分：

```solidity
Fuzz test
```

---

# 六、前端实验（Web3部分 ⭐）

学生实现简单页面：

```
[连接钱包]
余额：100 MTK

输入金额：[10]

[Deposit]
[Withdraw]
```

使用：

```bash
npm install ethers
```

调用：

```javascript
contract.deposit(amount)
```

学习：

👉 Web2 → Web3 转变。

---

# 七、部署到测试网（必须）

网络：

* Sepolia

学生提交：

* 合约地址
* Etherscan截图

---

# 八、提交成果

学生提交：

```
MiniDeFi/
│
├── contracts/
├── test/
├── frontend/
├── README.md
└── 部署截图
```

---

# 九、评分标准（教师直接用）

| 项目        | 分值 |
| --------- | -- |
| ERC20实现   | 15 |
| DeFi逻辑    | 25 |
| Foundry测试 | 15 |
| Hardhat部署 | 15 |
| 前端交互      | 20 |
| 创新功能      | 10 |

---

# 十、加分方向（优秀学生）

任选：

✅ 借贷功能（Borrow）
✅ 抵押率机制
✅ 清算逻辑
✅ 多Token支持
✅ APY动态变化


---

# 十一、教师隐藏教学目标（关键）

学生实际上会理解：

> **DeFi ≠ 金融**
>
> **DeFi = 可验证状态机 + 密码学账户控制**

这是区块链课程真正的终点。

