将实验项目从 Hardhat 迁移到 Foundry 是一个非常棒的决定。Foundry 基于 Rust 编写，运行速度极快，并且支持全程使用 Solidity 进行开发、测试和部署，是目前 Web3 开发领域的行业标准工具。

以下是为你量身定制的 Foundry 版去中心化众筹系统实验指南，完美适配原有的实验任务要求：

🛠️ 实验环境准备

首先，你需要安装 Foundry 工具链（Forge, Cast, Anvil, Chisel）。
在终端执行官方一键安装脚本：
curl -L https://foundry.paradigm.xyz | bash
安装完成后，重启终端或执行：
source ~/.bashrc
安装最新版 Foundry
foundryup

验证安装是否成功：forge --version

📂 项目目录结构

使用 Foundry 初始化项目后，标准的目录结构如下（比 Hardhat 更简洁）：
crowdfunding-dapp/
│
├── src/                # 存放智能合约源码 (原 contracts/)
│   └── Crowdfunding.sol
├── test/               # 存放 Solidity 测试文件
│   └── Crowdfunding.t.sol
├── script/             # 存放部署脚本
│   └── Deploy.s.sol
├── frontend/           # 前端保持不变 (React/Vite)
└── foundry.toml        # Foundry 核心配置文件

初始化命令：forge init crowdfunding-dapp

💻 Step 1：编写智能合约 (src/Crowdfunding.sol)

合约逻辑保持不变，只需将文件移动到 src/ 目录下。

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
        require(block.timestamp = goal, "Goal not reached");
        payable(owner).transfer(address(this).balance);
    }

    function refund() public {
        require(block.timestamp > deadline, "Not ended");
        require(totalFunds  0, "No contribution");
        contributions[msg.sender] = 0;
        payable(msg.sender).transfer(amount);
    }
}

🧪 Step 2：编写 Solidity 测试 (test/Crowdfunding.t.sol)

Foundry 的核心优势之一是用 Solidity 写测试，速度极快且原生支持模糊测试（Fuzzing）。

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/Crowdfunding.sol";

contract CrowdfundingTest is Test {
    Crowdfunding public crowd;
    address owner = address(1);
    address investor = address(2);

    function setUp() public {
        vm.startPrank(owner);
        // 目标 1 ETH，持续时间 1 天
        crowd = new Crowdfunding(1 ether, 1 days);
        vm.stopPrank();
    }

    function testContribute() public {
        vm.startPrank(investor);
        crowd.contribute{value: 0.5 ether}();
        assertEq(crowd.totalFunds(), 0.5 ether);
        assertEq(crowd.contributions(investor), 0.5 ether);
        vm.stopPrank();
    }

    // 模糊测试：随机金额投资
    function testFuzzContribute(uint256 amount) public {
        amount = bound(amount, 1, 1 ether); // 限制金额范围
        vm.startPrank(investor);
        crowd.contribute{value: amount}();
        assertEq(crowd.totalFunds(), amount);
        vm.stopPrank();
    }

    function testWithdrawSuccess() public {
        // 模拟投资达标
        vm.startPrank(investor);
        crowd.contribute{value: 1.5 ether}();
        vm.stopPrank();

        // 项目方提款
        vm.startPrank(owner);
        uint256 initialBalance = owner.balance;
        crowd.withdraw();
        assertEq(owner.balance, initialBalance + 1.5 ether);
        vm.stopPrank();
    }
}

运行测试命令：forge test -vvv（-vvv 可以查看详细日志）

🚀 Step 3：部署脚本与本地节点 (script/Deploy.s.sol)

Foundry 使用 Anvil 作为本地节点（替代 Ganache），使用 Solidity 脚本进行部署。

1. 编写部署脚本：
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/Crowdfunding.sol";

contract DeployScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // 目标 5 ETH，持续 1 小时
        Crowdfunding crowd = new Crowdfunding(5 ether, 1 hours);
        
        console.log("Crowdfunding deployed to:", address(crowd));

        vm.stopBroadcast();
    }
}

2. 启动本地链并部署：
打开两个终端窗口：
终端 1 (启动 Anvil 本地链)：
        anvil
    
   Anvil 会输出 10 个测试账户和私钥，以及 RPC 地址 http://127.0.0.1:8545。

终端 2 (执行部署)：
        # 导出 Anvil 输出的第一个账户私钥
    export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
    # 执行部署脚本
    forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --broadcast
    
   部署成功后，终端会打印出合约地址，请复制保存。

🌐 Step 4：前端连接 (frontend/)

前端逻辑基本不变，但需要更新合约地址。

MetaMask 设置：
    网络名称：Anvil
    RPC URL：http://127.0.0.1:8545
    Chain ID：31337
    导入 Anvil 输出的测试账户私钥。

修改 frontend/app.js：
    将 contractAddress 替换为你在上一步 forge script 部署后得到的地址。
   (JS 代码部分保持你原有的 ethers.js 写法即可，无需改动)

⚡ 完整运行顺序总结

编译与测试：
        forge build
    forge test
    
启动本地链：
        anvil
    
部署合约 (新终端)：
        export PRIVATE_KEY=
    forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --broadcast
    
启动前端：
        cd frontend
    npm run dev
    
浏览器交互：打开 http://localhost:5173，连接 MetaMask 进行投资和提款测试。

按照这个流程，你的实验不仅顺利完成了迁移，还体验到了 Foundry 极速编译和 Solidity 原生测试的强大优势！



这是为你准备的完整 frontend/app.js 代码。它包含了连接钱包、投资、项目方提款以及投资人退款的所有核心逻辑，并且完美适配你之前的 Crowdfunding.sol 合约。

你可以直接复制以下代码覆盖原有的 frontend/app.js：

import { ethers } from "ethers";

// ⚠️ 注意：请将这里替换为你使用 forge script 部署后得到的真实合约地址
const contractAddress = "YOUR_DEPLOYED_CONTRACT_ADDRESS";

// 合约的 ABI（包含了合约中所有的函数和事件）
const abi = [
  // 只读函数
  "function owner() view returns (address)",
  "function goal() view returns (uint)",
  "function deadline() view returns (uint)",
  "function totalFunds() view returns (uint)",
  "function contributions(address) view returns (uint)",
  
  // 写入函数
  "function contribute() payable",
  "function withdraw()",
  "function refund()"
];

let contract;
let signer;

// 1. 连接钱包并初始化合约
export async function connectWallet() {
  if (typeof window.ethereum === "undefined") {
    alert("请先安装 MetaMask 钱包！");
    return;
  }

  try {
    // 请求用户授权连接钱包
    await window.ethereum.request({ method: "eth_requestAccounts" });
    
    const provider = new ethers.BrowserProvider(window.ethereum);
    signer = await provider.getSigner();
    
    // 实例化合约
    contract = new ethers.Contract(contractAddress, abi, signer);
    
    console.log("钱包连接成功，合约已初始化");
    alert("钱包连接成功！");
    
    // 连接成功后，自动刷新页面上的众筹信息
    updateCrowdInfo();
  } catch (error) {
    console.error("连接失败:", error);
    alert("连接钱包失败，请查看控制台报错");
  }
}

// 2. 用户投资 (contribute)
export async function contribute() {
  if (!contract) return alert("请先连接钱包！");
  
  const amount = document.getElementById("ethAmount").value;
  if (!amount || amount 

  
  Foundry 众筹 DApp

  去中心化众筹系统
  
  🔗 连接钱包
  

  众筹状态
  目标金额：加载中...
  当前已筹：加载中...
  截止时间：加载中...
  我的投资：未连接
  

  操作区
  投资金额 (ETH): 
  
  💰 投资
  
  
  🏦 项目方提款 (仅Owner)
  ↩️ 申请退款 (未达标时)

  
  

💡 几个重要提示：

替换合约地址：记得把 app.js 顶部的 YOUR_DEPLOYED_CONTRACT_ADDRESS 换成你用 forge script 部署后打印出来的真实地址。
单位转换：Foundry 和 Ethers.js v6 中，前端传入的 ETH 需要用 ethers.parseEther() 转换，合约返回的 Wei 需要用 ethers.formatEther() 转换回可读的 ETH 单位。
错误提示：代码中加入了 error.reason 的捕获，这样如果 Solidity 里的 require("Ended") 或 require("Not owner") 触发时，前端会直接弹出对应的中文报错原因，非常便于调试。
