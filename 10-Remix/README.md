## 一、实验基本信息
- **实验名称**：基于 Remix IDE 开发与部署第一个 Solidity 智能合约
- **实验时长**：45 分钟
- **实验环境**：浏览器 + Remix IDE（https://remix.ethereum.org）
- **实验难度**：⭐（零基础可做）
- **核心目标**：认识 Remix、编写合约、编译、部署、调用函数

---

## 二、实验原理
1. Remix 是以太坊官方**浏览器版智能合约 IDE**，支持一键编译、部署、测试合约。
2. Solidity 是智能合约编程语言，本次实现最简单的**存值/取值/改值**功能。
3. 实验使用 Remix 内置的**测试虚拟机**，无需真实区块链、无任何费用。

---

## 三、实验步骤（老师/学生可直接跟着做）
### 步骤 1：打开 Remix IDE
在浏览器访问：
https://remix.ethereum.org
进入后直接点击 **Accept** 同意条款，进入主界面。

### 步骤 2：创建合约文件
1. 左侧点击 **Files** 面板
2. 在 `contracts` 文件夹上右键 → **New File**
3. 文件名输入：`StudentData.sol`

### 步骤 3：编写实验用智能合约（可直接复制）
这是专门为大学生设计的**极简合约**，功能：存储学生姓名、学号，可修改可查询。

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;  // 声明Solidity版本

// 学生信息智能合约
contract StudentInfo {
    // 定义状态变量（存在区块链上）
    string public studentName;  // 学生姓名
    uint public studentID;      // 学生学号

    // 构造函数：部署合约时自动执行一次
    constructor() {
        studentName = "TestStudent";
        studentID = 20260001;
    }

    // 函数1：修改学生信息
    function setStudentInfo(string memory _name, uint _id) public {
        studentName = _name;
        studentID = _id;
    }

    // 函数2：查询学生姓名
    function getName() public view returns(string memory) {
        return studentName;
    }

    // 函数3：查询学生学号
    function getID() public view returns(uint) {
        return studentID;
    }
}
```

### 步骤 4：编译合约
1. 左侧点击 **Solidity Compiler**（黄色图标）
2. 编译器版本选择：`0.8.18`（与代码一致）
3. 直接点击 **Compile StudentData.sol**
4. 出现绿色对勾 ✅ 表示编译成功

### 步骤 5：部署合约（核心步骤）
1. 左侧点击 **Deploy & Run Transactions**（小电脑图标）
2. **环境选择**：`Remix VM (London)`（内置测试虚拟机，免费）
3. 点击 **Deploy** 按钮
4. 下方控制台出现交易记录 → 部署成功

### 步骤 6：调用合约函数（实验操作）
部署后在下方 **Deployed Contracts** 里可以看到函数按钮：
1. **studentName / studentID**：点击直接查看初始值
2. **setStudentInfo**：输入你的姓名和学号，点击**Transact**执行
3. **getName / getID**：点击查看修改后的值

---

## 四、实验操作演示（学生必看）
1. 部署后点击 `studentName` → 显示：`TestStudent`
2. 输入：`"ZhangSan"` `202612345` → 点击 `setStudentInfo`
3. 再次点击 `studentName` → 显示你输入的名字
4. 点击 `studentID` → 显示你输入的学号

> 说明：数据已经**永久保存在测试区块链**上，实现了最简单的链上存储。

---

## 五、实验要求与提交内容
### 1. 必须完成的操作
- [ ] 创建合约文件
- [ ] 复制并编译代码
- [ ] 部署到 Remix VM
- [ ] 修改自己的姓名与学号
- [ ] 成功查询修改后的数据

### 2. 实验报告提交（3张截图即可）
1. Remix 中编写好的合约代码截图
2. 合约编译成功截图
3. 部署后调用函数、显示自己信息的截图

---

## 六、实验思考题（老师可用）
1. Remix VM 是什么？为什么不需要花钱？
2. `view` 函数为什么不需要消耗燃气费？
3. 合约中的状态变量和普通变量有什么区别？

---

## 七、极简教学话术（老师上课直接用）
1. Remix 是浏览器里的合约开发工具，打开就能用。
2. 我们写的合约可以存数据、改数据、读数据。
3. 点编译 → 点部署 → 点按钮调用函数，完成实验。

---

### 总结
这是**最适合大学生的 Remix 入门实验**：
- 零配置、零成本、零报错
- 代码简单、逻辑清晰、可直接运行
- 覆盖 Remix 核心操作：编写→编译→部署→调用
