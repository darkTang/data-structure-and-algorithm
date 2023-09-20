# 宽度优先搜索（BFS）
- 分层遍历
  - 一层一层的遍历一个图、树、矩阵
  - 题目
    - Binary Tree Level Order Traversal
    - [Clone Graph](https://leetcode.cn/problems/clone-graph/)
- 简单图最短路径
  - 简单图的定义是，图中所有的边长都一样
  - 题目
    - [Word Ladder](https://leetcode.cn/problems/om3reC/description/)
- 连通块问题
  - 通过图中一个点找到其他所有连通的点
  - 找到所有方案问题的一种非递归实现方式
  - 题目
    - [Number of Islands](https://leetcode.cn/problems/number-of-islands/description/)
- 拓扑排序
  - 求任意拓扑序
  - 求是否有拓扑序
  - 求字典序最小的拓扑序
  - 求是否唯一拓扑序

## 简单图
- 没有方向
- 没有权重
- 两点之间最多有一条边
- 不存在自己连自己

## BFS算法通用模板（图/树）

```go
package bfs

import (
  "container/list"
)


func bfs(node *UndirectedGraphNode) {
  st := list.New()
  distance := make(map[*UndirectedGraphNode]int)
  
  // 初始化，把初始节点放到st里，如果有多个就都放进去
  // 并标记初始节点的距离为0，记录在 distance 的hashmap里
  // distance作用：一是判断是否已经访问过 二是记录离起点的距离
  // 一旦入队，需要立马标记
  st.PushBack(node)
  distance[node] = 0
  var curNode *UndirectedGraphNode
  
  for st.Len() > 0 {
	  curNode = st.Remove(st.Front()).(*UndirectedGraphNode)
	  for _, neighbor := range curNode.getNeighbors() {
		  if _, ok := distance[neighbor]; ok {
			  continue
          }
		  distance[neighbor] = distance[node] + 1
		  st.PushBack(neighbor)
      }
  }
}
// 时间复杂度 O(n+m) n个点+m条边
// 当 m >> n 时，时间复杂度为 O(m)
```