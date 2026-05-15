注意：这里的代码是有点小问题的，需要学生主动解决

Foundry 是以太坊生态中**最热门、最高效的智能合约开发工具链**（由 Paradigm 开发），相比 Truffle/Hardhat 更轻量、更快，全程基于命令行，原生支持 Solidity 脚本/测试/部署，非常适合开发者高效迭代。

下面从**环境搭建→核心功能→实战示例**，给了一套新手友好的 Foundry 入门指南，覆盖开发全流程。

---

## 一、先搞懂：Foundry 核心优势
- **极速测试**：用 Rust 编写的执行引擎，测试速度比 Hardhat 快 10-100 倍；
- **原生 Solidity**：测试、部署脚本都可以用 Solidity 写（无需 JS/TS）；
- **一体化工具链**：内置 `forge`（开发/测试）、`cast`（链交互）、`anvil`（本地节点）、`chisel`（Solidity 调试）；
- **无缝兼容**：支持 Hardhat/Truffle 项目迁移，兼容 EVM 所有链（ETH/BSC/Polygon 等）。

---

## 二、环境搭建（Windows/macOS/Linux）
### 1. 安装 Foundry（一键脚本）
#### Linux/macOS
```bash
curl -L https://foundry.paradigm.xyz | bash
# 安装完成后，重启终端，执行：
foundryup
```

#### Windows（推荐 WSL2，原生 Windows 兼容差）
1. 先安装 WSL2 + Ubuntu（参考之前 Docker 配置的 WSL 步骤）；
2. 在 WSL Ubuntu 中执行上述 Linux 安装命令。

### 2. 验证安装
```bash
# 检查 forge 版本
forge --version
# 检查 anvil 本地节点
anvil --version
```
输出版本号即安装成功。

---

## 三、核心命令速览（必记）
Foundry 核心是 4 个工具，日常开发高频使用：

| 工具   | 作用                          | 常用命令示例                  |
|--------|-------------------------------|-------------------------------|
| `forge`| 开发/测试/编译/部署核心工具   | `forge init`/`forge test`     |
| `cast` | 链交互（查区块/调用合约/发交易）| `cast call`/`cast send`       |
| `anvil`| 本地测试节点（替代 Ganache）| `anvil`（启动节点）|
| `chisel`| Solidity 交互式调试器        | `chisel`（启动交互式环境）|

---

## 四、实战：从零创建第一个 Foundry 项目
### 1. 初始化项目
```bash
# 创建项目（命名为 my-foundry-project）
forge init my-foundry-project
# 进入项目目录
cd my-foundry-project
```

### 2. 项目结构解析（核心目录）
```
my-foundry-project/
├── src/          # 智能合约源码（.sol）
├── test/         # 测试文件（.sol，用 Solidity 写测试）
├── script/       # 部署脚本（.sol）
├── foundry.toml  # 配置文件（类似 hardhat.config.js）
└── lib/          # 依赖库（如 OpenZeppelin）
```

### 3. 编写简单合约（src/Counter.sol）
替换 `src/Counter.sol` 内容：
```solidity
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

contract Counter {
    uint256 public count;

    // 增加计数
    function increment() public {
        count += 1;
    }

    // 重置计数
    function reset() public {
        count = 0;
    }
}
```

### 4. 编写测试（test/Counter.t.sol）
Foundry 测试用 **Solidity 编写**，无需 JS：
```solidity
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

// 导入 Foundry 测试库
import "forge-std/Test.sol";
import "../src/Counter.sol";

// 测试合约继承 Test
contract CounterTest is Test {
    Counter public counter;

    // 每个测试前执行（类似 beforeEach）
    function setUp() public {
        counter = new Counter();
    }

    // 测试 increment 功能（函数名以 test 开头）
    function testIncrement() public {
        counter.increment();
        // 断言：count 应该等于 1
        assertEq(counter.count(), 1);
    }

    // 测试 reset 功能
    function testReset() public {
        counter.increment();
        counter.reset();
        assertEq(counter.count(), 0);
    }

    // 可选：测试失败场景（以 testFail 开头）
    function testFailUnderflow() public {
        counter.reset();
        // 执行 count -= 1，应该溢出失败
        counter.count(); // 这里仅示例，实际需写溢出逻辑
    }
}
```

### 5. 运行测试（核心命令）
```bash
# 运行所有测试
forge test

# 详细输出（显示每个测试步骤）
forge test -vvv

# 只运行指定测试（testIncrement）
forge test --match-test testIncrement
```
输出 `PASS` 即测试通过，Foundry 会显示测试耗时（毫秒级）。

### 6. 启动本地节点（anvil）
```bash
# 启动本地测试节点（默认端口 8545）
anvil
```
启动后会显示 10 个测试账户（私钥+地址），用于部署/测试。

### 7. 部署合约到本地节点
#### 步骤1：编写部署脚本（script/Counter.s.sol）
```solidity
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import "../src/Counter.sol";

contract CounterScript is Script {
    function run() public {
        // 加载私钥（anvil 第一个测试账户私钥）
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        // 开始广播交易
        vm.startBroadcast(deployerPrivateKey);
        // 部署合约
        Counter counter = new Counter();
        // 停止广播
        vm.stopBroadcast();

        // 打印合约地址
        console.log("Counter deployed to:", address(counter));
    }
}
```

#### 步骤2：设置私钥（WSL/Linux/macOS）
```bash
# 用 anvil 第一个测试账户私钥（默认：0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80）
export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

#### 步骤3：部署合约
```bash
# 部署到 anvil 本地节点
forge script script/Counter.s.sol:CounterScript --rpc-url http://localhost:8545 --broadcast
```
部署成功后会显示合约地址。

### 8. 用 cast 交互（链上调用合约）
```bash
# 1. 调用 increment 函数（发交易）
cast send <合约地址> "increment()" --private-key $PRIVATE_KEY --rpc-url http://localhost:8545

# 2. 查询 count 值（只读调用）
cast call <合约地址> "count()" --rpc-url http://localhost:8545
```
返回值为 `0x0000000000000000000000000000000000000000000000000000000000000001`（即 1），说明调用成功。

---

## 五、进阶技巧（提升开发效率）
### 1. 安装依赖（如 OpenZeppelin）
```bash
# 安装 OpenZeppelin 合约库
forge install OpenZeppelin/openzeppelin-contracts
```
在合约中导入：
```solidity
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
```

### 2. 调试合约（chisel）
```bash
# 启动交互式 Solidity 调试器
chisel
# 直接写 Solidity 代码，实时执行
> uint256 a = 10;
> a + 5; // 输出 15
```

### 3. 部署到测试网（如 Sepolia）
```bash
# 设置 RPC URL（需申请 Alchemy/Infura API Key）
export RPC_URL=https://sepolia.infura.io/v3/你的API_KEY
# 部署
forge script script/Counter.s.sol:CounterScript --rpc-url $RPC_URL --broadcast --verify
```

---

## 六、常见问题解决
1. **Windows 安装失败**：优先用 WSL2，原生 Windows 需手动配置环境变量；
2. **测试报错“合约不存在”**：检查 `setUp()` 是否正确初始化合约；
3. **部署失败“余额不足”**：anvil 测试账户有 10000 ETH，测试网需先领水龙头；
4. **cast 调用失败**：检查合约地址、函数签名、私钥是否正确。

---

### 总结
1. Foundry 核心是 `forge`（开发测试）、`anvil`（本地节点）、`cast`（链交互），全程命令行+Solidity 开发；
2. 入门流程：初始化项目 → 写合约 → 写 Solidity 测试 → anvil 启动节点 → forge 部署 → cast 交互；
3. 优势是速度快、原生 Solidity 支持，适合以太坊智能合约高效开发。

如果需要针对某个具体场景（如 ERC20 开发、测试网部署、合约调试）深入讲解，我可以补充对应的实战步骤。
