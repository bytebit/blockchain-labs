### 实验名称：PBFT 共识算法仿真与容错性分析实验

#### 1. 实验目的
* 理解 PBFT 算法的三个核心阶段及其状态机复制原理。
* 验证 PBFT 的容错边界：$n \ge 3f + 1$。
* 观察在存在“恶意节点”（不发消息、发伪造消息）时，系统如何达成一致性。

#### 2. 实验环境
* **编程语言**：Python (使用 `asyncio` 模拟异步网络) 或 Node.js。
* **仿真架构**：在一台机器上启动多个进程/线程，模拟 $N$ 个共识节点，通过网络套接字（Socket）或简单的消息队列通信。

---

### 3. 核心实验步骤设计

#### 第一阶段：基础原型实现（正常流程）
1.  **节点初始化**：启动 4 个节点（ID: 0, 1, 2, 3），节点 0 设为 **Primary (主节点)**。
2.  **Request**：客户端向主节点发送一条请求 `{"data": "Hello PBFT"}`。
3.  **Pre-prepare**：主节点广播预准备消息。
4.  **Prepare**：各节点接收后广播准备消息，并等待收集 $2f$ 个一致的准备消息。
5.  **Commit**：进入提交阶段，再次收集 $2f+1$ 个确认消息。
6.  **Reply**：节点执行指令并返回结果给客户端。

#### 第二阶段：容错性实验（恶意节点模拟）
设置变量 $f=1$（4个节点中允许1个恶意节点）：
1.  **静默攻击**：设置节点 3 为“宕机”状态（不回复任何消息），观察系统是否仍能达成共识。
2.  **拜占庭攻击（双花/伪造）**：设置节点 2 为恶意节点，使其在 Prepare 阶段向节点 A 发送“同意”，向节点 B 发送“反对”，观察 PBFT 如何通过法定人数（Quorum）过滤掉该错误。

#### 第三阶段：视图变更（View Change）实验
1.  **主节点作恶**：人为让主节点（节点 0）停止发送 Pre-prepare 消息。
2.  **超时触发**：从节点检测到超时，发起 `View-Change` 协议。
3.  **选举新主**：节点 1 成为新的 Primary，恢复系统运作。

---

### 4. 实验关键代码逻辑（Python 伪代码）

```python
class PBFTNode:
    def __init__(self, node_id, total_nodes):
        self.id = node_id
        self.f = (total_nodes - 1) // 3
        self.state = "IDLE" # 状态机控制
        self.prepare_msgs = {} # 收集准备消息
        
    async def handle_pre_prepare(self, msg):
        # 验证签名和序列号
        print(f"Node {self.id} received Pre-Prepare from {msg.leader}")
        self.broadcast_prepare(msg.data)
        
    async def handle_prepare(self, msg):
        self.prepare_msgs[msg.digest] += 1
        # 关键逻辑：判断是否达到 2f 个消息
        if self.prepare_msgs[msg.digest] >= 2 * self.f:
            self.broadcast_commit()

    def set_malicious(self, mode):
        # 设置攻击模式：'silent' 或 'lie'
        self.mode = mode
```

---

### 5. 实验观察与指标（用于展示）

| 实验场景 | 预期结果 | 核心指标 |
| :--- | :--- | :--- |
| **全员正常** | 快速达成共识 | 消息复杂度 $O(N^2)$ |
| **$f$ 个节点宕机** | 依然能正常达成共识 | 延迟略微增加 |
| **超过 $f$ 个节点宕机** | 系统陷入停滞（Liveness 失败） | 无法收集足够签名 |
| **主节点作恶** | 触发 View Change，新主上任 | 视图切换所需时间 |

---

### 6. 进阶思考（适合 PPT 结尾提问）
* **网络分区**：如果在 Prepare 阶段发生了网络分区，PBFT 的安全性（Safety）是如何保证的？
* **对比 HotStuff**：为什么现代区块链（如 Libra/Aptos）更倾向于使用优化版的 PBFT（HotStuff），其线性通信复杂度是如何实现的？
合作为**分布式系统**或**计算机网络**课程的一部分，在 Linux 环境下运行多节点仿真。

您需要我为您提供更详细的 **View Change** 状态转换逻辑图吗？
