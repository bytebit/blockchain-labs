# 三、实验任务

学生需要实现一个 **去中心化众筹系统**：

## 功能要求

### 1️⃣ 创建众筹

项目方设置：

* 众筹目标金额
* 截止时间
* 收款地址

---

### 2️⃣ 用户投资

投资人：

* 发送 ETH
* 自动记录投资金额

---

### 3️⃣ 众筹成功

若：

```
筹资 ≥ 目标金额
```

→ 项目方可提款

---

### 4️⃣ 众筹失败

若：

```
截止时间到 + 未达目标
```

→ 投资人可退款

---

# 四、系统架构

```
Frontend (React)
        ↓
Web3.js / Ethers.js
        ↓
Crowdfunding.sol
        ↓
Ethereum (Local/Ganache)
```

---

# 五、实验环境

| 工具       | 用途     |
| -------- | ------ |
| Node.js  | 前端运行   |
| Hardhat  | 合约开发   |
| MetaMask | 钱包     |
| Ganache  | 本地区块链  |
| React    | DApp界面 |

---

# 六、实验步骤

---

## Step 1：创建项目

```bash
mkdir crowdfunding-dapp
cd crowdfunding-dapp

npm init -y
npm install --save-dev hardhat
npx hardhat
```

选择：

```
Create JavaScript project
```

---

## Step 2：编写智能合约

### contracts/Crowdfunding.sol

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Crowdfunding {

    address public owner;
    uint public goal;
    uint public deadline;
    uint public totalFunds;

    mapping(address => uint) public contributions;

    constructor(uint _goal, uint _duration) {
        owner = msg.sender;
        goal = _goal;
        deadline = block.timestamp + _duration;
    }

    function contribute() public payable {
        require(block.timestamp < deadline, "Ended");

        contributions[msg.sender] += msg.value;
        totalFunds += msg.value;
    }

    function withdraw() public {
        require(msg.sender == owner);
        require(totalFunds >= goal);
        payable(owner).transfer(address(this).balance);
    }

    function refund() public {
        require(block.timestamp > deadline);
        require(totalFunds < goal);

        uint amount = contributions[msg.sender];
        contributions[msg.sender] = 0;

        payable(msg.sender).transfer(amount);
    }
}
```

---

## Step 3：部署脚本

`scripts/deploy.js`

```javascript
async function main() {
  const CF = await ethers.getContractFactory("Crowdfunding");

  const contract = await CF.deploy(
      ethers.parseEther("5"), // 5 ETH goal
      3600                     // 1 hour
  );

  await contract.waitForDeployment();

  console.log("Deployed:", await contract.getAddress());
}

main();
```

运行：

```bash
npx hardhat node
npx hardhat run scripts/deploy.js --network localhost
```

---

## Step 4：前端连接合约

安装：

```bash
npm install ethers
```

---

### frontend/app.js

```javascript
import { ethers } from "ethers";

const contractAddress = "YOUR_ADDRESS";
const abi = [...];

async function contribute() {
  const provider = new ethers.BrowserProvider(window.ethereum);
  const signer = await provider.getSigner();

  const contract = new ethers.Contract(
      contractAddress,
      abi,
      signer
  );

  await contract.contribute({
      value: ethers.parseEther("0.1")
  });
}
```

---

## Step 5：MetaMask 交互测试

测试：

* 多账户投资
* 达标提款
* 未达标退款

---

# 七、实验结果要求

学生需提交：

✅ 合约源码
✅ 部署截图
✅ 投资交易截图
✅ 成功或退款演示视频
✅ 实验报告

---

怎么运行前端？

---

# ✅ 一、先确认目录结构（标准实验结构）

建议你的项目长这样：

```
crowdfunding-dapp/
│
├── contracts/
├── scripts/
├── hardhat.config.js
│
└── frontend/
    ├── index.html
    ├── app.js
    └── package.json
```

---

# ✅ 二、创建前端环境（只做一次）

进入 frontend：

```bash
cd frontend
```

初始化：

```bash
npm init -y
```

安装依赖：

```bash
npm install ethers
npm install --save-dev vite
```

为什么用 **Vite**？

👉 比 webpack 简单
👉 学生几乎零配置
👉 Web3 官方推荐

---

# ✅ 三、修改 package.json

在 `frontend/package.json` 里加入：

```json
{
  "scripts": {
    "dev": "vite"
  }
}
```

---

# ✅ 四、创建 index.html（必须）

前端入口不是 JS，而是 HTML。

## frontend/index.html

```html
<!DOCTYPE html>
<html>
<head>
  <title>Crowdfunding DApp</title>
</head>

<body>

<h2>Crowdfunding</h2>

<button onclick="connect()">连接钱包</button>
<button onclick="contribute()">投资 0.1 ETH</button>

<script type="module" src="/app.js"></script>

</body>
</html>
```

⚠️ 注意：

```
type="module"
```

必须有，否则 import 不工作。

---

# ✅ 五、修改 app.js（关键）

## frontend/app.js

```javascript
import { ethers } from "ethers";

let contract;

const contractAddress = "你的部署地址";

const abi = [
  "function contribute() payable"
];

export async function connect() {

  await window.ethereum.request({
    method: "eth_requestAccounts"
  });

  const provider =
      new ethers.BrowserProvider(window.ethereum);

  const signer = await provider.getSigner();

  contract = new ethers.Contract(
      contractAddress,
      abi,
      signer
  );

  console.log("Wallet connected");
}

export async function contribute() {

  await contract.contribute({
      value: ethers.parseEther("0.1")
  });

  alert("Investment sent!");
}

window.connect = connect;
window.contribute = contribute;
```

---

# ✅ 六、启动前端（真正运行）

在 frontend 目录执行：

```bash
npm run dev
```

你会看到：

```
VITE v5.x ready in xxx ms

➜ Local: http://localhost:5173/
```

打开浏览器：

```
http://localhost:5173
```

---

# ✅ 七、MetaMask 设置（很多人忘）

确保：

### 1️⃣ Hardhat 本地链运行中

```bash
npx hardhat node
```

---

### 2️⃣ MetaMask 添加网络

```
Network Name: Hardhat
RPC URL: http://127.0.0.1:8545
Chain ID: 31337
Currency: ETH
```

---

### 3️⃣ 导入测试账户

复制 hardhat node 输出的 private key：

```
Account #0: 0xf39f...
Private Key: 0xac0974...
```

导入 MetaMask。

---

# ✅ 八、完整运行顺序（课堂标准流程）

按顺序：

```
① npx hardhat node
② npx hardhat run deploy.js --network localhost
③ cd frontend
④ npm run dev
⑤ 打开浏览器
⑥ 连接 MetaMask
⑦ 点击投资
```

---

# 🚨 常见错误（90%学生会遇到）

## ❌ Cannot use import outside module

👉 index.html 没写：

```html
type="module"
```

---

## ❌ window.ethereum undefined

👉 没安装 MetaMask

---

## ❌ transaction reverted

👉 合约地址写错

---

## ❌ insufficient funds

👉 MetaMask 账户不是 hardhat 测试账户

---

# 🎓 教学建议（非常重要）

课堂上建议让学生运行：

```
Hardhat node + Vite
```

不要用：

* ❌ 直接打开 html
* ❌ Live Server
* ❌ file:// 路径

否则 Web3 注入会失败。
